package haversine

import (
	"math"
)

const rEarth = 6372.8 // km

func haversine(θ float64) float64 {
	return .5 * (1 - math.Cos(θ))
}

type pos struct {
	latRad float64 // latitude, radians
	lonRad float64 // longitude, radians
}

func degPos(lat, lon float64) pos {
	return pos{lat * math.Pi / 180, lon * math.Pi / 180}
}

// Accept position in radians
func hDist(p1, p2 pos) float64 {
	return 2 * rEarth *
		math.Asin(math.Sqrt(haversine(p2.latRad-p1.latRad)+
			math.Cos(p1.latRad)*math.Cos(p2.latRad)*haversine(p2.lonRad-p1.lonRad)))
}

// Accept pair of lat lon
// Return haversine distance in kilometers
func HaversineDist(lat1, lon1, lat2, lon2 float64) float64 {
	return hDist(degPos(lat1, lon1), degPos(lat2, lon2))
}
