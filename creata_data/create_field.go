package creata_data

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
	"github.com/spf13/cast"
)

func CreateFields(table models.Table) {

	// ! create field for single line
	for fieldSlug, fieldType := range table.Fields {
		fmt.Println("1.>>>>>>>>>", fieldType)

		if fieldType == "uuid" && fieldSlug != "guid" {
			//Get table
			//Get fields
			tableInfo := strings.Split(fieldSlug, ".")
			fieldId := ""

			var (
				getFieldsUrl       = constants.GetFields + fmt.Sprintf(`table_id=%s`, table.ID)
				listTablesResponse models.FieldResponse
			)

			listTableResp, err := helper.DoRequest(getFieldsUrl, "GET", "")
			if err != nil {
				log.Fatal(err)
				return
			}

			if err := json.Unmarshal(listTableResp, &listTablesResponse); err != nil {
				log.Fatal(err)
				return
			}

			for _, field := range listTablesResponse.Data.Fields {
				if cast.ToString(field["slug"]) == tableInfo[1] {
					fieldId = cast.ToString(field["id"])
				}
			}

			var (
				createRelationRequest = constants.UrlRelation
				createRelationBody    = fmt.Sprintf(`{
					"table_from": "%s",
					"auto_filters": [],
					"action_relations": [],
					"attributes": {
					  "title_en": "%s",
					  "title_ru": "%s",
					  "title_uz": "%s",
					  "multiple_input": false
					},
					"table_to": "%s",
					"type": "Many2One",
					"default_limit": "",
					"multiple_insert": false,
					"filtersList": [],
					"columnsList": [],
					"view_fields": [
					  "%s"
					],
					"relation_table_slug": "%s",
					"columns": [],
					"quick_filters": [],
					"default_values": [],
					"title": "user"
				  }`,
					table.Name,
					tableInfo[1], tableInfo[1], tableInfo[1],
					tableInfo[0],
					fieldId,
					table.Name)
			)
			_, err = helper.DoRequest(createRelationRequest, "POST", createRelationBody)
			if err != nil {
				log.Fatal("><><>", err)
				return
			}
			fmt.Println("created relation to ", tableInfo[0], tableInfo[1])
			// createRelation

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
					"type": "FLOAT"
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