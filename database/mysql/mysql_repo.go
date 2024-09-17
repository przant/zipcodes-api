package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/przant/zipcodes-api/models"
)

type MySQLRepo struct {
	db *sql.DB
}

func NewMySQLRepo() (*MySQLRepo, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		Net:    "tcp",
		Addr:   "mysql-db:3306",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLRepo{db: db}, err
}

func (db *MySQLRepo) Close() {
	if err := db.db.Close(); err != nil {
		log.Fatalf("while closing the MySQL database connection: %s", err)
	}
}

func (mr *MySQLRepo) FetchByZipcode(zipcode string) (*models.Zipcode, error) {
	rows, err := mr.db.Query("SELECT state, state_abbr, zipcode, county, city FROM zipcodes WHERE zipcode = ?", zipcode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	zip := &models.Zipcode{}
	for rows.Next() {
		if err := rows.Scan(&zip.State, &zip.StateAbbr, &zip.Zipcode, &zip.County, &zip.City); err != nil {
			return nil, err
		} else {
			break
		}
	}
	return zip, nil
}

func (mr *MySQLRepo) FetchByCounty(county string) ([]models.Zipcode, error) {
	rows, err := mr.db.Query("SELECT state, state_abbr, zipcode, county, city FROM zipcodes WHERE county = ?", county)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	zips := make([]models.Zipcode, 0)

	for rows.Next() {
		zip := models.Zipcode{}
		if err := rows.Scan(&zip.State, &zip.StateAbbr, &zip.Zipcode, &zip.County, &zip.City); err != nil {
			return nil, err
		} else {
			zips = append(zips, zip)
		}
	}

	return zips, nil
}

func (mr *MySQLRepo) FetchByStateCounty(state, county string) ([]models.Zipcode, error) {
	rows, err := mr.db.Query("SELECT state, state_abbr, zipcode, county, city FROM zipcodes WHERE state = ? AND county = ?", state, county)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	zips := make([]models.Zipcode, 0)

	for rows.Next() {
		zip := models.Zipcode{}
		if err := rows.Scan(&zip.State, &zip.StateAbbr, &zip.Zipcode, &zip.County, &zip.City); err != nil {
			return nil, err
		} else {
			zips = append(zips, zip)
		}
	}

	return zips, nil
}

func (mr *MySQLRepo) FetchByStateCity(state, city string) ([]models.Zipcode, error) {
	rows, err := mr.db.Query("SELECT state, state_abbr, zipcode, county, city FROM zipcodes WHERE state = ? AND city = ?", state, city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	zips := make([]models.Zipcode, 0)

	for rows.Next() {
		zip := models.Zipcode{}
		if err := rows.Scan(&zip.State, &zip.StateAbbr, &zip.Zipcode, &zip.County, &zip.City); err != nil {
			return nil, err
		} else {
			zips = append(zips, zip)
		}
	}

	return zips, nil
}

func (mr *MySQLRepo) FetchByCountyCity(county, city string) ([]models.Zipcode, error) {
	rows, err := mr.db.Query("SELECT state, state_abbr, zipcode, county, city FROM zipcodes WHERE county = ? AND city = ?", county, city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	zips := make([]models.Zipcode, 0)

	for rows.Next() {
		zip := models.Zipcode{}
		if err := rows.Scan(&zip.State, &zip.StateAbbr, &zip.Zipcode, &zip.County, &zip.City); err != nil {
			return nil, err
		} else {
			zips = append(zips, zip)
		}
	}

	return zips, nil
}
