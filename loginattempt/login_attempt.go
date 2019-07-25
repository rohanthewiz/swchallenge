package loginattempt

import "swchallenge/geo"

type LoginAttempt struct {
	geo.Geo
	IP        string `json:"ip"`
	Timestamp int64  `json:"timestamp"`
}
