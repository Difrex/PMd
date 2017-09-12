package main

import (
	"flag"

	"github.com/Difrex/PMd/api"
	"github.com/Difrex/PMd/storage"
)

var (
	Listen, Storage, GpgID string
)

// init ...
func init() {
	flag.StringVar(&Listen, "listen", "127.0.0.1:2282", "Address to listen")
	flag.StringVar(&Storage, "storage", "./pmd.db", "Path to database")
	flag.StringVar(&GpgID, "gpgid", "", "Server GpgID")
	flag.Parse()
}

// main ...
func main() {

	db := storage.CheckAndInit(Storage)

	conf := api.ApiConf{
		Listen: Listen,
		Db:     db,
		GPGID:  GpgID,
	}

	conf.Serve()
}
