package create_data

/*

func CreateRelation() {
	if fieldType == "uuid" && fieldSlug != "guid" {
		//Get table
		//Get fields
		tableInfo := strings.Split(fieldSlug, ".")
		fieldId := ""

		var (
			getFieldsUrl       = constants.GetFields + fmt.Sprintf(`table_id=%s`, table.ID)
			listTablesResponse models.FieldResponse
		)

		listTableResp, err := helper.DoRequest(getFieldsUrl, "GET", "")
		if err != nil {
			log.Fatal(err)
			return
		}

		if err := json.Unmarshal(listTableResp, &listTablesResponse); err != nil {
			log.Fatal(err)
			return
		}

		for _, field := range listTablesResponse.Data.Fields {
			if cast.ToString(field["slug"]) == tableInfo[1] {
				fieldId = cast.ToString(field["id"])
			}
		}

		createRelationBody := fmt.Sprintf(`{
			"table_from": "%s",
			"auto_filters": [],
			"action_relations": [],
			"attributes": {
			  "title_en": "%s",
			  "title_ru": "%s",
			  "title_uz": "%s",
			  "multiple_input": false
			},
			"table_to": "%s",
			"type": "Many2One",
			"default_limit": "",
			"multiple_insert": false,
			"filtersList": [],
			"columnsList": [],
			"view_fields": [
			  "%s"
			],
			"relation_table_slug": "%s",
			"columns": [],
			"quick_filters": [],
			"default_values": [],
			"title": "user"
		  }`,
			table.Name,
			tableInfo[1], tableInfo[1], tableInfo[1],
			tableInfo[0],
			fieldId,
			table.Name,
		)
		_, err = helper.DoRequest(constants.UrlRelation, "POST", createRelationBody)
		if err != nil {
			log.Fatal("><><>", err)
			return
		}
		fmt.Println("created relation to ", tableInfo[0], tableInfo[1])
		// createRelation

	}
}

*/
