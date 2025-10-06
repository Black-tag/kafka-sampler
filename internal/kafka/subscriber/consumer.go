package subscriber

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"log"
	"time"

	logger "github.com/Black-tag/kafka-sampler/internal/logging"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader  *kafka.Reader
	metrics *metrics.Metrics
}

type ConsumerParams struct {
	Brokers []string
	Topic   string
	GroupId string
	Metrics *metrics.Metrics
}

func NewConsumer(brokers []string, topic string, groupID string, m *metrics.Metrics) *Consumer {
	logger.Log.Info("enetered Newconsumer")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		Topic:          topic,
		GroupID:        groupID,
		StartOffset:    kafka.FirstOffset,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
	})
	logger.Log.Info("exiting Newconsumer")
	return &Consumer{reader: r, metrics: m}
}

func (c *Consumer) StartConsuming(ctx context.Context, handle func(key, value string), id int, wg *sync.WaitGroup) {
	logger.Log.Info("entered Startconsuming function")
	defer wg.Done()

	for {
		msg, err := c.reader.ReadMessage(ctx)
		logger.Log.Info("start consuming msgs")
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}
			logger.Log.Error("cannot consume msg")
			log.Println("error reading the message", err)
			c.metrics.IncErrors()

		}
		logger.Log.Info("no error in Reding msg")
		c.metrics.IncConsumed()
		latency := time.Since(msg.Time)
		c.metrics.AddLatency(latency)

		handle(string(msg.Key), string(msg.Value))
		logger.Log.Info("handled msgs")
	}

}

func (c *Consumer) Close() {
	err := c.reader.Close()
	if err != nil {
		logger.Log.Error("cannot close Consumer")
		fmt.Printf("cannot close the consumer: %v", err)
	}
}
