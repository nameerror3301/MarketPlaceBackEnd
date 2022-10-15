package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ProductData struct {
	ID               string `json:"id,omitempty"`
	MarketName       string `json:"market_name,omitempty"`
	ProdManufacturer string `json:"prod_manufacturer,omitempty"`
	ProdName         string `json:"prod_name,omitempty"`
	Art              string `json:"art,omitempty"`
	Price            string `json:"price,omitempty"`
	Link             string `json:"link,omitempty"`
}

func (p *ProductData) ReadCSV(path string) ([]byte, error) {
	file, err := os.Open("./internal/models/csv/yandex.csv")
	if err != nil {
		return nil, fmt.Errorf("Err open csv - %s", err)
	}
	defer file.Close()

	var product []ProductData

	fileReader := csv.NewReader(file)
	for {
		line, err := fileReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Err read csv - %s", err)
		}

		product = append(product, ProductData{
			ID:               line[0],
			MarketName:       line[1],
			ProdManufacturer: line[2],
			ProdName:         line[3],
			Art:              line[4],
			Price:            line[5],
			Link:             line[6],
		})

	}
	out, err := json.Marshal(product)
	if err != nil {
		return nil, fmt.Errorf("Err marshal struct - %s", err)
	}

	return out, nil
}

/*
	For future scaling space under the database
*/
