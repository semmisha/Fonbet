package main

import (
	"Fonbet/api"
	"Fonbet/db/postgresql"
	"Fonbet/db/postgresql/connect"
	"Fonbet/db/postgresql/create"
	fonstruct "Fonbet/json"
	logging "Fonbet/logging"
	"Fonbet/utils"
	"fmt"
	"time"
)

var events *fonstruct.FonbetEvents
var result *fonstruct.FonbetResult
var dbConf = connect.DBClient{
	User:     "postgres",
	Password: "P@ssw0rd",
	Host:     "172.16.14.67",
	Port:     "5432",
	Dbname:   "postgres"}

//var once sync.Once

func main() {
	logger := logging.Logger()
	urlevents := "https://line06w.bkfon-resources.com/events/list?lang=ru&version=7987900598&scopeMarket=1600"
	urlresult := "https://clientsapi03w.bkfon-resources.com/results/results.json.php?locale=ru"

	for {

		db := connect.Connect(&dbConf, logger)
		//once.Do()
		create.DBStructure(db, logger)
		api.Parse(&events, urlevents, logger)
		err := Postgres.Sport(events, db, logger)
		if err != nil {
			fmt.Println(err)
		}
		Postgres.Events(events, db, logger)

		for i := 0; i <= 4; i++ {
			api.Parse(&result, utils.DayCount(urlresult, i), logger)
			//test := Postgres.Connect(&dbConf)
			//fmt.Printf(":%v", test)
			err := Postgres.Result(result, db, logger)
			if err != nil {
				fmt.Println(err)

			}
			Postgres.CompareResult(db, logger)

		}

		fmt.Println("|-------Done-------|")
		db.Close()

		time.Sleep(15 * time.Minute)

	}
}
