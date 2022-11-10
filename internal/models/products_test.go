package models

import (
	"strconv"
	"testing"
)

func TestReadCSV(t *testing.T) {
	t.Run("check-ReadCSV-valid", func(t *testing.T) {
		for idx := 1; idx <= 1000; idx++ {
			requ, err := ReadCSV(strconv.Itoa(idx))
			if requ == nil && err == nil {
				t.Error("ReadCSV() - Did not pass the test with the correct data")
			}
		}
	})

	t.Run("check-ReadCSV-invalid", func(t *testing.T) {
		var invalidData []string = []string{"-1", "0", "10000", "10001"}
		for _, val := range invalidData {
			requ, err := ReadCSV(val)
			if requ != nil && err == nil {
				t.Error("ReadCSV() - Didn't pass the test with the wrong data")
			}
		}
	})
}

func TestGetByIdProduct(t *testing.T) {
	t.Run("check-GetByIdProduct-valid", func(t *testing.T) {
		for idx := 1; idx <= 5; idx++ {
			requ, err := GetByIdProduct(strconv.Itoa(idx))
			if requ == nil && err == nil {
				t.Error("GetByIdProduct() - Did not pass the test with the correct data")
			}
		}
	})

	t.Run("check-GetByIdProduct-invalid", func(t *testing.T) {
		var invalidData []string = []string{"-1", "0", "10000", "10001"}
		for _, val := range invalidData {
			requ, err := GetByIdProduct(val)
			if requ != nil && err == nil {
				t.Error("ReadCSV() - Didn't pass the test with the wrong data")
			}
		}
	})
}
