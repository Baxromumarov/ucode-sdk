package constants

const (
	AppID           = "c57eedc3-a954-4262-a0af-376c65b5a280"
	PatternRelation = `\"([^"]+)\"`
	UrlTable        = "https://api.admin.u-code.io/v1/table"
	UrlMenu         = "https://api.admin.u-code.io/v2/menus"
	UrlField        = "https://api.admin.u-code.io/v2/fields/"
	UrlRelation     = "https://api.admin.u-code.io/v2/relations/"
	GetFields       = "https://api.admin.u-code.io/v1/field"
	Token           = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfcGxhdGZvcm1faWQiOiIiLCJjbGllbnRfdHlwZV9pZCI6ImY5ZGY5YjJjLTdmMDAtNGM4NC04NmVhLWY4YWUzZTgzOTY5YSIsImRhdGEiOiJhZGRpdGlvbmFsIGpzb24gZGF0YSIsImV4cCI6MTcxNzIyMjU1NywiaWF0IjoxNzE3MTM2MTU3LCJpZCI6ImY5NDBlMzZhLTFjNTItNDRiYy05NWIwLTY4YzY2NDMzZWE1OCIsImlwIjoiYWRkaXRpb25hbCBqc29uIGRhdGEiLCJsb2dpbl90YWJsZV9zbHVnIjoiXCJ1c2VyXCIiLCJwcm9qZWN0X2lkIjoiMTIyNDA5NWItZGRkMS00NzlmLWIyNTktY2M3NjI1ZmZmMTNkIiwicm9sZV9pZCI6IjhmMmM0NWE5LTE2MDctNDMyZC1hOTUwLWQ5YzE4Zjk4ODE3ZCIsInRhYmxlcyI6W10sInVzZXJfaWQiOiJkMzZlODhlZC04NWQwLTRhMzgtOTZiNC05NWQ5ODFkMzBmZmMifQ.0jpaiypf1yWBv09aDTJYY5IkSNQHum8seetM6HeFiMY"
)

var (
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
		"switch":    "SWITCH",
		"text[]":    "MULTISELECT",
		"varchar[]": "MULTISELECT",
		"serial":    "INCREMENT_NUMBER",
	}
)
