package main

import (
	"log"
	"net/http"
	"swchallenge/logger"
	"swchallenge/logindb"
	"swchallenge/server"
)

// Todo cite and note maxmind library

func main() {
	const httpPort = "8100"

	logger.SetupLog()
	logger.Log("Info", "Instance is starting")

	// Init DB
	logindb.DBHandle = logindb.NewDB("default")
	if logindb.DBHandle == nil {
		panic("dbHandle cannot be nil")
	}
	defer logindb.DBHandle.Close()

	// For the sake of this exercise, start with a fresh table
	logindb.DBHandle.CreateTable()

	// Start Web Server
	http.HandleFunc("/", server.HandleV1)
	logger.Log("Info", "Web server starting on port " + httpPort)
	log.Fatal(http.ListenAndServe(":" + httpPort, nil))
}
