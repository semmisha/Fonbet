package Postgres

import (
	"database/sql"
	"log"
)

func DbConnect2() *sql.DB {

	//fmt.Sprintf("postgres://%s:%s")
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=P@ssw0rd dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
