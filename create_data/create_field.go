package create_data

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
)

func CreateFields(table *models.Table) models.Table {
	// ! create field for single line
	for fieldSlug, fieldType := range table.FieldType {
		if fieldType == "uuid" {
			continue
		}

		if strings.Contains(fieldType, "varchar") || strings.Contains(fieldType, "string") {
			createSingleLineFieldBody := fmt.Sprintf(`{
			    	"attributes": {
			    	    "icon": "",
			    	    "label_en": "%s",
			    	    "label_ru": "%s",
			    	    "label_uz": "%s",
			    	    "show_label": true,
			    	    "defaultValue": "",
			    	    "showTooltip": false,
			    	    "number_of_rounds": null
			    	},
			    	"default": "",
			    	"index": "string",
			    	"label": "%s",
			    	"required": false,
			    	"slug": "%s",
			    	"table_id": "%s",
			    	"type": "SINGLE_LINE",
			    	"enable_multilanguage": false
				}`,
				fieldSlug, // label_en
				fieldSlug, // label_ru
				fieldSlug, // label_uz
				fieldSlug, // slug
				fieldSlug, // slug
				table.ID,  // table_id
			)
			respCreateSingleLine, err := helper.DoRequest(constants.UrlField, "POST", createSingleLineFieldBody)
			if err != nil {
				log.Fatal(err)
			}

			var responseField models.CreateResponse
			json.Unmarshal(respCreateSingleLine, &responseField)
			table.FieldID[fieldSlug] = responseField.Data.ID
			// os.Exit(1)

		} else if strings.Contains(fieldType, "float") || strings.Contains(fieldType, "number") || strings.Contains(fieldType, "int") || strings.Contains(fieldType, "integer") {
			createFloatFieldBody := fmt.Sprintf(`{
					"attributes": {
						"icon": "",
						"label_en": "%s",
						"label_ru": "%s",
						"label_uz": "%s",
						"show_label": true,
						"defaultValue": "",
						"showTooltip": false,
						"number_of_rounds": null
					},
					"default": "",
					"index": "string",
					"label": "%s",
					"required": false,
					"slug": "%s",
					"table_id": "%s",
					"type": "FLOAT"
				}`,
				fieldSlug, // label_en
				fieldSlug, // label_ru
				fieldSlug, // label_uz
				fieldSlug, // slug
				fieldSlug, // slug
				table.ID,  // table_id
			)

			respCreateFloat, err := helper.DoRequest(constants.UrlField, "POST", createFloatFieldBody)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Float field created successfully")
			var responseField models.CreateResponse
			json.Unmarshal(respCreateFloat, &responseField)
			table.FieldID[fieldSlug] = responseField.Data.ID

		} else {
			fmt.Print(fieldType, "\n")
			log.Fatal(" not found", " field type must be , float, number, integer, int, string,varchar")

		}
	}
	return *table
}