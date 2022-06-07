package api

import "github.com/sirupsen/logrus"

type Api interface {
	JsonToStruct(string, *logrus.Logger) error
}
