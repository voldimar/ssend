package main

import (
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.SetLevel(logrus.InfoLevel)

	godotenv.Load(os.Getenv("HOME") + "/.ssend")
	godotenv.Load()
	log.Info(os.Getenv("SENTRY_DSN"), os.Getenv("HOME"))
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   os.Getenv("SENTRY_DSN"),
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage(strings.Join(os.Args[1:], " "))
}
