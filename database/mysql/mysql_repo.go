package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type MySQLRepo struct {
	db *sql.DB
}

func NewMySQLRepo() (*MySQLRepo, error) {
	err := godotenv.Load("mysql.env")
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
