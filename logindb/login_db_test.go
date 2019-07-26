package logindb

import (
	"swchallenge/geo"
	"swchallenge/logger"
	"swchallenge/loginattempt"
	"testing"
)

func TestStoreAndQuery(t *testing.T) {
	laDB := NewDB("login_attempts_test.db")
	if laDB == nil {
		t.Fatal("Failed to get a DB handle")
	}
	defer laDB.Close()

	laDB.CreateTable()

	loginAtts := []loginattempt.LoginAttempt{
		{geo.Geo{1.1,  0.9, 100}, "logman", "abcd-efgh-ijk", "192.168.2.1", 1564115316},
		{geo.Geo{1.2,  1.0, 75}, "logman", "lmn-opq-rst", "192.168.2.13", 1564102716},
		{geo.Geo{1.3,  1.1, 10}, "sam", "cde-fgh-ijkl", "192.168.2.9", 1564117116},
		{geo.Geo{1.3,  1.1, 80}, "logman", "tuv-wxy-zabc", "192.168.2.20", 1564117116},
	}
	for _, lAtt := range loginAtts {
		err := laDB.StoreLoginAttempt(lAtt)
		if err != nil {
			logger.LogErr(err)
		}
	}

	// Get Previous Attempt
	lAtt, err := laDB.QueryLoginAttempts("bob", 1564115316, false)
	if err != nil {
		t.Error("Query previous login attempt failed")
		logger.LogErr(err)
	}
	t.Log("Previous login attempt ->", lAtt)

	// Get Next Attempt
	lAtt, err = laDB.QueryLoginAttempts("bob", 1564115316, true)
	if err != nil {
		t.Error("Query next login attempt failed")
		logger.LogErr(err)
	}
	t.Log("Next login attempt ->", lAtt)
}
