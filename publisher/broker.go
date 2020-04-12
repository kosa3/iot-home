package main

import (
	"log"
	"time"
)

type SensorData struct {
	DeviceId string  `json:"id"`
	Temp     float64 `json:"temp"`
	Lux      float64 `json:"lux"`
	Humidity float64 `json:"humidity"`
}

type Protocol struct {
	SensorData
	Timestamp time.Time
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
