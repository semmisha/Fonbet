package Events

import (
	"Fonbet/controllers/api/Events"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type DbEvents struct {
	Fonbet Events.EventStruct
	Db     *pgxpool.Pool
}

func (f *DbEvents) Insert(logger *logrus.Logger) (err error) {

	var (
		fonbet = f.Fonbet
		exist  = true
		count  = 0
		db     = f.Db
	)
	conn, err := db.Acquire(context.Background())
	if err != nil {
		logger.Errorf("Failed to Acauire connetcion, err: %v\n", err)
	}

	for i := 0; i < len(fonbet.Events); i++ {
		query := fmt.Sprint("SELECT EXISTS(select id from events where id = $1);")
		_ = conn.QueryRow(context.Background(), query, fonbet.Events[i].Id).Scan(&exist)

		query2 := fmt.Sprintf("INSERT INTO events (id, sportid, team1id, team2id, team1, team2, starttime) VALUES ($1, $2, $3,$4,$5,$6,$7)")
		_, err := conn.Exec(context.Background(), query2, fonbet.Events[i].Id, fonbet.Events[i].SportId, fonbet.Events[i].Team1Id, fonbet.Events[i].Team2Id, fonbet.Events[i].Team1, fonbet.Events[i].Team2, fonbet.Events[i].StartTime)

		if err != nil {
			logger.Warningf("Unable to Insert: %v error: %v\n", fonbet.Events[i].Id, err)

		} else {
			j := &count
			*j++
		}

	}
	defer conn.Release()
	logger.Infof("Total Events rows in JSON: %v New Events rows: %v\n", len(fonbet.Events), count)
	return
}

func (f *DbEvents) Update(logger *logrus.Logger) {
	//TODO implement me
	panic("implement me")
}

func (f *DbEvents) Delete() {
	//TODO implement me
	panic("implement me")
}
