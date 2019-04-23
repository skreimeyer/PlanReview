package planreview

type fieldAliases struct {
	ObjectID         string `json:"OBJECTID"`
	ParcelID         string `json:"PARCEL_ID"`
	GisPin           string `json:"GIS_PIN"`
	CamaPin          string `json:"CAMA_PIN"`
	CamaDate         string `json:"CAMA_DATE"`
	ParcelLgl        string `json:"PARCEL_LGL"`
	SrceDate         string `json:"SRCE_DATE"`
	OwName           string `json:"OW_NAME"`
	OwAdd            string `json:"OW_ADD"`
	OwAdd2           string `json:"OW_ADD2"`
	OwCity           string `json:"OW_CITY"`
	OwState          string `json:"OW_STATE"`
	OwZip            string `json:"OW_ZIP"`
	PhRdNum          string `json:"PH_RD_NUM"`
	PhPreDir         string `json:"PH_PRE_DIR"`
	PhRdNam          string `json:"PH_RD_NAM"`
	PhRdTyp          string `json:"PH_RD_TYP"`
	PhUnit           string `json:"PH_UNIT"`
	PhCtyNm          string `json:"PH_CTY_NM"`
	PhZip            string `json:"PH_ZIP"`
	PhAdd            string `json:"PH_ADD"`
	Type             string `json:"TYPE"`
	ImpCount         string `json:"IMP_COUNT"`
	AssessDat        string `json:"ASSESS_DAT"`
	AssessVal        string `json:"ASSESS_VAL"`
	ImpVal           string `json:"IMP_VAL"`
	LandVal          string `json:"LAND_VAL"`
	TotalVal         string `json:"TOTAL_VAL"`
	Nbhd             string `json:"NBHD"`
	STR              string `json:"S_T_R"`
	SchlCode         string `json:"SCHL_CODE"`
	SchlDesc         string `json:"SCHL_DESC"`
	SubName          string `json:"SUB_NAME"`
	Lot              string `json:"LOT"`
	Block            string `json:"BLOCK"`
	AcreArea         string `json:"ACRE_AREA"`
	CalcAcre         string `json:"CALC_ACRE"`
	Proplookup       string `json:"PROPLOOKUP"`
	ShapeXStareaXX   string `json:"Shape_STArea()"`
	ShapeXStlengthXX string `json:"Shape_STLength()"`
}

type spatialReference struct {
	Wkid       int `json:"wkid"`
	Latestwkid int `json:"latestWkid"`
}

type attributes struct {
	ObjectID         int     `json:"OBJECTID"`
	ParcelID         string  `json:"PARCEL_ID"`
	GisPin           string  `json:"GIS_PIN"`
	CamaPin          string  `json:"CAMA_PIN"`
	CamaDate         string  `json:"CAMA_DATE"`
	ParcelLgl        string  `json:"PARCEL_LGL"`
	SrceDate         string  `json:"SRCE_DATE"`
	OwName           string  `json:"OW_NAME"`
	OwAdd            string  `json:"OW_ADD"`
	OwAdd2           string  `json:"OW_ADD2"`
	OwCity           string  `json:"OW_CITY"`
	OwState          string  `json:"OW_STATE"`
	OwZip            string  `json:"OW_ZIP"`
	PhRdNum          string  `json:"PH_RD_NUM"`
	PhPreDir         string  `json:"PH_PRE_DIR"`
	PhRdNam          string  `json:"PH_RD_NAM"`
	PhRdTyp          string  `json:"PH_RD_TYP"`
	PhUnit           int     `json:"PH_UNIT"`
	PhCtyNm          string  `json:"PH_CTY_NM"`
	PhZip            string  `json:"PH_ZIP"`
	PhAdd            string  `json:"PH_ADD"`
	Type             string  `json:"TYPE"`
	ImpCount         int     `json:"IMP_COUNT"`
	AssessDat        string  `json:"ASSESS_DAT"`
	AssessVal        int     `json:"ASSESS_VAL"`
	ImpVal           int     `json:"IMP_VAL"`
	LandVal          int     `json:"LAND_VAL"`
	TotalVal         int     `json:"TOTAL_VAL"`
	Nbhd             int     `json:"NBHD"`
	STR              string  `json:"S_T_R"`
	SchlCode         int     `json:"SCHL_CODE"`
	SchlDesc         string  `json:"SCHL_DESC"`
	SubName          string  `json:"SUB_NAME"`
	Lot              string  `json:"LOT"`
	Block            string  `json:"BLOCK"`
	AcreArea         float64 `json:"ACRE_AREA"`
	CalcAcre         float64 `json:"CALC_ACRE"`
	Proplookup       string  `json:"PROPLOOKUP"`
	ShapeXStareaXX   float64 `json:"Shape_STArea()"`
	ShapeXStlengthXX float64 `json:"Shape_STLength()"`
}

type geometry struct {
	Rings [][][]float64 `json:"rings"`
}

type features struct {
	Attributes attributes `json:"attributes"`
	Geometry   geometry   `json:"geometry"`
}

type root struct {
	Displayfieldname string           `json:"displayFieldName"`
	Fieldaliases     fieldAliases     `json:"fieldAliases"`
	Geometrytype     string           `json:"geometryType"`
	Spatialreference spatialReference `json:"spatialReference"`
	Features         []features       `json:"features"`
}
