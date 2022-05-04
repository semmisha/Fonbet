package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

//type ConInfo struct {
//	host     string
//	port     string
//	user     string
//	password string
//	dbname   string
//	sslmode  string
//}

func DbConnect2() *sql.DB {

	//fmt.Sprintf("postgres://%s:%s")
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=P@ssw0rd dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Sport(fonbet *Fonbet2, db *sql.DB) (err error) {
	var sum int
	var i int = 0
	for i < len(fonbet.Sports) {
		exist, err := db.Query(`SELECT sum(CASE WHEN $1 IN ("sportid") THEN 1 ELSE 0 END) FROM sports ;`, fonbet.Sports[i].Id)
		if err != nil {
			log.Println(err)
		}
		for exist.Next() {
			err := exist.Scan(&sum)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println(sum)
		}

		if sum == 0 {
			_, err := db.Exec("INSERT INTO sports VALUES ($1, $2, $3)", fonbet.Sports[i].Id, fonbet.Sports[i].ParentId, fonbet.Sports[i].Name)
			if err != nil {
				fmt.Println(err)
			}

		}
		i++
	}
	fmt.Printf("Total Sports rows: %v\n ", i)
	return
}

func Events(fonbet *Fonbet2, db *sql.DB) (err error) {
	var sum, levels int = 0, 0
	CheckLevels(fonbet, &levels)
	CreateLevels(db, &levels)
	//time.Sleep(5 * time.Minute)
	var i int = 0
	for i < len(fonbet.Events) {
		query := fmt.Sprintf("SELECT sum(CASE WHEN $1 IN (id) THEN 1 ELSE 0 END) FROM events_level_%v ;", fonbet.Events[i].Level)
		exist, err := db.Query(query, fonbet.Events[i].Id)
		if err != nil {
			log.Println(err)
		}
		for exist.Next() {
			err := exist.Scan(&sum)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println(sum)
		}

		if sum == 0 {
			query := fmt.Sprintf("INSERT INTO events_level_%v VALUES ($1, $2, $3,$4,$5,$6)", fonbet.Events[i].Level)
			_, err := db.Exec(query, fonbet.Events[i].Id, fonbet.Events[i].ParentId, fonbet.Events[i].Name, fonbet.Events[i].SportId, fonbet.Events[i].Team1, fonbet.Events[i].Team2)
			if err != nil {
				fmt.Println(err)
			}

		}
		i++
	}
	fmt.Printf("Total Events rows: %v\n ", i)
	return
}

func Factor(fonbet *Fonbet2, db *sql.DB) (err error) {
	var sum int
	var i int = 0
	for i < len(fonbet.CustomFactors) {
		exist, err := db.Query(`SELECT sum(CASE WHEN $1 IN ("eventid") THEN 1 ELSE 0 END) FROM factors ;`, fonbet.CustomFactors[i].E)
		if err != nil {
			log.Println(err)
		}
		for exist.Next() {
			err := exist.Scan(&sum)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println(sum)
		}

		if sum == 0 {

			for b := 0; b < len(fonbet.CustomFactors[i].Factors); b++ {
				_, err := db.Exec("INSERT INTO factors VALUES ($1, $2, $3,$4,$5)", fonbet.CustomFactors[i].E, fonbet.CustomFactors[i].Factors[b].F, fonbet.CustomFactors[i].Factors[b].V, fonbet.CustomFactors[i].Factors[b].P, fonbet.CustomFactors[i].Factors[b].Pt)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		i++
	}
	fmt.Printf("Total Factor rows: %v\n ", i)
	return
}

func CheckLevels(fonbet *Fonbet2, levels *int) {
	for i := 0; i < len(fonbet.Events); i++ {

		if *levels < fonbet.Events[i].Level {
			*levels = fonbet.Events[i].Level
		}

	}

}
func CreateLevels(db *sql.DB, levels *int) {

	for i := 2; i <= *levels; i++ {
		currentlevel := fmt.Sprintf("events_level_%v", i)
		parentlevel := fmt.Sprintf("events_level_%v", i-1)

		query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (id INT PRIMARY KEY, parentid int, eventname VARCHAR(150), sportid INT, team1 VARCHAR(50), team2 VARCHAR(50), foreign key (parentid) references %v (id) )", currentlevel, parentlevel)
		_, err := db.Exec(query)
		if err != nil {
			fmt.Println(err)
		}

	}

}
