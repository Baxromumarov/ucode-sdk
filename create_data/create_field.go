package create_data

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

func CreateFields(table models.Table, allTables map[string]string) {

	// ! create field for single line
	for fieldSlug, fieldType := range table.FieldType {

		if fieldType == "uuid" && fieldSlug != "guid" {
			tableInfo := strings.Split(fieldSlug, ".")
			fieldId := ""
			tableId := allTables[tableInfo[0][:len(tableInfo[0])-3]]
			fmt.Println("MACWIN TABLE ID: ", tableId)
			var (
				getFieldsUrl       = constants.GetFields + fmt.Sprintf(`?table_id=%s`, tableId)
				listTablesResponse models.FieldResponse
			)

			listTableResp, err := helper.DoRequest(getFieldsUrl, "GET", "")
			if err != nil {
				log.Fatal("error2 ", err)
			}

			if err := json.Unmarshal(listTableResp, &listTablesResponse); err != nil {
				log.Fatal("error1  ", err)
			}

			for _, singleField := range listTablesResponse.Data.Fields {
				if cast.ToString(singleField["slug"]) == tableInfo[1] {
					fieldId = cast.ToString(singleField["id"])
				}
			}

			var (
				createRelationBody = fmt.Sprintf(`{
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
					tableInfo[1],
					tableInfo[1],
					tableInfo[1],
					tableInfo[0][:len(tableInfo[0])-3],
					fieldId,
					table.Name)
			)
			_, err = helper.DoRequest(constants.UrlRelation, "POST", createRelationBody)
			if err != nil {
				log.Fatal("><><>", err)
			}
			fmt.Println("created relation to ", tableInfo[0], tableInfo[1])
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
			fmt.Println("Single Line field created successfully")
			var responseField models.CreateResponse
			err = json.Unmarshal(respCreateSingleLine, &responseField)
			if err != nil {
				fmt.Println("Here ", err)
				log.Fatal(err)
			}
			// table.FieldID[fieldSlug] = responseField.Data.ID
			// os.Exit(1)

		} else if strings.Contains(fieldType, "float") || strings.Contains(fieldType, "number") || strings.Contains(fieldType, "int") || strings.Contains(fieldType, "integer") {
			fmt.Println("table id", table.ID, fieldSlug)
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
			fmt.Println("Float field created successfully", fieldSlug, fieldType)
			var responseField models.CreateResponse
			err = json.Unmarshal(respCreateFloat, &responseField)
			if err != nil {
				fmt.Println("Here ", err)
				log.Fatal(err)
			}
			// table.FieldID[fieldSlug] = responseField.Data.ID

		} else if fieldType == "uuid" {
			continue
		} else {
			fmt.Print(fieldType, "\n")
			log.Fatal(" not found", " field type must be , float, number, integer, int, string,varchar")

		}
	}
}
