package main

import (
	"fmt"
	"go-data-read/internal/reader"
)

func main() {
	data, err := reader.ReadCsv("data/conjunto-dados.csv")
	if err != nil {
		fmt.Println(err)
	}

	for _, record := range data {
		fmt.Println(record.Nome)
	}
}
