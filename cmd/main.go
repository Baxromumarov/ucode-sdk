package main

import (
	"fmt"
	"log"
	"os"

	"github.com/baxromumarov/ucode-sdk/constants"
	reader "github.com/baxromumarov/ucode-sdk/erd_reader"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("Getwd: ", err)
	}

	file, err := os.Open(path + "/erd.sql")
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer file.Close()

	err = reader.Reader(file, constants.Token)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
