package Postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Database interface {
	Insert(db *pgxpool.Pool, logger *logrus.Logger)
	Select(db *pgxpool.Pool, logger *logrus.Logger)
	Delete()
}
