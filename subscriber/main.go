package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
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
	Message   SensorData
	Timestamp time.Time
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// rabbitmqからキューを受け取りelasticsearchにデータを投入するsub
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"remo", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		var p Protocol
		for d := range msgs {
			err := json.Unmarshal([]byte(d.Body), &p)
			if err != nil {
				failOnError(err, "Failed to unmarshal json")
			}
			log.Printf("Received a message: %v", p.Message)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
