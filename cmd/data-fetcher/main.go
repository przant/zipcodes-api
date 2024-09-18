package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/przant/zipcodes-api/models"
	"github.com/przant/zipcodes-api/utils"
)

var (
	dbType string
)

const (
	INSERT_MYSQL = "USE %s;\nINSERT INTO %s(state_fips, state, state_abbr, zipcode, county, city) VALUES\n"
	INSERT_MONGO = "const conn = new Mongo();\n\nconst db = conn.getDB(\"%s\");\n\ndb.%s.drop();\n"
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

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("while loading %q file: %s", ".env", err)
	}

	switch {
	case dbType == MYSQLDB:
		createMySQlInsert(zips)
	case dbType == MONGODB:
		createMONGOInsert(zips)
	}
}

func createMySQlInsert(zips []models.Zipcode) {
	file, err := os.OpenFile("insert.sh", os.O_RDWR|os.O_CREATE, 0650)
	if err != nil {
		log.Fatalf("while creating the file %q: %s", "insert.sql", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "SQL=$(cat <<-EOF\n")
	fmt.Fprintf(file, INSERT_MYSQL, os.Getenv("MYSQL_DATABASE"), os.Getenv("MYSQL_TABLE"))
	for _, v := range zips[1 : len(zips)-1] {
		fmt.Fprintf(file, "(%q,%q,%q,%q,%q,%q),\n", v.StateFIPS, v.State, v.StateAbbr, v.Zipcode, v.County, v.City)
	}
	v := zips[len(zips)-1]
	fmt.Fprintf(file, "(%q,%q,%q,%q,%q,%q);\n", v.StateFIPS, v.State, v.StateAbbr, v.Zipcode, v.County, v.City)
	fmt.Fprintf(file, "EOF\n)\necho $SQL | mysql -u$MYSQL_USER -p$MYSQL_PASSWORD\n")
}

func createMONGOInsert(zips []models.Zipcode) {
	file, err := os.OpenFile("seed.js", os.O_RDWR|os.O_CREATE, 0650)
	if err != nil {
		log.Fatalf("while creating the file %q: %s", "seed.js", err)
	}
	defer file.Close()

	fmt.Fprintf(file, INSERT_MONGO, os.Getenv("MONGODB_DATABASE"), os.Getenv("MONGODB_COLLECTION"))
	fmt.Fprintf(file, "\ndb.%s.insertMany(", os.Getenv("MONGODB_COLLECTION"))
	bs, err := json.MarshalIndent(zips[1:], "", "    ")
	if err != nil {
		log.Fatalf("while marshaling the data to insert: %s", err)
	}
	fmt.Fprintf(file, "%s", bs)
	fmt.Fprintf(file, ");\n")
}
