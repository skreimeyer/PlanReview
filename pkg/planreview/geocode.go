package planreview

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type spatialReference struct {
	Wkid       int `json:"wkid"`
	Latestwkid int `json:"latestWkid"`
}

// Location is a coordinate pair representing the centroid of an addressed parcel.
type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// type attributes struct {
// }

type candidates struct {
	Address    string     `json:"address"`
	Location   Location   `json:"location"`
	Score      int        `json:"score"`
	Attributes attributes `json:"attributes"`
}

type gcResponse struct {
	Spatialreference spatialReference `json:"spatialReference"`
	Candidates       []candidates     `json:"candidates"`
}

// Geocode takes an address string and returns the location matched by the PAGIS server
func Geocode(addr string) Location {
	geoURL, err := url.Parse("https://www.pagis.org/arcgis/rest/services/LOCATORS/CompositeAddressPtsRoadCL/GeocodeServer/findAddressCandidates")
	if err != nil {
		panic(err)
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
		panic(err)
	}
	geoData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	var geoJSON gcResponse

	jsonErr := json.Unmarshal(geoData, &geoJSON)
	if jsonErr != nil {
		panic(err)
	}
	return geoJSON.Candidates[0].Location
}
