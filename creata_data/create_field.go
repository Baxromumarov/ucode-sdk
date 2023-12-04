package creata_data

import (
	"fmt"
	"log"
	"strings"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
)

func CreateFields(table models.Table) {
	// ! create field for single line
	for fieldSlug, fieldType := range table.Fields {
		// if fieldType == "uuid"{
		// 	createRelation
		// }

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
			    	"enable_multilanguage": false,
			    	"id": "d8a11b50-baf4-4749-bf13-3a5bed72de14"
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
			fmt.Println("Single Line field created successfully", string(respCreateSingleLine))

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
					"type": "FLOAT",
					"id": "0b15c00f-8222-4c30-9f76-c66ba9ff622b"
				}`,
				fieldSlug, // label_en
				fieldSlug, // label_ru
				fieldSlug, // label_uz
				fieldSlug, // slug
				fieldSlug, // slug
				table.ID,  // table_id
			)

			respCreateSingleLine, err := helper.DoRequest(constants.UrlField, "POST", createFloatFieldBody)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Float field created successfully", string(respCreateSingleLine))

		} else {
			fmt.Print(fieldType, "\n")
			log.Fatal(" not found", " field type must be , float, number, integer, int, string,varchar")

		}
	}
}
