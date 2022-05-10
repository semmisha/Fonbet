package connect

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func DbConnect2(logger *logrus.Logger) *sql.DB {

	//fmt.Sprintf("postgres://%s:%s")
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=P@ssw0rd dbname=postgres sslmode=disable")
	if err != nil {
		logger.Fatalf("Unable to connect to DB, err:", err)

	}

	return db
}
