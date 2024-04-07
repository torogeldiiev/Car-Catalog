package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

var DB *sql.DB // Exported db variable

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

	// Run migrations
	if err := RunMigrations(); err != nil {
		return nil, err
	}

	return DB, nil
}

func RunMigrations() error {
	migrationDir := "./database"
	files, err := ioutil.ReadDir(migrationDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrationPath := filepath.Join(migrationDir, file.Name())
			migrationSQL, err := ioutil.ReadFile(migrationPath)
			if err != nil {
				return err
			}

			if _, err := DB.Exec(string(migrationSQL)); err != nil {
				return err
			}
		}
	}

	return nil
}
