package main

import (
	"log"

	db "github.com/przant/zipcodes-api/database/mysql"
)

func main() {

	db, err := db.NewMySQLRepo()
	if err != nil {
		log.Fatalf("while connecting to the db: %s", err)
	}

	log.Print("Conexion attempt succesfully to the mysql database")

	defer db.Close()
}
