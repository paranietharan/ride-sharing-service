package models

// hard coded loaction
var rideLocations = []struct {
	start string
	end   string
	fare  float64
}{
	{"Place-A", "Place-D", 74.00},
	{"Place-B", "Place-E", 87.5},
	{"Place-C", "Place-F", 100.00},
}

func GetHardcodedRideLocation(index int) (string, string, float64, bool) {
	if index >= 0 && index < len(rideLocations) {
		return rideLocations[index].start, rideLocations[index].end, rideLocations[index].fare, true
	}
	return "", "", 0, false
}
