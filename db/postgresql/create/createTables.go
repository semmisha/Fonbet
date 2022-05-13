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
	sportid int4 NULL,
	team1 varchar(50) NULL,
	team2 varchar(50) NULL,
	starttime int4 NULL,
	result_bool bool NULL DEFAULT false,
	"result" varchar(50) NULL,
	"team1result" int NULL,
	"team2result" int NULL,
    "921" float8 NULL,
	"922" float8 NULL,
	"923" float8 NULL,
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
	stringname varchar(100) NOT NULL,
	team1 varchar(100) NULL,
	team2 varchar(100) NULL,
	starttime int4 NOT NULL,
	score varchar(50) NOT NULL,
	CONSTRAINT result_constraint PRIMARY KEY (stringname, starttime)
);



`)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant create table results",
			"query reply": results.String(),
		}).Error(err)

	}

	return nil
}
