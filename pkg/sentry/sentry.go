package pkg_sentry

import (
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

type Sentry struct{}

func NewSentry() *Sentry {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_URL"),
	})
	if err != nil {
		logrus.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	return &Sentry{}
}
