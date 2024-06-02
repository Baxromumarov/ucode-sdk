package structure

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
	"github.com/fatih/color"
)

func CreateFields(table models.Table) {

	// ! create field for single line
	for _, field := range table.Fields {

		createFieldBody := ""
		if _, ok := constants.FieldTypes[field.Type]; !ok {

			reqBody := helper.MultiSelectBody(field.Name, table.ID, field.Slug, field.Type)

			byteReq, err := json.Marshal(reqBody)
			if err != nil {
				log.Fatal("error while unmarshalling multi select body ", err)
			}
			createFieldBody = string(byteReq)

		} else {
			createFieldBody = helper.FieldBody(field.Name, field.Slug, table.ID, constants.FieldTypes[field.Type])
		}

		respField, err := helper.DoRequest(constants.UrlField+table.Slug, "POST", createFieldBody)
		if err != nil {
			log.Fatal("error while creating field ", err)
			return
		}

		var responseField models.CreateResponse
		fmt.Println(">>>>>>>>>>>> ", string(respField))
		if err := json.Unmarshal(respField, &responseField); err != nil {
			log.Fatal("error while unmarshalling responseField ", err)
			return
		}
		color.Green("Field successfully created")
	}
}
