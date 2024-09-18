package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	repo "github.com/przant/zipcodes-api/repository"
)

// @Summary Get Zipcode Information
// @Description get info by zipcode
// @ID get-info-by-zipcode
// @Accept json
// @Produce json
// @Param zipcode path string true "Zipcode value"
// @Success 200 {object} models.Zipcode
// @Router /zipcodes/{zipcode} [get]
func GetByZipcode(c echo.Context) error {
	zc := c.Param("zipcode")
	db := repo.GetZipcodesService()
	record, err := db.FetchByZipcode(zc)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, record)
}

// @Summary Get Zipcode Information
// @Description get info by county
// @ID get-info-by-county
// @Accept json
// @Produce json
// @Param county path string true "County name"
// @Success 200 {array} models.Zipcode
// @Router /counties/{county} [get]
func GetByCounty(c echo.Context) error {
	co := c.Param("county")
	db := repo.GetZipcodesService()
	records, err := db.FetchByCounty(co)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, records)
}

// @Summary Get Zipcode Information
// @Description get info by state and county
// @ID get-info-by-state-county
// @Accept json
// @Produce json
// @Param state path string true "State name"
// @Param county path string true "County name"
// @Success 200 {array} models.Zipcode
// @Router /states/{state}/counties/{county} [get]
func GetByStateCounty(c echo.Context) error {
	st := c.Param("state")
	co := c.Param("county")
	db := repo.GetZipcodesService()
	records, err := db.FetchByStateCounty(st, co)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, records)
}

// @Summary Get Zipcode Information
// @Description get info by state and city
// @ID get-info-by-state-city
// @Accept json
// @Produce json
// @Param state path string true "State name"
// @Param city path string true "City name"
// @Success 200 {array} models.Zipcode
// @Router /states/{state}/cities/{city} [get]
func GetByStateCity(c echo.Context) error {
	st := c.Param("state")
	city := c.Param("city")
	db := repo.GetZipcodesService()
	records, err := db.FetchByStateCity(st, city)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, records)
}

// @Summary Get Zipcode Information
// @Description get info by county and city
// @ID get-info-by-county-city
// @Accept json
// @Produce json
// @Param county path string true "County name"
// @Param city path string true "City name"
// @Success 200 {array} models.Zipcode
// @Router /counties/{county}/cities/{city} [get]
func GetByCountyCity(c echo.Context) error {
	co := c.Param("county")
	city := c.Param("city")
	db := repo.GetZipcodesService()
	records, err := db.FetchByCountyCity(co, city)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, records)
}
