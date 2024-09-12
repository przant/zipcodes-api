package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/przant/zipcodes-api/models"
)

type LocalDBRepo struct {
	db *LocalDB
}

func NewLocalDBRepo() (*LocalDBRepo, error) {
	lr := &LocalDBRepo{
		db: &LocalDB{
			StateTable: make(map[string]map[string]map[string]string),
		},
	}

	if err := lr.db.InitLocalStorage(); err != nil {
		return nil, err
	}

	return lr, nil
}

func (lr *LocalDBRepo) Close() {
	if err := lr.db.Close(); err != nil {
		log.Printf("while closing the LocalDBRepo: %s", err)
	}
}

func (lr *LocalDBRepo) FetchByZipcode(zipcode string) (*models.Zipcode, error) {
	for st, zips := range lr.db.StateTable {
		if _, exist := zips[zipcode]; exist {
			return &models.Zipcode{
				State:     st,
				Zipcode:   zipcode,
				StateFIPS: zips[zipcode]["StateFIPS"],
				StateAbbr: zips[zipcode]["StateAbbr"],
				County:    zips[zipcode]["CountyName"],
				City:      zips[zipcode]["CityName"],
			}, nil
		}
	}

	return nil, fmt.Errorf("the zipcode %q does not exist in the database", zipcode)
}

func (lr *LocalDBRepo) FetchByCounty(county string) ([]models.Zipcode, error) {
	rs := make([]models.Zipcode, 0)

	for st, zips := range lr.db.StateTable {
		for zipcode, attrs := range zips {
			if strings.EqualFold(attrs["CountyName"], county) {
				r := models.Zipcode{
					State:     st,
					Zipcode:   zipcode,
					County:    county,
					StateFIPS: attrs["StateFIPS"],
					City:      attrs["CityName"],
				}

				rs = append(rs, r)
			}
		}
	}

	if len(rs) == 0 {
		return nil, fmt.Errorf("the county %q does not exist in the database", county)
	}

	return rs, nil
}

func (lr *LocalDBRepo) FetchByStateCounty(state, county string) ([]models.Zipcode, error) {
	rs := make([]models.Zipcode, 0)

	for st, zips := range lr.db.StateTable {
		if strings.EqualFold(st, state) {
			for zip, attrs := range zips {
				if strings.EqualFold(attrs["CountyName"], county) {
					r := models.Zipcode{
						State:     state,
						Zipcode:   zip,
						County:    county,
						StateFIPS: attrs["StateFIPS"],
						StateAbbr: attrs["StateAbbr"],
						City:      attrs["CityName"],
					}

					rs = append(rs, r)
				}
			}
		}
	}

	if len(rs) == 0 {
		return nil, fmt.Errorf("the state-county pair %q does not exist in the database", state+"-"+county)
	}

	return rs, nil
}

func (lr *LocalDBRepo) FetchByStateCity(state, city string) ([]models.Zipcode, error) {
	rs := make([]models.Zipcode, 0)

	for st, zips := range lr.db.StateTable {
		if strings.EqualFold(st, state) {
			for zip, attrs := range zips {
				if strings.EqualFold(attrs["CityName"], city) {
					r := models.Zipcode{
						State:     state,
						Zipcode:   zip,
						City:      city,
						StateFIPS: attrs["StateFIPS"],
						StateAbbr: attrs["StateAbbr"],
						County:    attrs["CountyName"],
					}

					rs = append(rs, r)
				}
			}
		}
	}

	if len(rs) == 0 {
		return nil, fmt.Errorf("the state-city pair %q does not exist in the database", state+"-"+city)
	}
	return rs, nil
}

func (lr *LocalDBRepo) FetchByCountyCity(county, city string) ([]models.Zipcode, error) {
	rs := make([]models.Zipcode, 0)

	for st, zips := range lr.db.StateTable {
		for zip, attrs := range zips {
			if strings.EqualFold(attrs["CountyName"], county) && strings.EqualFold(attrs["CityName"], city) {
				r := models.Zipcode{
					State:     st,
					Zipcode:   zip,
					City:      city,
					County:    county,
					StateFIPS: attrs["StateFIPS"],
					StateAbbr: attrs["StateAbbr"],
				}

				rs = append(rs, r)
			}
		}
	}

	if len(rs) == 0 {
		return nil, fmt.Errorf("the county-city pair %q does not exist in the database", county+"-"+city)
	}
	return rs, nil
}
