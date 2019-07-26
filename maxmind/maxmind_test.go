package maxmind

import (
	"reflect"
	"swchallenge/geo"
	"swchallenge/logger"
	"testing"
)

var testData = []struct {
	IP  string
	Lat float64
	Lon float64
}{
	// TODO - A few more cases
	{"81.2.69.142", 51.418, -0.1752},
}

func TestIPToLatLon(t *testing.T) {
	lnTestData := len(testData)

	for i := 0; i < lnTestData; i++ {
		loc, err := IPToLatLon("GeoLite2-City/GeoLite2-City.mmdb", testData[i].IP)
		if err != nil {
			logger.LogErr(err)
		}
		if testData[i].Lat != loc.Latitude || testData[i].Lon != loc.Longitude {
			t.Error("Lat (want):", testData[i].Lat, "Lat (got):", loc.Latitude)
			t.Error("Lon (want):", testData[i].Lon, "Lon (got):", loc.Longitude)
		}
	}
}

func TestMaxMindMapWrap(t *testing.T) {
	const ip = "206.81.252.6"
	currGeo := geo.Geo{Latitude: 39.211, Longitude: -76.8362, AccuracyRadius: 5}
	MaxMindMap.SetGeo(ip, currGeo)

	g, err := MaxMindMap.GetGeo(ip)
	if err != nil && !reflect.DeepEqual(g, currGeo) {
		t.Error("Maxmind map cache failed")
	}
	t.Log(g)
}
