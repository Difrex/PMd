package api

import (
	"errors"

	"github.com/Difrex/gpg"
	log "github.com/Sirupsen/logrus"
)

// importPubKey ...
func importPubKey(pubkey string) (string, error) {
	gpgid, err := gpg.ImportPubkey(pubkey)
	if err != nil {
		e := errors.New(gpgid)
		return "", e
	}

	key, err := gpg.ShowKey(gpgid)
	if err != nil {
		return "", err
	}

	log.Warn(key.ID)
	return key.ID, nil
}

// verifySignature ...
func verifyAndDetach(data string) (string, []byte, error) {
	// Verify
	gpgid, err := gpg.Verify(data)
	if err != nil {
		log.Error(err.Error())
		e := errors.New("Cant verify signature")
		return "", []byte(""), e
	}

	// Extract data
	content, err := gpg.ExtractDataFromSigned(data)
	if err != nil {
		log.Info(content.String())
		log.Error(err.Error())
		e := errors.New("Cant detach signature")
		return "", []byte(""), e
	}

	// Get long gpgid
	key, err := gpg.ShowKey(gpgid)
	if err != nil {
		log.Error(err.Error())
		e := errors.New("Cannot show key")
		return "", []byte(""), e
	}

	return key.ID, content.Bytes(), nil
}

// encryptArmour gpg -e -a --recipient <gpgid>
func (conf ApiConf) encryptArmour(data, gpgid string) (string, error) {
	var encrypted string
	if gpgid == "" {
		gpgid = conf.GPGID
	}

	encrypted, err := gpg.EncryptArmorData(gpgid, data)
	if err != nil {
		log.Error(err.Error())
		return encrypted, err
	}

	return encrypted, nil
}
