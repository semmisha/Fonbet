package main

import (
	"Fonbet/controllers/api"
	ApiEvents "Fonbet/controllers/api/Events"
	Factors2 "Fonbet/controllers/api/Factors"
	ApiResults "Fonbet/controllers/api/Results"
	ApiSports "Fonbet/controllers/api/Sports"
	"Fonbet/logging"
	dbConnect "Fonbet/repository/postgres/Connect"
	dbCreate "Fonbet/repository/postgres/Create"
	DbEvents "Fonbet/repository/postgres/Manipulate"
	"Fonbet/usecases/Compare"
	UcEvents "Fonbet/usecases/Convert"
	"time"
)

var dbConf = dbConnect.DBClient{
	User:     "postgres",
	Password: "password",
	Host:     "localhost",
	Port:     "5432",
	Dbname:   "postgres"}

const urls = "https://www.fon.bet/urls.json"
const urlsfile = "data"

func main() {

	var (
		//TODO ------- Main
		Logger = logging.Logger()
		db     = dbConnect.Connect(&dbConf, Logger)
	)
	dbCreate.DBStructure(db, Logger)
	for {
		//TODO ------- JsonToStruct
		var (
			fonUrl     = api.NewListURLStruct()
			apiSports  = ApiSports.NewApiSports()
			apiEvents  = ApiEvents.NewApiEvents()
			apiFactors = Factors2.NewApiFactors()
			apiResults = ApiResults.NewApiResults()

			//TODO ------- UseCases
			ucSports  = UcEvents.NewUcSports()
			ucEvents  = UcEvents.NewUcEvents()
			ucFactors = UcEvents.NewUcFactors()
			ucResults = UcEvents.NewUcResults()
		)
		Logger.Println("|-------Start-------| Time:", time.Now().Format(time.RFC3339))
		err := fonUrl.FileToStruct(urlsfile, Logger)
		if err != nil {
			Logger.Fatalf("Cant retieve List of APi after 5 retries, error:%v", err)
		}

		// TODO -----Sports ----- //
		apiSports.JsonToStruct(fonUrl, Logger)
		ucSports.ReAssign(*apiSports)
		var dbSports = DbEvents.DbSports{
			UcSportsStruct: ucSports.UcSportsStruct,
		}
		dbSports.Insert(db, Logger)

		// TODO ----- Events ----- //
		apiEvents.Parse(fonUrl, Logger)
		ucEvents.ReAssign(*apiEvents)
		var dbEvents = DbEvents.DbEvents{
			UcEventStruct: ucEvents.UcEventStruct,
		}
		dbEvents.Insert(db, Logger)

		// TODO ----- Factors ----- //
		apiFactors.JsonToStruct(fonUrl, Logger)
		ucFactors.ReAssign(*apiFactors)
		var dbFactors = DbEvents.DbFactors{
			UcFactorsStruct: ucFactors.UcFactorsStruct,
		}
		dbFactors.Insert(db, Logger)

		// TODO ----- Results ----- //

		apiResults.JsonToStruct(fonUrl, Logger)
		ucResults.ReAssign(*apiResults, Logger)
		var dbResults = DbEvents.DbResults{
			UcResultsStruct: ucResults.UcResultsStruct,
		}
		dbResults.Insert(db, Logger)

		// TODO ----- Compare ----- //

		var (
			compareEvent  DbEvents.DbEvents
			compareResult DbEvents.DbResults
		)
		compareEvent.Select(db, Logger)
		compareResult.Select(db, Logger)
		dbResultId := Compare.CompareResult(compareEvent, compareResult, Logger)
		dbResultId.Update(db, Logger)

		Logger.Println("|-------Done-------| Time:", time.Now().Format(time.RFC3339))

		//TODO ---- Sleep ----- //
		timeSleep := api.RandomNum(40)
		Logger.Infof("Sleep for %v minutes", timeSleep)
		time.Sleep(time.Duration(timeSleep) * time.Minute)

	}

}
