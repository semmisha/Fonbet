package Postgres

import (
	fonstruct "Fonbet/json"
	"Fonbet/utils"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"strings"
)

//

func Result(fonbet *fonstruct.FonbetResult, db *pgxpool.Pool, logger *logrus.Logger) (err error) {
	var sum, count = 0, 0

	for i := 0; i < len(fonbet.Events); i++ {
		sportid := utils.SearchSportId(fonbet, i, logger)
		exist, err := db.Query(context.Background(), `SELECT coalesce((sum(CASE WHEN $1 IN ("stringname") and $2 in ("starttime") and $3 in ("sportid") THEN 1 ELSE 0 END)),0) FROM results ;`, fonbet.Events[i].Name, fonbet.Events[i].StartTime, sportid)
		if err != nil {
			logger.Errorf("Unable to SELECT sum, error:%v\n ", err)

		}
		for exist.Next() {
			err := exist.Scan(&sum)
			if err != nil {
				logger.Errorf("Unable to SCAN sum, error:%v\n ", err)

			}
		}

		if strings.ContainsAny(fonbet.Events[i].Name, "-––") && fonbet.Events[i].Status == 3 && strings.Contains(fonbet.Events[i].Name, "очковые") != true {

			strarray := strings.Split(fonbet.Events[i].Score, " ")
			resultarray := strings.Split(strarray[0], ":")

			if len(resultarray) >= 2 && resultarray[0] != " " && resultarray[1] != " " {
				if sum == 0 {
					_, err = db.Exec(context.Background(), "INSERT INTO results (stringname, starttime, score,team1,team2,sportid) VALUES ($1, $2, $3, $4, $5,$6)", fonbet.Events[i].Name, fonbet.Events[i].StartTime, fonbet.Events[i].Score, resultarray[0], resultarray[1], sportid)
					j := &count
					*j++
					if err != nil {
						logger.Warningf("Unable to insert into Results: %v sum:%v  error:%v\n", fonbet.Events[i].Name, sum, err)
					}
				}

			}
		}

	}
	logger.Infof("Total Result rows: %v", count)
	return
}
