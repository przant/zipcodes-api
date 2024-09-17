package main

import (
	"flag"
	"log"

	localdb "github.com/przant/zipcodes-api/database/local"
	mgodb "github.com/przant/zipcodes-api/database/mongo"
	sqldb "github.com/przant/zipcodes-api/database/mysql"
	"github.com/przant/zipcodes-api/http/rest"
	repo "github.com/przant/zipcodes-api/repository"
)

var (
	service string
)

const (
	LOCALDBSVC = "local"
	MYSQLDBSVC = "mysql"
	MONGODBSVC = "mongo"
)

func init() {
	flag.StringVar(&service, "database", "local", "The database name to use to store and fetch the US zipcodes")
	flag.StringVar(&service, "d", "local", "The database name to use to store and fetch the US zipcodes(shorthand)")
}

func main() {
	flag.Parse()
	var db repo.ZipcodesRepo
	var err error

	switch {
	case service == LOCALDBSVC:
		db, err = localdb.NewLocalDBRepo()
		if err != nil {
			log.Fatalf("while connecting to the db: %s", err)
		}
	case service == MYSQLDBSVC:
		db, err = sqldb.NewMySQLRepo()
		if err != nil {
			log.Fatalf("while connecting to the db: %s", err)
		}
	case service == MONGODBSVC:
		db, err = mgodb.NewMongoRepo()
		if err != nil {
			log.Fatalf("while connecting to the db: %s", err)
		}
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
