package structure

import (
	"encoding/json"
	"fmt"
	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
	"log"
)

func CreateFields(table models.Table) {

	// ! create field for single line
	for _, field := range table.Fields {
		fmt.Println(field.Name)
		fmt.Println(field.Slug)
		fmt.Println(table.ID)
		fmt.Println(field.Type)
		fmt.Println(constants.FieldTypes[field.Type])
		createFieldBody := helper.FieldBody(field.Name, field.Slug, table.ID, constants.FieldTypes[field.Type])
		// fmt.Println(createFieldBody)
		respField, err := helper.DoRequest(constants.UrlField+table.Slug, "POST", createFieldBody)
		fmt.Println(">>>>>>>>", constants.UrlField+table.Slug)
		if err != nil {
			log.Fatal(err)
		}

		var responseField models.CreateResponse
		fmt.Println(">>>>>>123412 ", string(respField))

		if err := json.Unmarshal(respField, &responseField); err != nil {
			fmt.Println("Here12 ", err)
			log.Fatal(err)
		}
		fmt.Println("Single Line field created successfully")
	}
}
