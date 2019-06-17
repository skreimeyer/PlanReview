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
	GisLrGisplanZNumberObjectid1                string `json:"GIS_LR.GISPLAN.Z_Number.OBJECTID_1"`
	GisLrGisplanZNumberObjectid                 string `json:"GIS_LR.GISPLAN.Z_Number.OBJECTID"`
	GisLrGisplanZNumberLabel                    string `json:"GIS_LR.GISPLAN.Z_Number.LABEL"`
	GisLrGisplanZNumberLinkznum                 string `json:"GIS_LR.GISPLAN.Z_Number.LinkZnum"`
	GisLrGisplanZNumberLastupdate               string `json:"GIS_LR.GISPLAN.Z_Number.LastUpdate"`
	GisLrGisplanZNumberEditorname               string `json:"GIS_LR.GISPLAN.Z_Number.EditorName"`
	GisLrGisplanZoningInputObjectid             string `json:"GIS_LR.GISPLAN.Zoning_Input.OBJECTID"`
	GisLrGisplanZoningInputZNumber              string `json:"GIS_LR.GISPLAN.Zoning_Input.Z_Number"`
	GisLrGisplanZoningInputZonedFrom            string `json:"GIS_LR.GISPLAN.Zoning_Input.Zoned_From"`
	GisLrGisplanZoningInputZonedTo              string `json:"GIS_LR.GISPLAN.Zoning_Input.Zoned_To"`
	GisLrGisplanZoningInputOtherZNumbers        string `json:"GIS_LR.GISPLAN.Zoning_Input.Other_Z_Numbers"`
	GisLrGisplanZoningInputConditionalUse       string `json:"GIS_LR.GISPLAN.Zoning_Input.Conditional_Use"`
	GisLrGisplanZoningInputAnyConditions        string `json:"GIS_LR.GISPLAN.Zoning_Input.Any_Conditions"`
	GisLrGisplanZoningInputOrdinanceNumber      string `json:"GIS_LR.GISPLAN.Zoning_Input.Ordinance_Number"`
	GisLrGisplanZoningInputAreaZoned            string `json:"GIS_LR.GISPLAN.Zoning_Input.Area_Zoned"`
	GisLrGisplanZoningInputMultipleActions      string `json:"GIS_LR.GISPLAN.Zoning_Input.Multiple_Actions"`
	GisLrGisplanZoningInputIssueActionRequested string `json:"GIS_LR.GISPLAN.Zoning_Input.Issue_Action_Requested"`
	GisLrGisplanZoningInputLocation             string `json:"GIS_LR.GISPLAN.Zoning_Input.Location"`
	GisLrGisplanZoningInputOtherZoneCategories  string `json:"GIS_LR.GISPLAN.Zoning_Input.Other_Zone_Categories"`
	GisLrGisplanZoningInputVariance             string `json:"GIS_LR.GISPLAN.Zoning_Input.Variance"`
	GisLrGisplanZoningInputApproved             string `json:"GIS_LR.GISPLAN.Zoning_Input.Approved"`
	GisLrGisplanZoningInputDateOfPcAction       string `json:"GIS_LR.GISPLAN.Zoning_Input.Date_of_PC_Action"`
	GisLrGisplanZoningInputDateOfBoaAction      string `json:"GIS_LR.GISPLAN.Zoning_Input.Date_of_BOA_Action"`
	GisLrGisplanZoningInputDateOfBoardAction    string `json:"GIS_LR.GISPLAN.Zoning_Input.Date_of_Board_Action"`
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
	GisLrGisplanZNumberObjectid1                int    `json:"GIS_LR.GISPLAN.Z_Number.OBJECTID_1"`
	GisLrGisplanZNumberObjectid                 int    `json:"GIS_LR.GISPLAN.Z_Number.OBJECTID"`
	GisLrGisplanZNumberLabel                    string `json:"GIS_LR.GISPLAN.Z_Number.LABEL"`
	GisLrGisplanZNumberLinkznum                 string `json:"GIS_LR.GISPLAN.Z_Number.LinkZnum"`
	GisLrGisplanZNumberLastupdate               int    `json:"GIS_LR.GISPLAN.Z_Number.LastUpdate"`
	GisLrGisplanZNumberEditorname               string `json:"GIS_LR.GISPLAN.Z_Number.EditorName"`
	GisLrGisplanZoningInputObjectid             int    `json:"GIS_LR.GISPLAN.Zoning_Input.OBJECTID"`
	GisLrGisplanZoningInputZNumber              string `json:"GIS_LR.GISPLAN.Zoning_Input.Z_Number"`
	GisLrGisplanZoningInputZonedFrom            string `json:"GIS_LR.GISPLAN.Zoning_Input.Zoned_From"`
	GisLrGisplanZoningInputZonedTo              string `json:"GIS_LR.GISPLAN.Zoning_Input.Zoned_To"`
	GisLrGisplanZoningInputOtherZNumbers        int    `json:"GIS_LR.GISPLAN.Zoning_Input.Other_Z_Numbers"`
	GisLrGisplanZoningInputConditionalUse       int    `json:"GIS_LR.GISPLAN.Zoning_Input.Conditional_Use"`
	GisLrGisplanZoningInputAnyConditions        int    `json:"GIS_LR.GISPLAN.Zoning_Input.Any_Conditions"`
	GisLrGisplanZoningInputOrdinanceNumber      string `json:"GIS_LR.GISPLAN.Zoning_Input.Ordinance_Number"`
	GisLrGisplanZoningInputAreaZoned            string `json:"GIS_LR.GISPLAN.Zoning_Input.Area_Zoned"`
	GisLrGisplanZoningInputMultipleActions      int    `json:"GIS_LR.GISPLAN.Zoning_Input.Multiple_Actions"`
	GisLrGisplanZoningInputIssueActionRequested string `json:"GIS_LR.GISPLAN.Zoning_Input.Issue_Action_Requested"`
	GisLrGisplanZoningInputLocation             string `json:"GIS_LR.GISPLAN.Zoning_Input.Location"`
	GisLrGisplanZoningInputOtherZoneCategories  int    `json:"GIS_LR.GISPLAN.Zoning_Input.Other_Zone_Categories"`
	GisLrGisplanZoningInputVariance             int    `json:"GIS_LR.GISPLAN.Zoning_Input.Variance"`
	GisLrGisplanZoningInputApproved             int    `json:"GIS_LR.GISPLAN.Zoning_Input.Approved"`
	GisLrGisplanZoningInputDateOfPcAction       string `json:"GIS_LR.GISPLAN.Zoning_Input.Date_of_PC_Action"`
	GisLrGisplanZoningInputDateOfBoaAction      int    `json:"GIS_LR.GISPLAN.Zoning_Input.Date_of_BOA_Action"`
	GisLrGisplanZoningInputDateOfBoardAction    string `json:"GIS_LR.GISPLAN.Zoning_Input.Date_of_Board_Action"`
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
	GisLrGisplanZoningPolyObjectid   string `json:"GIS_LR.GISPLAN.Zoning_Poly.OBJECTID"`
	GisLrGisplanZoningPolyZoning     string `json:"GIS_LR.GISPLAN.Zoning_Poly.ZONING"`
	GisLrGisplanZoningPolySymbol     string `json:"GIS_LR.GISPLAN.Zoning_Poly.SYMBOL"`
	GisLrGisplanZoningPolyCup        string `json:"GIS_LR.GISPLAN.Zoning_Poly.CUP"`
	GisLrGisplanZoningPolyAcres      string `json:"GIS_LR.GISPLAN.Zoning_Poly.ACRES"`
	GisLrGisplanZoningPolyLastupdate string `json:"GIS_LR.GISPLAN.Zoning_Poly.LastUpdate"`
	GisLrGisplanZoningPolyEditorname string `json:"GIS_LR.GISPLAN.Zoning_Poly.EditorName"`
	ShapeArea                        string `json:"Shape.area"`
	ShapeLen                         string `json:"Shape.len"`
	ZoningDefinitionsObjectid        string `json:"Zoning_Definitions.OBJECTID"`
	ZoningDefinitionsLabel           string `json:"Zoning_Definitions.LABEL"`
	ZoningDefinitionsCategory        string `json:"Zoning_Definitions.CATEGORY"`
	ZoningDefinitionsDetails1        string `json:"Zoning_Definitions.DETAILS_1"`
	ZoningDefinitionsDetails2        string `json:"Zoning_Definitions.DETAILS_2"`
	ZoningDefinitionsLinks           string `json:"Zoning_Definitions.LINKS"`
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
	GisLrGisplanZoningPolyObjectid   int     `json:"GIS_LR.GISPLAN.Zoning_Poly.OBJECTID"`
	GisLrGisplanZoningPolyZoning     string  `json:"GIS_LR.GISPLAN.Zoning_Poly.ZONING"`
	GisLrGisplanZoningPolySymbol     int     `json:"GIS_LR.GISPLAN.Zoning_Poly.SYMBOL"`
	GisLrGisplanZoningPolyCup        int     `json:"GIS_LR.GISPLAN.Zoning_Poly.CUP"`
	GisLrGisplanZoningPolyAcres      float64 `json:"GIS_LR.GISPLAN.Zoning_Poly.ACRES"`
	GisLrGisplanZoningPolyLastupdate int     `json:"GIS_LR.GISPLAN.Zoning_Poly.LastUpdate"`
	GisLrGisplanZoningPolyEditorname string  `json:"GIS_LR.GISPLAN.Zoning_Poly.EditorName"`
	ShapeArea                        float64 `json:"Shape.area"`
	ShapeLen                         float64 `json:"Shape.len"`
	ZoningDefinitionsObjectid        int     `json:"Zoning_Definitions.OBJECTID"`
	ZoningDefinitionsLabel           string  `json:"Zoning_Definitions.LABEL"`
	ZoningDefinitionsCategory        string  `json:"Zoning_Definitions.CATEGORY"`
	ZoningDefinitionsDetails1        string  `json:"Zoning_Definitions.DETAILS_1"`
	ZoningDefinitionsDetails2        string  `json:"Zoning_Definitions.DETAILS_2"`
	ZoningDefinitionsLinks           string  `json:"Zoning_Definitions.LINKS"`
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
		return "",err
	}
	params := url.Values{}
	params.Add("f", "json")
	params.Add("returnGeometry", "true")
	params.Add("outFields", "GIS_LR.GISPLAN.Zoning_Poly.OBJECTID,GIS_LR.GISPLAN.Zoning_Poly.ZONING,GIS_LR.GISPLAN.Zoning_Poly.SYMBOL,GIS_LR.GISPLAN.Zoning_Poly.CUP,GIS_LR.GISPLAN.Zoning_Poly.ACRES,GIS_LR.GISPLAN.Zoning_Poly.LastUpdate,GIS_LR.GISPLAN.Zoning_Poly.EditorName,GIS_LR.GISPLAN.Zoning_Poly.Shape,Shape.area,Shape.len,Zoning_Definitions.OBJECTID,Zoning_Definitions.LABEL,Zoning_Definitions.CATEGORY,Zoning_Definitions.DETAILS_1,Zoning_Definitions.DETAILS_2,Zoning_Definitions.LINKS")
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
		return "",err
	}
	zData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "",err
	}

	var zJSON zoning

	err = json.Unmarshal(zData, &zJSON)
	if err != nil {
		return "",err
	}
	if len(zJSON.Features) > 1 {
		return zJSON.Features[0].Attributes.GisLrGisplanZoningPolyZoning, nil
	}
	return "",nil

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
	params.Add("outFields", "GIS_LR.GISPLAN.Z_Number.OBJECTID_1,GIS_LR.GISPLAN.Z_Number.OBJECTID,GIS_LR.GISPLAN.Z_Number.LABEL,GIS_LR.GISPLAN.Z_Number.LinkZnum,GIS_LR.GISPLAN.Z_Number.LastUpdate,GIS_LR.GISPLAN.Z_Number.EditorName,GIS_LR.GISPLAN.Z_Number.Shape,GIS_LR.GISPLAN.Zoning_Input.OBJECTID,GIS_LR.GISPLAN.Zoning_Input.Z_Number,GIS_LR.GISPLAN.Zoning_Input.Zoned_From,GIS_LR.GISPLAN.Zoning_Input.Zoned_To,GIS_LR.GISPLAN.Zoning_Input.Other_Z_Numbers,GIS_LR.GISPLAN.Zoning_Input.Conditional_Use,GIS_LR.GISPLAN.Zoning_Input.Any_Conditions,GIS_LR.GISPLAN.Zoning_Input.Ordinance_Number,GIS_LR.GISPLAN.Zoning_Input.Area_Zoned,GIS_LR.GISPLAN.Zoning_Input.Multiple_Actions,GIS_LR.GISPLAN.Zoning_Input.Issue_Action_Requested,GIS_LR.GISPLAN.Zoning_Input.Location,GIS_LR.GISPLAN.Zoning_Input.Other_Zone_Categories,GIS_LR.GISPLAN.Zoning_Input.Variance,GIS_LR.GISPLAN.Zoning_Input.Approved,GIS_LR.GISPLAN.Zoning_Input.Date_of_PC_Action,GIS_LR.GISPLAN.Zoning_Input.Date_of_BOA_Action,GIS_LR.GISPLAN.Zoning_Input.Date_of_Board_Action")
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
	if len(cJSON.Features) > 1 {
		return cJSON.Features[0].Attributes.GisLrGisplanZNumberLabel, nil
	}
	return "", nil
}
