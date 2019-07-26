package logindb

import (
	"errors"
	"fmt"
	"github.com/rohanthewiz/serr"
	"swchallenge/logger"
	"swchallenge/loginattempt"
)

func (d DB) CreateTable() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	d.dropTable()
	query :=
		`		CREATE TABLE IF NOT EXISTS login_attempts (
			username TEXT NOT NULL,
			uuid TEXT NOT NULL,
			latitude REAL NOT NULL,
			longitude REAL NOT NULL,
			radius INTEGER NOT NULL,
			ip TEXT NOT NULL,
			created_at INTEGER NOT NULL);`
	_, err := d.db.Exec(query)
	if err != nil {
		logger.LogErr(serr.Wrap(err, "error creating login_attempts table"))
		panic(err)
	}
}

func (d DB) dropTable() {
	query := `DROP TABLE IF EXISTS login_attempts;`
	_, err := d.db.Exec(query)
	if err != nil {
		logger.LogErr(serr.Wrap(err, "error dropping login_attempts table"))
		panic(err)
	}
}

// Store login attempts into the DB.
// Only allow unique UUIDs to be saved
func (d DB) StoreLoginAttempt(lAtt loginattempt.LoginAttempt) error {
	if exists, err := d.UUIDExists(lAtt.EventUUID); exists || err != nil {
		if exists {
			logger.Log("Warn", "UUID already exists in DB")
		} else {
			logger.LogErr(serr.Wrap(err))
		}
		return nil
	}
	stage := "when storing login attempt"

	if d.db == nil {
		return serr.Wrap(errors.New("db object is nil"), "stage", stage)
	}

	query :=
		`	INSERT OR REPLACE INTO login_attempts(
				username, uuid, latitude, longitude, radius, ip, created_at)
				values(?, ?, ?, ?, ?, ?, ?)`
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return serr.Wrap(err, "Error preparing insert query", "stage", stage)
	}
	defer stmt.Close()

	fmt.Printf("%#v\n", lAtt)

	_, err = stmt.Exec(lAtt.Username, lAtt.EventUUID, lAtt.Latitude, lAtt.Longitude,
		lAtt.AccuracyRadius, lAtt.IP, lAtt.Timestamp)
	if err != nil {
		return serr.Wrap(err, "Error executing insert statement", "stage", stage)
	}
	return nil
}
