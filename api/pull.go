package api

import (
	"encoding/json"
	"net/http"

	"github.com/Difrex/PMd/storage"
)

type PullRequest struct {
	Version string `json:"version"`
}

// pullHandler ...
func (conf ApiConf) pullHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(w, NOT_ALLOWED, ErrorResponse{
			Error: "Method " + r.Method + " not allowed.",
			State: "error",
		})
		return
	}

	body, err := readRequestBody(r)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{
			Error: err.Error(),
			State: "error",
		})
		return
	}

	gpgid, content, err := verifyAndDetach(string(body))
	if err != nil {
		jsonResponse(w, ACCESS_DENIED, ErrorResponse{
			Error: err.Error(),
			State: "error",
		})
		return
	}

	var req PullRequest
	err = json.Unmarshal(content, &req)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{
			Error: err.Error(),
			State: "error",
		})
		return
	}

	data, err := conf.Db.GetUserData(storage.User{
		GPGID: gpgid,
	}, req.Version)
	if err != nil {
		jsonResponse(w, INTERNAL_ERROR, ErrorResponse{
			Error: "Smth went wrong. Try again later.",
			State: "error",
		})
		return
	}

	w.WriteHeader(OK)
	w.Write([]byte(data))
}
