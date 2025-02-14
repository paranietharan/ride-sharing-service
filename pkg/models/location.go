package models

var rideLocations = []struct {
	start string
	end   string
	fare  float64
}{
	{"Place-A", "Place-D", 74.00},
	{"Place-B", "Place-E", 87.5},
	{"Place-C", "Place-F", 100.00},
}

func GetHardcodedRideLocation(pickup string, drop string) (float64, bool) {
	for _, location := range rideLocations {
		if location.start == pickup && location.end == drop {
			return location.fare, true
		}
	}
	return 0, false
}
