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
	}
	query, _ := db.Query(context.Background(), "Select id, team1,team2, starttime from events where result_bool = false")
	var tempslice []temp
	for query.Next() {
		var tempstruct temp
		b := &tempslice
		if err := query.Scan(&tempstruct.id, &tempstruct.team1, &tempstruct.team2, &tempstruct.starttime); err != nil {
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
				_, err := db.Exec(context.Background(), "UPDATE events set result = $1, team1result = $2, team2result = $3 result_bool = true where id = $4", result.Events[j].Score, result.Events[j].Team1Score, result.Events[j].Team2Score, i.id)

				if err != nil {
					logger.Warningf("Cant update result: %v  in ID: %v   error: %v", result.Events[j].Score, i.id, err)

				}
			}

		}

	}
	logger.Infof("New copmare entries: %v", count)
}

func CompareFactor(event *fonstruct.FonbetEvents, db *pgxpool.Pool, logger *logrus.Logger) {
	type temp struct {
		id     int
		factor int
		bet    float32
	}
	query, _ := db.Query(context.Background(), "Select eventid, factor,bet from factors where factor_bool = false")
	var tempslice []temp
	for query.Next() {
		var tempstruct temp
		b := &tempslice
		if err := query.Scan(&tempstruct.id, &tempstruct.factor, &tempstruct.bet); err != nil {
			fmt.Println(err)
		}

		//fmt.Println(tempstruct)
		*b = append(*b, tempstruct)

	}

	var count = 0
	for _, i := range tempslice {

		for j := 0; j < len(event.Events); j++ {

			if i.id == event.Events[j].Id && i.factor >= 921 && i.factor <= 923 {
				query := fmt.Sprintf(`UPDATE events set "%v" = %v where id = %v`, i.factor, i.bet, i.id)
				fmt.Println(query)
				_, err := db.Exec(context.Background(), query)
				if err != nil {
					logger.Warningf("Cant update result: %v  in ID: %v   error: %v", i.factor, i.id, err)
				}
				b := &count
				*b++
			}

		}
		logger.Infof("New copmare entries: %v", count)

	}
}
