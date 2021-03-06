package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Status codes
const (
	OK             = 200
	ACCESS_DENIED  = 403
	NOT_ALLOWED    = 405
	BAD_REQUEST    = 400
	INTERNAL_ERROR = 500
)

type ErrorResponse struct {
	State string `json:"state"`
	Error string `json:"error"`
}

// readRequestBody ...
func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, err
}

// errorResponse ...
func jsonResponse(w http.ResponseWriter, code int, resp interface{}) {
	j, _ := json.Marshal(resp)
	w.WriteHeader(code)
	w.Write(j)
}
