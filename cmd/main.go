package main

import (
	reader "github.com/baxromumarov/ucode-sdk/erd_reader"
)

func main() {
	// fmt.Println(helper.RelationParser(`ALTER TABLE "Мошина.car" ADD FOREIGN KEY ("guid") REFERENCES "Клиент.customer" ("car_id");`))
	reader.Reader()
}
