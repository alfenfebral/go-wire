package pkg_logger

import (
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Error(err error)
}

type LoggerImpl struct{}

func NewLogger() Logger {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_URL"),
	})
	if err != nil {
		logrus.Error(err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	return &LoggerImpl{}
}

func (logger *LoggerImpl) Error(err error) {
	if os.Getenv("ENABLE_SENTRY_LOG") == "true" {
		sentry.CaptureException(err)
	}
	logrus.Error(err)
}
