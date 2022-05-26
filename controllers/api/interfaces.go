package api

import "github.com/sirupsen/logrus"

type Api interface {
	Parse(string, *logrus.Logger) error
}
