package logindb

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rohanthewiz/serr"
	"log"
	"strconv"
	"swchallenge/logger"
	"swchallenge/loginattempt"
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

func (d DB) CreateTable() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	d.dropTable()
	query :=
	`		CREATE TABLE IF NOT EXISTS login_attempts (
			latitude REAL NOT NULL,
			longitude REAL NOT NULL,
			radius INTEGER,
			ip TEXT NOT NULL,
			created_at INTEGER);`
	_, err := d.db.Exec(query)
	if err != nil {
		log.Println("Error creating table.")
		panic(err)
	}
}

func (d DB) dropTable() {
	query := `DROP TABLE IF EXISTS login_attempts;`
	_, err := d.db.Exec(query)
	if err != nil {
		log.Println("Error dropping table.")
		panic(err)
	}
}

func (d DB) StoreLoginAttempt(lAtt loginattempt.LoginAttempt) error {
	stage := "when storing login attempt"

	if d.db == nil {
		return serr.Wrap(errors.New("db object is nil"), "stage", stage)
	}

	query :=
	`	INSERT OR REPLACE INTO login_attempts(
				latitude, longitude, radius, ip, created_at)
				values(?, ?, ?, ?, ?)`
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return serr.Wrap(err, "Error preparing insert query", "stage", stage)
	}
	defer stmt.Close()

	fmt.Printf("%#v\n", lAtt)

	_, err = stmt.Exec(lAtt.Latitude, lAtt.Longitude, lAtt.AccuracyRadius, lAtt.IP, lAtt.Timestamp)
	if err != nil {
		return serr.Wrap(err, "Error executing insert statement", "stage", stage)
	}
	return nil
}

// Get the previous or next login attempt relative to the given timestamp
func (d DB) QueryLoginAttempts(next bool, timestamp int64) (loginattempt.LoginAttempt, error) {
	stage := "when obtaining login attempt"
	loginAtt := loginattempt.LoginAttempt{}

	if d.db == nil {
		return loginAtt, serr.Wrap(errors.New("database object is nil"), "stage", stage)
	}

	query := `SELECT latitude, longitude, radius, ip, created_at FROM login_attempts`

	op := "<"
	direction := "DESC"
	if next {
		op = ">"
		direction = "ASC"
	}

	condition := "WHERE created_at " + op + " " + strconv.FormatInt(timestamp, 10)

	orderBy := "created_at " + direction

	query = fmt.Sprintf("%s %s ORDER BY %s LIMIT 1", query, condition, orderBy)
	logger.Log("Debug", "query: "+query)

	rs, err := d.db.Query(query)
	if err != nil {
		return loginAtt, serr.Wrap(err, "Error in DB query", "stage", stage)
	}
	defer rs.Close()

	if rs.Next() {
		err = rs.Scan(&loginAtt.Latitude, &loginAtt.Longitude, &loginAtt.AccuracyRadius,
			&loginAtt.IP, &loginAtt.Timestamp)
		if err != nil {
			return loginAtt, serr.Wrap(err, "Error scanning recordset", "stage", stage)
		}
	}
	return loginAtt, nil
}
