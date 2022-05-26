package main

import (
	"Fonbet/logging"
	"Fonbet/repository/postgres/Connect"
	"Fonbet/repository/postgres/Create"
	"Fonbet/repository/postgres/Insert/Events"
	"Fonbet/repository/postgres/Insert/Factors"
	"Fonbet/repository/postgres/Insert/Results"
	"Fonbet/repository/postgres/Insert/Sports"
	Events2 "Fonbet/usecases/Events"
	Factors2 "Fonbet/usecases/Factors"
	Results2 "Fonbet/usecases/Results"
	Sports2 "Fonbet/usecases/Sports"
	"time"
)

var dbConf = Connect.DBClient{
	User:     "postgres",
	Password: "P@ssw0rd",
	Host:     "127.0.0.1",
	Port:     "5432",
	Dbname:   "postgres"}

//var once sync.Once

func main() {
	const urlevents = "https://line06w.bkfon-resources.com/events/list?lang=ru&version=7987900598&scopeMarket=1600"
	const urlresult = "https://clientsapi03w.bkfon-resources.com/results/results.json.php?locale=ru"

	var (
		//------- Main
		logger = logging.Logger()
		Db     = Connect.Connect(&dbConf, logger)
		//------- Parse
		sports  = Sports.DbSports{Db: Db}
		events  = Events.DbEvents{Db: Db}
		factors = Factors.DbFactors{Db: Db}
		results = Results.DbResults{Db: Db}
		//------- UseCases
		ucSports  = Sports2.UcSportsSlice{}
		ucEvents  = Events2.UcEventSlice{}
		ucFactors = Factors2.UcFactorsSlice{}
		ucResults = Results2.UcResultSlice{}
	)
	Create.DBStructure(Db, logger)
	for {
		logger.Println("|-------Start-------| Time:", time.Now().Format(time.RFC3339))

		sports.Fonbet.Parse(urlevents, logger)
		ucSports.ReAssign(sports.Fonbet)

		events.Fonbet.Parse(urlevents, logger)
		ucEvents.ReAssign(events.Fonbet)

		factors.Fonbet.Parse(urlevents, logger)
		ucFactors.ReAssign(factors.Fonbet)
		results.Fonbet.Parse(urlresult, logger)

		ucResults.ReAssign(results.Fonbet, logger)
		logger.Println("|-------Done-------| Time:", time.Now().Format(time.RFC3339))
		time.Sleep(5 * time.Minute)
	}
	Db.Close()

}
