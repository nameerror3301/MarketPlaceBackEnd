package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

/*
	IF -> You want to run the service without using the containerization technology,
			you should pass this path to
			the desired function - "./internal/models/csv/yandex.csv"

	ELSE -> os.Getenv("PATH_TO_FILE")
*/

// Structure for generating goods json before passing it to the user
type ProductData struct {
	ID               string `json:"id,omitempty"`
	MarketName       string `json:"market_name,omitempty"`
	ProdManufacturer string `json:"prod_manufacturer,omitempty"`
	ProdName         string `json:"prod_name,omitempty"`
	Art              string `json:"art,omitempty"`
	Price            string `json:"price,omitempty"`
	Link             string `json:"link,omitempty"`
}

// Constants for determining the minimum and maximum values for the parameters
const (
	min = 0
	max = 10000
)

// WORK: Checking the resulting value
func checkParamAndQuery(total string) (bool, int) {
	num, _ := strconv.Atoi(total)
	if num <= min || num > max {
		return false, 0
	}
	return true, num
}

// WORK: Needed to read a csv file of goods
func ReadCSV(total string) ([]ProductData, error) {
	file, err := os.Open(os.Getenv("PATH_TO_FILE"))

	if err != nil {
		return nil, fmt.Errorf("Err open products file - %s", err)
	}
	defer file.Close()

	var product []ProductData

	if total == "" {
		requ := readAllFile(product, file)
		return requ, nil
	}

	valid, num := checkParamAndQuery(total)
	if !valid {
		return nil, nil
	}

	requ := readToTotal(product, file, num)

	return requ, nil
}

// WORK: Returns one item by the specified identifier
func GetByIdProduct(id string) ([]ProductData, error) {
	request, err := ReadCSV("")
	if err != nil {
		return nil, err
	}

	for _, val := range request {
		if val.ID == id {
			return []ProductData{val}, nil
		}
	}
	return nil, nil
}

// WORK: Returns all items if the value total was not received
func readAllFile(product []ProductData, file *os.File) []ProductData {
	fileReader := csv.NewReader(file)
	for {
		line, err := fileReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Errorf("Err read product file - %s", err)
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

	return product
}

// WORK: Returns all items up to the value specified on the total
func readToTotal(product []ProductData, file *os.File, total int) []ProductData {
	fileReader := csv.NewReader(file)
	for idx := 0; idx != total; {
		line, err := fileReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Errorf("Err read product file - %s", err)
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
		idx++
	}
	return product
}
