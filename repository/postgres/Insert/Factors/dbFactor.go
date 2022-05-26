package Factors

import (
	Factors2 "Fonbet/controllers/api/Factors"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type DbFactors struct {
	Fonbet Factors2.CustomFactorsStruct
	Db     *pgxpool.Pool
}

func (f *DbFactors) Insert(db *pgxpool.Pool, logger *logrus.Logger) (err error) {

	var (
		fonbet = f.Fonbet
		exist  = false
		exist2 = true
		count  = 0
	)

	conn, err := db.Acquire(context.Background())
	if err != nil {
		logger.Errorf("Failed to Acauire connetcion, err: %v\n", err)
	}

	type Factors struct {
		id        int
		firstwin  float64
		draw      float64
		secondwin float64
	}
	var factor Factors

	for i := 0; i < len(fonbet.CustomFactors); i++ {
		_ = conn.QueryRow(context.Background(), `SELECT exists(Select id from events where id = $1);`, fonbet.CustomFactors[i].E).Scan(&exist)
		_ = conn.QueryRow(context.Background(), `SELECT exists(Select eventid from factors where eventid = $1);`, fonbet.CustomFactors[i].E).Scan(&exist2)
		if exist == true && exist2 == false {

			query := fmt.Sprintf(`INSERT INTO factors (eventid, "921", "922", "923") Values ( %v, %v , %v, %v )`, factor.id, factor.firstwin, factor.draw, factor.secondwin)
			_, err = conn.Exec(context.Background(), query)
			if err != nil {
				logger.Warningf("Unable to Insert factors: %v error: %v", factor.id, err)
			} else {
				j := &count
				*j++
			}
		}

	}

	logger.Infof("Total Factors row in JSON:%v New Factor rows: %v\n", len(fonbet.CustomFactors), count)
	defer conn.Release()
	return

}

func (f *DbFactors) Update() {
	//TODO implement me
	panic("implement me")
}

func (f *DbFactors) Delete() {
	//TODO implement me
	panic("implement me")
}
