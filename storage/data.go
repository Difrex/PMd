package storage

import (
	"encoding/base64"
	"time"

	"crypto/sha256"

	log "github.com/Sirupsen/logrus"
)

type AddState struct {
	Version string `json:"version"`
	Time    string `json:"time"`
	State   string `json:"state"`
	Error   string `json:"error"`
}

// AddData ...
func (db *DB) AddData(user User, data string) (AddState, error) {
	var state AddState
	currentTime := time.Now()

	// Find user id
	userID, err := db.findUserID(user)
	if err != nil {
		log.Error(err.Error())
		return state, err
	}

	version := calculateSHA(data)

	addDataSQL := `INSERT INTO keys(version, userid, time, data) values(?, ?, ?, ?)`
	addDataStmt, err := db.Conn.Prepare(addDataSQL)
	if err != nil {
		log.Error(err.Error())
		return state, err
	}

	_, err = addDataStmt.Exec(version, userID, currentTime.Unix(), data)
	if err != nil {
		log.Error(err.Error())
		return state, err
	}

	textTime, err := currentTime.MarshalText()
	if err != nil {
		log.Error(err.Error())
		return state, err
	}

	state.Version = version
	state.Time = string(textTime)
	state.State = "OK"

	return state, nil
}

// List user uploaded versions
type List struct {
	Version string    `json:"version"`
	Time    time.Time `json:"time"`
}

// ListVersions list user data versions
func (db *DB) ListVersions(user User) ([]List, error) {
	var list []List

	userID, err := db.findUserID(user)
	if err != nil {
		log.Error(err.Error())
		return list, err
	}

	listSQL := `SELECT version, time FROM keys WHERE userid=?`
	rows, err := db.Conn.Query(listSQL, userID)
	if err != nil {
		log.Error()
		return list, err
	}

	for rows.Next() {
		var v string
		var t time.Time
		err := rows.Scan(&v, &t)
		if err != nil {
			log.Error(err.Error())
			return list, err
		}
		list = append(list, List{v, t})
	}

	return list, nil
}

// GetUserData ...
func (db *DB) GetUserData(user User, version string) (string, error) {
	var data string

	userID, err := db.findUserID(user)
	if err != nil {
		log.Error(err.Error())
		return data, err
	}

	getSQL := `SELECT data FROM keys WHERE userid=? and version=?`
	rows, err := db.Conn.Query(getSQL, userID, version)
	if err != nil {
		log.Error(err.Error())
		return data, err
	}

	for rows.Next() {
		err := rows.Scan(&data)
		log.Warn(data)
		if err != nil {
			log.Error(err.Error())
			return data, err
		}
	}

	return data, nil
}

// calculateSHA ...
func calculateSHA(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	b64 := base64.URLEncoding.EncodeToString(sum)

	return b64
}
