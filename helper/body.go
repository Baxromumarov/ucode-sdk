package helper

import (
	"fmt"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/models"
	"github.com/google/uuid"
)

func FieldBody(fieldName, fieldSlug, tableID, fieldType string) string {
	return fmt.Sprintf(`{
					"attributes": {
						"label": "",
						"defaultValue": "",
						"label_en": "%s",
						"number_of_rounds": null
					},
					"default": "",
    				"index": "string",
    				"label": "%s",
    				"required": false,
    				"slug": "%s",
    				"table_id": "%s",
    				"type": "%s",
    				"enable_multilanguage": false,
    				"id": "%s",
    				"show_label": true
				}`,
		fieldName,           // label_en
		fieldName,           // label
		fieldSlug,           // slug
		tableID,             // table_id
		fieldType,           // type
		uuid.New().String(), // id
	)

}

func MenuBody(appId string, label string) string {
	return fmt.Sprintf(`{
		"app_id": "%s",
		"icon": "folder.svg",
		"attributes": {
			"label":"",
			"label_en": "%s"
		},
		"parent_id": "c57eedc3-a954-4262-a0af-376c65b5a284",
		"type": "FOLDER",
		"label": "%s"
		}`,
		appId, // app_id
		label, // label_en
		label, // label
	)
}

func TableBody(menuID, tableLabel, tableSlug string) string {

	return fmt.Sprintf(`
		{
			"show_in_menu": true,
			"fields": [],
			"app_id": "%s",
			"summary_section": {
				"id": "%s",
				"label": "Summary",
				"fields": [],
				"icon": "",
				"order": 1,
				"column": "SINGLE",
				"is_summary_section": true
			},
			"label": "%s",
			"description": "",
			"slug": "%s",
			"icon": "folder.svg",
			"attributes": {
				"label": "",
				"label_en": "%s"
			},
			"is_login_table": false,
			"is_cached": false,
			"soft_delete": false,
			"order_by": false,
			"layoutRelations": []
		}`,
		menuID,              // app_id
		uuid.New().String(), //summary section id
		tableLabel,          // label
		tableSlug,           // slug
		tableSlug,           // label_en
	)
}

func TableToMenuBody(menuID, tableID, tableName string) string {
	return fmt.Sprintf(`
		{
			"icon": "folder.svg",
			"attributes": {
				"label_en": "%s"
			},
			"table_id": "%s",
			"parent_id": "%s",
			"type": "TABLE",
			"label": "%s"
		}`,
		tableName,
		tableID,
		menuID,
		tableName,
	)
}

// func RelationBody(labelEn, labelToEn, oneTable, manyTable string) string {
// 	return fmt.Sprintf(`{
// 			"attributes": {
// 				"math": {
// 					"label": "plus",
// 					"value": "+"
// 				},
// 				"format": "RELATION",
// 				"options": [],
// 				"label": "",
// 				"label_to": "",
// 				"label_en": "%s",
// 				"label_to_en": "%s"
// 			},
// 			"type": "Many2One",
// 			"relation_type": "Many2One",
// 			"table_to": "%s,
// 			"view_fields": ["910affd1-c2d2-45b0-a92a-14968b7aac75"],
// 			"table_from": "%s",
// 			"filtersList": [],
// 			"columnsList": [],
// 			"relation_table_slug": "%s",
// 			"required": false,
// 			"multiple_insert": false,
// 			"show_label": true,
// 			"id": "%s"
// 		}`,
// 		labelEn,
// 		labelToEn,
// 		oneTable,
// 		manyTable,
// 		manyTable,
// 		uuid.New().String(),
// 	)
// }

func RelationBody(labelEn, labelToEn, oneTable, manyTable string) RelationCreateBody {
	return RelationCreateBody{
		Attributes: Attributes{
			Math: Math{
				Label: "plus",
				Value: "+",
			},
			Format:    "RELATION",
			Options:   []string{},
			Label:     "",
			LabelTo:   "",
			LabelEn:   labelEn,
			LabelToEn: labelToEn,
		},
		Type:              "Many2One",
		RelationType:      "Many2One",
		TableTo:           oneTable,
		ViewFields:        []string{},
		TableFrom:         manyTable,
		FiltersList:       []string{},
		ColumnsList:       []string{},
		RelationTableSlug: manyTable,
		Required:          false,
		MultipleInsert:    false,
		ShowLabel:         true,
		ID:                uuid.NewString(),
	}
	// return fmt.Sprintf(`{
	// 		"attributes": {
	// 			"math": {
	// 				"label": "plus",
	// 				"value": "+"
	// 			},
	// 			"format": "RELATION",
	// 			"options": [],
	// 			"label": "",
	// 			"label_to": "",
	// 			"label_en": "%s",
	// 			"label_to_en": "%s"
	// 		},
	// 		"type": "Many2One",
	// 		"relation_type": "Many2One",
	// 		"table_to": "%s,
	// 		"view_fields": ["910affd1-c2d2-45b0-a92a-14968b7aac75"],
	// 		"table_from": "%s",
	// 		"filtersList": [],
	// 		"columnsList": [],
	// 		"relation_table_slug": "%s",
	// 		"required": false,
	// 		"multiple_insert": false,
	// 		"show_label": true,
	// 		"id": "%s"
	// 	}`,
	// 	labelEn,
	// 	labelToEn,
	// 	oneTable,
	// 	manyTable,
	// 	manyTable,
	// 	uuid.New().String(),
	// )
}
func MultiSelectBody(label, tableId, slug, enumName string) models.MultiSelectRequestBody {
	var options = []models.Options{}

	for _, value := range constants.Enums[enumName] {
		options = append(options, models.Options{
			ID:    uuid.NewString(),
			Value: value,
			Icon:  "",
			Color: "",
			Label: value,
		})
	}

	var result = models.MultiSelectRequestBody{
		Attributes: models.Attributes{
			Label:          "",
			DefaultValue:   "",
			LabelEn:        label,
			HasColor:       false,
			IsMultiselect:  true,
			Options:        options,
			NumberOfRounds: nil,
		},
		Default:   "",
		Index:     "string",
		Label:     label,
		Required:  false,
		Slug:      slug,
		TableID:   tableId,
		Type:      "MULTISELECT",
		ID:        uuid.NewString(),
		ShowLabel: true,
	}
	return result
}

type RelationCreateBody struct {
	Attributes        Attributes `json:"attributes"`
	Type              string     `json:"type"`
	RelationType      string     `json:"relation_type"`
	TableTo           string     `json:"table_to"`
	ViewFields        []string   `json:"view_fields"`
	TableFrom         string     `json:"table_from"`
	FiltersList       []string   `json:"filtersList"`
	ColumnsList       []string   `json:"columnsList"`
	RelationTableSlug string     `json:"relation_table_slug"`
	Required          bool       `json:"required"`
	MultipleInsert    bool       `json:"multiple_insert"`
	ShowLabel         bool       `json:"show_label"`
	ID                string     `json:"id"`
}
type Math struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type Attributes struct {
	Math      Math     `json:"math"`
	Format    string   `json:"format"`
	Options   []string `json:"options"`
	Label     string   `json:"label"`
	LabelTo   string   `json:"label_to"`
	LabelEn   string   `json:"label_en"`
	LabelToEn string   `json:"label_to_en"`
}
