package esri

import (
	"testing"
)

// TODO Convert all of these to table-based tests

func TestGeocode(t *testing.T) {
	precision := 1.0
	loc := Location{
		X: 1229504.418500146,
		Y: 151453.58653979315,
	}
	gc, _ := Geocode("500 W Markham")
	if gc.X-loc.X > precision || gc.Y-loc.Y > precision {
		t.Errorf("geocode failed\nexpected: %f, %f\t got: %f, %f", loc.X, loc.Y, gc.X, gc.Y)
	}
}

func TestParcel(t *testing.T) {
	precision := 1.0
	ring := [][]float64{{1229652.12, 151364.85}, {1229614.06, 151350.96}, {1229353.22, 151395.22}, {1229376.18, 151533.24}, {1229671.79, 151483.09}, {1229652.12, 151364.85}}
	cityHall := Geocode("500 W Markham")
	pr, _ := FetchParcel(cityHall)
	r := GetRing(pr)
	for i := range r {
		for j := range r[i] {
			if ring[i][j]-r[i][j] > precision {
				t.Errorf("fetching rings failed\nindex:%d,%d\texpected:%f\tfound:%f", i, j, ring[i][j], r[i][j])
			}
		}
	}
}

func TestFlood(t *testing.T) {
	gc, _ := Geocode("1500 Westpark Dr")
	par, _ := FetchParcel(gc)
	ring := GetRing(par)
	env := MakeEnvelope(ring, 0.05)
	zones, _ := FloodData(env)
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
	gc, _ := Geocode("12800 Chenal Parkway")
	zone, _ := FetchZone(gc)
	if zone != "PCD" {
		t.Errorf("zoning test failed. Target: %s\t Fetched: %s", "PCD", zone)
	}
}

func TestCaseFile(t *testing.T) {
	gc, _ := Geocode("12800 Chenal Parkway")
	par, _ := FetchParcel(gc)
	ring := GetRing(par)
	env := MakeEnvelope(ring, 0.01)
	zone, _ := FetchCases(env)
	target := "Z-6199-A"
	if zone != target {
		t.Errorf("zoning test failed. Target: %s\t Fetched: %s", target, zone)
	}
}

func TestTrans(t *testing.T) {
	gc, _ := Geocode("2724 Fair Park Blvd")
	par, _ := FetchParcel(gc)
	ring := GetRing(par)
	env := MakeEnvelope(ring, 0.2)
	streets, _ := FetchRoads(env)
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
