package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/baxromumarov/ucode-sdk/create_data"
	"github.com/baxromumarov/ucode-sdk/helper"
	"github.com/baxromumarov/ucode-sdk/models"
)

var (
	moduleID = ""
	// tableID  = ""
)

func main() {

	filePath := "./sdk-ucode.sql"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		table = models.Table{
			FieldType: make(map[string]string),
		}
		allTables = map[string]string{}
	)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, ");") {
			create_data.CreateFields(table, allTables)
			// time.Sleep(time.Second)
			table = models.Table{
				FieldType: map[string]string{},
			}
		}

		if strings.Contains(line, "ALTER") {
			break
		}

		if strings.Contains(line, "CREATE TABLE") {
			tableName := getTableName(line)

			//! create a new table
			tableID := create_data.CreateTable(tableName, moduleID)
			fmt.Println("TABLEID, TABLE NAME: ", tableName, tableID)

			table.ID = tableID
			table.Name = tableName
			allTables[tableName] = tableID
			continue
		}
		splitedRow := strings.Split(line, " ")

		if len(splitedRow) == 1 {
			continue
		}

		finalRow := helper.CleaningFields(splitedRow)

		table.FieldType[finalRow[0]] = finalRow[1]

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		log.Fatal(err)
	}
}

func getTableName(line string) string {
	line = strings.ReplaceAll(line, `CREATE TABLE`, "")
	line = strings.ReplaceAll(line, ` `, "")
	line = strings.ReplaceAll(line, `"`, "")
	line = strings.ReplaceAll(line, `(`, "")

	row := strings.Split(line, ".")

	if len(row) == 1 {
		return row[0]
	} else {
		// ! Create module
		moduleID = create_data.CreateModule(row[0])
		return row[1]
	}

}
