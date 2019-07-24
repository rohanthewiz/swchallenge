package logger

import (
	"github.com/rohanthewiz/serr"
	"github.com/sirupsen/logrus"
	"strings"
)

// Logging for structured errors (SErr)
// Take an err (preferrably an SErr), and optional (single!) message argument
// Example:
//	err := serr.Wrap(errors.New("Just testing an error"), "attribute1", "value1", "attribute2", "value2")
//	logger.LogErr(err, "Testing out LogErr()")
func LogErr(err error, mesg ...string) {
	msgs := []string{} // accumulate "msg" fields
	errs := []string{} // accumulate "error" fields

	flds := logrus.Fields{}

	// Add optional
	if len(mesg) > 0 {
		msgs = []string{mesg[0]}
	}
	if len(mesg) > 1 {
		flds["extras"] = strings.Join(mesg[1:], ", ")
	}

	// Add error string from original error
	if er := err.Error(); er != "" {
		errs = []string{er}
	}

	// If error is structured error, get key vals
	if ser, ok := err.(serr.SErr); ok {
		for key, val := range ser.FieldsMap() {
			if key != "" {
				switch strings.ToLower(key) {
				case "error":
					errs = append(errs, val)
				case "msg":
					msgs = append(msgs, val)
				default:
					flds[key] = val
				}
			}
		}
	}
	// message is required by logrus so use the original error string if msgs empty
	if len(msgs) == 0 {
		msgs = []string{err.Error()}
	}
	// Populate the "error" field
	if len(errs) > 0 {
		flds["error"] = strings.Join(errs, " - ")
	}

	msg := strings.Join(msgs, " - ")
	msg = logEnvironment + msg // Always prepend environment
	logrus.WithFields(flds).Error(msg)
}
