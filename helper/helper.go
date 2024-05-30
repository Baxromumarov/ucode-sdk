package helper

import (
	"bytes"
	"fmt"
	"github.com/baxromumarov/ucode-sdk/constants"
	"io"
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
