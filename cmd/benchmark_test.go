package main

import (
	"Fonbet/controllers/api"
	"github.com/sirupsen/logrus"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	var logger *logrus.Logger
	var bList api.ListURLStruct
	bList.JsonToStruct("https://www.fon.bet/urls.json", logger)

}
