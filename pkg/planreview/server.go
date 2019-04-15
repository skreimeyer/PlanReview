// Package planReview makes REST API calls to multiple GIS mapping servers and templates review comments for Public Works staff review.
package planreview

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Main is my main man
func main() {
	findParcel("7801 Kanis Road")
}

// findParcel requests geocoding data from pagis.org for a given address string
func findParcel(addr string) {
	var result map[string]interface{}
	geoCodeUrl, err := url.Parse("https://www.pagis.org/arcgis/rest/services/LOCATORS/CompositeAddressPtsRoadCL/GeocodeServer/findAddressCandidates")
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	params.Add("SingleLine", addr)
	params.Add("f", "json")
	geoCodeUrl.RawQuery = params.Encode()
	res, err := http.Get(geoCodeUrl.String())
	if err != nil {
		panic(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	jsonErr := json.Unmarshal(data, &result)
	if jsonErr != nil {
		panic(jsonErr)
	}
	fmt.Println(result)
}
