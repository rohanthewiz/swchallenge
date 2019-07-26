package superman

import (
	"swchallenge/geo"
	"swchallenge/logger"
	"swchallenge/loginattempt"
	"swchallenge/logindb"
	"swchallenge/maxmind"

	"github.com/rohanthewiz/serr"
)

type Response struct {
	CurrentGeo            geo.Geo          `json:"currentGeo"`
	PrecedingLoginAttempt *PreviousEvent   `json:"precedingIpAccess,omitempty"`
	NextLoginAttempt      *SubsequentEvent `json:"subsequentIpAccess,omitempty"`
}

type PreviousEvent struct {
	TravelToCurrGeoSuspicious      bool                      `json:"travelToCurrentGeoSuspicious"`
	PrecedingLoginAttempt          struct {
		loginattempt.LoginAttempt
		Speed int64 `json:"speed"`
	} `json:"precedingIpAccess,omitempty"`
}

type SubsequentEvent struct {
	TravelFromCurrentGeoSuspicious bool                      `json:"travelFromCurrentGeoSuspicious"`
	SubsequentLoginAttempt         struct {
		loginattempt.LoginAttempt
		Speed int64 `json:"speed"`
	} `json:"subsequentIpAccess,omitempty"`
}

func ProcessLoginAttempt(eventUUID, ipAddr, username string, timestamp int64) (resp Response, err error) {
	// Get CurrentGeo
	currentGeo, err := maxmind.IPToLatLon("default", ipAddr)
	if err != nil {
		return resp, serr.Wrap(err, "Error obtaining current Geo from ip address", "ipAddr", ipAddr)
	}
	resp.CurrentGeo = currentGeo

	// Store the login Attempt
	currLA := loginattempt.LoginAttempt{ IP: ipAddr, Timestamp: timestamp }
	currLA.Latitude = currentGeo.Latitude
	currLA.Longitude = currentGeo.Longitude
	currLA.AccuracyRadius = currentGeo.AccuracyRadius

	err = logindb.DBHandle.StoreLoginAttempt(currLA)
	if err != nil {
		return resp, serr.Wrap(err, "unable to store current login attempt")
	}

	// Get previous login attempt
	prevLA, err := logindb.DBHandle.QueryLoginAttempts(false, currLA.Timestamp)
	if err != nil {
		logger.LogErr(serr.Wrap(err, "Unable to obtain previous login attempt"))
	} else {
		prevEvent := PreviousEvent{}
		prevEvent.TravelToCurrGeoSuspicious, prevEvent.PrecedingLoginAttempt.Speed = isSuspiciousTravel(currLA, prevLA)

		// Get suspicious travel info
		// Populate resp
	}


	// Get next login attempt
	// Get the HvDist (subtract off the radius)
	//if dist > threshold then set TravelToCurrentGeoSusp
	// Populate resp

	return
}
