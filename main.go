package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pellizzetti/fsfc-simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	topic := os.Getenv("KAFKA_READ_TOPIC")
	kafka.Publish("ola", topic, producer)
	for {
		_ = 1
	}
	// route := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// route.LoadPositions()
	// stringJSON, _ := route.ExportJSONPositions()
	// fmt.Println(stringJSON[1])
}
