package Postgres

import (
	fonstruct "Fonbet/json"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

//

func Result(fonbet *fonstruct.FonbetResult, db *sql.DB) (err error) {
	var sum int

	var i int = 0
	for i < len(fonbet.Events) {
		exist, err := db.Query(`SELECT coalesce((sum(CASE WHEN $1 IN ("stringname") AND $2 IN ("starttime") THEN 1 ELSE 0 END)),0) FROM results ;`, fonbet.Events[i].Name, fonbet.Events[i].StartTime)
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

		if sum == 0 && strings.Contains(fonbet.Events[i].Name, "-") {

			fontime, err := time.Parse(time.RFC3339, strconv.Itoa(fonbet.Events[i].StartTime))
			fmt.Printf("Name: %v StartTime: %v  Result: %v \n", fonbet.Events[i].Name, fontime, fonbet.Events[i].Score)
			_, err = db.Exec("INSERT INTO results (stringname, starttime, score) VALUES ($1, $2, $3)", fonbet.Events[i].Name, fonbet.Events[i].StartTime, fonbet.Events[i].Score)
			if err != nil {
				fmt.Println(err)
			}

		}
		i++
	}
	fmt.Printf("Total Result rows: %v\n ", i)
	return
}
