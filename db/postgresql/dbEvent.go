package Postgres

import (
	fonstruct "Fonbet/json"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

func Events(fonbet *fonstruct.FonbetEvents, db *pgxpool.Pool, logger *logrus.Logger) (err error) {

	var sum, count = 0, 0
	for i := 0; i < len(fonbet.Events); i++ {
		query := fmt.Sprintf("SELECT coalesce((sum(CASE WHEN $1 IN (id) THEN 1 ELSE 0 END) ),0) FROM events;")
		exist, err := db.Query(context.Background(), query, fonbet.Events[i].Id)
		if err != nil {
			log.Println(err)
		}
		for exist.Next() {
			err := exist.Scan(&sum)
			if err != nil {
				fmt.Println(err)
			}

		}
		//var fontime = time.Now().Add(7 * time.Hour).Unix()
		if sum == 0 && fonbet.Events[i].Team1Id != 0 && fonbet.Events[i].Team2Id != 0 && fonbet.Events[i].Level == 1 {
			query := fmt.Sprintf("INSERT INTO events (id, sportid, team1id, team2id, team1, team2, starttime) VALUES ($1, $2, $3,$4,$5,$6,$7)")
			_, err := db.Exec(context.Background(), query, fonbet.Events[i].Id, fonbet.Events[i].SportId, fonbet.Events[i].Team1Id, fonbet.Events[i].Team2Id, fonbet.Events[i].Team1, fonbet.Events[i].Team2, fonbet.Events[i].StartTime)
			j := &count
			*j++
			if err != nil {
				fmt.Println(err)
			}
			for _, b := range fonbet.CustomFactors {
				for _, c := range b.Factors {
					if b.E == fonbet.Events[i].Id && (c.F == 921 || c.F == 922 || c.F == 923) {
						query := fmt.Sprintf(`UPDATE events  set "%v" = %v where id = $1`, c.F, c.V)
						_, err := db.Exec(context.Background(), query, fonbet.Events[i].Id)
						if err != nil {
							fmt.Println(err)
						}

					}

				}
			}
		}
	}

	logger.Infof("New Events rows: %v", count)
	return
}
