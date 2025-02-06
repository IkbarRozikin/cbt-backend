package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	var err error

	connStr := "user=ikbar password=Ikbar123 dbname=cbt-backend sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	fmt.Println("Database connected")

	return err
}
