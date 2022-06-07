package Sports

import (
	"Fonbet/controllers/api/Sports"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type DbSports struct {
	Fonbet Sports.ApiSports
	Db     *pgxpool.Pool
}

func (f DbSports) Insert(logger *logrus.Logger) (err error) {
	var (
		fonbet = f.Fonbet
		db     = f.Db

		exist = true
		count = 0
	)
	conn, err := db.Acquire(context.Background())
	if err != nil {
		logger.Errorf("Failed to Acauire connetcion, err: %v\n", err)
	}

	for i := 0; i < len(fonbet.Sports); i++ {
		_ = db.QueryRow(context.Background(), `SELECT exists(Select sportid from sports where sportid = $1);`, fonbet.Sports[i].Id).Scan(&exist)

		if exist == false {
			_, err := conn.Exec(context.Background(), "INSERT INTO sports (sportid, parentid, name) VALUES ($1, $2, $3)", fonbet.Sports[i].Id, fonbet.Sports[i].ParentId, fonbet.Sports[i].Name)

			if err != nil {
				logger.Warningf("Unable to Insert into Sports: %v exist:%v  error:%v\n", fonbet.Sports[i].Id, exist, err)
			} else {
				j := &count
				*j++
			}

		}
	}
	logger.Infof("Total Sports rows in JSON:%v New Sports rows: %v\n", len(fonbet.Sports), count)
	defer conn.Release()
	return

}

func (f DbSports) Update() {
	//TODO implement me
	panic("implement me")
}

func (f DbSports) Delete() {
	//TODO implement me
	panic("implement me")
}
