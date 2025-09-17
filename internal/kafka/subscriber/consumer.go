package subscriber

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)


type Consumer struct {
	reader *kafka.Reader
}



func NewConsumer(brokers []string, topic string, groupID string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic: topic,
		GroupID: groupID,
	})
	return &Consumer{reader: r}
}


func (c *Consumer) StartConsuming(ctx context.Context, handle func(key, value string)) {
	go func() {
		for {
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Println("error reading the message", err)
				continue
			}
			handle(string(msg.Key), string(msg.Value))
		}
	}()
}


func (c *Consumer) Close() {
	c.reader.Close()
}