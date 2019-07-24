package main

import (
	"fmt"
	"swchallenge/logger"
	"swchallenge/maxmind"
)

// Todo cite and note maxmind library
// Todo move logger package here
// Todo no lfs hook

func main() {
	logger.Log("Info", "Instance is starting")

	loc, err := maxmind.IPToLatLon("81.2.69.142")
	if err != nil {
		logger.LogErr(err)
	} else {
		fmt.Println("location:", loc)
	}
}
