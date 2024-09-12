package utils

import (
	"encoding/csv"
	"io"
	"net/http"

	"github.com/przant/zipcodes-api/models"
)

const (
	source = `https://raw.githubusercontent.com/scpike/us-state-county-zip/master/geo-data.csv`
)

func FetchData() ([]models.Zipcode, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	csvReader := csv.NewReader(resp.Body)
	csvReader.LazyQuotes = true
	records := make([]models.Zipcode, 0)

	for {
		data, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		record := models.Zipcode{
			StateFIPS: data[0],
			State:     data[1],
			StateAbbr: data[2],
			Zipcode:   data[3],
			County:    data[4],
			City:      data[5],
		}
		records = append(records, record)
	}

	return records, nil
}
