package main

import (
	"Fonbet/api"
	"Fonbet/db/postgresql"
	fonstruct "Fonbet/json"
	"fmt"
	"time"
)

var events *fonstruct.FonbetEvents
var result *fonstruct.FonbetResult

func main() {
	urlevents := "https://line06w.bkfon-resources.com/events/list?lang=ru&version=7987900598&scopeMarket=1600"
	urlresult := "https://clientsapi03w.bkfon-resources.com/results/results.json.php?locale=ru&lineDate=2022-05-03"

	for {

		api.Parse(&events, urlevents)
		api.Parse(&result, urlresult)
		db := Postgres.DbConnect2()

		err := Postgres.Result(result, db)
		if err != nil {
			fmt.Println(err)

		}

		//err := Postgres.Sport(events, db)
		//if err != nil {
		//	fmt.Println(err)
		//
		//}
		//err = Postgres.Events(events, db)
		//if err != nil {
		//	fmt.Println(err)
		//
		//}
		//err = Postgres.Factor(events, db)
		//if err != nil {
		//	fmt.Println(err)
		//
		//}

		err = db.Close()
		if err != nil {
			fmt.Println(err)

		}
		time.Sleep(30 * time.Minute)

	}
}
