package api

import (
	"encoding/json"
	"net/http"

	"github.com/Difrex/PMd/storage"
)

type DeleteDataRequest struct {
	Version string `json:"version"`
}

// deleteDataHandler ...
func (conf ApiConf) deleteDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(w, NOT_ALLOWED, ErrorResponse{"error", "Method " + r.Method + " not allowed"})
		return
	}

	data, err := readRequestBody(r)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{"error", err.Error()})
		return
	}

	gpgid, content, err := verifyAndDetach(string(data))
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{"error", err.Error()})
		return
	}

	var req DeleteDataRequest
	err = json.Unmarshal([]byte(content), &req)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{"error", err.Error()})
		return
	}

	err = conf.Db.DeleteData(storage.User{
		GPGID: gpgid,
	}, req.Version)
	if err != nil {
		jsonResponse(w, INTERNAL_ERROR, ErrorResponse{"error", "Cant delete data. Try again later."})
		return
	}

	jsonResponse(w, OK, map[string]string{
		"state":   "ok",
		"version": req.Version,
	})
}
