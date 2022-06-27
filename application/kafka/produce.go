package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/pellizzetti/fsfc-simulator/application/route"
	"github.com/pellizzetti/fsfc-simulator/infra/kafka"
	kafkago "github.com/segmentio/kafka-go"
)

func Produce(msg *kafkago.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJSONPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KAFKA_PRODUCE_TOPIC"), producer)
		time.Sleep(time.Millisecond * 300)
	}
}
