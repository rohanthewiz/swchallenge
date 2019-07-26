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

	//loc, err := maxmind.IPToLatLon("maxmind/GeoLite2-City/GeoLite2-City.mmdb", "81.2.69.142")
	//if err != nil {
	//	logger.LogErr(err)
	//} else {
	//	fmt.Println("location:", loc)
	//}

	// Init DB
	logindb.DBHandle = logindb.NewDB("default")
	if logindb.DBHandle == nil {
		panic("dbHandle cannot be nil")
	}
	defer logindb.DBHandle.Close()

	//laDB := logindb.NewDB("logindb/login_attempts_test.db")
	//defer laDB.Close()
	//
	//la, err := laDB.QueryLoginAttempts(false, 60000)
	//if err != nil {
	//	logger.LogErr(err)
	//}
	//fmt.Println("Login attempt:", la)
	//
	//la, err = laDB.QueryLoginAttempts(true, 60000)
	//if err != nil {
	//	logger.LogErr(err)
	//}
	//fmt.Println("Login attempt:", la)

	// Start Web Server
	http.HandleFunc("/", server.HandleV1)
	logger.Log("Info", "Web server starting on port " + httpPort)
	log.Fatal(http.ListenAndServe(":" + httpPort, nil))
}
