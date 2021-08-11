package common

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
	Info:
		common.Log.Info("Info")

	Debug:
		common.Log.Debug("Debug")

	Error:
		common.Log.Error("Error")

	Warn:
		common.Log.Warn("Warn")
*/
var Log = &logrus.Logger{
	Out:       os.Stdout,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}
