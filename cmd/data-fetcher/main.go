package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/przant/zipcodes-api/models"
	"github.com/przant/zipcodes-api/utils"
)

var (
	dbType string
)

const (
	INSERT_MYSQL = "USE us_zipcodes;\nINSERT INTO zipcodes(state, state_abbr, zipcode, county, city) VALUES"
	INSERT_MONGO = "const conn = new Mongo();\n\nconst db = conn.getDB(\"us_zipcodes\");\n\ndb.zipcodes.drop();\n"
	MYSQLDB      = "mysql"
	MONGODB      = "mongo"
)

func init() {
	flag.StringVar(&dbType, "database", "", "The database type to use. Valid values are mysql and mongo")
	flag.StringVar(&dbType, "d", "", "The database type to use. Valid values are mysql and mongo(shorthand)")
}

func main() {
	flag.Parse()

	zips, err := utils.FetchData()
	if err != nil {
		log.Fatalf("while feching the zipcodes data: %s", err)
	}

	switch {
	case dbType == MYSQLDB:
		createMySQlInsert(zips)
	case dbType == MONGODB:
		createMONGOInsert(zips)
	}
}

func createMySQlInsert(zips []models.Zipcode) {
	file, err := os.OpenFile("insert.sql", os.O_RDWR|os.O_CREATE, 0650)
	if err != nil {
		log.Fatalf("while creating the file %q: %s", "insert.sql", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\n", INSERT_MYSQL)
	for _, v := range zips[1 : len(zips)-1] {
		fmt.Fprintf(file, "(%q,%q,%q,%q,%q),\n", v.State, v.StateAbbr, v.Zipcode, v.County, v.City)
	}
	v := zips[len(zips)-1]
	fmt.Fprintf(file, "(%q,%q,%q,%q,%q);\n", v.State, v.StateAbbr, v.Zipcode, v.County, v.City)
}

func createMONGOInsert(zips []models.Zipcode) {
	file, err := os.OpenFile("seed.js", os.O_RDWR|os.O_CREATE, 0650)
	if err != nil {
		log.Fatalf("while creating the file %q: %s", "seed.js", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\ndb.zipcodes.insertMany(", INSERT_MONGO)
	bs, err := json.MarshalIndent(zips[1:], "", "    ")
	if err != nil {
		log.Fatalf("while marshaling the data to insert: %s", err)
	}
	fmt.Fprintf(file, "%s", bs)
	fmt.Fprintf(file, ");\n")
}
