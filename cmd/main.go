package main

import (
	"Fonbet/controllers/api"
	ApiEvents "Fonbet/controllers/api/Events"
	Factors2 "Fonbet/controllers/api/Factors"
	ApiSports "Fonbet/controllers/api/Sports"
	"Fonbet/logging"
	dbConnect "Fonbet/repository/postgres/Connect"
	dbCreate "Fonbet/repository/postgres/Create"
	DbEvents "Fonbet/repository/postgres/Insert/Events"
	DbFactors "Fonbet/repository/postgres/Insert/Factors"
	DbSports "Fonbet/repository/postgres/Insert/Sports"
	UcEvents "Fonbet/usecases/Events"
	UcFactors "Fonbet/usecases/Factors"
	UcSports "Fonbet/usecases/Sports"

	"time"
)

var dbConf = dbConnect.DBClient{
	User:     "postgres",
	Password: "P@ssw0rd",
	Host:     "172.16.14.67",
	Port:     "5432",
	Dbname:   "postgres"}

//var once sync.Once

func main() {
	const urls = "https://www.fon.bet/urls.json"

	var (
		//------- Main
		logger = logging.Logger()
		db     = dbConnect.Connect(&dbConf, logger)
		//------- JsonToStruct
		fonUrl     = api.ListURLStruct{}
		apiSports  = ApiSports.ApiSports{}
		apiEvents  = ApiEvents.ApiEvents{}
		apiFactors = Factors2.ApiFactors{}
		//apiResults = Results2.ApiResults{}
		////------- UseCases
		ucSports  = UcSports.UcSports{}
		ucEvents  = UcEvents.UcEvents{}
		ucFactors = UcFactors.UcFactors{}
		//	ucResults = Results3.UcResults{}
	)
	dbCreate.DBStructure(db, logger)
	for {
		logger.Println("|-------Start-------| Time:", time.Now().Format(time.RFC3339))
		fonUrl.JsonToStruct(urls, logger)

		// -----Sports ----- //
		apiSports.JsonToStruct(&fonUrl, logger)
		ucSports.ReAssign(apiSports)
		var dbSports = DbSports.DbSports{
			UcSportsStruct: ucSports.UcSportsStruct,
		}
		dbSports.Insert(db, logger)
		// ----- Events ----- //
		apiEvents.Parse(&fonUrl, logger)
		ucEvents.ReAssign(apiEvents)
		var dbEvents = DbEvents.DbEvents{
			UcEventStruct: ucEvents.UcEventStruct,
		}
		dbEvents.Insert(db, logger)

		// ----- Factors ----- //
		apiFactors.JsonToStruct(&fonUrl, logger)
		ucFactors.ReAssign(apiFactors)
		var dbFactors = DbFactors.DbFactors{
			UcFactorsStruct: ucFactors.UcFactorsStruct,
		}
		dbFactors.Insert(db, logger)
		//apiResults.JsonToStruct(urlresult, logger)
		//
		//ucResults.ReAssign(apiResults.Fonbet, logger)

		logger.Println("|-------Done-------| Time:", time.Now().Format(time.RFC3339))
		time.Sleep(5 * time.Minute)
	}
	db.Close()

}
