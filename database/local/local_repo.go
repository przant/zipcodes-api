package database

import (
	"fmt"

	"github.com/przant/zipcodes-api/models"
)

type LocalDBRepo struct {
	db *LocalDB
}

func NewLocalDBRepo() (*LocalDBRepo, error) {
	lr := &LocalDBRepo{
		db: &LocalDB{
			ZipcodesTable:    make(map[string]map[string]string),
			CountyTable:      make(map[string]map[string]map[string]string),
			StateCountyTable: make(map[string]map[string]map[string]map[string]string),
			StateCityTable:   make(map[string]map[string]map[string]map[string]string),
			CountyCityTable:  make(map[string]map[string]map[string]map[string]string),
		},
	}

	if err := lr.db.InitLocalStorage(); err != nil {
		return nil, err
	}

	return lr, nil
}

func (lr *LocalDBRepo) Close() error {
	return lr.db.Close()
}

func (lr *LocalDBRepo) FetchByZipcode(zipcode string) (*models.Zipcode, error) {
	if _, exist := lr.db.ZipcodesTable[zipcode]; !exist {
		return nil, fmt.Errorf("the zipcode %q does not exist in the database", zipcode)
	}

	return &models.Zipcode{
		Zipcode:   zipcode,
		StateFIPS: lr.db.ZipcodesTable[zipcode]["StateFIPS"],
		StateAbbr: lr.db.ZipcodesTable[zipcode]["StateAbbr"],
		State:     lr.db.ZipcodesTable[zipcode]["StateName"],
		County:    lr.db.ZipcodesTable[zipcode]["CountyName"],
		City:      lr.db.ZipcodesTable[zipcode]["CityName"],
	}, nil
}

func (lr *LocalDBRepo) FetchByCounty(county string) ([]models.Zipcode, error) {
	if _, exist := lr.db.CountyTable[county]; !exist {
		return nil, fmt.Errorf("the county %q does not exist in the database", county)
	}

	rs := make([]models.Zipcode, 0)
	for zc, attrs := range lr.db.CountyTable[county] {
		r := models.Zipcode{
			Zipcode:   zc,
			County:    county,
			StateFIPS: attrs["StateFIPS"],
			StateAbbr: attrs["StateAbbr"],
			State:     attrs["StateName"],
			City:      attrs["CityName"],
		}

		rs = append(rs, r)
	}

	return rs, nil
}

func (lr *LocalDBRepo) FetchByStateCounty(state, county string) ([]models.Zipcode, error) {
	if _, exist := lr.db.StateCountyTable[state][county]; !exist {
		return nil, fmt.Errorf("the state-county pair %q does not exist in the database", state+"-"+county)
	}

	rs := make([]models.Zipcode, 0)

	for zc, attrs := range lr.db.StateCountyTable[state][county] {
		r := models.Zipcode{
			Zipcode:   zc,
			State:     state,
			County:    county,
			StateFIPS: attrs["StateFIPS"],
			StateAbbr: attrs["StateAbbr"],
			City:      attrs["CityName"],
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (lr *LocalDBRepo) FetchByStateCity(state, city string) ([]models.Zipcode, error) {
	if _, exist := lr.db.StateCityTable[state][city]; !exist {
		return nil, fmt.Errorf("the state-city pair %q does not exist in the database", state+"-"+city)
	}

	rs := make([]models.Zipcode, 0)

	for zc, attrs := range lr.db.StateCountyTable[state][city] {
		r := models.Zipcode{
			Zipcode:   zc,
			State:     state,
			City:      city,
			StateFIPS: attrs["StateFIPS"],
			StateAbbr: attrs["StateAbbr"],
			County:    attrs["CityName"],
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (lr *LocalDBRepo) FecthByCountyCity(county, city string) ([]models.Zipcode, error) {
	if _, exist := lr.db.StateCityTable[county][city]; !exist {
		return nil, fmt.Errorf("the state-city pair %q does not exist in the database", county+"-"+city)
	}

	rs := make([]models.Zipcode, 0)

	for zc, attrs := range lr.db.StateCountyTable[county][city] {
		r := models.Zipcode{
			Zipcode:   zc,
			City:      city,
			County:    county,
			StateFIPS: attrs["StateFIPS"],
			StateAbbr: attrs["StateAbbr"],
			State:     attrs["StateName"],
		}
		rs = append(rs, r)
	}
	return rs, nil
}
