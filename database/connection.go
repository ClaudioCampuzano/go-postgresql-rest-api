package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	connStr := "postgres://-:-*@database-omia.ccco8vwbpupr.us-west-2.rds.amazonaws.com/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
