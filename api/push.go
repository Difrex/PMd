package api

import (
	"encoding/json"
	"net/http"

	"github.com/Difrex/PMd/storage"
	log "github.com/Sirupsen/logrus"
)

type AddData struct {
	Data string `json:"data"`
}

type AddResponse struct {
	Error string `json:"error"`
	State string `json:"state"`
}

// addDataHandler ...
func (conf ApiConf) addDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(w, NOT_ALLOWED, ErrorResponse{"Method not allowed", "error"})
		return
	}

	var req AddData

	data, err := readRequestBody(r)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{err.Error(), "error"})
		return
	}

	// Verify signature
	gpgid, content, err := verifyAndDetach(string(data))
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{err.Error(), "error"})
		return
	}

	err = json.Unmarshal(content, &req)
	if err != nil {
		jsonResponse(w, BAD_REQUEST, ErrorResponse{err.Error(), "error"})
		return
	}

	var encrypted string
	if conf.Password == "" {
		encrypted, err = conf.encryptArmour(req.Data, gpgid)
		if err != nil {
			log.Error(err.Error())
			jsonResponse(w, INTERNAL_ERROR, ErrorResponse{"Smth went wrong! Cant encrypt data! Abort...", "error"})
			return
		}
	}
	state, err := conf.Db.AddData(storage.User{
		GPGID: gpgid,
	}, encrypted)
	if err != nil {
		jsonResponse(w, INTERNAL_ERROR, ErrorResponse{"Data already present", "error"})
		return
	}

	jsonResponse(w, OK, state)
}
