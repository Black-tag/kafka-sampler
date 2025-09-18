package subscriber

import (
	"context"
	"errors"

	"log"
	"time"

	"github.com/Black-tag/kafka-sampler/internal/metrics"
	"github.com/segmentio/kafka-go"
)


type Consumer struct {
	reader *kafka.Reader
	metrics *metrics.Metrics
}



func NewConsumer(brokers []string, topic string, groupID string, m *metrics.Metrics) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic: topic,
		GroupID: groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
		CommitInterval: time.Second,
	})
	return &Consumer{reader: r, metrics: m}
}


func (c *Consumer) StartConsuming(ctx context.Context, handle func(key, value string)) {
	
		for {
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				 if errors.Is(err, context.Canceled) {
            		return 
        		}

				log.Println("error reading the message", err)
				c.metrics.IncErrors()
				continue
			}
			c.metrics.IncConsumed()
			latency := time.Since(msg.Time)
			c.metrics.AddLatency(latency)
			
			handle(string(msg.Key), string(msg.Value))
		}
		
	
}


func (c *Consumer) Close() {
	c.reader.Close()
}