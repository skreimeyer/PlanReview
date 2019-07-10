package esri

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// STRUCTS FOR CASE FILE ONLY

type casefieldAliases struct {
}

type casespatialReference struct {
	Wkid       int `json:"wkid"`
	Latestwkid int `json:"latestWkid"`
}

type casefields struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Alias string `json:"alias"`
}

type caseattributes struct {
	GisLrGisplanZNumberLabel string `json:"GIS_LR.GISPLAN.Z_Number.LABEL"`
}

type casegeometry struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type casefeatures struct {
	Attributes caseattributes `json:"attributes"`
	Geometry   casegeometry   `json:"geometry"`
}

type caseFiles struct {
	Displayfieldname string               `json:"displayFieldName"`
	Fieldaliases     casefieldAliases     `json:"fieldAliases"`
	Geometrytype     string               `json:"geometryType"`
	Spatialreference casespatialReference `json:"spatialReference"`
	Fields           []casefields         `json:"fields"`
	Features         []casefeatures       `json:"features"`
}

// STRUCTS FOR ZONING QUERY

type zoningfieldAliases struct {
}

type zoningspatialReference struct {
	Wkid       int `json:"wkid"`
	Latestwkid int `json:"latestWkid"`
}

type zoningfields struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Alias string `json:"alias"`
}

type zoningattributes struct {
	GisLrGisplanZoningPolyZoning string `json:"GIS_LR.GISPLAN.Zoning_Poly.ZONING"`
}

type zoninggeometry struct {
	Rings [][][]float64 `json:"rings"`
}

type zoningfeatures struct {
	Attributes zoningattributes `json:"attributes"`
	Geometry   zoninggeometry   `json:"geometry"`
}

type zoning struct {
	Displayfieldname string                 `json:"displayFieldName"`
	Fieldaliases     zoningfieldAliases     `json:"fieldAliases"`
	Geometrytype     string                 `json:"geometryType"`
	Spatialreference zoningspatialReference `json:"spatialReference"`
	Fields           []zoningfields         `json:"fields"`
	Features         []zoningfeatures       `json:"features"`
}

// FetchZone takes an envelope (two points defining a rectangle which enclose a
// parcel) and returns a Zone, which is a code for land uses permitted by
// municipal ordinance. example:
//	R2 - single family residential
//	PRD - Planned Residential Development
func FetchZone(l Location) (string, error) {
	zURL, err := url.Parse("https://maps.littlerock.state.ar.us/arcgis/rest/services/Zoning/MapServer/32/query")
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("f", "json")
	params.Add("returnGeometry", "true")
	params.Add("outFields", "GIS_LR.GISPLAN.Zoning_Poly.ZONING")
	params.Add("maxAllowableOffset", "1")
	geomString := fmt.Sprintf("{\"xmin\":%f,\"ymin\":%f,\"xmax\":%f,\"ymax\":%f,\"spatialReference\":{\"wkid\":102651,\"latestWkid\":3433}}", l.X, l.Y, l.X+13, l.Y+13) // spatialReference will not change
	params.Add("geometry", geomString)
	params.Add("geometryType", "esriGeometryEnvelope")
	params.Add("spatialRel", "esriSpatialRelIntersects")
	params.Add("inSR", "102651")
	params.Add("outSR", "102651")
	zURL.RawQuery = params.Encode()

	res, err := http.Get(zURL.String())
	if err != nil {
		return "", err
	}
	zData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	var zJSON zoning

	err = json.Unmarshal(zData, &zJSON)
	if err != nil {
		return "", err
	}
	if len(zJSON.Features) > 0 {
		return zJSON.Features[0].Attributes.GisLrGisplanZoningPolyZoning, nil
	}
	return "", nil

}

// IsMultifam is a trivial function to determine if zoning is multifamily (ie)
// anything zoned M24, etc. Does not capture planned residential developments
// which would require deep inspection of zoning files.
func IsMultifam(z string) bool {
	return strings.HasPrefix(z, "M")
}

// FetchCases returns all the case files associated with a parcel.
// TODO: fetch a list of cases instead of the first
func FetchCases(e Envelope) (string, error) {
	cURL, err := url.Parse("https://maps.littlerock.state.ar.us/arcgis/rest/services/Zoning/MapServer/7/query")
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("f", "json")
	params.Add("returnGeometry", "true")
	params.Add("outFields", "GIS_LR.GISPLAN.Z_Number.LABEL,")
	params.Add("maxAllowableOffset", "2")
	geomString := fmt.Sprintf("{\"xmin\":%f,\"ymin\":%f,\"xmax\":%f,\"ymax\":%f,\"spatialReference\":{\"wkid\":102651,\"latestWkid\":3433}}", e.Min.X, e.Min.Y, e.Max.X, e.Max.Y) // spatialReference will not change
	params.Add("geometry", geomString)
	params.Add("geometryType", "esriGeometryEnvelope")
	params.Add("spatialRel", "esriSpatialRelIntersects")
	params.Add("inSR", "102651")
	params.Add("outSR", "102651")
	cURL.RawQuery = params.Encode()

	res, err := http.Get(cURL.String())
	if err != nil {
		return "", err
	}
	cData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	var cJSON caseFiles

	err = json.Unmarshal(cData, &cJSON)
	if err != nil {
		return "", err
	}
	if len(cJSON.Features) > 0 {
		return cJSON.Features[0].Attributes.GisLrGisplanZNumberLabel, nil
	}
	return "", nil
}
