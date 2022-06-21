package kafka

import (
	"context"
	"crypto/tls"
	"log"
	"os"

	kafkago "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func NewKafkaProducer() *kafkago.Writer {
	mechanism, err := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_SASL_USERNAME"), os.Getenv("KAFKA_SASL_PASSWORD"))
	if err != nil {
		log.Fatalln(err)
	}
	dialer := &kafkago.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}
	p := kafkago.NewWriter(kafkago.WriterConfig{
		Brokers: []string{os.Getenv("KAFKA_BOOTSTRAP_SERVERS")},
		Dialer:  dialer,
	})
	return p
}

func Publish(msg string, topic string, producer *kafkago.Writer) error {
	err := producer.WriteMessages(context.Background(), kafkago.Message{
		Topic: topic,
		Value: []byte(msg),
	})
	if err != nil {
		return err
	}
	return nil
}
