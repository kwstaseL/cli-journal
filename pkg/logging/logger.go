package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

var Logger = slog.New(slog.NewTextHandler(os.Stderr, nil))

func LogInfo(msg string, args ...interface{}) {
    Logger.Info(msg, args...)
}

func LogError(msg string, args ...interface{}) {
    Logger.Error(msg, args...)
}

func LogDebug(msg string, args ...interface{}) {
	Logger.Debug(msg, args...)
}

func LogWarn(msg string, args ...interface{}) {
	Logger.Warn(msg, args...)
}
