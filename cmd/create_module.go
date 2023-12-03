package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
)

func CreateModule(moduleName string) string {
	createModuleBody := fmt.Sprintf(`{
		"app_id": "29a8af0c-7582-419a-b12f-53cc9e567564",
		"icon": "bear-toy.svg",
		"attributes": {
			"label_en": "%s",
			"label_ru": "%s",
			"label_uz": "%s"
		},
		"parent_id": "c57eedc3-a954-4262-a0af-376c65b5a284",
		"type": "FOLDER",
		"label": "%s"
	}`,
		moduleName, // label_en
		moduleName, // label_ru
		moduleName, // label_uz
		moduleName, // label
	)

	respCreateModule, err := helper.DoRequest(constants.UrlModule, "POST", createModuleBody)
	if err != nil {

		log.Fatal(err)
	}
	var responseModule CreateResponse
	json.Unmarshal(respCreateModule, &responseModule)
	fmt.Println("Table created successfully", responseModule.Data.ID)
	return responseModule.Data.ID
}
