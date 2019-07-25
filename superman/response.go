package superman

import (
	"swchallenge/geo"
	"swchallenge/loginattempt"
)

type Response struct {
	CurrentGeo geo.Geo `json:"currentGeo"`
	TravelToCurrGeoSuspicious bool `json:"travelToCurrentGeoSuspicious"`
	TravelFromCurrentGeoSuspicious bool `json:"travelFromCurrentGeoSuspicious"`
	PrecedingLoginAttempt loginattempt.LoginAttempt `json:"precedingIpAccess"`
	SubsequentLoginAttempt loginattempt.LoginAttempt `json:"subsequentIpAccess"`
}

func ProcessLoginAttempt(eventUUID, ipAddr, username string, timestamp int64) (resp Response, err error) {
	// Get the HvDist
	// Subtract off the radius

	//if dist > threshold - do the right thing

	return
}