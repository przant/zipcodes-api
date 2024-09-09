package database

import (
	"encoding/csv"
	"io"
	"net/http"

	"github.com/przant/zipcodes-api/models"
)

const (
	source = `https://raw.githubusercontent.com/scpike/us-state-county-zip/master/geo-data.csv`
)

type LocalDB struct {
	StateTable map[string]map[string]map[string]string
}

func getInitialData() ([]models.Zipcode, error) {
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

func (ldb *LocalDB) Close() error {
	clear(ldb.StateTable)
	return nil
}

func (ldb *LocalDB) InitLocalStorage() error {
	records, err := getInitialData()
	if err != nil {
		return err
	}

	if err := ldb.createStateTable(records); err != nil {
		return err
	}

	return nil
}

func (ldb *LocalDB) createStateTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		if _, exist := ldb.StateTable[r.State]; !exist {
			ldb.StateTable[r.State] = map[string]map[string]string{
				r.Zipcode: {
					"StateFIPS":  r.StateFIPS,
					"StateAbbr":  r.StateAbbr,
					"CountyName": r.County,
					"CityName":   r.City,
				},
			}
		} else {
			ldb.StateTable[r.State][r.Zipcode] = map[string]string{
				"StateFIPS":  r.StateFIPS,
				"StateAbbr":  r.StateAbbr,
				"CountyName": r.County,
				"CityName":   r.City,
			}
		}
	}

	return nil
}
