package kafka

import (
	"context"

	"github.com/kanhaiya15/gopf/utils"
	"github.com/segmentio/kafka-go"
)

var w *kafka.Writer

// InitProducer InitProducer
func InitProducer(brokers []string) {
	logger.Debug(" Kafka Producer -starts")
	topic, err := utils.GetConfValue("KAFKA_TOPIC")
	if err != nil {
		panic(err.Error())
	}
	w = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

// WriteMessage WriteMessage
func WriteMessage(message interface{}) {
	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
	)
}
