package common

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
	Info:
		Log.Info("Info")

	Debug:
		Log.Debug("Debug")

	Error:
		Log.Error("Error")

	Warn:
		Log.Warn("Warn")
*/
var Log = &logrus.Logger{
	Out:       os.Stdout,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

func init() {
	Log.SetReportCaller(true)
}
