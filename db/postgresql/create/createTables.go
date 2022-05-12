package create

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func DBStructure(db *pgxpool.Pool, logger *logrus.Logger) {

	//Check if Db structure is created
	sportid, err := db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS sports (sportid INT PRIMARY KEY not null, parentid int, name VARCHAR(150))")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant create table sports",
			"query reply": sportid.String(),
		}).Error(err)

	}

	factors, err := db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS factors (eventid INT not null, factor int, bet float, p int, pt text, Constraint factors_id Primary Key (eventid, factor) ,foreign key (eventid) references events(id) )")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"message":     "cant create table sports",
			"query reply": factors.String(),
		}).Error(err)

	}

}
