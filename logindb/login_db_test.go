package logindb

import (
	"swchallenge/geo"
	"swchallenge/logger"
	"swchallenge/loginattempt"
	"sync"
	"testing"
)

var mutex = &sync.Mutex{}

func TestStoreAndQuery(t *testing.T) {
	mutex.Lock() // in case of parallel tests
	laDB := NewDB("login_attempts_test.db")
	mutex.Unlock()
	if laDB == nil {
		t.Fatal("Failed to get a DB handle")
	}
	defer laDB.Close()

	laDB.CreateTable()

	loginAtts := []loginattempt.LoginAttempt{
		{geo.Geo{1.1,  0.9, 500}, "192.168.2.1", 50000},
		{geo.Geo{1.2,  1.0, 600}, "192.168.2.2", 60000},
		{geo.Geo{1.3,  1.1, 700}, "192.168.2.3", 70000},
	}
	for _, lAtt := range loginAtts {
		err := laDB.StoreLoginAttempt(lAtt)
		if err != nil {
			logger.LogErr(err)
		}
	}

	// Get Previous Attempt
	lAtt, err := laDB.QueryLoginAttempts(false, 60000)
	if err != nil {
		t.Error("Query previous login attempt failed")
		logger.LogErr(err)
	}
	t.Log("Previous login attempt ->", lAtt)

	// Get Next Attempt
	lAtt, err = laDB.QueryLoginAttempts(true, 60000)
	if err != nil {
		t.Error("Query next login attempt failed")
		logger.LogErr(err)
	}
	t.Log("Next login attempt ->", lAtt)
}
