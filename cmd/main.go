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
	Password: "P@ssw0rd",
	Host:     "172.16.14.67",
	Port:     "5432",
	Dbname:   "postgres"}

const urls = "https://www.fon.bet/urls.json"

var (
	//TODO ------- Main
	Logger = logging.Logger()
	db     = dbConnect.Connect(&dbConf, Logger)
)

func main() {

	dbCreate.DBStructure(db, Logger)
	for {
		//TODO ------- JsonToStruct
		var (
			fonUrl     = api.ListURLStruct{}
			apiSports  = ApiSports.ApiSports{}
			apiEvents  = ApiEvents.ApiEvents{}
			apiFactors = Factors2.ApiFactors{}
			apiResults = ApiResults.ApiResults{}

			//TODO ------- UseCases
			ucSports  = UcEvents.UcSports{}
			ucEvents  = UcEvents.UcEvents{}
			ucFactors = UcEvents.UcFactors{}
			ucResults = UcEvents.UcResults{}
		)
		Logger.Println("|-------Start-------| Time:", time.Now().Format(time.RFC3339))
		fonUrl.JsonToStruct(urls, Logger)

		// TODO -----Sports ----- //
		apiSports.JsonToStruct(&fonUrl, Logger)
		ucSports.ReAssign(apiSports)
		var dbSports = DbEvents.DbSports{
			UcSportsStruct: ucSports.UcSportsStruct,
		}
		dbSports.Insert(db, Logger)

		// TODO ----- Events ----- //
		apiEvents.Parse(&fonUrl, Logger)
		ucEvents.ReAssign(apiEvents)
		var dbEvents = DbEvents.DbEvents{
			UcEventStruct: ucEvents.UcEventStruct,
		}
		dbEvents.Insert(db, Logger)

		// TODO ----- Factors ----- //
		apiFactors.JsonToStruct(&fonUrl, Logger)
		ucFactors.ReAssign(apiFactors)
		var dbFactors = DbEvents.DbFactors{
			UcFactorsStruct: ucFactors.UcFactorsStruct,
		}
		dbFactors.Insert(db, Logger)

		// TODO ----- Results ----- //

		apiResults.JsonToStruct(&fonUrl, Logger)
		ucResults.ReAssign(apiResults, Logger)
		var dbResults = DbEvents.DbResults{
			UcResultsStruct: ucResults.UcResultsStruct,
		}
		dbResults.Insert(db, Logger)

		Logger.Println("|-------Done-------| Time:", time.Now().Format(time.RFC3339))

		// TODO ----- Compare ----- //

		var (
			testEvent  DbEvents.DbEvents
			testResult DbEvents.DbResults
		)

		testEvent.Select(db, Logger)
		testResult.Select(db, Logger)

		dbResultId := Compare.CompareResult(testEvent, testResult, Logger)
		dbResultId.Update(db, Logger)

		time.Sleep(20 * time.Minute)

	}
	db.Close()

}
