package main

import (
	"log"

	db "github.com/przant/zipcodes-api/database/mysql"
	"github.com/przant/zipcodes-api/http/rest"
	repo "github.com/przant/zipcodes-api/repository"
)

func main() {

	db, err := db.NewMySQLRepo()
	if err != nil {
		log.Fatalf("while connecting to the db: %s", err)
	}
	defer db.Close()

	repo.NewZipcodesService(db)

	api := rest.NewAPI()
	api.GET("/zipcodes/:zipcode", rest.GetByZipcode)
	api.GET("/counties/:county", rest.GetByCounty)
	api.GET("/states/:state/counties/:county", rest.GetByStateCounty)
	api.GET("/states/:state/cities/:city", rest.GetByStateCity)
	api.GET("/counties/:county/cities/:city", rest.GetByCountyCity)

	api.Logger.Fatal(api.Start(rest.PORT))
}
