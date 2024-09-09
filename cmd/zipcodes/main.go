package main

import (
	"log"

	db "github.com/przant/zipcodes-api/database/local"
	"github.com/przant/zipcodes-api/http/rest"
	repo "github.com/przant/zipcodes-api/repository"
)

func main() {

	db, err := db.NewLocalDBRepo()
	if err != nil {
		log.Fatalf("while connecting to the db: %s", err)
	}
	defer db.Close()

	repo.NewZipcodesService(db)

	api := rest.NewAPI()

	api.GET("/zipcode", rest.GetByZipcode)
	api.GET("/county", rest.GetByCounty)
	api.GET("/state-county", rest.GetByStateCounty)
	api.GET("/state-city", rest.GetByStateCity)
	api.GET("/county-city", rest.GetByCountyCity)

	api.Logger.Fatal(api.Start(rest.PORT))
}
