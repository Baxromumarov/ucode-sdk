package structure

import (
	"encoding/json"
	"log"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
	"github.com/fatih/color"
)

// NOTE app_id ---> menu_id.
// This functions returns table id
func CreateTable(tableSlug string, menuID string, tableLabel string) string {
	// ! create table body
	createTableBody := helper.TableBody(menuID, tableLabel, tableSlug)
	respCreateTable, err := helper.DoRequest(constants.UrlTable, "POST", createTableBody)
	if err != nil {
		log.Fatal("error while creating table: ", err)
		return ""
	}

	var responseTable models.CreateResponse

	if err = json.Unmarshal(respCreateTable, &responseTable); err != nil {
		log.Fatal("error while unmarshalling table response:1 ", err)
		return ""
	}

	addTableToModule := helper.TableToMenuBody(menuID, responseTable.Data.ID, tableLabel)

	_, err = helper.DoRequest(constants.UrlMenu, "POST", addTableToModule)
	if err != nil {
		log.Fatal("error while creating Table ", err)
		return ""
	}

	color.Green("Table successfully created")
	return responseTable.Data.ID
}
