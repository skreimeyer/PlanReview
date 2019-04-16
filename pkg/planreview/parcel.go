package planreview

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// borrowed from geocode.go
type location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type fieldAliases struct {
	Objectid     string `json:"OBJECTID"`
	Parcel_Id    string `json:"PARCEL_ID"`
	Parcel_Lgl   string `json:"PARCEL_LGL"`
	Srce_Date    string `json:"SRCE_DATE"`
	Ow_Name      string `json:"OW_NAME"`
	Ow_Add       string `json:"OW_ADD"`
	Ow_Add2      string `json:"OW_ADD2"`
	Ow_City      string `json:"OW_CITY"`
	Ow_State     string `json:"OW_STATE"`
	Ow_Zip       string `json:"OW_ZIP"`
	Ph_Rd_Num    string `json:"PH_RD_NUM"`
	Ph_Pre_Dir   string `json:"PH_PRE_DIR"`
	Ph_Rd_Nam    string `json:"PH_RD_NAM"`
	Ph_Rd_Typ    string `json:"PH_RD_TYP"`
	Ph_Unit      string `json:"PH_UNIT"`
	Ph_Cty_Nm    string `json:"PH_CTY_NM"`
	Ph_Zip       string `json:"PH_ZIP"`
	Ph_Add       string `json:"PH_ADD"`
	Type         string `json:"TYPE"`
	Assess_Val   string `json:"ASSESS_VAL"`
	Imp_Val      string `json:"IMP_VAL"`
	Land_Val     string `json:"LAND_VAL"`
	Total_Val    string `json:"TOTAL_VAL"`
	Assess_Dat   string `json:"ASSESS_DAT"`
	Nbhd         string `json:"NBHD"`
	S_T_R        string `json:"S_T_R"`
	Schl_Code    string `json:"SCHL_CODE"`
	Acre_Area    string `json:"ACRE_AREA"`
	Sub_Name     string `json:"SUB_NAME"`
	Lot          string `json:"LOT"`
	Block        string `json:"BLOCK"`
	Cama_Date    string `json:"CAMA_DATE"`
	Calc_Acre    string `json:"CALC_ACRE"`
	Shape_Length string `json:"Shape_Length"`
	Shape_Area   string `json:"Shape_Area"`
	Gis_Pin      string `json:"GIS_PIN"`
	Cama_Pin     string `json:"CAMA_PIN"`
	Imp_Count    string `json:"IMP_COUNT"`
	Schl_Desc    string `json:"SCHL_DESC"`
	Proplookup   string `json:"PROPLOOKUP"`
}

type spatialReference struct {
	Wkid       int `json:"wkid"`
	Latestwkid int `json:"latestWkid"`
}

type attributes struct {
	Objectid     int     `json:"OBJECTID"`
	Parcel_Id    string  `json:"PARCEL_ID"`
	Parcel_Lgl   string  `json:"PARCEL_LGL"`
	Srce_Date    string  `json:"SRCE_DATE"`
	Ow_Name      string  `json:"OW_NAME"`
	Ow_Add       string  `json:"OW_ADD"`
	Ow_Add2      string  `json:"OW_ADD2"`
	Ow_City      string  `json:"OW_CITY"`
	Ow_State     string  `json:"OW_STATE"`
	Ow_Zip       string  `json:"OW_ZIP"`
	Ph_Rd_Num    string  `json:"PH_RD_NUM"`
	Ph_Pre_Dir   string  `json:"PH_PRE_DIR"`
	Ph_Rd_Nam    string  `json:"PH_RD_NAM"`
	Ph_Rd_Typ    string  `json:"PH_RD_TYP"`
	Ph_Unit      int     `json:"PH_UNIT"`
	Ph_Cty_Nm    string  `json:"PH_CTY_NM"`
	Ph_Zip       string  `json:"PH_ZIP"`
	Ph_Add       string  `json:"PH_ADD"`
	Type         string  `json:"TYPE"`
	Assess_Val   int     `json:"ASSESS_VAL"`
	Imp_Val      int     `json:"IMP_VAL"`
	Land_Val     int     `json:"LAND_VAL"`
	Total_Val    int     `json:"TOTAL_VAL"`
	Assess_Dat   string  `json:"ASSESS_DAT"`
	Nbhd         int     `json:"NBHD"`
	S_T_R        string  `json:"S_T_R"`
	Schl_Code    int     `json:"SCHL_CODE"`
	Acre_Area    int     `json:"ACRE_AREA"`
	Sub_Name     string  `json:"SUB_NAME"`
	Lot          string  `json:"LOT"`
	Block        string  `json:"BLOCK"`
	Cama_Date    string  `json:"CAMA_DATE"`
	Calc_Acre    float64 `json:"CALC_ACRE"`
	Shape_Length float64 `json:"Shape_Length"`
	Shape_Area   float64 `json:"Shape_Area"`
	Gis_Pin      string  `json:"GIS_PIN"`
	Cama_Pin     string  `json:"CAMA_PIN"`
	Imp_Count    int     `json:"IMP_COUNT"`
	Schl_Desc    string  `json:"SCHL_DESC"`
	Proplookup   string  `json:"PROPLOOKUP"`
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

func FetchParcel(loc location) [][]float64 {
	parcelUrl, err := url.Parse("https://pagis.org/arcgis/rest/services/APPS/OperationalLayers/MapServer/52/query")
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

	parcelUrl.RawQuery = params.Encode()

	res, err := http.Get(parcelUrl.String())
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
