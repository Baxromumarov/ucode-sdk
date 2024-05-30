package main

import (
	"bufio"
	"fmt"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
	"github.com/baxromumarov/ucode-sdk/structure"
	"github.com/fatih/color"
	"log"
	"os"
	"strings"
)

var (
	menuID = ""
	// tableID  = ""
)

func main() {
	defer func() {
		color.Green("Congratulations, BAXROM!!! You did id. GO and drink  😄 ")

	}()
	filePath := "./erd.sql"

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
			continue
		}

		if strings.Contains(line, "ALTER TABLE ") {
			fmt.Println("<<<<<<<<< relation creating ")
			structure.CreateRelation(line)
			continue
		}

		splitLine := strings.Split(line, " ")

		if strings.Contains(line, "CREATE TABLE") {

			data := strings.Split(splitLine[2], ".")
			if len(data) == 3 {
				menuName, tableName, tableSlug := data[0], data[1], data[2]
				menuName = menuName[1:]
				tableSlug = tableSlug[:len(tableSlug)-1]

				// fmt.Println(menuName)
				// fmt.Println(tableName)
				// fmt.Println(tableSlug)

				menuID = structure.CreateMenu(menuName)
				tableID := structure.CreateTable(tableSlug, menuID, tableName)
				fmt.Println("IF CASE TABLE ID ", tableID)

				// ! table data collecting
				table.ID = tableID
				table.Slug = tableSlug
				table.Name = tableName

				// just create table with menu id
				fmt.Println(">>> ", splitLine[2])
			} else {
				// just create a table with exist menu id
				tableName, tableSlug := data[0], data[1]
				tableName = tableName[1:]
				tableSlug = tableSlug[:len(tableSlug)-1]

				// fmt.Println(tableName)
				// fmt.Println(tableSlug)

				tableID := structure.CreateTable(tableSlug, menuID, tableName)
				fmt.Println("ELSE CASE TABLE ID ", tableID)
				// ! table data collecting
				table.ID = tableID
				table.Slug = tableSlug
				table.Name = tableName
			}
			tables[table.Name] = table.ID
			continue
		}
		// fmt.Println(line)
		if strings.Contains(line, ".") {
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
