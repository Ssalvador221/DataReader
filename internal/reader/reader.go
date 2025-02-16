package reader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CamposCsv struct {
	Organizacao          string
	Nome                 string
	Descricao            string
	Tags                 string
	QuantidadeRecursos   int
	QuantidadeReusos     int
	QuantidadeDownloads  int
	QuantidadeSeguidores int
}

func ReadCsv(filePath string) ([]CamposCsv, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Unable to parse file as CSV for "+filePath, err)
	}

	var camposCsv []CamposCsv

	for i, record := range records {
		if i == 0 {
			continue
		}

		quantidadeRecursos, _ := strconv.Atoi(record[4])
		quantidadeReusos, _ := strconv.Atoi(record[5])
		quantidadeDownloads, _ := strconv.Atoi(record[6])
		quantidadeSeguidores, _ := strconv.Atoi(record[7])

		camposCsv = append(camposCsv, CamposCsv{
			Organizacao:          record[0],
			Nome:                 record[1],
			Descricao:            record[2],
			Tags:                 record[3],
			QuantidadeRecursos:   quantidadeRecursos,
			QuantidadeReusos:     quantidadeReusos,
			QuantidadeDownloads:  quantidadeDownloads,
			QuantidadeSeguidores: quantidadeSeguidores,
		})
	}

	return camposCsv, nil
}
