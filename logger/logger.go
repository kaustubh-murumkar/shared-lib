package logger

import (
	"io"
	"io/fs"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	osOpenFile = os.OpenFile
)

func Init(filePath string, isProduction bool) (*logrus.Logger, error) {
	logFile, err := osOpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, fs.FileMode(644))
	if err != nil {
		return nil, err
	}

	logWriter := io.MultiWriter(logFile, os.Stdout)

	var formatter logrus.Formatter = &logrus.JSONFormatter{}
	level := logrus.DebugLevel
	if isProduction {
		formatter = &logrus.TextFormatter{}
		level = logrus.ErrorLevel
	}
	logger := &logrus.Logger{
		Out:       logWriter,
		Formatter: formatter,
		Level:     level,
	}

	return logger, nil
}
