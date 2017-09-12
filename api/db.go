package api

import (
	"github.com/Difrex/PMd/storage"
)

func (conf ApiConf) userAdd(gpgid string) (int, error) {
	user := storage.User{
		GPGID: gpgid,
	}
	userID, err := conf.Db.AddUser(user)
	if err != nil {
		return -1, err
	}

	return userID, nil
}
