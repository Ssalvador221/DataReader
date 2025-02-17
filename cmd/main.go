package main

import (
	"context"
	"fmt"
	"go-data-read/internal/db/store"
	"go-data-read/internal/handler"
	"go-data-read/internal/reader"
	"go-data-read/pkg"
)

func main() {

	ctx := context.Background()

	db, err := pkg.NewDB() // db connection
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	csv, err := reader.ReadCsv("data/conjunto-dados.csv")
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	h := handler.NewHandler(db, csv)

	for _, record := range csv {
		_, err := h.AddNewRecord(ctx, store.CreateRecordParams{
			Organizacao:   record.Organizacao,
			Nome:          record.Nome,
			Descricao:     pkg.NewNullString(record.Descricao),
			Tags:          record.Tags,
			Qtdrecursos:   pkg.NewNullInt32(int32(record.QuantidadeRecursos)),
			Qtdreusos:     pkg.NewNullInt32(int32(record.QuantidadeReusos)),
			Qtddownloads:  pkg.NewNullInt32(int32(record.QuantidadeDownloads)),
			Qtdseguidores: pkg.NewNullInt32(int32(record.QuantidadeSeguidores)),
		})
		if err != nil {
			fmt.Println("Error adding new record:", err)
			continue
		}
	}

	fmt.Println("All records added successfully.")

}
