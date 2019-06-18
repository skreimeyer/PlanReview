// Package esri handles interaction with GIS servers. Specifically, it provides
// functions which act on PAGIS.org web services and littlerock.gov servers.
// This includes:
// - Geocoding
// - Obtaining land parcel geometry
// - Finding flood hazard areas within a geometric envelope
// - Finding streets and their classification within an envelope
// - Finding zoning classification of a parcel
// There are a large number of mostly redundant structs which exist to keep some
// coherence of the highly varied JSON responses which the GIS servers provide
// to the queries.
package esri

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type spatialReference struct {
	Wkid       int `json:"wkid"`
	Latestwkid int `json:"latestWkid"`
}

// Location is a coordinate pair representing the centroid of an addressed
// parcel.
type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type candidates struct {
	Address    string     `json:"address"`
	Location   Location   `json:"location"`
	Score      float64    `json:"score"`
	Attributes attributes `json:"attributes"`
}

type gcResponse struct {
	Spatialreference spatialReference `json:"spatialReference"`
	Candidates       []candidates     `json:"candidates"`
}

// Geocode takes an address string and returns the location matched by the
// PAGIS server
func Geocode(addr string) (Location, error) {
	var loc Location
	geoURL, err := url.Parse(
		"https://www.pagis.org/arcgis/rest/services/LOCATORS/CompositeAddressPtsRoadCL/GeocodeServer/findAddressCandidates")
	if err != nil {
		return loc, err
	}
	params := url.Values{}
	params.Add("SingleLine", addr)
	params.Add("f", "json")
	params.Add("outFields", "*")
	params.Add("maxLocations", "3")
	params.Add("outSR", "{\"wkid\":102651,\"latestWkid\":3433}")
	geoURL.RawQuery = params.Encode()

	res, err := http.Get(geoURL.String())
	if err != nil {
		return loc, err
	}
	geoData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return loc, err
	}

	var geoJSON gcResponse

	err = json.Unmarshal(geoData, &geoJSON)
	if err != nil {
		return loc, err
	}
	if len(geoJSON.Candidates) >= 1 {
		return geoJSON.Candidates[0].Location, nil
	}
	return loc, errors.New("No candidates found")
}
