package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// Build flds and msg for logrus
// Level can be one of "debug", "info", "warn", "error", "fatal"
// `args` should be a list of argument pairs
// Example:
//		logger.Log("Info", "App is initializing...")
//		logger.Log("Warn", "Weird things are happening", "thing1", "value1", "thing2", "value2")
func Log(level, msg string, args ...string) {
	flds := logrus.Fields{}

	// Gather the other keys and values
	key := ""
	keys := []string{}
	for i, arg := range args {
		if i%2 == 0 { // arg is a key
			key = arg
			keys = append(keys, key)
		} else {
			flds[string(key)] = arg
		}
	}
	// Fixup / Validate
	if len(args)%2 != 0 {
		logrus.Warn("Even number of arguments required to Log() function. Odd argument will be paired with a blank")
	}

	// Always prepend environment
	msg = logEnvironment + msg

	// Call the logger
	lg := logrus.WithFields(flds)
	switch strings.ToLower(level) {
	case "debug":
		lg.Debug(msg)
	case "info":
		lg.Info(msg)
	case "warn":
		lg.Warn(msg)
	case "error":
		lg.Error(msg) // Log error, but don't quit
	case "fatal":
		lg.Fatal(msg) // Calls os.Exit() after logging
	}
}
