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
	logger = logging.Logger()
	db     = dbConnect.Connect(&dbConf, logger)

	//TODO ------- JsonToStruct
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

func main() {

	dbCreate.DBStructure(db, logger)
	for {
		logger.Println("|-------Start-------| Time:", time.Now().Format(time.RFC3339))
		fonUrl.JsonToStruct(urls, logger)

		go func() {
			time.Sleep(1 * time.Minute)
			// TODO -----Sports ----- //
			apiSports.JsonToStruct(&fonUrl, logger)
			ucSports.ReAssign(apiSports)
			var dbSports = DbEvents.DbSports{
				UcSportsStruct: ucSports.UcSportsStruct,
			}
			dbSports.Insert(db, logger)

			// TODO ----- Events ----- //
			apiEvents.Parse(&fonUrl, logger)
			ucEvents.ReAssign(apiEvents)
			var dbEvents = DbEvents.DbEvents{
				UcEventStruct: ucEvents.UcEventStruct,
			}
			dbEvents.Insert(db, logger)

			// TODO ----- Factors ----- //
			apiFactors.JsonToStruct(&fonUrl, logger)
			ucFactors.ReAssign(apiFactors)
			var dbFactors = DbEvents.DbFactors{
				UcFactorsStruct: ucFactors.UcFactorsStruct,
			}
			dbFactors.Insert(db, logger)

			// TODO ----- Results ----- //

			apiResults.JsonToStruct(&fonUrl, logger)
			ucResults.ReAssign(apiResults, logger)
			var dbResults = DbEvents.DbResults{
				UcResultsStruct: ucResults.UcResultsStruct,
			}
			dbResults.Insert(db, logger)

			logger.Println("|-------Done-------| Time:", time.Now().Format(time.RFC3339))
		}()
		// TODO ----- Compare ----- //
		go func() {
			var (
				testEvent  DbEvents.DbEvents
				testResult DbEvents.DbResults
			)

			testEvent.Select(db, logger)
			testResult.Select(db, logger)

			dbResultId := Compare.CompareResult(testEvent, testResult, logger)
			dbResultId.Update(db, logger)

		}()
		time.Sleep(15 * time.Minute)

	}
	db.Close()

}
