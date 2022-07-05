package Postgres

import (
	Sports2 "Fonbet/usecases/Convert"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type DbSports Sports2.UcSports

func (f *DbSports) Insert(db *pgxpool.Pool, logger *logrus.Logger) {
	var (
		fonbet = f.UcSportsStruct
		exist  = true
		count  = 0
	)

	conn, err := db.Acquire(context.Background())
	if err != nil {
		logger.Errorf("Failed to Acauire connetcion, err: %v\n", err)
	}

	for i := 0; i < len(fonbet); i++ {
		_ = db.QueryRow(context.Background(), `SELECT exists(Select sportid from sports where sportid = $1);`, fonbet[i].Id).Scan(&exist)

		if !exist {
			_, err := conn.Exec(context.Background(), "INSERT INTO sports (sportid, parentid, name) VALUES ($1, $2, $3)", fonbet[i].Id, fonbet[i].ParentId, fonbet[i].Name)

			if err != nil {
				logger.Warningf("Unable to Manipulate into Sports: %v exist:%v  error:%v\n", fonbet[i].Id, exist, err)
			} else {
				j := &count
				*j++
			}

		}
	}
	logger.Infof("Total Sports rows in JSON:%v New Sports rows: %v\n", len(fonbet), count)
	defer conn.Release()

}

func (f *DbSports) Select(db *pgxpool.Pool, logger *logrus.Logger) {
	//TODO implement me
	panic("implement me")
}

func (f DbSports) Delete() {
	//TODO implement me
	panic("implement me")
}
