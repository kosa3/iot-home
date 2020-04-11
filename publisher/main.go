package main

import (
	"context"
	"fmt"
	"github.com/tenntenn/natureremo"
	"log"
	"os"
)

func main() {
	cli := natureremo.NewClient(os.Getenv("NATURE_ACCESS_TOKEN"))
	ctx := context.Background()
	devices, err := cli.DeviceService.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices {
		fmt.Println("Temperature:", device.NewestEvents[natureremo.SensorTypeTemperature].Value, "°C")
		fmt.Println("Humidity:", device.NewestEvents[natureremo.SensorTypeHumidity].Value, "%")
		fmt.Println("illumination:", device.NewestEvents[natureremo.SensortypeIllumination].Value)
	}
}
