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

func (ls *LocalDB) Close() error {
	clear(ls.ZipcodesTable)
	clear(ls.CountyTable)
	clear(ls.StateCountyTable)
	clear(ls.StateCityTable)
	clear(ls.CountyCityTable)
	return nil
}

func (ls *LocalDB) InitLocalStorage() error {
	records, err := getInitialData()
	if err != nil {
		return err
	}

	if err := ls.createZipcodeTable(records); err != nil {
		return err
	}

	if err := ls.createCountyTable(records); err != nil {
		return err
	}

	if err := ls.createStateCountyTable(records); err != nil {
		return err
	}

	if err := ls.createStateCityTable(records); err != nil {
		return err
	}

	if err := ls.createCountyCityTable(records); err != nil {
		return err
	}

	return nil
}

func (ls *LocalDB) createZipcodeTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		ls.ZipcodesTable[r.Zipcode] = map[string]string{
			"StateFIPS":  r.StateFIPS,
			"StateAbbr":  r.StateAbbr,
			"StateName":  r.State,
			"CountyName": r.County,
			"CityName":   r.City,
		}
	}

	return nil
}

func (ls *LocalDB) createCountyTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		if _, exist := ls.CountyTable[r.County]; !exist {
			ls.CountyTable[r.County] = map[string]map[string]string{
				r.Zipcode: {
					"StateFIPS": r.StateFIPS,
					"StateAbbr": r.StateAbbr,
					"StateName": r.State,
					"CityName":  r.City,
				},
			}
		} else {
			ls.CountyTable[r.County][r.Zipcode] = map[string]string{
				"StateFIPS": r.StateFIPS,
				"StateAbbr": r.StateAbbr,
				"StateName": r.State,
				"CityName":  r.City,
			}
		}
	}
	return nil
}

func (ls *LocalDB) createStateCountyTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		if _, exist := ls.StateCountyTable[r.State]; !exist {
			ls.StateCountyTable[r.State] = map[string]map[string]map[string]string{
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

		if _, exist := ls.StateCountyTable[r.State][r.County]; !exist {
			ls.StateCountyTable[r.State][r.County] = map[string]map[string]string{
				r.Zipcode: {
					"StateFIPS": r.StateFIPS,
					"StateAbbr": r.StateAbbr,
					"CityName":  r.City,
				},
			}
			continue
		}

		ls.StateCountyTable[r.State][r.County][r.Zipcode] = map[string]string{
			"StateFIPS": r.StateFIPS,
			"StateAbbr": r.StateAbbr,
			"CityName":  r.City,
		}
	}

	return nil
}

func (ls *LocalDB) createStateCityTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		if _, exist := ls.StateCityTable[r.State]; !exist {
			ls.StateCountyTable[r.State] = map[string]map[string]map[string]string{
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

		if _, exist := ls.StateCityTable[r.State][r.City]; !exist {
			ls.StateCityTable[r.State][r.City] = map[string]map[string]string{
				r.Zipcode: {
					"StateFIPS":  r.StateFIPS,
					"StateAbbr":  r.StateAbbr,
					"CountyName": r.County,
				},
			}
			continue
		}

		ls.StateCityTable[r.State][r.City][r.Zipcode] = map[string]string{
			"StateFIPS":  r.StateFIPS,
			"StateAbbr":  r.StateAbbr,
			"CountyName": r.County,
		}
	}

	return nil
}

func (ls *LocalDB) createCountyCityTable(records []models.Zipcode) error {
	for _, r := range records[1:] {
		if _, exist := ls.CountyCityTable[r.County]; !exist {
			ls.CountyCityTable[r.County] = map[string]map[string]map[string]string{
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

		if _, exist := ls.CountyCityTable[r.County][r.City]; !exist {
			ls.CountyCityTable[r.County][r.City] = map[string]map[string]string{
				r.Zipcode: {
					"StateFIPS": r.StateFIPS,
					"StateAbbr": r.StateAbbr,
					"StateName": r.State,
				},
			}
			continue
		}

		ls.CountyCityTable[r.County][r.City][r.Zipcode] = map[string]string{
			"StateFIPS": r.StateFIPS,
			"StateAbbr": r.StateAbbr,
			"StateName": r.State,
		}
	}

	return nil
}
