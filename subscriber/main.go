package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// rabbitmqからキューを受け取りelasticsearchにデータを投入するsub
	conn, err := amqp.Dial("amqp://guest:guest@" + os.Getenv("RABBITMQ_ENDPOINT"))
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

	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ES_ENDPOINT"),
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	failOnError(err, "Failed to connect elasticsearch")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			req := esapi.IndexRequest{
				Index:   "natureremo",
				Body:    strings.NewReader(string(d.Body)),
				Refresh: "true",
			}
			res, err := req.Do(context.Background(), es)
			failOnError(err, "Error getting response")
			defer res.Body.Close()

			log.Printf("Received a message: %v", string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
