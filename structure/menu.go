package structure

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
)

func CreateMenu(menuName string) string {

	createModuleBody := helper.MenuBody(constants.AppID, menuName)

	respCreateModule, err := helper.DoRequest(constants.UrlMenu, "POST", createModuleBody)
	if err != nil {
		log.Fatal("error while creating menu", err)
	}

	var responseModule models.CreateMenuResponse
	fmt.Println(">>>>>>> ", string(respCreateModule))
	if err := json.Unmarshal(respCreateModule, &responseModule); err != nil {
		log.Fatal("Error while unmarshalling menu ", err)
	}

	fmt.Println("Menu created successfully. Menu id: ", responseModule.Data.ID)
	return responseModule.Data.ID
}
