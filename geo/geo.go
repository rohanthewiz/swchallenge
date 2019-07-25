package geo

type Geo struct {
	Latitude       float64 `maxminddb:"latitude" json:"lat"`
	Longitude      float64 `maxminddb:"longitude" json:"lon"`
	AccuracyRadius int64   `maxminddb:"accuracy_radius" json:"radius"`
}
