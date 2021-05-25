package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   "https://dbe0d18ea06a4c7eac9f06ebf5dfaf45:0a47b933c9c74bbab720537a139140d3@log.gosp.cloud/2",
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage(strings.Join(os.Args[1:], " "))
}
