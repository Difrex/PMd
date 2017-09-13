package main

import (
	"flag"

	"github.com/Difrex/PMd/api"
	"github.com/Difrex/PMd/storage"
	log "github.com/Sirupsen/logrus"
)

var (
	Listen, Storage, GpgID, Passfile string
)

// init ...
func init() {
	flag.StringVar(&Listen, "listen", "127.0.0.1:2282", "Address to listen")
	flag.StringVar(&Storage, "storage", "./pmd.db", "Path to database")
	flag.StringVar(&GpgID, "gpgid", "", "Server GpgID")
	flag.StringVar(&Passfile, "passwordfile", "", "Path to file with private key password")
	flag.Parse()
}

// main ...
func main() {
	db := storage.CheckAndInit(Storage)

	if GpgID == "" {
		log.Fatal("gpgid connot be empty")
	}
	var password string
	if Passfile != "" {
		password = readPassword(Passfile)
	}

	conf := api.ApiConf{
		Listen:   Listen,
		Db:       db,
		GPGID:    GpgID,
		Password: password,
	}

	conf.Serve()
}
