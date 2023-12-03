package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/baxromumarov/ucode-sdk/helper"
)

var (
	moduleID = ""
)

func main() {

	filePath := "/Users/macbookpro/go/src/github.com/baxromumarov/ucode-sdk/sdk-ucode.sql"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var fieldMap = map[string]string{}

		if strings.Contains(line, "CREATE TABLE") {
			tableName := getTableName(line)
			fmt.Println("TABLE NAME:   ", tableName)
			continue
		}

		splitedRow := strings.Split(line, " ")
		if len(splitedRow) == 1 {
			continue
		}
		finalRow := helper.CleaningRow(splitedRow)
		fmt.Println("final row: ", finalRow, fieldMap)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
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
		// Create module
		moduleID = CreateModule(row[0])
		return row[1]
	}

}

// CREATE TABLE "sdk_table" (

// func getTableName(line string) string {
// 	line = strings.ReplaceAll(line, `CREATE TABLE`, "")
// 	line = strings.ReplaceAll(line, ` `, "")
// 	line = strings.ReplaceAll(line, `"`, "")
// 	line = strings.ReplaceAll(line, `(`, "")
// 	row := strings.Split(line, ".")
// 	if len(row) == 1 {
// 		return row[0]
// 	} else {
// 		// Create module
// 		moduleID = createModule(row[0])
// 		return row[1]
// 	}

// }
// func cleaningRow(input []string) []string {
// 	var result []string

// 	for _, str := range input {
// 		if len(str) > 4 {
// 			// if strings.Contains(str, ".") && !strings.Contains(str, "id") {
// 			// 	splitedStr := strings.Split(str, ".")
// 			// 	moduleName := splitedStr[0]
// 			// 	moduleName = strings.ReplaceAll(moduleName, `"`, "")

// 			// 	tableName := splitedStr[1]
// 			// 	str = strings.ReplaceAll(tableName, `"`, "")

// 			// 	fmt.Println(moduleName, tableName)

// 			// } else {
// 			str = strings.ReplaceAll(str, `"`, "")
// 			str = strings.ReplaceAll(str, `,`, "")

// 			// }
// 			result = append(result, str)
// 		}
// 	}

// 	return result
// }
// func DoRequest(url string, method string, body string) ([]byte, error) {

// 	request, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	request.Header.Set("Content-Type", "application/json")
// 	request.Header.Set("Authorization", token)

// 	client := &http.Client{
// 		Timeout: time.Duration(10 * time.Second),
// 	}
// 	resp, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	respByte, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return respByte, nil
// }

type CreateResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		ID string `json:"id"`
	} `json:"data"`
	CustomMessage string `json:"custom_message"`
}
