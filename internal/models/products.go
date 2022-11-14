package models

import (
	conn "MarketPlaceBackEnd/internal/database"
	"database/sql"
	"strconv"
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

// WORK: Checking the resulting value
func checkParamAndQuery(total string) (bool, int) {
	num, _ := strconv.Atoi(total)
	if num <= 0 || num > 10000 {
		return false, 0
	}
	return true, num
}

// WORK: Needed to read a csv file of goods (rename to findall)
func FindAll(total string) ([]ProductData, error) {
	var data []ProductData
	db, err := conn.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	status, num := checkParamAndQuery(total)
	if status {
		if data, err := findToTotal(db, data, num); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	} else {
		// The use of * is not appropriate here because the fields may change in future migrations
		rows, err := db.Query(`SELECT id, market_name, prod_manufacturer, prod_name, art, price, link FROM products`)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var p ProductData

			err := rows.Scan(&p.ID, &p.MarketName, &p.ProdManufacturer, &p.ProdName, &p.Art, &p.Price, &p.Link)
			if err != nil {
				return nil, err
			}
			data = append(data, p)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}
		return data, nil
	}
}

// WORK: Returns one item by the specified identifier
func FindById(id string) ([]ProductData, error) {
	var data []ProductData
	db, err := conn.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if status, _ := checkParamAndQuery(id); !status {
		return nil, nil
	}

	// The use of * is not appropriate here because the fields may change in future migrations
	rows, err := db.Query(`SELECT id, market_name, prod_manufacturer, prod_name, art, price, link FROM products WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p ProductData

		err := rows.Scan(&p.ID, &p.MarketName, &p.ProdManufacturer, &p.ProdName, &p.Art, &p.Price, &p.Link)
		if err != nil {
			return nil, err
		}

		data = append(data, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

// WORK: Find products to total
func findToTotal(db *sql.DB, data []ProductData, num int) ([]ProductData, error) {
	rows, err := db.Query(`SELECT id, market_name, prod_manufacturer, prod_name, art, price, link FROM products LIMIT $1`, num)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p ProductData

		err := rows.Scan(&p.ID, &p.MarketName, &p.ProdManufacturer, &p.ProdName, &p.Art, &p.Price, &p.Link)
		if err != nil {
			return nil, err
		}

		data = append(data, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
