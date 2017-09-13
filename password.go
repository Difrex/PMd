package main

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

// readPassword ...
func readPassword(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	return string(f)
}
