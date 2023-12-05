package create_data

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
)

// !NOTE app_id ---> module_id
func CreateTable(tableName string, moduleID string) string {
	// fmt.Println("modul id>>> ", moduleID)
	// ! create table body
	createTableBody := fmt.Sprintf(`{
			"show_in_menu": true,
			"fields": [],
			"app_id": "%s",
			"summary_section": {
				"id": "667045d2-0958-4f75-8cfb-f535ceecf0ac",
				"label": "Summary",
				"fields": [],
				"icon": "",
				"order": 1,
				"column": "SINGLE",
				"is_summary_section": true
			},
			"label": "%s",
			"description": "",
			"slug": "%s",
			"icon": "bear-toy.svg",
			"attributes": {
				"label_en": "%s",
				"label_ru": "%s",
				"label_uz": "%s",
				"description_en": "",
				"description_ru": "",
				"description_uz": ""
			},
			"subtitle_field_slug": "",
			"is_login_table": false,
			"is_cached": false,
			"soft_delete": false,
			"order_by": false,
			"layoutRelations": []
		}`,
		moduleID,  // app_id
		tableName, // label
		tableName, // slug
		tableName, // label_en
		tableName, // label_ru
		tableName, // label_uz
	)
	respCreateTable, err := helper.DoRequest(constants.UrlTable, "POST", createTableBody)
	if err != nil {
		log.Fatal(err)
	}
	var responseTable models.CreateResponse
	err = json.Unmarshal(respCreateTable, &responseTable)
	if err != nil {
		log.Fatal("eror while unmarshalling", err)
	}
	// fmt.Println("Table created successfully", responseTable.Data.ID)

	addTableToModule := fmt.Sprintf(`
		{
			"parent_id": "%s",
			"type": "TABLE",
			"table_id": "%s",
			"label": "new",
			"attributes": {
				"label_en": "%s",
				"label_ru": "%s",
				"label_uz": "%s",
				"description_en": "",
				"description_ru": "",
				"description_uz": ""
			},
			"icon": "bear-toy.svg"
		}`,
		moduleID,
		responseTable.Data.ID,
		tableName,
		tableName,
		tableName,
	)

	_, err = helper.DoRequest(constants.UrlModule, "POST", addTableToModule)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Add table response>> ", string(respAddTable))
	return responseTable.Data.ID
}
