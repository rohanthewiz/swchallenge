package loginattempt

import "swchallenge/geo"

type LoginAttempt struct {
	geo.Geo
	Username  string `json:"-"`
	EventUUID string `json:"-"`
	IP        string `json:"ip"`
	Timestamp int64  `json:"timestamp"`
}
