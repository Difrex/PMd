package api

import (
	"errors"

	"github.com/Difrex/gpg"
)

// importPubKey ...
func importPubKey(pubkey string) (string, error) {
	gpgid, err := gpg.ImportPubkey(pubkey)
	if err != nil {
		e := errors.New(gpgid)
		return "", e
	}
	return gpgid, nil
}
