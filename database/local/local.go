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
	ZipcodesTable    map[string]map[string]string
	CountyTable      map[string]map[string]map[string]string
	StateCountyTable map[string]map[string]map[string]map[string]string
	StateCityTable   map[string]map[string]map[string]map[string]string
	CountyCityTable  map[string]map[string]map[string]map[string]string
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
	clear(ldb.ZipcodesTable)
	clear(ldb.CountyTable)
	clear(ldb.StateCountyTable)
	clear(ldb.StateCityTable)
	clear(ldb.CountyCityTable)
	return nil
}

func (ldb *LocalDB) InitLocalStorage() error {
	records, err := getInitialData()
	if err != nil {
		return err
	}

	if err := ldb.createZipcodeTable(records); err != nil {
		return err
	}

	if err := ldb.createCountyTable(records); err != nil {
		return err
	}

	if err := ldb.createStateCountyTable(records); err != nil {
		return err
	}

	if err := ldb.createStateCityTable(records); err != nil {
		return err
	}

	if err := ldb.createCountyCityTable(records); err != nil {
		return err
	}

	return nil
}

func (ldb *LocalDB) createZipcodeTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		ldb.ZipcodesTable[r.Zipcode] = map[string]string{
			"StateFIPS":  r.StateFIPS,
			"StateAbbr":  r.StateAbbr,
			"StateName":  r.State,
			"CountyName": r.County,
			"CityName":   r.City,
		}
	}

	return nil
}

func (ldb *LocalDB) createCountyTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		if _, exist := ldb.CountyTable[r.County]; !exist {
			ldb.CountyTable[r.County] = map[string]map[string]string{
				r.Zipcode: {
					"StateFIPS": r.StateFIPS,
					"StateAbbr": r.StateAbbr,
					"StateName": r.State,
					"CityName":  r.City,
				},
			}
		} else {
			ldb.CountyTable[r.County][r.Zipcode] = map[string]string{
				"StateFIPS": r.StateFIPS,
				"StateAbbr": r.StateAbbr,
				"StateName": r.State,
				"CityName":  r.City,
			}
		}
	}
	return nil
}

func (ldb *LocalDB) createStateCountyTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		if _, exist := ldb.StateCountyTable[r.State]; !exist {
			ldb.StateCountyTable[r.State] = map[string]map[string]map[string]string{
				r.County: {
					r.Zipcode: {
						"StateFIPS": r.StateFIPS,
						"StateAbbr": r.StateAbbr,
						"CityName":  r.City,
					},
				},
			}
			continue
		}

		if _, exist := ldb.StateCountyTable[r.State][r.County]; !exist {
			ldb.StateCountyTable[r.State][r.County] = map[string]map[string]string{
				r.Zipcode: {
					"StateFIPS": r.StateFIPS,
					"StateAbbr": r.StateAbbr,
					"CityName":  r.City,
				},
			}
			continue
		}

		ldb.StateCountyTable[r.State][r.County][r.Zipcode] = map[string]string{
			"StateFIPS": r.StateFIPS,
			"StateAbbr": r.StateAbbr,
			"CityName":  r.City,
		}
	}

	return nil
}

func (ldb *LocalDB) createStateCityTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		if _, exist := ldb.StateCityTable[r.State]; !exist {
			ldb.StateCountyTable[r.State] = map[string]map[string]map[string]string{
				r.City: {
					r.Zipcode: {
						"StateFIPS":  r.StateFIPS,
						"StateAbbr":  r.StateAbbr,
						"CountyName": r.County,
					},
				},
			}
			continue
		}

		if _, exist := ldb.StateCityTable[r.State][r.City]; !exist {
			ldb.StateCityTable[r.State][r.City] = map[string]map[string]string{
				r.Zipcode: {
					"StateFIPS":  r.StateFIPS,
					"StateAbbr":  r.StateAbbr,
					"CountyName": r.County,
				},
			}
			continue
		}

		ldb.StateCityTable[r.State][r.City][r.Zipcode] = map[string]string{
			"StateFIPS":  r.StateFIPS,
			"StateAbbr":  r.StateAbbr,
			"CountyName": r.County,
		}
	}

	return nil
}

func (ldb *LocalDB) createCountyCityTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		if _, exist := ldb.CountyCityTable[r.County]; !exist {
			ldb.CountyCityTable[r.County] = map[string]map[string]map[string]string{
				r.City: {
					r.Zipcode: {
						"StateFIPS": r.StateFIPS,
						"StateAbbr": r.StateAbbr,
						"StateName": r.State,
					},
				},
			}
			continue
		}

		if _, exist := ldb.CountyCityTable[r.County][r.City]; !exist {
			ldb.CountyCityTable[r.County][r.City] = map[string]map[string]string{
				r.Zipcode: {
					"StateFIPS": r.StateFIPS,
					"StateAbbr": r.StateAbbr,
					"StateName": r.State,
				},
			}
			continue
		}

		ldb.CountyCityTable[r.County][r.City][r.Zipcode] = map[string]string{
			"StateFIPS": r.StateFIPS,
			"StateAbbr": r.StateAbbr,
			"StateName": r.State,
		}
	}

	return nil
}
