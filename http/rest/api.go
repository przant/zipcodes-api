package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	PORT = "20790"
)

func NewAPI() (api *echo.Echo) {
	api = echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	return api
}
