package main

import (
	logging "Fonbet/Logging"
	"Fonbet/api"
	"Fonbet/db/postgresql"
	"Fonbet/db/postgresql/connect"
	fonstruct "Fonbet/json"
	"Fonbet/utils"
	"fmt"
	"time"
)

var events *fonstruct.FonbetEvents
var result *fonstruct.FonbetResult
var dbConf = connect.DBClient{
	User:     "postgres",
	Password: "P@ssw0rd",
	Host:     "localhost",
	Port:     "5432",
	Dbname:   "postgres"}

func main() {
	logger := logging.Logger()
	urlevents := "https://line06w.bkfon-resources.com/events/list?lang=ru&version=7987900598&scopeMarket=1600"
	urlresult := "https://clientsapi03w.bkfon-resources.com/results/results.json.php?locale=ru"

	for {
		db := connect.DbConnect2(logger)
		api.Parse(&events, urlevents, logger)
		err := Postgres.Sport(events, db, logger)
		if err != nil {
			fmt.Println(err)

		}
		err = Postgres.Events(events, db, logger)
		if err != nil {
			fmt.Println(err)

		}
		err = Postgres.Factor(events, db, logger)
		if err != nil {
			fmt.Println(err)

		}
		for i := 0; i <= 5; i++ {

			api.Parse(&result, utils.DayCount(urlresult, i), logger)
			//test := Postgres.Connect(&dbConf)
			//fmt.Printf(":%v", test)
			err := Postgres.Result(result, db, logger)
			if err != nil {
				fmt.Println(err)

			}
			Postgres.CompareResult(result, db, logger)

		}
		err = db.Close()
		if err != nil {
			fmt.Println(err)

		}
		time.Sleep(5 * time.Minute)

	}
}
