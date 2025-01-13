package main

import (
	"time"

	"github.com/moondance-labs/tanssi-bridge-relayer/cmd"
	log "github.com/sirupsen/logrus"
)

func configureLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: log.FieldMap{
			log.FieldKeyTime: "@timestamp",
			log.FieldKeyMsg:  "message",
		},
	})
}

func main() {
	configureLogger()
	cmd.Execute()
}
