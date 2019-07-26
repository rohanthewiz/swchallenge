package superman

import (
	"math"
	"swchallenge/haversine"
	"swchallenge/loginattempt"
)

const (
	speedThreshold = 500 // mph
	kmToMile = 0.6214
)

// TODO !!! Unit Test
// Given two login attempts determine speed and if suspicious travel
func IsSuspiciousTravel(la1, la2 loginattempt.LoginAttempt) (isSuspicious bool, speed int64) {
	// Get the HvDist
	dist := haversine.HaversineDist(la1.Latitude, la1.Longitude, la2.Latitude, la2.Longitude)
	// subtract radii
	dist -= float64(la1.AccuracyRadius + la2.AccuracyRadius)

	timeDiffHrs := math.Abs(float64(la1.Timestamp - la2.Timestamp)) / 3600
	speed = int64(dist * kmToMile / timeDiffHrs)

	if speed > speedThreshold {
		isSuspicious = true
	}

	return
}
