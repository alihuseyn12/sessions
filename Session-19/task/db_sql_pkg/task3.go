package db_sql_pkg

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Connectiondb() *sql.DB {
	connStr := "host=localhost user=student port=5432 password=securepass dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("error pinging db", err)
	}
	fmt.Println("Connection successful")
	DB = db

	return DB
}
