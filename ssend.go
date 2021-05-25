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
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   os.Getenv("SENTRY_DSN"),
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown host"
	}
	log.Info(hostname)
	message := hostname + ": " + strings.Join(os.Args[1:], " ")
	log.Info(message)
	sentry.CaptureMessage(message)
}
