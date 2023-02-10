package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Dbconnection() *sql.DB {
	connectionStr := "user=postgres dbname=gostore password=241917 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
