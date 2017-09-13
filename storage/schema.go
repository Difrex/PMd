package storage

import (
	"database/sql"

	"sync"

	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
	Path string
	mux  sync.Mutex
}

type DbSchema struct {
	Users []User
	Keys  []Key
}

type User struct {
	ID    int
	GPGID string
}

type Key struct {
	ID      int
	UserID  int
	Time    int
	Version int
	Data    string
}

// NewDbSchema ...
func NewDbSchema() *DbSchema {
	var schema DbSchema

	return &schema
}

// CheckAndInit ...
func CheckAndInit(path string) *DB {
	db, err := InitConnection(path)
	checkErr(err)
	db.CreateSchema()

	return db
}

// InitConnection ...
func InitConnection(path string) (*DB, error) {
	db, err := sql.Open("sqlite3", path)
	checkErr(err)

	var database DB
	database.Conn = db
	database.Path = path

	return &database, nil
}

// Create ...
func (d *DB) CreateSchema() error {
	usersTableSQL := `CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY AUTOINCREMENT,
                                                       gpgid varchar(42) NOT NULL UNIQUE);`

	keysTableSQL := `CREATE TABLE IF NOT EXISTS keys(version varchar(42) PRIMARY KEY NOT NULL UNIQUE,
                                                     userid INTEGER NOT NULL,
                                                     time TIMESTAMP NOT NULL,
                                                     data TEXT NOT NULL);`

	usersStmt, err := d.Conn.Prepare(usersTableSQL)
	checkErr(err)
	keysStmt, err := d.Conn.Prepare(keysTableSQL)
	checkErr(err)

	_, err = usersStmt.Exec()
	checkErr(err)
	_, err = keysStmt.Exec()
	checkErr(err)

	return nil
}

// checkErr ...
func checkErr(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
	}
	return true
}
