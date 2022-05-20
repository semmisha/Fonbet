package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

func Logger() (logger *logrus.Logger) {
	logger = logrus.New()
	logger.ReportCaller = true

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier: func(frm *runtime.Frame) (function string, file string) {
			file = path.Base(frm.File)
			return fmt.Sprintf("%s", frm.Function), fmt.Sprintf("file:%v , line:%v", file, frm.Line)

		},
	})
	logger.Out = os.Stdout
	return
}
