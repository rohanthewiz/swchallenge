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
	TravelToCurrGeoSuspicious bool                  `json:"travelToCurrentGeoSuspicious"`
	PrecedingLoginAttempt     PrecedingLoginAttempt `json:"precedingIpAccess,omitempty"`
}

type SubsequentEvent struct {
	TravelFromCurrentGeoSuspicious bool                   `json:"travelFromCurrentGeoSuspicious"`
	SubsequentLoginAttempt         SubsequentLoginAttempt `json:"subsequentIpAccess,omitempty"`
}

type PrecedingLoginAttempt struct {
	loginattempt.LoginAttempt
	Speed int64 `json:"speed"`
}

type SubsequentLoginAttempt struct {
	loginattempt.LoginAttempt
	Speed int64 `json:"speed"`
}

func ProcessLoginAttempt(uuid, ipAddr, username string, timestamp int64) (resp Response, err error) {
	// Get CurrentGeo
	currentGeo, err := maxmind.IPToLatLon("default", ipAddr)
	if err != nil {
		return resp, serr.Wrap(err, "Error obtaining current Geo from ip address", "ipAddr", ipAddr)
	}
	resp.CurrentGeo = currentGeo

	// Store the login Attempt
	currLA := loginattempt.LoginAttempt{IP: ipAddr, Timestamp: timestamp, Username: username, EventUUID: uuid}
	currLA.Latitude = currentGeo.Latitude
	currLA.Longitude = currentGeo.Longitude
	currLA.AccuracyRadius = currentGeo.AccuracyRadius

	err = logindb.DBHandle.StoreLoginAttempt(currLA)
	if err != nil {
		return resp, serr.Wrap(err, "unable to store current login attempt")
	}

	// Get previous login attempt
	prevLA, err := logindb.DBHandle.QueryLoginAttempts(currLA.Username, currLA.Timestamp, false)
	if err != nil {
		if err.Error() == logindb.NotFoundMessage {
			logger.Log("Warn", "No previous login found")
		} else {
			logger.LogErr(serr.Wrap(err, "error obtaining previous login attempt"))
		}
		err = nil // keep going
	} else {
		prevEvent := PreviousEvent{}
		prevEvent.PrecedingLoginAttempt.Timestamp = prevLA.Timestamp
		prevEvent.PrecedingLoginAttempt.IP = prevLA.IP
		prevEvent.PrecedingLoginAttempt.Latitude = prevLA.Latitude
		prevEvent.PrecedingLoginAttempt.Longitude = prevLA.Longitude
		prevEvent.PrecedingLoginAttempt.AccuracyRadius = prevLA.AccuracyRadius
		prevEvent.TravelToCurrGeoSuspicious, prevEvent.PrecedingLoginAttempt.Speed = IsSuspiciousTravel(currLA, prevLA)
		resp.PrecedingLoginAttempt = &prevEvent // Populate resp
	}

	// Next
	nextLA, err := logindb.DBHandle.QueryLoginAttempts(currLA.Username, currLA.Timestamp, true)
	if err != nil {
		if err.Error() == logindb.NotFoundMessage {
			logger.Log("Warn", "No subsequent login found")
		} else {
			logger.LogErr(serr.Wrap(err, "error obtaining subsequent login attempt"))
		}
		err = nil // keep going
	} else {
		nextEvent := SubsequentEvent{}
		nextEvent.SubsequentLoginAttempt.Timestamp = nextLA.Timestamp
		nextEvent.SubsequentLoginAttempt.IP = nextLA.IP
		nextEvent.SubsequentLoginAttempt.Latitude = nextLA.Latitude
		nextEvent.SubsequentLoginAttempt.Longitude = nextLA.Longitude
		nextEvent.SubsequentLoginAttempt.AccuracyRadius = nextLA.AccuracyRadius
		nextEvent.TravelFromCurrentGeoSuspicious, nextEvent.SubsequentLoginAttempt.Speed =
			IsSuspiciousTravel(currLA, nextLA)
		resp.NextLoginAttempt = &nextEvent
	}

	return
}
