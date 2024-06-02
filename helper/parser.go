package helper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/baxromumarov/ucode-sdk/constants"
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

	re := regexp.MustCompile(`"([^"]+)\.([^"]+)"\s+([\w\[\]]+)`)

	// Find the matches
	matches := re.FindStringSubmatch(line)
	if len(matches) != 4 {
		log.Fatal("No matches found(FieldParser)")
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
		log.Fatal("No matches found(RelationParser)")
		return
	}

	labelToEn = matches[1]
	to = matches[2]
	labelEn = matches[4]
	from = matches[5]
// fmt.Println(from, to, labelEn, labelToEn)
	return from, to, labelEn, labelToEn
}

func EnumNameParser(line string) (multiSelect string) {
	re := regexp.MustCompile(`'\s*([^']*)\s*'\s*,?`)

	matches := re.FindStringSubmatch(line)
	if len(matches) != 2 {
		log.Fatal("No matches found or invalid format(EnumNameParser)")
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
		log.Fatal("No matches found or invalid format(EnumParser)")
		return
	}

	enumName = matches[1]

	return enumName
}

func TableParser(line string) (tableName, tableSlug, menuName string) {
	// Trim the prefix and suffix
	trimmed := strings.TrimPrefix(line, `CREATE TABLE "`)
	trimmed = strings.TrimSuffix(trimmed, `" (`)

	// Split the components
	parts := strings.Split(trimmed, ".")

	if len(parts) == 3 {
		menuName = parts[0]
		tableName = parts[1]
		tableSlug = parts[2]
	} else if len(parts) == 2 {
		tableName = parts[0]
		tableSlug = parts[1]
	} else {
		fmt.Println("Invalid format")
		return
	}

	fmt.Printf("menuName = %s\n", menuName)
	fmt.Printf("tableName = %s\n", tableName)
	fmt.Printf("tableSlug = %s\n", tableSlug)
	return tableName, tableSlug, menuName
}
