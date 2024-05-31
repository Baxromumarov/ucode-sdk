package constants

const (
	AppID           = "c57eedc3-a954-4262-a0af-376c65b5a280"
	PatternRelation = `\"([^"]+)\"`
	UrlTable        = "https://api.admin.u-code.io/v1/table"
	UrlMenu         = "https://api.admin.u-code.io/v2/menus"
	UrlField        = "https://api.admin.u-code.io/v2/fields/"
	UrlRelation     = "https://api.admin.u-code.io/v2/relations/"
	GetFields       = "https://api.admin.u-code.io/v1/field"
	Token           = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfcGxhdGZvcm1faWQiOiIiLCJjbGllbnRfdHlwZV9pZCI6ImI3NDVlNjY0LThmMTEtNGRhYS1iNWQ5LTQ1NWNiZGFjODRjNSIsImRhdGEiOiJhZGRpdGlvbmFsIGpzb24gZGF0YSIsImV4cCI6MTcxNzIxNDQyOSwiaWF0IjoxNzE3MTI4MDI5LCJpZCI6IjdhZGUwNTQwLWZmZWQtNDY0Yy05M2Q4LTA4NDIzMDY2ODdkZiIsImlwIjoiYWRkaXRpb25hbCBqc29uIGRhdGEiLCJsb2dpbl90YWJsZV9zbHVnIjoiXCJ1c2VyXCIiLCJwcm9qZWN0X2lkIjoiNjUyOGIwNmMtMjg0Mi00OTA1LWEwYWEtYmM0MDY1M2FmYzU4Iiwicm9sZV9pZCI6ImI0YmI4YTlhLTNjNDctNGNjNC1hNmJhLTZiMGU4YjA5MTIyMCIsInRhYmxlcyI6W10sInVzZXJfaWQiOiI0MWFlYzc3OS02ZWIwLTQ4YjMtYmMzMy05Mzk0ZTFiZjYwNDcifQ.RwV7wog9aR8VsLIxHFVbxs9-IUfTX311BQlx7S1pL2c"
)

var (
	MenuID = ""
	Enums  = map[string][]string{}

	FieldTypes = map[string]string{
		"varchar":   "SINGLE_LINE",
		"text":      "MULTI_LINE",
		"email":     "EMAIL",
		"photo":     "PHOTO",
		"phone":     "PHONE",
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
		"switch":    "SWITCH",
		"text[]":    "MULTISELECT",
		"varchar[]": "MULTISELECT",
		"serial":    "INCREMENT_NUMBER",
	}
)
