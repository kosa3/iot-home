package main

import (
	"log"
	"time"
)

type Protocol struct {
	Message   SensorData
	Timestamp time.Time
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
