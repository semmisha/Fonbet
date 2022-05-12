package connect

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
)

type DBClient struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func Connect(db *DBClient, logger *logrus.Logger) (pool *pgxpool.Pool) {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", db.User, db.Password, db.Host, db.Port, db.Dbname)
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Cant connect to DB, error: ", err)
	}

	return

}
