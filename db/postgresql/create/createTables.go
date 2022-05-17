package create

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
			"message":     "cant create table sports",
			"query reply": sportid.String(),
		}).Error(err)

	}

	events, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS events (
	id int4 NOT NULL,
	parentid int4 NULL,
	eventname varchar(150) NULL,
	sportid int4 NOT NULL,
	team1id int4 NOT NULL,
    team2id int4 NOT NULL,
    team1 varchar(50) NULL,
	team2 varchar(50) NULL,
	starttime int4 NOT NULL,
    FOREIGN KEY (sportid) References sports (sportid),
	CONSTRAINT events_pkey PRIMARY KEY (id)
);
`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant create table events",
			"query reply": events.String(),
		}).Error(err)

	}

	results, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS results (
	eventid int4 NULL UNIQUE,
    sportid int NOT NULL, 
    stringname varchar(100) NOT NULL,
	team1 int NULL,
	team2 int NULL,
	starttime int4 NOT NULL,
	score varchar(50) NOT NULL,
    FOREIGN KEY (eventid) References events (id),
	FOREIGN KEY (sportid) References sports (sportid),
    CONSTRAINT result_constraint PRIMARY KEY (stringname, starttime, sportid)
);
`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant create table results",
			"query reply": results.String(),
		}).Error(err)

	}

	factors, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS factors (
	eventid int4 NOT NULL UNIQUE,
    "921" float8 NULL,
	"922" float8 NULL,
	"923" float8 NULL,
    FOREIGN KEY (eventid) References events (id)
	
);
`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant create table results",
			"query reply": factors.String(),
		}).Error(err)

	}

	return nil
}
