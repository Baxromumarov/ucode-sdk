package structure

import (
	"encoding/json"

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

	createRelationBody := helper.RelationBody(labelEn, labelToEn, from, to)

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
	if err := json.Unmarshal(respRelationCreate, &relationResponse); err != nil {
		log.Fatal("error while unmarshalling relation: ", err)
		return
	}
	
	color.Green("Relation successfully created")

}
