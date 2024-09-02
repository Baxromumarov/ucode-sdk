package constants

const (
	AppID           = "c57eedc3-a954-4262-a0af-376c65b5a280"
	PatternRelation = `\"([^"]+)\"`
	// UrlTable        = "https://api.admin.u-code.io/v1/table"
	// UrlMenu         = "https://api.admin.u-code.io/v2/menus"
	// UrlField        = "https://api.admin.u-code.io/v2/fields/"
	// UrlRelation     = "https://api.admin.u-code.io/v2/relations/"
	// GetFields       = "https://api.admin.u-code.io/v1/field"
	UrlTable    = "https://api.admin.u-code.io/v1/table"
	UrlMenu     = "https://api.admin.u-code.io/v2/menus"
	UrlField    = "https://api.admin.u-code.io/v2/fields/"
	UrlRelation = "https://api.admin.u-code.io/v2/relations/"
	GetFields   = "https://api.admin.u-code.io/v1/field"
)

var (
	Token  = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfcGxhdGZvcm1faWQiOiIiLCJjbGllbnRfdHlwZV9pZCI6IjljMzhiYWM2LTRmYjItNDM2My05YjUyLWU3NjdhNGEzYzQxNyIsImRhdGEiOiJhZGRpdGlvbmFsIGpzb24gZGF0YSIsImV4cCI6MTcyNDIyODExOCwiaWF0IjoxNzI0MTQxNzE4LCJpZCI6ImIxY2QwZTMzLTJmYTMtNGUxNi04YjM0LWIyYzZhMzU0YWU0NSIsImlwIjoiYWRkaXRpb25hbCBqc29uIGRhdGEiLCJsb2dpbl90YWJsZV9zbHVnIjoiXCJ1c2VyXCIiLCJwcm9qZWN0X2lkIjoiMjM2NGQ1MWUtMDdmNy00MDFhLTk0OGYtZTMyNzI5MTIxMzY1Iiwicm9sZV9pZCI6ImJmYTM5MDI5LThmYjctNDkwMi05NjdmLTRlZTM5MjA1MGI1ZCIsInRhYmxlcyI6W10sInVzZXJfaWQiOiIwNzU5ZWQxYS1jMjE2LTQwYzAtYTI5NC02MjMwMzMwM2MyYmIifQ.X7lY5vUtUfhpN_2ZQkzjth-mcpQrewvKTLtO69qU7IU"
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
