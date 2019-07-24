package maxmind

import "testing"

var testData = []struct {
	IP string
	Lat float64
	Lon float64
}{
	{"81.2.69.142", 51.418, -0.1752},
}

func TestIPToLatLon(t *testing.T) {
	lnTestData := len(testData)

	for i := 0; i < lnTestData; i++ {

	}
}