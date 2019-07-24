package maxmind

import (
	"net"
	"github.com/oschwald/maxminddb-golang"
	"github.com/rohanthewiz/serr"
)

type Location struct {
	AccuracyRadius int64 `maxminddb:"accuracy_radius"`
	Latitude float64 `maxminddb:"latitude"`
	Longitude float64 `maxminddb:"longitude"`
}

func IPToLatLon(strIP string) (loc Location, err error) {
	db, err := maxminddb.Open("maxmind/GeoLite2-City/GeoLite2-City.mmdb")
	if err != nil {
		return loc, serr.Wrap(err, "Error opening maxmind database")
	}
	defer db.Close()

	ip := net.ParseIP(strIP)

	var record struct {
		Loc Location `maxminddb:"location"`
	//	Country struct {
	//		ISOCode string `maxminddb:"iso_code"`
	//	} `maxminddb:"country"`
	}

	err = db.Lookup(ip, &record)
	if err != nil {
		return loc, serr.Wrap(err, "Error obtaining lat/lon for IP", "ip", strIP)
		//return loc, serr.Wrap(err, "Error obtaining lat/lon for IP", "ip", strIP)
	}
	// fmt.Printf("%#v\n", record)

	return record.Loc, err
}
