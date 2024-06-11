package main

import (
	"log"
	"os"

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

	reader.Reader(file, "")
}
