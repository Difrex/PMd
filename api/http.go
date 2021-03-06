package api

import (
	"net/http"
	"time"

	"github.com/Difrex/PMd/storage"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type ApiConf struct {
	Listen   string
	Db       *storage.DB
	GPGID    string
	Password string
}

// Server ...
func (conf ApiConf) Serve() {
	r := mux.NewRouter()

	r.HandleFunc("/register", conf.registrationHandler)
	r.HandleFunc("/push", conf.addDataHandler)
	r.HandleFunc("/pull", conf.pullHandler)
	r.HandleFunc("/list", conf.listDataHandler)
	r.HandleFunc("/purge", conf.purgeUserHandler)
	r.HandleFunc("/delete", conf.deleteDataHandler)

	http.Handle("/", r)

	srv := http.Server{
		Handler:      r,
		Addr:         conf.Listen,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Listening API on ", conf.Listen)
	log.Fatal(srv.ListenAndServe())
}
