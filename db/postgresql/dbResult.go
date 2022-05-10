package Postgres

import (
	fonstruct "Fonbet/json"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
)

//

func Result(fonbet *fonstruct.FonbetResult, db *sql.DB, logger *logrus.Logger) (err error) {
	var sum, count int = 0, 0

	for i := 0; i < len(fonbet.Events); i++ {
		exist, err := db.Query(`SELECT coalesce((sum(CASE WHEN $1 IN ("stringname") and $2 in ("starttime")THEN 1 ELSE 0 END)),0) FROM results ;`, fonbet.Events[i].Name, fonbet.Events[i].StartTime)
		if err != nil {
			log.Println(err)
		}
		for exist.Next() {
			err := exist.Scan(&sum)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println(sum)
		}

		if sum == 0 && strings.Contains(fonbet.Events[i].Name, "-") || sum == 0 && strings.Contains(fonbet.Events[i].Name, "–") {

			_, err = db.Exec("INSERT INTO results (stringname, starttime, score) VALUES ($1, $2, $3)", fonbet.Events[i].Name, fonbet.Events[i].StartTime, fonbet.Events[i].Score)
			j := &count
			*j++
			if err != nil {
				fmt.Println(err)
			}

		}

	}
	logger.Infof("Total Result rows: %v", count)
	return
}
