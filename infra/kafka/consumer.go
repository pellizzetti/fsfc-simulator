package kafka

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	kafkago "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type KafkaConsumer struct {
	MsgChan chan *kafkago.Message
}

func NewKafkaConsumer(msgChan chan *kafkago.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

func (k *KafkaConsumer) Consume() {
	mechanism, err := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_SASL_USERNAME"), os.Getenv("KAFKA_SASL_PASSWORD"))
	if err != nil {
		log.Fatalln(err)
	}
	dialer := &kafkago.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}
	c := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_BOOTSTRAP_SERVERS")},
		GroupID: os.Getenv("KAFKA_CONSUMER_GROUP_ID"),
		Topic:   os.Getenv("KAFKA_READ_TOPIC"),
		Dialer:  dialer,
	})
	fmt.Println("kafka consumer has been started")
	for {
		msg, err := c.FetchMessage(context.Background())
		if err == nil {
			k.MsgChan <- &msg
		}
	}
}
