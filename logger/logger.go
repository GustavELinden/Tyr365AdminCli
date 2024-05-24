package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
    Log = logrus.New()
    Log.SetFormatter(&logrus.JSONFormatter{})
    Log.SetOutput(os.Stdout)
    Log.SetLevel(logrus.InfoLevel)
}

func SetupLogging(logFile string) {
    if logFile != "" {
        file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
            Log.WithError(err).Info("Failed to log to file, using default stderr")
            Log.SetOutput(os.Stderr)
        } else {
            Log.SetOutput(io.MultiWriter(os.Stdout, file))
        }
    } else {
        Log.SetOutput(os.Stdout)
    }
}