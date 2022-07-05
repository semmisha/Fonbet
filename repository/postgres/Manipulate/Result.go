package Postgres

import (
	UcResults "Fonbet/usecases/Convert"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"time"
)

type DbResults UcResults.UcResults

func (f *DbResults) Insert(db *pgxpool.Pool, logger *logrus.Logger) {

	var (
		fonbet = f.UcResultsStruct

		exist = true
		count = 0
	)
	conn, err := db.Acquire(context.Background())
	if err != nil {
		logger.Errorf("Unable to acuire DB connection, error: %v", err)

	}
	for i := 0; i < len(fonbet); i++ {

		query := `SELECT EXISTS(Select stringname from results where stringname=$1 and starttime = $2 and sportid= $3);`

		_ = conn.QueryRow(context.Background(), query, fonbet[i].Name, fonbet[i].StartTime, fonbet[i].SportId).Scan(&exist)

		if !exist {
			_, err = conn.Exec(context.Background(), "INSERT INTO results (stringname, starttime, score,team1score,team2score,sportid) VALUES ($1, $2, $3, $4, $5,$6)", fonbet[i].Name, fonbet[i].StartTime, fonbet[i].TotalScore, fonbet[i].Team1Score, fonbet[i].Team2Score, fonbet[i].SportId)

			if err != nil {
				logger.Warningf("Unable to Manipulate into Results: %v exist:%v  error:%v\n", fonbet[i].ResultId, exist, err)
			} else {
				j := &count
				*j++
			}
		}

	}

	logger.Infof("Total Result rows in JSON:%v New Result rows: %v\n", len(fonbet), count)
	defer conn.Release()

}

func (f *DbResults) Select(db *pgxpool.Pool, logger *logrus.Logger) {
	var (
		fonbet UcResults.Result
		count  = 0
	)
	conn, err := db.Acquire(context.Background())
	if err != nil {
		logger.Errorf("Failed to Acquire connetcion, err: %v\n", err)
	}
	Timer := time.Now().Add(-6 * time.Hour)
	data, _ := conn.Query(context.Background(), "Select resultid, sportid, stringname, starttime from results where eventid is null and starttime > $1", Timer)

	for data.Next() {
		err = data.Scan(&fonbet.ResultId, &fonbet.SportId, &fonbet.Name, &fonbet.StartTime)
		if err != nil {
			logger.Errorf("Unable top select Events from DB, error:%v", err)
		} else {
			f.UcResultsStruct = append(f.UcResultsStruct, fonbet)
			count++

		}
	}

	defer conn.Release()
	logger.Infof("Total Results from DB: %v\n", count)

}

func (f *DbResults) Update(db *pgxpool.Pool, logger *logrus.Logger) {

	var (
		fonbet = f.UcResultsStruct

		count = 0
	)
	conn, err := db.Acquire(context.Background())
	if err != nil {

		logger.Errorf("Unable to Acquire connection:%v\n", err)

	}

	for _, i := range fonbet {
		query := "Update results set eventid = $1 where resultid = $2"
		_, err = conn.Exec(context.Background(), query, i.EventId, i.ResultId)
		if err != nil {
			logger.Warningf("Unable to update result, ResultId:%v EventId:%v error:%v\n", i.ResultId, i.EventId, err)

		} else {
			count++
		}
	}

	logger.Infof("Total Updated Result rows count: %v\n", count)
	defer conn.Release()

}

func (f DbResults) Delete() {
	//TODO implement me
	panic("implement me")
}
