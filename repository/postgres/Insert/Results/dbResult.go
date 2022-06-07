package DbResults

import (
	UcResults "Fonbet/usecases/Results"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type DbResults UcResults.UcResults

func (f *DbResults) Insert(db *pgxpool.Pool, logger *logrus.Logger) (err error) {

	var (
		fonbet = f.UcResultsStruct

		exist = true
		count = 0
	)
	conn, err := db.Acquire(context.Background())

	for i := 0; i < len(fonbet); i++ {

		query := fmt.Sprintf(`SELECT EXISTS(Select stringname from results where stringname='%v' and starttime = %v and sportid= %v);`, fonbet[i].Name, fonbet[i].StartTime, fonbet[i].SportId)

		_ = conn.QueryRow(context.Background(), query).Scan(&exist)

		//TODO move to use cases	firstTeam, secondTeam, ok := f.Split(i, logger)

		if exist != true {
			_, err = conn.Exec(context.Background(), "INSERT INTO results (stringname, starttime, score,team1score,team2score,sportid) VALUES ($1, $2, $3, $4, $5,$6)", fonbet[i].Name, fonbet[i].StartTime, fonbet[i].TotalScore, fonbet[i].Team1Score, fonbet[i].Team2Score, fonbet[i].SportId)

			if err != nil {
				logger.Warningf("Unable to Insert into Results: %v exist:%v  error:%v\n", fonbet[i].Id, exist, err)
			} else {
				j := &count
				*j++
			}
		}

	}

	logger.Infof("Total Result rows in JSON:%v New Result rows: %v\n", len(fonbet), count)
	defer conn.Release()
	return
}

func (f DbResults) Update() {
	//TODO implement me
	panic("implement me")
}

func (f DbResults) Delete() {
	//TODO implement me
	panic("implement me")
}
