package helper

import (
	"bytes"
	"fmt"
	"github.com/baxromumarov/ucode-sdk/constants"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

func DoRequest(url string, method string, body string) ([]byte, error) {

	request, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", constants.Token)

	client := &http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respByte, nil
}

func FieldParser(line string) (fieldName, fieldSlug, fieldType string) {
	fmt.Println("CollectFieldData: ", line)
	// re := regexp.MustCompile(`"([^"]+)\.([^"]+)"\s+(\w+),`)
	re := regexp.MustCompile(`"([^"]+)\.([^"]+)"\s+([\w\[\]]+)`)

	// Find the matches
	matches := re.FindStringSubmatch(line)
	if len(matches) != 4 {
		fmt.Println("No matches found")
		return
	}

	fieldName = matches[1]
	fieldSlug = matches[2]
	fieldType = matches[3]

	return fieldName, fieldSlug, fieldType
}

func RelationParser(line string) (from, to, labelEn, labelToEn string) {
	re := regexp.MustCompile(`ALTER TABLE "([^"]+)\.([^"]+)" ADD FOREIGN KEY \("([^"]+)"\) REFERENCES "([^"]+)\.([^"]+)" \("([^"]+)"\);`)

	// Find the matches
	matches := re.FindStringSubmatch(line)
	if len(matches) != 7 {
		fmt.Println("No matches found")
		return
	}
	labelEn = matches[1]
	from = matches[2]
	labelToEn = matches[4]
	to = matches[5]

	fmt.Printf("label_en = %s\n", labelEn)
	fmt.Printf("from = %s\n", from)
	fmt.Printf("label_to_en = %s\n", labelToEn)
	fmt.Printf("to = %s\n", to)
	return from, to, labelEn, labelToEn
}

func EnumFieldParser(line string) {
	// Regular expression to match and capture the field name, slug, and type
	re := regexp.MustCompile(`"([^"]+)\.([^"]+)"\s+(\w+)(?:\s+NOT NULL)?`)

	// Find the matches
	matches := re.FindStringSubmatch(line)
	if len(matches) != 4 {
		fmt.Println("No matches found or invalid format")
		return
	}

	fieldName := matches[1]
	fieldSlug := matches[2]
	fieldType := matches[3]

	fmt.Printf("fieldName = %s\n", fieldName)
	fmt.Printf("fieldSlug = %s\n", fieldSlug)
	fmt.Printf("fieldType = %s\n", fieldType)

}
func EnumNameParser(line string) (multiSelect string) {
	re := regexp.MustCompile(`'\s*([^']*)\s*'\s*,?`)

	matches := re.FindStringSubmatch(line)
	if len(matches) != 2 {
		fmt.Println("No matches found or invalid format")
		return multiSelect
	}

	multiSelect = matches[1]

	return multiSelect
}

func EnumParser(line string) (enumName string) {
	// Regular expression to match and capture the enum name
	re := regexp.MustCompile(`CREATE TYPE "([^"]+)" AS ENUM`)

	// Find the matches
	matches := re.FindStringSubmatch(line)
	if len(matches) != 2 {
		fmt.Println("No matches found or invalid format")
		return
	}

	enumName = matches[1]

	return enumName
}

func TableParser(line string) (tableName, tableSlug, menuName string) {
	re := regexp.MustCompile(`CREATE TABLE "(?:(\w+)\.)?(?:(\w+)\.)?([^\.]+)\.([^"]+)" \(`)

	matches := re.FindStringSubmatch(line)
	if len(matches) != 5 {
		log.Fatal("No matches found or invalid format")
	}

	menuName = matches[1]
	tableName = matches[2]
	if tableName == "" {
		tableName = matches[3]
		tableSlug = matches[4]
	} else {
		tableName = matches[3]
		tableSlug = matches[4]
	}

	fmt.Printf("menuName = %s\n", menuName)
	fmt.Printf("tableName = %s\n", tableName)
	fmt.Printf("tableSlug = %s\n", tableSlug)
	fmt.Println("----------")
	return tableName, tableSlug, menuName
}
