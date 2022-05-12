package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger() (logger *logrus.Logger) {
	logger = logrus.New()
	logger.ReportCaller = true

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  time.Kitchen,
		DisableTimestamp: false,
		DataKey:          "",
		FieldMap:         nil,
		CallerPrettyfier: nil,
		PrettyPrint:      false,
	})
	logger.Out = os.Stdout
	return
}
