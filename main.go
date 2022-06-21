package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	akafka "github.com/pellizzetti/fsfc-simulator/application/kafka"
	"github.com/pellizzetti/fsfc-simulator/infra/kafka"
	kafkago "github.com/segmentio/kafka-go"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}
}

func main() {
	msgChan := make(chan *kafkago.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go akafka.Produce(msg)
	}
}
