package storage

import (
	log "github.com/Sirupsen/logrus"
)

// AddUser ...
func (db *DB) AddUser(user User) (int, error) {

	userID, err := db.findUserID(user)
	if err != nil {
		return -1, err
	}
	if userID != 0 {
		return userID, err
	}

	addUserSQL := `INSERT INTO users(gpgid) values(?)`
	addStmt, err := db.Conn.Prepare(addUserSQL)
	if err != nil {
		return -1, err
	}

	res, err := addStmt.Exec(user.GPGID)
	if err != nil {
		return -1, err
	}

	log.Info("User(", user.GPGID, ") added")
	rows, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	log.Debug("Rows afected: ", rows)

	// Find user id
	userID, err = db.findUserID(user)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

// PurgeUser - remove all user data and user
func (db *DB) PurgeUser(user User) error {
	err := db.deleteUserData(user)
	if err != nil {
		return err
	}

	log.Warn("User(" + user.GPGID + ") data deleted")

	err = db.deleteUser(user)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Warn("User(" + user.GPGID + ") data")

	return nil
}

// Deluser and delete all user data
func (db *DB) deleteUser(user User) error {
	// Find user id
	userID, err := db.findUserID(user)
	if err != nil {
		return err
	}

	// Delete user
	deleteUserSQL := `DELETE FROM users WHERE id=?`
	deleteUserStmt, err := db.Conn.Prepare(deleteUserSQL)
	if err != nil {
		return err
	}
	_, err = deleteUserStmt.Exec(userID)
	if err != nil {
		return err
	}
	log.Warn("User(", user.GPGID, ") deleted")

	return nil
}

// deleteUserData ...
func (db *DB) deleteUserData(user User) error {
	userID, err := db.findUserID(user)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	// Delete user data
	deleteDataSQL := `DELETE FROM keys WHERE userid=?`
	deleteDataStmt, err := db.Conn.Prepare(deleteDataSQL)
	if err != nil {
		return err
	}

	_, err = deleteDataStmt.Exec(userID)
	if err != nil {
		return err
	}

	return nil
}

// findUserID ...
func (db *DB) findUserID(user User) (int, error) {
	userIDSQL := `SELECT id FROM users WHERE gpgid=?`
	rows, err := db.Conn.Query(userIDSQL, user.GPGID)
	if err != nil {
		return -1, err
	}

	var userID int
	for rows.Next() {
		err = rows.Scan(&userID)
		if err != nil {
			return userID, err
		}
	}
	return userID, nil
}
