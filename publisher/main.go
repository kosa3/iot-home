package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/tenntenn/natureremo"
	"log"
	"os"
	"time"
)

func main() {
	cli := natureremo.NewClient(os.Getenv("NATURE_ACCESS_TOKEN"))
	ctx := context.Background()
	devices, err := cli.DeviceService.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

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

	for _, device := range devices {
		fmt.Println("Temperature:", device.NewestEvents[natureremo.SensorTypeTemperature].Value, "°C")
		fmt.Println("Humidity:", device.NewestEvents[natureremo.SensorTypeHumidity].Value, "%")
		fmt.Println("illumination:", device.NewestEvents[natureremo.SensortypeIllumination].Value)

		p := &Protocol{
			SensorData: SensorData{
				DeviceId: device.ID,
				Temp:     device.NewestEvents[natureremo.SensorTypeTemperature].Value,
				Lux:      device.NewestEvents[natureremo.SensortypeIllumination].Value,
				Humidity: device.NewestEvents[natureremo.SensorTypeHumidity].Value,
			},
			Timestamp: time.Now(),
		}

		body, err := json.Marshal(p)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
		}

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
	}
}
