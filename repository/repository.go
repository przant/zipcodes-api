package repository

import (
	"github.com/przant/zipcodes-api/models"
)

type ZipcodesRepo interface {
	FetchByZipcode(zipcode string) (*models.Zipcode, error)
	FetchByCounty(county string) ([]models.Zipcode, error)
	FetchByStateCounty(state, county string) ([]models.Zipcode, error)
	FetchByStateCity(state, city string) ([]models.Zipcode, error)
	FetchByCountyCity(county, city string) ([]models.Zipcode, error)
	Close()
}

type ZipcodesService struct {
	repo ZipcodesRepo
}

var (
	instance *ZipcodesService
)

func NewZipcodesService(zr ZipcodesRepo) {
	if instance == nil {
		instance = &ZipcodesService{
			repo: zr,
		}
	}
}

func GetZipcodesService() *ZipcodesService {
	return instance
}

func (zs *ZipcodesService) FetchByZipcode(zipcode string) (*models.Zipcode, error) {
	zcs, err := zs.repo.FetchByZipcode(zipcode)
	if err != nil {
		return nil, err
	}

	return zcs, nil
}

func (zs *ZipcodesService) FetchByCounty(county string) ([]models.Zipcode, error) {
	zcs, err := zs.repo.FetchByCounty(county)
	if err != nil {
		return nil, err
	}

	return zcs, nil
}

func (zs *ZipcodesService) FetchByStateCounty(state, county string) ([]models.Zipcode, error) {
	zcs, err := zs.repo.FetchByStateCounty(state, county)
	if err != nil {
		return nil, err
	}

	return zcs, nil
}

func (zs *ZipcodesService) FetchByStateCity(state, city string) ([]models.Zipcode, error) {
	zcs, err := zs.repo.FetchByStateCity(state, city)
	if err != nil {
		return nil, err
	}

	return zcs, nil
}

func (zs *ZipcodesService) FetchByCountyCity(county, city string) ([]models.Zipcode, error) {
	zcs, err := zs.repo.FetchByCountyCity(county, city)
	if err != nil {
		return nil, err
	}

	return zcs, nil
}

func (zs *ZipcodesService) Close() {
	zs.repo.Close()
}
