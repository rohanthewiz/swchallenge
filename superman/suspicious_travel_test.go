package superman

import (
	"swchallenge/geo"
	"swchallenge/loginattempt"
	"testing"
)

func TestIsSuspiciousTravel(t *testing.T) {
	la1 := loginattempt.LoginAttempt{
		geo.Geo{36.12, -86.67, 82}, "bob", "abcd-efgh-ijk", "192.168.2.1", 1564115316,
	}
	la2 := loginattempt.LoginAttempt{
		geo.Geo{33.94, -118.40, 30}, "bob", "lmn-opq-rst", "192.168.2.13", 1564102716,
	}
	la3 := loginattempt.LoginAttempt{
		geo.Geo{40.807330104, -74.072833042, 40}, "bob", "tuv-wxy-zabc", "192.168.2.20", 1564117116,
	}

	isSuspicious, speed := IsSuspiciousTravel(la1, la2)
	t.Log(isSuspicious, " - ", speed)

	isSuspicious, speed = IsSuspiciousTravel(la2, la3)
	t.Log(isSuspicious, " - ", speed)

	isSuspicious, speed = IsSuspiciousTravel(la1, la3)
	t.Log(isSuspicious, " - ", speed)
}
