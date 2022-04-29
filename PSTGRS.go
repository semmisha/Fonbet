package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type ConInfo map[string]string

func DbConnect() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=P@ssw0rd dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("INSERT INTO main VALUES(56,10);")
	if err != nil {
		log.Fatal(err)
	}

	rstr, _ := rows.Columns()
	fmt.Printf("%+v", rstr)

	defer db.Close()
}

func DbInject() {

}
