package database

import (
	"github.com/przant/zipcodes-api/models"
	"github.com/przant/zipcodes-api/utils"
)

type LocalDB struct {
	StateTable map[string]map[string]map[string]string
}

func (ldb *LocalDB) Close() error {
	clear(ldb.StateTable)
	return nil
}

func (ldb *LocalDB) InitLocalStorage() error {
	records, err := utils.FetchData()
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
