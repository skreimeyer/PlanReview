package planreview

type fields struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Alias string `json:"alias"`
}

type znAttributes struct {
	GisLrGisplanZNumberObjectid1                int    `json:"GIS_LR.GISPLAN.Z_Number.OBJECTID_1"`
	GisLrGisplanZNumberObjectid                 int    `json:"GIS_LR.GISPLAN.Z_Number.OBJECTID"`
	GisLrGisplanZNumberLabel                    string `json:"GIS_LR.GISPLAN.Z_Number.LABEL"`
	GisLrGisplanZNumberLinkznum                 string `json:"GIS_LR.GISPLAN.Z_Number.LinkZnum"`
	GisLrGisplanZNumberLastupdate               int    `json:"GIS_LR.GISPLAN.Z_Number.LastUpdate"`
	GisLrGisplanZNumberEditorname               string `json:"GIS_LR.GISPLAN.Z_Number.EditorName"`
	GisLrGisplanZoningInputObjectid             int    `json:"GIS_LR.GISPLAN.Zoning_Input.OBJECTID"`
	GisLrGisplanZoningInputZNumber              string `json:"GIS_LR.GISPLAN.Zoning_Input.Z_Number"`
	GisLrGisplanZoningInputZonedFrom            int    `json:"GIS_LR.GISPLAN.Zoning_Input.Zoned_From"`
	GisLrGisplanZoningInputZonedTo              int    `json:"GIS_LR.GISPLAN.Zoning_Input.Zoned_To"`
	GisLrGisplanZoningInputOtherZNumbers        int    `json:"GIS_LR.GISPLAN.Zoning_Input.Other_Z_Numbers"`
	GisLrGisplanZoningInputConditionalUse       int    `json:"GIS_LR.GISPLAN.Zoning_Input.Conditional_Use"`
	GisLrGisplanZoningInputAnyConditions        int    `json:"GIS_LR.GISPLAN.Zoning_Input.Any_Conditions"`
	GisLrGisplanZoningInputOrdinanceNumber      int    `json:"GIS_LR.GISPLAN.Zoning_Input.Ordinance_Number"`
	GisLrGisplanZoningInputAreaZoned            int    `json:"GIS_LR.GISPLAN.Zoning_Input.Area_Zoned"`
	GisLrGisplanZoningInputMultipleActions      int    `json:"GIS_LR.GISPLAN.Zoning_Input.Multiple_Actions"`
	GisLrGisplanZoningInputIssueActionRequested string `json:"GIS_LR.GISPLAN.Zoning_Input.Issue_Action_Requested"`
	GisLrGisplanZoningInputLocation             string `json:"GIS_LR.GISPLAN.Zoning_Input.Location"`
	GisLrGisplanZoningInputOtherZoneCategories  int    `json:"GIS_LR.GISPLAN.Zoning_Input.Other_Zone_Categories"`
	GisLrGisplanZoningInputVariance             int    `json:"GIS_LR.GISPLAN.Zoning_Input.Variance"`
	GisLrGisplanZoningInputApproved             int    `json:"GIS_LR.GISPLAN.Zoning_Input.Approved"`
	GisLrGisplanZoningInputDateOfPcAction       int    `json:"GIS_LR.GISPLAN.Zoning_Input.Date_of_PC_Action"`
	GisLrGisplanZoningInputDateOfBoaAction      string `json:"GIS_LR.GISPLAN.Zoning_Input.Date_of_BOA_Action"`
	GisLrGisplanZoningInputDateOfBoardAction    int    `json:"GIS_LR.GISPLAN.Zoning_Input.Date_of_Board_Action"`
}

type znFeatures struct {
	Attributes znAttributes `json:"attributes"`
	Geometry   geometry     `json:"geometry"`
}

type root struct {
	Displayfieldname string           `json:"displayFieldName"`
	Fieldaliases     fieldAliases     `json:"fieldAliases"`
	Geometrytype     string           `json:"geometryType"`
	Spatialreference spatialReference `json:"spatialReference"`
	Fields           []fields         `json:"fields"`
	Features         []znFeatures     `json:"features"`
}
