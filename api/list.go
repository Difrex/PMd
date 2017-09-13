package api

import (
	"net/http"

	"github.com/Difrex/PMd/storage"
)

type ListResponse struct {
	State    string         `json:"state"`
	Versions []storage.List `json:"versions"`
}

// listDataHandler ...
func (conf ApiConf) listDataHandler(w http.ResponseWriter, r *http.Request) {
	var resp ListResponse

	body, err := readRequestBody(r)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ListResponse{
			State: "error",
		})
		return
	}

	// Verify
	gpgid, _, err := verifyAndDetach(string(body))
	if err != nil {
		jsonResponse(w, ACCESS_DENIED, ListResponse{
			State: "error",
		})
		return
	}

	list, err := conf.Db.ListVersions(storage.User{
		GPGID: gpgid,
	})
	if err != nil {
		jsonResponse(w, INTERNAL_ERROR, ListResponse{
			State: "error",
		})
		return
	}

	resp.State = "ok"
	resp.Versions = list
	jsonResponse(w, OK, resp)
}
