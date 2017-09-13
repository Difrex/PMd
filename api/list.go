package api

import (
	"encoding/json"
	"net/http"

	"github.com/Difrex/PMd/storage"
)

type ListResponse struct {
	State    string         `json:"state"`
	Versions []storage.List `json:"versions"`
}

// listDataHandler ...
func (conf ApiConf) listDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(w, NOT_ALLOWED, ErrorResponse{
			State: "error",
			Error: "Method " + r.Method + " not allowed.",
		})
	}

	var resp ListResponse

	body, err := readRequestBody(r)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{
			State: "error",
			Error: "Cant read response.",
		})
		return
	}

	// Verify
	gpgid, _, err := verifyAndDetach(string(body))
	if err != nil {
		jsonResponse(w, ACCESS_DENIED, ErrorResponse{
			State: "error",
			Error: "Access denied. May be forgot register?",
		})
		return
	}

	list, err := conf.Db.ListVersions(storage.User{
		GPGID: gpgid,
	})
	if err != nil {
		jsonResponse(w, INTERNAL_ERROR, ErrorResponse{
			State: "error",
			Error: "Internal error. Try again later.",
		})
		return
	}

	resp.State = "ok"
	resp.Versions = list
	js := Marshal(resp)

	encrypted, err := conf.encryptArmour(string(js), gpgid)
	if err != nil {
		jsonResponse(w, INTERNAL_ERROR, ErrorResponse{
			State: "error",
			Error: "Cant encrypt data for " + gpgid,
		})
		return
	}
	w.WriteHeader(OK)
	w.Write([]byte(encrypted))
}

// Marshal ...
func Marshal(data interface{}) []byte {
	j, _ := json.Marshal(data)
	return j
}
