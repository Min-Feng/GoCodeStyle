package part

import (
	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter(address []string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(address...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		MaxAttempts:  0,
		BatchSize:    0,
		BatchBytes:   0,
		BatchTimeout: 0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		RequiredAcks: kafka.RequireOne,
		Async:        false,
		Completion:   nil,
		Compression:  0,
		Logger:       nil,
		ErrorLogger:  nil,
		Transport:    nil,
	}
}

func NewKafkaReader(address []string, topic string) *kafka.Reader {

	panic("")
}
