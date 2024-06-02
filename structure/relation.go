package structure

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"

	"log"
)

// second_table many table
// test_table one table

func CreateRelation(line string) {
	from, to, labelEn, labelToEn := helper.RelationParser(line)
	// fmt.Println(from)
	// fmt.Println(to)
	// fmt.Println(labelEn)
	// fmt.Println(labelToEn)
	createRelationBody := helper.RelationBody(labelEn, labelToEn, from, to)

	fmt.Println("Create relation request body: ", createRelationBody)

	req, err := json.Marshal(createRelationBody)
	if err != nil {
		log.Fatal("error while Marshalling createRelationBody ", err)
		return
	}
	respRelationCreate, err := helper.DoRequest(constants.UrlRelation+from, "POST", string(req))
	if err != nil {
		log.Fatal("error while creating relation: ", err)
		return
	}

	var relationResponse models.RelationResponse
	fmt.Println(">>>>>>>>>>>>>234>>>>  ", string(respRelationCreate))
	if err := json.Unmarshal(respRelationCreate, &relationResponse); err != nil {
		log.Fatal("error while unmarshalling relation: ", err)
		return
	}

	color.Green("Relation successfully created")

}
