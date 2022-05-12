package Postgres

import (
	fonstruct "Fonbet/json"
	"Fonbet/utils"
	"github.com/sirupsen/logrus"
	"time"

	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func Sport(fonbet *fonstruct.FonbetEvents, db *sql.DB, logger *logrus.Logger) (err error) {
	var sum, count = 0, 0
	for i := 0; i < len(fonbet.Sports); i++ {
		exist, err := db.Query(`SELECT coalesce((sum(CASE WHEN $1 IN ("sportid") THEN 1 ELSE 0 END)),0) FROM sports ;`, fonbet.Sports[i].Id)
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
			_, err := db.Exec("INSERT INTO sports VALUES ($1, $2, $3)", fonbet.Sports[i].Id, fonbet.Sports[i].ParentId, fonbet.Sports[i].Name)
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

func Events(fonbet *fonstruct.FonbetEvents, db *sql.DB, logger *logrus.Logger) (err error) {
	var levels = 0

	utils.CheckLevels(fonbet, &levels)
	utils.CreateLevels(db, &levels)

	var sum, count = 0, 0
	for i := 0; i < len(fonbet.Events); i++ {
		query := fmt.Sprintf("SELECT coalesce((sum(CASE WHEN $1 IN (id) THEN 1 ELSE 0 END) ),0) FROM events_level_%v ;", fonbet.Events[i].Level)
		exist, err := db.Query(query, fonbet.Events[i].Id)
		if err != nil {
			log.Println(err)
		}
		for exist.Next() {
			err := exist.Scan(&sum)
			if err != nil {
				fmt.Println(err)
			}

		}
		var fontime = time.Now().Add(2 * time.Hour).Unix()
		if sum == 0 && fonbet.Events[i].StartTime >= fontime && fonbet.Events[i].Team1Id != 0 && fonbet.Events[i].Team2Id != 0 {
			query := fmt.Sprintf("INSERT INTO events_level_%v VALUES ($1, $2, $3,$4,$5,$6,$7)", fonbet.Events[i].Level)
			_, err := db.Exec(query, fonbet.Events[i].Id, fonbet.Events[i].ParentId, fonbet.Events[i].Name, fonbet.Events[i].SportId, fonbet.Events[i].Team1, fonbet.Events[i].Team2, fonbet.Events[i].StartTime)
			j := &count
			*j++
			if err != nil {
				fmt.Println(err)
			}
			for _, b := range fonbet.CustomFactors {
				for _, c := range b.Factors {
					if b.E == fonbet.Events[i].Id && (c.F == 921 || c.F == 922 || c.F == 923) {
						query := fmt.Sprintf(`UPDATE events_level_1  set "%v" = %v where id = $1`, c.F, c.V)
						_, err := db.Exec(query, fonbet.Events[i].Id)
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

//func Factor(fonbet *fonstruct.FonbetEvents, db *sql.DB, logger *logrus.Logger) (err error) {
//	var sum, count int = 0, 0
//
//	for i := 0; i < len(fonbet.CustomFactors); i++ {
//		exist, err := db.Query(`SELECT coalesce ((sum(CASE WHEN $1 IN ("eventid") THEN 1 ELSE 0 END)),0) FROM factors ;`, fonbet.CustomFactors[i].E)
//		if err != nil {
//			log.Println(err)
//		}
//		for exist.Next() {
//			err := exist.Scan(&sum)
//			if err != nil {
//				fmt.Println(err)
//			}
//			//fmt.Println(sum)
//		}
//
//		if sum == 0 {
//
//			for b := 0; b < len(fonbet.CustomFactors[i].Factors); b++ {
//				_, err := db.Exec("INSERT INTO factors VALUES ($1, $2, $3,$4,$5)", fonbet.CustomFactors[i].E, fonbet.CustomFactors[i].Factors[b].F, fonbet.CustomFactors[i].Factors[b].V, fonbet.CustomFactors[i].Factors[b].P, fonbet.CustomFactors[i].Factors[b].Pt)
//				j := &count
//				*j++
//				if err != nil {
//					fmt.Println(err)
//				}
//			}
//		}
//
//	}
//	logger.Infof("New Factor rows: %v", count)
//	return
//}
