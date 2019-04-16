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

type location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type attributes struct {
}

type candidates struct {
	Address    string     `json:"address"`
	Location   location   `json:"location"`
	Score      int        `json:"score"`
	Attributes attributes `json:"attributes"`
}

type gcResponse struct {
	Spatialreference spatialReference `json:"spatialReference"`
	Candidates       []candidates     `json:"candidates"`
}

// Geocode takes an address string and returns the location matched by the PAGIS server
func Geocode(addr string) location {
	geoUrl, err := url.Parse("https://www.pagis.org/arcgis/rest/services/LOCATORS/CompositeAddressPtsRoadCL/GeocodeServer/findAddressCandidates")
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	params.Add("SingleLine", addr)
	params.Add("f", "json")
	params.Add("outFields", "*")
	params.Add("maxLocations", "3")
	params.Add("outSR", "{\"wkid\":102651,\"latestWkid\":3433}")
	geoUrl.RawQuery = params.Encode()

	res, err := http.Get(geoUrl.String())
	if err != nil {
		panic(err)
	}
	geoData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	var geoJson gcResponse

	jsonErr := json.Unmarshal(geoData, &geoJson)
	if jsonErr != nil {
		panic(err)
	}
	return geoJson.Candidates[0].Location
}
