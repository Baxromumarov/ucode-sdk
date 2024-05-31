package structure

import (
	"encoding/json"
	"log"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
	"github.com/fatih/color"
)

func CreateMenu(menuName string) string {

	createModuleBody := helper.MenuBody(constants.AppID, menuName)

	respCreateModule, err := helper.DoRequest(constants.UrlMenu, "POST", createModuleBody)
	if err != nil {
		log.Fatal("error while creating menu", err)
	}

	var responseModule models.CreateMenuResponse

	if err := json.Unmarshal(respCreateModule, &responseModule); err != nil {
		log.Fatal("Error while unmarshalling menu ", err)
		return ""
	}

	color.Green("Menu successfully created. ")

	return responseModule.Data.ID
}
