package structure

import (
	"encoding/json"
	"fmt"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"

	"log"
)

// second_table many table
// test_table one table

func CreateRelation(line string) {
	from, to, labelEn, labelToEn := helper.RelationParser(line)
	fmt.Println("from: ", from)
	fmt.Println("to: ", to)
	fmt.Println("labelEn: ", labelEn)
	fmt.Println("labelToEn: ", labelToEn)

	createRelationBody := helper.RelationBody(labelEn, labelToEn, from, to)
	fmt.Println("createRelationBody: ", createRelationBody)
	req, _ := json.Marshal(createRelationBody)
	respRelationCreate, err := helper.DoRequest(constants.UrlRelation+from, "POST", string(req))
	if err != nil {
		log.Fatal("error while creating relation: ", err)
	}

	var relationResponse models.RelationResponse
	fmt.Println(">>>>>>1234 ", string(respRelationCreate))
	if err := json.Unmarshal(respRelationCreate, &relationResponse); err != nil {
		log.Fatal("error while unmarshalling relation: ", err)
	}

}
