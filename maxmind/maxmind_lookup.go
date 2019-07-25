package maxmind

import (
	"github.com/oschwald/maxminddb-golang"
	"github.com/rohanthewiz/serr"
	"net"
	"swchallenge/geo"
)

func IPToLatLon(dbPath, strIP string) (loc geo.Geo, err error) {
	db, err := maxminddb.Open(dbPath)
	if err != nil {
		return loc, serr.Wrap(err, "Error opening maxmind database")
	}
	defer db.Close()

	ip := net.ParseIP(strIP)

	var record struct {
		Loc geo.Geo `maxminddb:"location"`
	}

	err = db.Lookup(ip, &record)
	if err != nil {
		return loc, serr.Wrap(err, "Error obtaining lat/lon for IP", "ip", strIP)
		//return loc, serr.Wrap(err, "Error obtaining lat/lon for IP", "ip", strIP)
	}
	// fmt.Printf("%#v\n", record)

	return record.Loc, err
}
