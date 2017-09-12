package api

import (
	"encoding/json"
	"net/http"
)

type RegRequest struct {
	PubKey string `json:"pubkey"`
}

type RegResponse struct {
	Error  string `json:"error"`
	UserID int    `json:"userid"`
}

// RegistrationHandler ...
func (conf ApiConf) registrationHandler(w http.ResponseWriter, r *http.Request) {
	var req RegRequest
	if r.Method != "POST" {
		jsonResponse(w, NOT_ALLOWED, RegResponse{"Method not allowed", -1})
		return
	}

	body, err := readRequestBody(r)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, RegResponse{err.Error(), -1})
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, RegResponse{err.Error(), -1})
		return
	}

	gpgid, err := importPubKey(req.PubKey)
	if err != nil {
		jsonResponse(w, INTERNAL_ERROR, RegResponse{err.Error(), -1})
		return
	}

	userID, err := conf.userAdd(gpgid)
	if err != nil {
		jsonResponse(w, INTERNAL_ERROR, RegResponse{err.Error(), -1})
		return
	}

	jsonResponse(w, OK, RegResponse{"", userID})
}
