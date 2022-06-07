package Results

import (
	"Fonbet/controllers/api/Results"
	"github.com/jackc/pgx/v4/pgxpool"
)

//
//import (
//	"Fonbet/controllers/api/Results"
//	"context"
//	"fmt"
//	"github.com/jackc/pgx/v4/pgxpool"
//	"github.com/sirupsen/logrus"
//)
//
type DbResults struct {
	Fonbet Results.ApiResults
	Db     *pgxpool.Pool
}

//
//func (f *DbResults) Insert(logger *logrus.Logger) (err error) {
//
//	var (
//		fonbet = f.Fonbet
//		db     = f.Db
//		exist  = true
//		count  = 0
//	)
//	conn, err := db.Acquire(context.Background())
//
//	for i := 0; i < len(fonbet.Events); i++ {
//
//
//		if ok == true {
//			query := fmt.Sprintf(`SELECT EXISTS(Select stringname from results where stringname='%v' and starttime = %v and sportid= %v);`, fonbet.Events[i].Name, fonbet.Events[i].StartTime, sportId)
//
//			_ = conn.QueryRow(context.Background(), query).Scan(&exist)
//
//			//TODO move to use cases	firstTeam, secondTeam, ok := f.Split(i, logger)
//
//			if exist == false && ok == true {
//				_, err = conn.Exec(context.Background(), "INSERT INTO results (stringname, starttime, score,team1,team2,sportid) VALUES ($1, $2, $3, $4, $5,$6)", fonbet.Events[i].Name, fonbet.Events[i].StartTime, fonbet.Events[i].Score, firstTeam, secondTeam, sportId)
//
//				if err != nil {
//					logger.Warningf("Unable to Insert into Results: %v exist:%v  error:%v\n", fonbet.Events[i].Name, exist, err)
//				} else {
//					j := &count
//					*j++
//				}
//			}
//
//		}
//	}
//	logger.Infof("Total Result rows in JSON:%v New Result rows: %v\n", len(fonbet.Events), count)
//	defer conn.Release()
//	return
//}
//
//func (f DbResults) Update() {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (f DbResults) Delete() {
//	//TODO implement me
//	panic("implement me")
//}
