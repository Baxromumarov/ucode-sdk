package structure

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
)

// NOTE app_id ---> menu_id.
// This functions returns table id
func CreateTable(tableSlug string, menuID string, tableLabel string) string {
	fmt.Println("MENU ID: ", menuID, tableSlug, tableLabel)
	// ! create table body
	createTableBody := helper.TableBody(menuID, tableLabel, tableSlug)
	respCreateTable, err := helper.DoRequest(constants.UrlTable, "POST", createTableBody)
	if err != nil {
		log.Fatal("error while creating table: ", err)
	}

	var responseTable models.CreateResponse

	if err = json.Unmarshal(respCreateTable, &responseTable); err != nil {
		fmt.Println(string(respCreateTable))
		log.Fatal("error while unmarshalling table response:1 ", err)
	}

	fmt.Println("Table created successfully.", responseTable.Data.ID)

	addTableToModule := helper.TableToMenuBody(menuID, responseTable.Data.ID, tableLabel)

	_, err = helper.DoRequest(constants.UrlMenu, "POST", addTableToModule)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Table successfully added to menu")
	return responseTable.Data.ID
}
