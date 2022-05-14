package Postgres

import (
	fonstruct "Fonbet/json"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
)

func Sport(fonbet *fonstruct.FonbetEvents, db *pgxpool.Pool, logger *logrus.Logger) (err error) {
	var sum, count = 0, 0
	for i := 0; i < len(fonbet.Sports); i++ {
		exist, err := db.Query(context.Background(), `SELECT coalesce((sum(CASE WHEN $1 IN ("sportid") THEN 1 ELSE 0 END)),0) FROM sports ;`, fonbet.Sports[i].Id)
		if err != nil {
			log.Println(err)
		}
		for exist.Next() {

			err := exist.Scan(&sum)
			if err != nil {
				logger.Warningf("Unable to scan sum for ID: %v error: %v", fonbet.Sports[i].Id, err)
			}
			//fmt.Println(sum)
		}

		if sum == 0 {
			_, err := db.Exec(context.Background(), "INSERT INTO sports (sportid, parentid, name) VALUES ($1, $2, $3)", fonbet.Sports[i].Id, fonbet.Sports[i].ParentId, fonbet.Sports[i].Name)
			j := &count
			*j++
			if err != nil {
				fmt.Println(err)
			}

		}
	}
	logger.Infof("New Sports rows: %v", count)
	return
}
