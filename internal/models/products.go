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

func ReadAllCSV(path string) ([]byte, error) {
	// if you run this project not use docker paste this path in func os.Open("./internal/models/csv/yandex.csv")
	file, err := os.Open(path)
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

// All products ...
func GetAllProducts(totalQuery string) ([]byte, error) {
	// var product []ProductData "./internal/models/csv/yandex.csv"
	request, err := ReadAllCSV(os.Getenv("PATH_TO_FILE"))
	if err != nil {
		return nil, err
	}

	/*
		Add the ability to get a limited number of items
	*/

	return request, nil
}

// Get prodect by id ...
func GetByIdProduct(id string) ([]byte, error) {
	var product []ProductData

	request, err := ReadAllCSV(os.Getenv("PATH_TO_FILE"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(request, &product)
	if err != nil {
		return nil, fmt.Errorf("Err unmarshal product to struct - %s", err)
	}

	for _, val := range product {
		if val.ID == id {
			if out, err := json.Marshal(val); err != nil {
				return nil, fmt.Errorf("Err marshal struct - %s", err)
			} else {
				return out, nil
			}
		}
	}
	return nil, nil
}
