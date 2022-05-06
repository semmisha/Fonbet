package Postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type DBClient struct {
	user     string
	password string
	host     string
	port     string
	dbname   string
}

func Connect(db *DBClient) (pool *pgxpool.Pool) {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", db.user, db.password, db.host, db.port, db.dbname)
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Cant connect to DB, error: ", err)
	}

	return

}
