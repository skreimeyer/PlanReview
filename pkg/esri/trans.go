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

type StreetClass int

// Street classifications. Residential and minor residential aren't queryable
// so there may be no way to get this to work.
const (
	Residential StreetClass = iota
	MinorResidential
	Collector
	Commercial
	MinorArterial
	Arterial
)

//go:generate stringer -type=StreetClass
// Street refers to a specific road
type Street struct {
	Name  string
	Class StreetClass
	Row   int
	Alt   bool
	ARDOT bool
}

// FetchRoads takes an envelope as an argument and returns a list of street
// names and their classifications.
// TODO: this is HIGHLY sensitive to the envelope buffer. Absolute widths are probably necessary.
func FetchRoads(e Envelope) ([]Street, error) {
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
		return []Street{Street{}},err
	}
	tData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return []Street{Street{}},err
	}

	var tJSON trans

	err = json.Unmarshal(tData, &tJSON)
	if err != nil {
		return []Street{Street{}},err
	}
	var result []Street
	for _, f := range tJSON.Features {
		class := deString(f.Attributes.ScaddType)
		st := Street{
		Name: f.Attributes.Mapname,
		Class: class,
		Row: calcRow(class),
		Alt: isAlt(f.Attributes.Altdes),
		ARDOT: isAlt(f.Attributes.Altdes), // this is a VERY weak inference
	}
		result = append(result, st)
	}
	return result, nil

}

func deString(s string) StreetClass {
	switch s {
	case "MINOR RESIDENTIAL":
		return MinorResidential
	case "RESIDENTIAL":
		return Residential
	case "COLLECTOR":
		return Collector
	case "COMMERCIAL":
		return Commercial // not sure this exists
	case "MINOR ARTERIAL":
		return MinorArterial
	case "PRINCIPAL ARTERIAL":
		return Arterial
	default:
		return MinorResidential
	}
}

func calcRow(sc StreetClass) int {
	switch sc {
	case Residential:
		return 50
	case MinorResidential:
		return 45
	case Collector:
		return 60
	case Commercial:
		return 60
	case MinorArterial:
		return 90
	case Arterial:
		return 110
	default:
		return 50
	}
}

func isAlt(s string) bool {
	if s == "YES" {
		return true
	}
	return false
}