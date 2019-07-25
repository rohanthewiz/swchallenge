package haversine

import (
	"math"
	"testing"
)

func TestHaversineDist(t *testing.T) {
	want := 2887.2599506
	got := HaversineDist(36.12, -86.67, 33.94, -118.40)

	if math.Abs(got - want) > 1e6 {
		t.Error("Got", got, "Want", want)
	}
}
