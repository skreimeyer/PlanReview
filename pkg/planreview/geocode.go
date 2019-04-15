package planreview

import (
	"encoding/json"
	"fmt"
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

type root struct {
	Spatialreference spatialReference `json:"spatialReference"`
	Candidates       []candidates     `json:"candidates"`
}

// Main is my main man
func main() {
	// var result map[string]interface{}
	var result root
	testdata := []byte(`{"spatialReference":{"wkid":102651,"latestWkid":3433},"candidates":[{"address":"7801 KANIS RD, 72204","location":{"x":1204066.358307957,"y":147937.39795626889},"score":100,"attributes":{}},{"address":"7801 KANIS RD, 72204","location":{"x":1204169.6110534209,"y":148586.87696869907},"score":100,"attributes":{}},{"address":"7802 KANIS RD, 72204","location":{"x":1204172.8705882574,"y":148627.42274624179},"score":79,"attributes":{}}]}`)
	jsonErr := json.Unmarshal(testdata, &result)
	if jsonErr != nil {
		panic(jsonErr)
	}
	fmt.Println("\n struct = ")
	fmt.Println(result)
	fmt.Println("\nBold")
	fmt.Println(result.Candidates)
	fmt.Println("\nDaring")
	fmt.Println(result.Candidates[0].Location)
}
