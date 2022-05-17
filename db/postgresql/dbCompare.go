package Postgres

import (
	fonstruct "Fonbet/json"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"strings"
)

func CompareResult(result *fonstruct.FonbetResult, db *pgxpool.Pool, logger *logrus.Logger) {
	type temp struct {
		id        int
		team1     string
		team2     string
		starttime int64
		sportid   int64
	}
	query, _ := db.Query(context.Background(), "Select id, team1,team2, starttime,sportid from events")
	var tempslice []temp
	for query.Next() {
		var tempstruct temp
		b := &tempslice
		if err := query.Scan(&tempstruct.id, &tempstruct.team1, &tempstruct.team2, &tempstruct.starttime, &tempstruct.sportid); err != nil {
			fmt.Println(err)
		}

		//	fmt.Println(tempstruct)
		*b = append(*b, tempstruct)

	}

	var count = 0
	for _, i := range tempslice {

		for j := 0; j < len(result.Events); j++ {

			if strings.Contains(result.Events[j].Name, i.team1) &&
				strings.Contains(result.Events[j].Name, i.team2) &&
				result.Events[j].StartTime == i.starttime &&
				result.Events[j].Status == 3 {
				b := &count
				*b++
				_, err := db.Exec(context.Background(), "UPDATE results set eventid = $1 where stringname = $2 and starttime = $3 ", i.id, result.Events[j].Name, result.Events[j].StartTime)

				if err != nil {
					logger.Warningf("Cant update result: %v  in ID: %v   error: %v", result.Events[j].Score, i.id, err)

				}
			}

		}

	}
	logger.Infof("New copmare entries: %v", count)
}
