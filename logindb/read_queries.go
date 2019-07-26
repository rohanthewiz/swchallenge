package logindb

import (
	"errors"
	"fmt"
	"github.com/rohanthewiz/serr"
	"strconv"
	"swchallenge/logger"
	"swchallenge/loginattempt"
)

const NotFoundMessage = "Not Found"

// Get the previous or next login attempt relative to the given timestamp
func (d DB) QueryLoginAttempts(username string, timestamp int64, next bool) (loginattempt.LoginAttempt, error) {
	stage := "when obtaining login attempt"
	loginAtt := loginattempt.LoginAttempt{}

	if d.db == nil {
		return loginAtt, serr.Wrap(errors.New("database object is nil"), "stage", stage)
	}

	query := `SELECT username, uuid, latitude, longitude, radius, ip, created_at FROM login_attempts`

	op := "<"
	direction := "DESC"
	if next {
		op = ">"
		direction = "ASC"
	}

	condition := "WHERE username = '" + username + "' AND created_at " +
		op + " " + strconv.FormatInt(timestamp, 10)

	orderBy := "created_at " + direction

	query = fmt.Sprintf("%s %s ORDER BY %s LIMIT 1", query, condition, orderBy)
	logger.Log("Debug", "query: "+query)

	rs, err := d.db.Query(query)
	if err != nil {
		return loginAtt, serr.Wrap(err, "Error in DB query", "stage", stage)
	}
	defer rs.Close()

	if rs.Next() {
		err = rs.Scan(&loginAtt.Username, &loginAtt.EventUUID, &loginAtt.Latitude, &loginAtt.Longitude,
			&loginAtt.AccuracyRadius, &loginAtt.IP, &loginAtt.Timestamp)
		if err != nil {
			return loginAtt, serr.Wrap(err, "Error scanning recordset", "stage", stage)
		}
	} else {
		return loginAtt, serr.Wrap(errors.New(NotFoundMessage))
	}
	return loginAtt, nil
}

// See if UUID already exist to avoid multiple writes
func (d DB) UUIDExists(uuid string) (exists bool, err error) {
	stage := " when obtaining uuid"

	if d.db == nil {
		return exists, serr.Wrap(errors.New("database object is nil"), "stage", stage)
	}

	query := `SELECT uuid FROM login_attempts WHERE uuid = '` + uuid + `' LIMIT 1`
	rs, err := d.db.Query(query)
	if err != nil {
		return exists, serr.Wrap(err, "Error in DB query", "stage", stage)
	}
	defer rs.Close()

	if rs.Next() {
		dbUUID := ""
		err = rs.Scan(&dbUUID)
		if err != nil {
			return exists, serr.Wrap(err, "Error scanning recordset", "stage", stage)
		}
		return uuid == dbUUID, err
	}

	return
}
