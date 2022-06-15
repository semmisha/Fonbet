package dbCreate

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func DBStructure(db *pgxpool.Pool, logger *logrus.Logger) func() {

	//Check if Db structure is created
	sportid, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS sports (
	sportid int4 NOT NULL,
	parentid int4 NULL,
	"name" text NULL,
	CONSTRAINT sports_pkey PRIMARY KEY (sportid)
);
`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant Create table sports",
			"query reply": sportid.String(),
		}).Error(err)

	}

	events, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS events (
	id int NOT NULL,
	parentid int NULL,
	eventname varchar(150) NULL,
	sportid int NOT NULL,
	team1id int NOT NULL,
    team2id int NOT NULL,
    team1 varchar(50) NULL,
	team2 varchar(50) NULL,
	starttime timestamp NOT NULL,
    FOREIGN KEY (sportid) References sports (sportid),
	CONSTRAINT events_pkey PRIMARY KEY (id)
);
`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant Create table events",
			"query reply": events.String(),
		}).Error(err)

	}

	results, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS results (
	resultid  serial ,
    eventid int NULL UNIQUE,
    sportid int NOT NULL, 
    stringname varchar(100) NOT NULL,
	team1score int NULL,
	team2score int NULL,
	starttime timestamp NOT NULL,
	score varchar(50) NOT NULL,
    FOREIGN KEY (eventid) References events (id),
    CONSTRAINT result_constraint PRIMARY KEY (resultid, stringname, starttime, sportid)
);
`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant Create table results",
			"query reply": results.String(),
		}).Error(err)

	}

	factors, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS factors (
	eventid int NOT NULL UNIQUE,
    "921" float8 NULL,
	"922" float8 NULL,
	"923" float8 NULL,
    FOREIGN KEY (eventid) References events (id)
	
);
`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant Create table results",
			"query reply": factors.String(),
		}).Error(err)

	}

	view, err := db.Exec(context.Background(), `
CREATE OR REPLACE VIEW public.orange
AS SELECT e.id,
    e.sportid,
    s.name,
    e.team1id,
    e.team2id,
    e.team1,
    e.team2,
    e.starttime,
    f."921",
    f."922",
    f."923",
    r.team1score,
    r.team2score,
    r.score
   FROM events e
     JOIN results r ON e.id = r.eventid
     JOIN factors f ON e.id = f.eventid
     LEFT JOIN sports s ON e.sportid = s.sportid;
`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant Create view",
			"query reply": view.String(),
		}).Error(err)

	}

	viewpredict, err := db.Exec(context.Background(), `-- public.orangepredict source

CREATE OR REPLACE VIEW public.orangepredict
AS SELECT e.id,
    e.sportid,
    s.name,
    e.team1id,
    e.team2id,
    e.team1,
    e.team2,
    e.starttime,
    f."921",
    f."922",
    f."923"
   FROM events e
     LEFT JOIN factors f ON e.id = f.eventid
     LEFT JOIN sports s ON e.sportid = s.sportid
  WHERE e.starttime > CURRENT_TIMESTAMP;
`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant Create viewpredict",
			"query reply": viewpredict.String(),
		}).Error(err)

	}

	return nil
}
