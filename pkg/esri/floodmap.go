package esri

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type fmFieldAliases struct {
	Objectid                                       string `json:"OBJECTID"`
	FldZone                                        string `json:"FLD_ZONE"`
	Legend                                         string `json:"LEGEND"`
	Panel                                          string `json:"PANEL"`
	FirmPan                                        string `json:"FIRM_PAN"`
	FloodplainadministratorcontactsCityname        string `json:"FloodplainAdministratorContacts_CityName"`
	FloodplainadministratorcontactsFloodplainadmin string `json:"FloodplainAdministratorContacts_FloodplainAdmin"`
	FloodplainadministratorcontactsPhone           string `json:"FloodplainAdministratorContacts_Phone"`
	FloodplainadministratorcontactsEmail           string `json:"FloodplainAdministratorContacts_Email"`
	OwName                                         string `json:"OW_NAME"`
	Proplookup                                     string `json:"PROPLOOKUP"`
	OrigFid                                        string `json:"ORIG_FID"`
	FidUnionfloodzonespanelsparcels                string `json:"FID_UnionFloodZonesPanelsParcels"`
}

type fmSpatialReference struct {
	Wkid       int `json:"wkid"`
	Latestwkid int `json:"latestWkid"`
}

type fmAttributes struct {
	Objectid                                       int    `json:"OBJECTID"`
	FldZone                                        string `json:"FLD_ZONE"`
	Legend                                         string `json:"LEGEND"`
	Panel                                          string `json:"PANEL"`
	FirmPan                                        string `json:"FIRM_PAN"`
	FloodplainadministratorcontactsCityname        string `json:"FloodplainAdministratorContacts_CityName"`
	FloodplainadministratorcontactsFloodplainadmin string `json:"FloodplainAdministratorContacts_FloodplainAdmin"`
	FloodplainadministratorcontactsPhone           string `json:"FloodplainAdministratorContacts_Phone"`
	FloodplainadministratorcontactsEmail           string `json:"FloodplainAdministratorContacts_Email"`
	OwName                                         string `json:"OW_NAME"`
	Proplookup                                     string `json:"PROPLOOKUP"`
	OrigFid                                        int    `json:"ORIG_FID"`
	FidUnionfloodzonespanelsparcels                int    `json:"FID_UnionFloodZonesPanelsParcels"`
}

type fmGeometry struct {
	Rings [][][]float64 `json:"rings"`
}

type fmFeatures struct {
	Attributes fmAttributes `json:"attributes"`
	Geometry   fmGeometry   `json:"geometry"`
}

type fmResponse struct {
	Displayfieldname string             `json:"displayFieldName"`
	Fieldaliases     fmFieldAliases     `json:"fieldAliases"`
	Geometrytype     string             `json:"geometryType"`
	Spatialreference fmSpatialReference `json:"spatialReference"`
	Features         []fmFeatures       `json:"features"`
}

// FloodHaz is an enumeration of valid flood hazard area designations. Its use
// is an alternative to a hashmap
type FloodHaz int

// Flood Hazard Area classifications
const (
	X    FloodHaz = iota
	FIVE          // 0.2% annual chance
	A
	AE
	FLOODWAY
	LEVEE
)

func (f FloodHaz) String() string {
	return [...]string{"X", "FIVE", "A", "AE", "FLOODWAY", "LEVEE"}[f]
}

// FloodData takes an envelope as an argument, queries the PAGIS DFIRM map
// server and returns an array of stirings of flood
func FloodData(e Envelope) []FloodHaz {
	floodURL, err := url.Parse("https://www.pagis.org/arcgis/rest/services/APPS/Apps_DFIRM/MapServer//dynamicLayer/query")
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	params.Add("f", "json")
	params.Add("returnGeometry", "true")
	params.Add("spatialRel", "esriSpatialRelIntersects")
	params.Add("maxAllowableOffset", "1")
	params.Add("geometry", fmt.Sprintf("{\"xmin\":%f,\"ymin\":%f,\"xmax\":%f,\"ymax\":%f,\"spatialReference\":{\"wkid\":102651,\"latestWkid\":3433}}", e.Min.X, e.Min.Y, e.Max.X, e.Max.Y))
	params.Add("esriGeometryType", "esriGeometryEnvelope")
	params.Add("inSR", "102651")
	params.Add("outFields", "OBJECTID,FLD_ZONE,LEGEND,PANEL,FIRM_PAN,FloodplainAdministratorContacts_CityName,FloodplainAdministratorContacts_FloodplainAdmin,FloodplainAdministratorContacts_Phone,FloodplainAdministratorContacts_Email,OW_NAME,PROPLOOKUP,ORIG_FID,FID_UnionFloodZonesPanelsParcels")
	params.Add("layer", "{\"source\":{\"type\":\"mapLayer\",\"mapLayerId\":21}}")
	floodURL.RawQuery = params.Encode()

	res, err := http.Get(floodURL.String())
	if err != nil {
		panic(err)
	}
	floodData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	var floodJSON fmResponse

	jsonErr := json.Unmarshal(floodData, &floodJSON)
	if jsonErr != nil {
		panic(jsonErr)
	}

	zName := make(map[string]FloodHaz)
	zName["The Point Clicked has a 1% Annual Chance Flood Hazard, Outside Floodway, Zone AE"] = AE
	zName["The Point Clicked has a 1% Annual Chance Flood Hazard, Inside Floodway, Zone AE"] = FLOODWAY
	zName["The Point Clicked has a 0.2% Annual Chance Flood Hazard"] = FIVE
	zName["The Point Clicked is Outside 1% Annual Floodplain, Zone X"] = X
	zName["The Point Clicked is Outside 1% Annual Floodplain, Zone X (Protected by Levee)"] = LEVEE
	zName["The Point Clicked has a 1% Annual Chance Flood Hazard, Zone A"] = A

	var hazards []FloodHaz

	for _, feat := range floodJSON.Features {
		hazards = append(hazards, zName[feat.Attributes.Legend])
	}

	return hazards
}
