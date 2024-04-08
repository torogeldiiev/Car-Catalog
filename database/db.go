package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/torogeldiiev/car_catalog/config"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Database connection successful")

	return DB, nil
}
