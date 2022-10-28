package models

import "testing"

func TestReadAllCSV(t *testing.T) {
	t.Run("check-ReadAllCSV-valid", func(t *testing.T) {
		req, err := ReadAllCSV("./csv/yandex.csv")
		if req == nil && err != nil {
			t.Error("ReadAllCSV() - Did not pass the test with the correct data")
		}
	})

	t.Run("check-ReadAllCSV-invalid", func(t *testing.T) {
		req, err := ReadAllCSV("")
		if req != nil && err == nil {
			t.Error("ReadAllCSV() - Didn't pass the test with the wrong data")
		}
	})
}
