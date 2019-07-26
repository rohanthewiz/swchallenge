package logindb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var DBHandle *DB // DB singleton
var dbMutex *sync.Mutex

func init() {
	dbMutex = &sync.Mutex{}
}

// SQLite DB wrapper
type DB struct {
	path string
	db   *sql.DB
}

// Create a new sqlite DB wrapper
func NewDB(path string) *DB {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	d := new(DB)
	if path == "" || path == "default" {
		path = "loginattempt/login_attempts.db"
	}
	d.path = path
	db, err := sql.Open("sqlite3", d.path)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db is nil")
	}
	d.db = db
	return d
}

func (d DB) Close() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if d.db != nil {
		_ = d.db.Close()
		d.db = nil
	}
}
