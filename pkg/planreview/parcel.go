package planreview

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type fieldAliases struct {
	Objectid    string `json:"OBJECTID"`
	ParcelID    string `json:"PARCEL_ID"`
	ParcelLgl   string `json:"PARCEL_LGL"`
	SrceDate    string `json:"SRCE_DATE"`
	OwName      string `json:"OW_NAME"`
	OwAdd       string `json:"OW_ADD"`
	OwAdd2      string `json:"OW_ADD2"`
	OwCity      string `json:"OW_CITY"`
	OwState     string `json:"OW_STATE"`
	OwZip       string `json:"OW_ZIP"`
	PhRdNum     string `json:"PH_RD_NUM"`
	PhPreDir    string `json:"PH_PRE_DIR"`
	PhRdNam     string `json:"PH_RD_NAM"`
	PhRdTyp     string `json:"PH_RD_TYP"`
	PhUnit      string `json:"PH_UNIT"`
	PhCtyNm     string `json:"PH_CTY_NM"`
	PhZip       string `json:"PH_ZIP"`
	PhAdd       string `json:"PH_ADD"`
	Type        string `json:"TYPE"`
	AssessVal   string `json:"ASSESS_VAL"`
	ImpVal      string `json:"IMP_VAL"`
	LandVal     string `json:"LAND_VAL"`
	TotalVal    string `json:"TOTAL_VAL"`
	AssessDat   string `json:"ASSESS_DAT"`
	Nbhd        string `json:"NBHD"`
	STR         string `json:"S_T_R"`
	SchlCode    string `json:"SCHL_CODE"`
	AcreArea    string `json:"ACRE_AREA"`
	SubName     string `json:"SUB_NAME"`
	Lot         string `json:"LOT"`
	Block       string `json:"BLOCK"`
	CamaDate    string `json:"CAMA_DATE"`
	CalcAcre    string `json:"CALC_ACRE"`
	ShapeLength string `json:"Shape_Length"`
	ShapeArea   string `json:"Shape_Area"`
	GisPin      string `json:"GIS_PIN"`
	CamaPin     string `json:"CAMA_PIN"`
	ImpCount    string `json:"IMP_COUNT"`
	SchlDesc    string `json:"SCHL_DESC"`
	Proplookup  string `json:"PROPLOOKUP"`
}

// type spatialReference struct {
// 	Wkid       int `json:"wkid"`
// 	Latestwkid int `json:"latestWkid"`
// }

type attributes struct {
	Objectid    int     `json:"OBJECTID"`
	ParcelID    string  `json:"PARCEL_ID"`
	ParcelLgl   string  `json:"PARCEL_LGL"`
	SrceDate    string  `json:"SRCE_DATE"`
	OwName      string  `json:"OW_NAME"`
	OwAdd       string  `json:"OW_ADD"`
	OwAdd2      string  `json:"OW_ADD2"`
	OwCity      string  `json:"OW_CITY"`
	OwState     string  `json:"OW_STATE"`
	OwZip       string  `json:"OW_ZIP"`
	PhRdNum     string  `json:"PH_RD_NUM"`
	PhPreDir    string  `json:"PH_PRE_DIR"`
	PhRdNam     string  `json:"PH_RD_NAM"`
	PhRdTyp     string  `json:"PH_RD_TYP"`
	PhUnit      int     `json:"PH_UNIT"`
	PhCtyNm     string  `json:"PH_CTY_NM"`
	PhZip       string  `json:"PH_ZIP"`
	PhAdd       string  `json:"PH_ADD"`
	Type        string  `json:"TYPE"`
	AssessVal   int     `json:"ASSESS_VAL"`
	ImpVal      int     `json:"IMP_VAL"`
	LandVal     int     `json:"LAND_VAL"`
	TotalVal    int     `json:"TOTAL_VAL"`
	AssessDat   string  `json:"ASSESS_DAT"`
	Nbhd        int     `json:"NBHD"`
	STR         string  `json:"S_T_R"`
	SchlCode    int     `json:"SCHL_CODE"`
	AcreArea    int     `json:"ACRE_AREA"`
	SubName     string  `json:"SUB_NAME"`
	Lot         string  `json:"LOT"`
	Block       string  `json:"BLOCK"`
	CamaDate    string  `json:"CAMA_DATE"`
	CalcAcre    float64 `json:"CALC_ACRE"`
	ShapeLength float64 `json:"Shape_Length"`
	ShapeArea   float64 `json:"Shape_Area"`
	GisPin      string  `json:"GIS_PIN"`
	CamaPin     string  `json:"CAMA_PIN"`
	ImpCount    int     `json:"IMP_COUNT"`
	SchlDesc    string  `json:"SCHL_DESC"`
	Proplookup  string  `json:"PROPLOOKUP"`
}

type geometry struct {
	Rings [][][]float64 `json:"rings"`
}

type features struct {
	Attributes attributes `json:"attributes"`
	Geometry   geometry   `json:"geometry"`
}

type pResponse struct {
	Displayfieldname string           `json:"displayFieldName"`
	Fieldaliases     fieldAliases     `json:"fieldAliases"`
	Geometrytype     string           `json:"geometryType"`
	Spatialreference spatialReference `json:"spatialReference"`
	Features         []features       `json:"features"`
}

// FetchParcel takes a location ie {x:float64,Y:float64} and returns the first ESRI "ring" object given by the PAGIS REST API. A ring is a 2-dimensional array of x,y coordinates which describe the points of a (irregular) polygon.
func FetchParcel(loc Location) [][]float64 {
	parcelURL, err := url.Parse("https://pagis.org/arcgis/rest/services/APPS/OperationalLayers/MapServer/52/query")
	if err != nil {
		panic(err)
	}
	// TODO: discard all the cruft
	params := url.Values{}
	params.Add("f", "json")
	params.Add("spatialRel", "esriSpatialRelIntersects")
	params.Add("maxAllowableOffset", "1")
	geomString := fmt.Sprintf("{\"xmin\":%f,\"ymin\":%f,\"xmax\":%f,\"ymax\":%f,\"spatialReference\":{\"wkid\":102651,\"latestWkid\":3433}}", loc.X, loc.Y, loc.X+13, loc.Y+13) // spatialReference will not change
	params.Add("geometry", geomString)
	params.Add("geometryType", "esriGeometryEnvelope")
	params.Add("inSR", "102651")
	params.Add("outFields", "OBJECTID,PARCEL_ID,PARCEL_LGL,SRCE_DATE,OW_NAME,OW_ADD,OW_ADD2,OW_CITY,OW_STATE,OW_ZIP,PH_RD_NUM,PH_PRE_DIR,PH_RD_NAM,PH_RD_TYP,PH_UNIT,PH_CTY_NM,PH_ZIP,PH_ADD,TYPE,ASSESS_VAL,IMP_VAL,LAND_VAL,TOTAL_VAL,ASSESS_DAT,NBHD,S_T_R,SCHL_CODE,ACRE_AREA,SUB_NAME,LOT,BLOCK,CAMA_DATE,CALC_ACRE,Shape_Length,Shape_Area,GIS_PIN,CAMA_PIN,IMP_COUNT,SCHL_DESC,PROPLOOKUP")
	params.Add("outSR", "102651")
	params.Add("callback", "dojo_request_script_callbacks.dojo_request_script76")

	parcelURL.RawQuery = params.Encode()

	res, err := http.Get(parcelURL.String())
	if err != nil {
		panic(err)
	}
	parcelData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	// scrape off the dojo junk from our JSON
	i := bytes.IndexRune(parcelData, '{')
	t := bytes.IndexRune(parcelData, ';') - 1

	parcelData = parcelData[i:t]

	var parcel pResponse

	jsonErr := json.Unmarshal(parcelData, &parcel)
	if jsonErr != nil {
		panic(err)
	}
	return parcel.Features[0].Geometry.Rings[0]
}
