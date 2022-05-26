package postgres

import (
	"github.com/sirupsen/logrus"
)

type Database interface {
	Insert(logger *logrus.Logger) error
	Update(logger *logrus.Logger)
	Delete()
}

type DbMaintenance interface {
	Connect()
	CreateTables()
	Close()
}
