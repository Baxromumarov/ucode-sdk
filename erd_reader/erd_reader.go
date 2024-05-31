package structure

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strings"

	"github.com/baxromumarov/ucode-sdk/constants"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
	"github.com/baxromumarov/ucode-sdk/structure"
	"github.com/fatih/color"
)

func Reader() {
	if constants.AppID == "" || constants.Token == "" {
		log.Fatal("app_id or token is empty")
	}

	defer func() {
		color.Green("Congratulations, BAXROM!!! You did it. GO and drink. 😄")

	}()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error while getting directory ", err)
	}

	filePath := dir + "/erd.sql"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		tables = map[string]string{} // tableName:id, Филиал:31131c67-1d74-4a85-8473-eaa5d10240e5
		table  = models.Table{
			Fields: []models.Field{},
		}
	)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, "guid") && strings.Contains(line, "uuid") {
			// this is checking for default table uuid
			// we just skip this part
			continue
		}

		if strings.Contains(line, "CREATE TYPE ") {
			// then this is enum (multi select). Store it "enums" map
			// And if later face this type then create a multi select
			enumName := helper.EnumParser(line)
			// fmt.Println(enumName)
			var multiSelectValues []string
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, ");") {
					break
				}

				multiSelectValue := helper.EnumNameParser(line)
				multiSelectValues = append(multiSelectValues, multiSelectValue)

				// fmt.Println(">>>>> ", multiSelectValue)
			}
			constants.Enums[enumName] = multiSelectValues
		}

		// if strings.Contains(line, ".") {
		// 	//! this vulnerable because other lines have also '.' symbol
		// 	fmt.Println(">>>>>>>>> ", line)
		// 	fieldName, fieldSlug, fieldType := helper.FieldParser(line)
		// 	fmt.Println(fieldName)
		// 	fmt.Println(fieldSlug)
		// 	fmt.Println(fieldType)
		// }
		fmt.Println(constants.Enums)
		// continue

		if strings.Contains(line, "ALTER TABLE ") {
			// *this means relation creating started in .sql file
			fmt.Println("<<<<<<<<< relation creating ")
			structure.CreateRelation(line)
			continue
		}

		// splitLine := strings.Split(line, " ")

		if strings.Contains(line, "CREATE TABLE") {

			tableName, tableSlug, menuName := helper.TableParser(line)

			if menuName != "" {
			// this means 
				constants.MenuID = structure.CreateMenu(menuName)
				tableID := structure.CreateTable(tableSlug, constants.MenuID, tableName)

				// table data collecting
				table.ID = tableID
				table.Slug = tableSlug
				table.Name = tableName

			} else {
				tableID := structure.CreateTable(tableSlug, constants.MenuID, tableName)

				// table data collecting
				table.ID = tableID
				table.Slug = tableSlug
				table.Name = tableName
			}

			// fmt.Println("SPLITLINE: ", splitLine[2])
			// data := strings.Split(splitLine[2], ".")
			// if len(data) == 3 {
			// 	menuName, tableName, tableSlug := data[0], data[1], data[2]
			// 	menuName = menuName[1:]
			// 	tableSlug = tableSlug[:len(tableSlug)-1]

			// 	constants.MenuID = structure.CreateMenu(menuName)
			// 	tableID := structure.CreateTable(tableSlug, constants.MenuID, tableName)

			// 	//  table data collecting
			// 	table.ID = tableID
			// 	table.Slug = tableSlug
			// 	table.Name = tableName

			// } else {
			// 	fmt.Println(">>>>>>> ", data)
			// 	tableName, tableSlug := data[0], data[1]
			// 	tableName = tableName[1:]
			// 	tableSlug = tableSlug[:len(tableSlug)-1]

			// 	tableID := structure.CreateTable(tableSlug, constants.MenuID, tableName)

			// 	//  table data collecting
			// 	table.ID = tableID
			// 	table.Slug = tableSlug
			// 	table.Name = tableName
			// }

			tables[table.Name] = table.ID

			continue
		}

		if strings.Contains(line, ".") {
			//! this vulnerable because other lines have also '.' symbol

			fieldName, fieldSlug, fieldType := helper.FieldParser(line)

			table.Fields = append(table.Fields, models.Field{
				Name: fieldName,
				Slug: fieldSlug,
				Type: fieldType,
			})
		}

		if strings.Contains(line, ");") {
			structure.CreateFields(table)

			table = models.Table{
				Fields: []models.Field{},
			}
		}

		if strings.Contains(line, "ALTER") {
			break
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		log.Fatal(err)
	}
}
