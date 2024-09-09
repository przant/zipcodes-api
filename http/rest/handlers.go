package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	repo "github.com/przant/zipcodes-api/repository"
)

func GetByZipcode(c echo.Context) error {
	zc := c.QueryParam("zipcode")
	db := repo.GetZipcodesService()
	record, err := db.FetchByZipcode(zc)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, record)
}

func GetByCounty(c echo.Context) error {
	co := c.QueryParam("county")
	db := repo.GetZipcodesService()
	records, err := db.FetchByCounty(co)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, records)
}

func GetByStateCounty(c echo.Context) error {
	st := c.QueryParam("state")
	co := c.QueryParam("county")
	db := repo.GetZipcodesService()
	records, err := db.FetchByStateCounty(st, co)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, records)
}

func GetByStateCity(c echo.Context) error {
	st := c.QueryParam("state")
	city := c.QueryParam("city")
	db := repo.GetZipcodesService()
	records, err := db.FetchByStateCity(st, city)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, records)
}

func GetByCountyCity(c echo.Context) error {
	co := c.QueryParam("county")
	city := c.QueryParam("city")
	db := repo.GetZipcodesService()
	records, err := db.FetchByCountyCity(co, city)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, records)
}
