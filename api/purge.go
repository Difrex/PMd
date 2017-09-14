package api

import (
	"net/http"

	"github.com/Difrex/PMd/storage"
)

type PurgeResponse struct {
	State string `json:"state"`
	GPGID string `json:"gpgid"`
}

// purgeUserHandler ...
func (conf ApiConf) purgeUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(w, NOT_ALLOWED, ErrorResponse{"error", "Method " + r.Method + " not allowed"})
		return
	}

	data, err := readRequestBody(r)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{"error", err.Error()})
		return
	}

	gpgid, _, err := verifyAndDetach(string(data))
	if err != nil {
		jsonResponse(w, ACCESS_DENIED, ErrorResponse{"error", "Cant verify signature!"})
		return
	}

	err = conf.Db.PurgeUser(storage.User{
		GPGID: gpgid,
	})
	if err != nil {
		jsonResponse(w, INTERNAL_ERROR, ErrorResponse{"error", "USER NOT DELETED. PANIC!!!"})
		return
	}
	jsonResponse(w, OK, PurgeResponse{"ok", gpgid})
}
