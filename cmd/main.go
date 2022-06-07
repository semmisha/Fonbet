package main

import (
	"Fonbet/controllers/api"
	Events2 "Fonbet/controllers/api/Events"
	"Fonbet/logging"
	"Fonbet/repository/postgres/Connect"
	"Fonbet/repository/postgres/Create"
	Events3 "Fonbet/usecases/Events"
	"time"
)

var dbConf = Connect.DBClient{
	User:     "postgres",
	Password: "P@ssw0rd",
	Host:     "172.16.14.67",
	Port:     "5432",
	Dbname:   "postgres"}

//var once sync.Once

func main() {
	const urls = "https://www.fon.bet/urls.json"
	const urlevents = "https://line06w.bkfon-resources.com/apiEvents/list?lang=ru&version=7987900598&scopeMarket=1600"
	const urlresult = "https://clientsapi03w.bkfon-resources.com/apiResults/apiResults.json.php?locale=ru"

	var (
		//------- Main
		logger = logging.Logger()
		db     = Connect.Connect(&dbConf, logger)
		//------- JsonToStruct
		fonUrl = api.ListURLStruct{}
		//apiSports = Sports2.ApiSports{}
		apiEvents = Events2.ApiEvents{}
		//apiFactors = Factors2.ApiFactors{}
		//apiResults = Results2.ApiResults{}
		////------- UseCases
		//	ucSports  = Sports3.UcSports{}
		ucEvents = Events3.UcEvents{}
		//	ucFactors = Factors3.UcFactors{}
		//	ucResults = Results3.UcResults{}
	)
	Create.DBStructure(db, logger)
	for {
		logger.Println("|-------Start-------| Time:", time.Now().Format(time.RFC3339))
		fonUrl.JsonToStruct(urls, logger)

		//apiSports.JsonToStruct(urlevents, logger)
		//ucSports.ReAssign(apiSports)

		apiEvents.Parse(&fonUrl, logger)
		ucEvents.ReAssign(apiEvents)
		dbEvents := ucEvents.CreateDbVar(logger)
		dbEvents.Insert(db, logger)

		//apiFactors.JsonToStruct(urlevents, logger)
		//ucFactors.ReAssign(apiFactors)
		//apiResults.JsonToStruct(urlresult, logger)
		//
		//ucResults.ReAssign(apiResults.Fonbet, logger)

		logger.Println("|-------Done-------| Time:", time.Now().Format(time.RFC3339))
		time.Sleep(5 * time.Minute)
	}
	db.Close()

}
