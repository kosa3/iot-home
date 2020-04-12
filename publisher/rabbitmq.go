package main

import (
	"log"
	"time"
)

type Protocol struct {
	SensorData
	Timestamp time.Time
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
