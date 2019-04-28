package esri

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type transfieldAliases struct {
	Objectid   string `json:"OBJECTID"`
	Mapname    string `json:"MapName"`
	Altdes     string `json:"AltDes"`
	ScaddType  string `json:"SCADD_Type"`
	Editdate   string `json:"EditDate"`
	Editorname string `json:"EditorName"`
	ShapeLen   string `json:"Shape.len"`
}

type transspatialReference struct {
	Wkid       int `json:"wkid"`
	Latestwkid int `json:"latestWkid"`
}

type transfields struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Alias string `json:"alias"`
}

type transattributes struct {
	Objectid   int     `json:"OBJECTID"`
	Mapname    string  `json:"MapName"`
	Altdes     string  `json:"AltDes"`
	ScaddType  string  `json:"SCADD_Type"`
	Editdate   int     `json:"EditDate"`
	Editorname int     `json:"EditorName"`
	ShapeLen   float64 `json:"Shape.len"`
}

type transgeometry struct {
	Paths [][][]float64 `json:"paths"`
}

type transfeatures struct {
	Attributes transattributes `json:"attributes"`
	Geometry   transgeometry   `json:"geometry"`
}

type trans struct {
	Displayfieldname string                `json:"displayFieldName"`
	Fieldaliases     transfieldAliases     `json:"fieldAliases"`
	Geometrytype     string                `json:"geometryType"`
	Spatialreference transspatialReference `json:"spatialReference"`
	Fields           []transfields         `json:"fields"`
	Features         []transfeatures       `json:"features"`
}

// Street pairs the name of a street with its classification.
type Street struct {
	Name  string
	Class string
}

// FetchRoads takes an envelope as an argument and returns a list of street names and their classifications
// TODO: this is HIGHLY sensitive to the envelope buffer. Absolute widths are probably necessary.
func FetchRoads(e Envelope) []Street {
	tURL, err := url.Parse("https://maps.littlerock.state.ar.us/arcgis/rest/services/Master_Street_Plan/MapServer/0/query")
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	params.Add("f", "json")
	params.Add("spatialRel", "esriSpatialRelIntersects")
	geomString := fmt.Sprintf("{\"xmin\":%f,\"ymin\":%f,\"xmax\":%f,\"ymax\":%f,\"spatialReference\":{\"wkid\":102651,\"latestWkid\":3433}}", e.Min.X, e.Min.Y, e.Max.X, e.Max.Y) // spatialReference will not change
	params.Add("geometry", geomString)
	params.Add("geometryType", "esriGeometryEnvelope")
	params.Add("inSR", "102651")
	params.Add("outFields", "OBJECTID,MapName,AltDes,SCADD_Type,EditDate,EditorName,Shape.len")
	params.Add("outSR", "102651")
	tURL.RawQuery = params.Encode()

	res, err := http.Get(tURL.String())
	if err != nil {
		panic(err)
	}
	tData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	var tJSON trans

	jsonErr := json.Unmarshal(tData, &tJSON)
	if jsonErr != nil {
		panic(jsonErr)
	}
	var result []Street
	for _, f := range tJSON.Features {
		st := Street{f.Attributes.Mapname, f.Attributes.ScaddType}
		result = append(result, st)
	}
	return result

}
