package main

import (
	"fmt"
	"swchallenge/maxmind"
	"github.com/rohanthewiz/logger"
)

// Todo cite and note maxmind library
// Todo move logger package here
// Todo no lfs hook

func main() {
	loc, err := maxmind.IPToLatLon("81.2.69.142")
	if err != nil {
		logger.LogErr(err)
	} else {
		fmt.Println("location:", loc)
	}
}
