package helper

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/baxromumarov/ucode-sdk/constants"
)

func CleaningFields(input []string) []string {
	var result []string

	for _, str := range input {
		if len(str) > 4 {

			str = strings.ReplaceAll(str, `"`, "")
			str = strings.ReplaceAll(str, `,`, "")

			result = append(result, str)
		}
	}

	return result
}

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
