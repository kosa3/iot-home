package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/streadway/amqp"
	"log"
)

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

	ctx := context.Background()
	es, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
	)
	failOnError(err, "Failed to connect elasticsearch")
	defer es.Stop()

	forever := make(chan bool)

	go func() {
		var p Protocol
		for d := range msgs {
			err := json.Unmarshal([]byte(d.Body), &p)
			if err != nil {
				failOnError(err, "Failed to unmarshal json")
			}
			_, err = es.Index().
				Index("natureremo").
				BodyJson(p).
				Do(ctx)

			if err != nil {
				failOnError(err, "Failed to post data")
			}

			log.Printf("Received a message: %v", p.SensorData)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
