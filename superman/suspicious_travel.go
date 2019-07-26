package superman

import (
	"swchallenge/haversine"
	"swchallenge/loginattempt"
)

// Given distance in km and two unix timestamps
// TODO !!! Unit Test
func isSuspiciousTravel(la1, la2 loginattempt.LoginAttempt) (isSuspicious bool, speed int64) {
	// Get the HvDist
	dist := haversine.HaversineDist(la1.Latitude, la1.Longitude, la2.Latitude, la2.Longitude)
	// subtract off the radius
	// dist -= float64(currLA.AccuracyRadius + prevLA.AccuracyRadius) // TODO !!! - make sure the radius is km
	// Is suspicious?
	if
	//if dist > threshold then set TravelFromCurrentGeoSusp


	return
}
