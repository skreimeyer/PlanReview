package planreview

import (
	"testing"
)

func TestGeocode(t *testing.T) {
	precision := 1.0
	loc := Location{
		X: 1229504.418500146,
		Y: 151453.58653979315,
	}
	gc := Geocode("500 W Markham")
	if gc.X-loc.X > precision || gc.Y-loc.Y > precision {
		t.Errorf("geocode failed\nexpected: %f, %f\t got: %f, %f", loc.X, loc.Y, gc.X, gc.Y)
	}
}

func TestParcel(t *testing.T) {
	precision := 1.0
	ring := [][]float64{{1229652.12, 151364.85}, {1229614.06, 151350.96}, {1229353.22, 151395.22}, {1229376.18, 151533.24}, {1229671.79, 151483.09}, {1229652.12, 151364.85}}
	cityHall := Geocode("500 W Markham")
	pr := FetchParcel(cityHall)
	for i := range pr {
		for j := range pr[i] {
			if ring[i][j]-pr[i][j] > precision {
				t.Errorf("fetching rings failed\nindex:%d,%d\texpected:%f\tfound:%f", i, j, ring[i][j], pr[i][j])
			}
		}
	}
}

func TestFlood(t *testing.T) {
	gc := Geocode("1500 Westpark Dr")
	ring := FetchParcel(gc)
	env := MakeEnvelope(ring, 0.05)
	zones := FloodData(env)
	knownZones := []FloodHaz{AE, FLOODWAY, FIVE}
	for _, k := range knownZones {
		found := false
		for _, z := range zones {
			if z == k {
				found = true
			}
		}
		if found == false {
			t.Errorf("failed to locate target zone: %s within zones:%v", k, zones)
		}
	}
}

func TestZone(t *testing.T) {
	gc := Geocode("12800 Chenal Parkway")
	zone := FetchZone(gc)
	if zone != "PCD" {
		t.Errorf("zoning test failed. Target: %s\t Fetched: %s", "PCD", zone)
	}
}

func TestCaseFile(t *testing.T) {
	gc := Geocode("12800 Chenal Parkway")
	ring := FetchParcel(gc)
	env := MakeEnvelope(ring, 0.01)
	zone := FetchCases(env)
	target := "Z-6199-A"
	if zone != target {
		t.Errorf("zoning test failed. Target: %s\t Fetched: %s", target, zone)
	}
}

func TestTrans(t *testing.T) {
	gc := Geocode("2724 Fair Park Blvd")
	ring := FetchParcel(gc)
	env := MakeEnvelope(ring, 0.2)
	streets := FetchRoads(env)
	target := []Street{
		Street{
			"FAIR PARK BLVD",
			"MINOR ARTERIAL",
		},
		Street{
			"W 20TH ST",
			"COLLECTOR",
		},
	}
	if !func(s []Street, t Street) bool {
		for _, x := range s {
			if x == t {
				return true
			}
		}
		return false
	}(streets, target[0]) {
		t.Errorf("transportation test failed. Target: %s\t Fetched: %s", target, streets)
	}
}
