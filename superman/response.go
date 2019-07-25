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
