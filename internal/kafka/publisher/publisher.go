package publisher

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)





type Producer struct {
	writer *kafka.Writer
}



func NewProducer(brokers []string, topic string) *Producer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic: topic,
	})
	return &Producer{writer: w}
}


func (p *Producer) SendMessage(key,value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg := kafka.Message{
		Key: []byte(key),
		Value: []byte(value),
	}
	return p.writer.WriteMessages(ctx, msg)
}


func (p *Producer) Close() error {
	return p.writer.Close()
}