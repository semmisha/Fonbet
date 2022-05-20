package Postgres

import (
	"Fonbet/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func CompareResult(db *pgxpool.Pool, logger *logrus.Logger) {
	type tempEvent struct {
		id        int
		team1     string
		team2     string
		starttime int64
		sportid   int64
	}
	queryEvent, _ := db.Query(context.Background(), "Select id, team1,team2, starttime,sportid from events")
	var eventSlice []tempEvent
	for queryEvent.Next() {
		var eventStruct tempEvent
		b := &eventSlice
		if err := queryEvent.Scan(&eventStruct.id, &eventStruct.team1, &eventStruct.team2, &eventStruct.starttime, &eventStruct.sportid); err != nil {
			fmt.Println(err)
		}

		//	fmt.Println(eventStruct)
		*b = append(*b, eventStruct)

	}
	//fmt.Println(eventSlice)
	type tempResult struct {
		eventid    int
		stringname string
		starttime  int64
		sportid    int64
	}
	queryResult, _ := db.Query(context.Background(), "Select stringname, starttime,sportid from results where eventid is null ")
	var resultSlice []tempResult
	for queryResult.Next() {
		var resultStruct tempResult
		b := &resultSlice
		if err := queryResult.Scan(&resultStruct.stringname, &resultStruct.starttime, &resultStruct.sportid); err != nil {
			fmt.Println(err)
		}

		//	fmt.Println(resultStruct)
		*b = append(*b, resultStruct)

	}

	var count = 0
	for _, i := range resultSlice {
		resultString := utils.Replacer(i.stringname)
		for _, j := range eventSlice {
			eventString := utils.Replacer(fmt.Sprintf("%v%v", j.team1, j.team2))
			//fmt.Println(resultString, eventString)
			if resultString == eventString &&
				i.starttime == j.starttime &&
				i.sportid == j.sportid {
				b := &count
				*b++

				_, err := db.Exec(context.Background(), "UPDATE results set eventid = $1 where stringname = $2 and starttime = $3 and sportid = $4", j.id, i.stringname, i.starttime, i.sportid)

				if err != nil {
					logger.Warningf("Cant update result: %v  in ID: %v   error: %v", i.stringname, j.id, err)

				}
				break
			}

		}

	}
	logger.Infof("New copmare entries: %v", count)
}
