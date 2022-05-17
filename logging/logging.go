package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Logger() (logger *logrus.Logger) {
	logger = logrus.New()
	logger.ReportCaller = true

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          true,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})
	logger.Out = os.Stdout
	return
}
