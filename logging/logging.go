package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

func Logger() (logger *logrus.Logger) {
	logger = logrus.New()
	logger.ReportCaller = true

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier: func(frm *runtime.Frame) (function string, file string) {
			file = path.Base(frm.File)
			return frm.Function, fmt.Sprintf("file:%v , line:%v", file, frm.Line)
		},
		PrettyPrint: false,
	})
	file, err := os.OpenFile("/app/logging/logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		logger.Errorf("Unable to Create/open file: %v", file.Name())
	}

	logger.Out = io.MultiWriter(os.Stdout, file)
	return
}
