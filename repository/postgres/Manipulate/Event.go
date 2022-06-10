package Postgres

import (
	UcEvents "Fonbet/usecases/Convert"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type DbEvents UcEvents.UcEvents

func (f *DbEvents) Insert(db *pgxpool.Pool, logger *logrus.Logger) (err error) {

	var (
		fonbet = f.UcEventStruct
		exist  = true
		count  = 0
	)

	conn, err := db.Acquire(context.Background())
	if err != nil {
		logger.Errorf("Failed to Acquire connetcion, err: %v\n", err)
	}

	for i := 0; i < len(fonbet); i++ {
		query := fmt.Sprint("SELECT EXISTS(select id from events where id = $1);")
		_ = conn.QueryRow(context.Background(), query, fonbet[i].Id).Scan(&exist)
		if exist != true {
			query2 := fmt.Sprintf("INSERT INTO events (id, sportid, team1id, team2id, team1, team2, starttime,eventname) VALUES ($1, $2, $3,$4,$5,$6,$7,$8)")
			_, err := conn.Exec(context.Background(), query2, fonbet[i].Id, fonbet[i].SportId, fonbet[i].Team1Id, fonbet[i].Team2Id, fonbet[i].Team1, fonbet[i].Team2, fonbet[i].StartTime, fonbet[i].Name)

			if err != nil {
				logger.Warningf("Unable to Manipulate: %v error: %v\n", fonbet[i].Id, err)

			} else {
				j := &count
				*j++
			}

		}
	}
	defer conn.Release()
	logger.Infof("Total Events rows in JSON: %v New Events rows: %v\n", len(fonbet), count)
	return
}

func (f *DbEvents) Select(db *pgxpool.Pool, logger *logrus.Logger) {
	var (
		fonbet UcEvents.Event
		count  = 0
	)
	conn, err := db.Acquire(context.Background())
	if err != nil {
		logger.Errorf("Failed to Acquire connetcion, err: %v\n", err)
	}
	data, _ := conn.Query(context.Background(), "Select id, team1, team2, starttime, sportid from events ")

	for data.Next() {
		err = data.Scan(&fonbet.Id, &fonbet.Team1, &fonbet.Team2, &fonbet.StartTime, &fonbet.SportId)
		if err != nil {
			logger.Errorf("Unable top select Events from DB, error:%v", err)
		} else {
			f.UcEventStruct = append(f.UcEventStruct, fonbet)
			count++

		}
	}

	defer conn.Release()
	logger.Infof("Total Events from DB: %v\n", count)
	return
}

func (f *DbEvents) Delete() {
	//TODO implement me
	panic("implement me")
}
