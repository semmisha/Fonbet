package Postgres

import (
	UcFactors "Fonbet/usecases/Convert"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type DbFactors UcFactors.UcFactors

func (f *DbFactors) Insert(db *pgxpool.Pool, logger *logrus.Logger) (err error) {

	var (
		fonbet = f.UcFactorsStruct
		exist  = false
		exist2 = true
		count  = 0
	)

	conn, err := db.Acquire(context.Background())
	if err != nil {
		logger.Errorf("Failed to Acauire connetcion, err: %v\n", err)
	}

	for i := 0; i < len(fonbet); i++ {
		_ = conn.QueryRow(context.Background(), `SELECT exists(Select id from events where id = $1);`, fonbet[i].Id).Scan(&exist)
		_ = conn.QueryRow(context.Background(), `SELECT exists(Select eventid from factors where eventid = $1);`, fonbet[i].Id).Scan(&exist2)
		if exist == true && exist2 == false {

			query := fmt.Sprintf(`INSERT INTO factors (eventid, "921", "922", "923") Values ( %v, %v , %v, %v )`, fonbet[i].Id, fonbet[i].FrstWn, fonbet[i].Drw, fonbet[i].ScndWn)
			_, err = conn.Exec(context.Background(), query)
			if err != nil {
				logger.Warningf("Unable to Manipulate factors: %v error: %v", fonbet, err)
			} else {
				j := &count
				*j++
			}
		}

	}

	logger.Infof("Total Factors row in JSON:%v New Factor rows: %v\n", len(fonbet), count)
	defer conn.Release()
	return

}

func (f *DbFactors) Select(db *pgxpool.Pool, logger logrus.Logger) {
	//TODO implement me
	panic("implement me")
}

func (f *DbFactors) Delete() {
	//TODO implement me
	panic("implement me")
}
