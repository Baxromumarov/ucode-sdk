package constants

const (
	AppID           = "c57eedc3-a954-4262-a0af-376c65b5a280"
	PatternRelation = `\"([^"]+)\"`
	UrlTable        = "https://api.admin.u-code.io/v1/table"
	UrlMenu         = "https://api.admin.u-code.io/v2/menus"
	UrlField        = "https://api.admin.u-code.io/v2/fields/"
	UrlRelation     = "https://api.admin.u-code.io/v2/relations/"
	GetFields       = "https://api.admin.u-code.io/v1/field"
)

var (
	Token  = ""
	MenuID = ""
	Enums  = map[string][]string{}

	FieldTypes = map[string]string{
		"varchar":   "SINGLE_LINE",
		"text":      "MULTI_LINE",
		"email":     "EMAIL",
		"photo":     "PHOTO",
		"phone":     "INTERNATIONAL_PHONE",
		"file":      "FILE",
		"date":      "DATE",
		"time":      "TIME",
		"date_time": "DATE_TIME",
		"timestamp": "DATE_TIME_WITHOUT_TIME_ZONE",
		"float":     "FLOAT",
		"number":    "FLOAT",
		"int":       "FLOAT",
		"integer":   "FLOAT",
		"checkbox":  "CHECKBOX",
		"bool":      "CHECKBOX",
		"switch":    "SWITCH",
		"text[]":    "MULTISELECT",
		"varchar[]": "MULTISELECT",
		"serial":    "INCREMENT_NUMBER",
	}
)
