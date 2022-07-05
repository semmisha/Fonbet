package dbConnect

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type DBClient struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func NewDBClient() *DBClient {
	return &DBClient{}
}

func Connect(db *DBClient, logger *logrus.Logger) (pool *pgxpool.Pool) {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", db.User, db.Password, db.Host, db.Port, db.Dbname)

	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		logger.Fatalf("\nCant Connect to DB by pool, error:%v\n", err)
	}

	return

}
