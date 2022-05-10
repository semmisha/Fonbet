package Postgres

import (
	fonstruct "Fonbet/json"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

func CompareResult(result *fonstruct.FonbetResult, db *sql.DB, logger *logrus.Logger) {
	type temp struct {
		id        int
		team1     string
		team2     string
		starttime int64
	}
	query, _ := db.Query("Select id, team1,team2, starttime from events_level_1 where result_bool = false")
	var tempslice []temp
	for query.Next() {
		var tempstruct temp
		b := &tempslice
		if err := query.Scan(&tempstruct.id, &tempstruct.team1, &tempstruct.team2, &tempstruct.starttime); err != nil {
			fmt.Println(err)
		}

		fmt.Println(tempstruct)
		*b = append(*b, tempstruct)

	}

	var count int = 0
	for _, i := range tempslice {

		for j := 0; j < len(result.Events); j++ {

			if strings.Contains(result.Events[j].Name, i.team1) &&
				strings.Contains(result.Events[j].Name, i.team2) &&
				result.Events[j].StartTime == i.starttime &&
				result.Events[j].Status == 3 {
				b := &count
				*b++
				_, err := db.Exec("UPDATE events_level_1 set result = $1, result_bool = true where id = $2", result.Events[j].Score, i.id)

				if err != nil {
					logger.Warningf("Cant update result: %v  in ID: %v   error: %v", result.Events[j].Score, i.id, err)

				}
			}

		}

	}
	logger.Infof("New copmare entries: %v", count)
}

func CompareFactor(db *sql.DB) {

}
