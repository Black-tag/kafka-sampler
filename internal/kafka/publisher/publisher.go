package publisher

import (
	"context"
	"fmt"

	"time"

	logger "github.com/Black-tag/kafka-sampler/internal/logging"
	"github.com/Black-tag/kafka-sampler/internal/messages"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer  *kafka.Writer
	metrics *metrics.Metrics
}

func NewProducer(brokers []string, topic string, m *metrics.Metrics) *Producer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})
	logger.Log.Info("produced a new produce")
	return &Producer{writer: w, metrics: m}
}

func (p *Producer) SendMessage(key, value string) error {

	logger.Log.Info("entered Sendmessege function")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message, _, err := messages.GenerateRandomMessage()
	if err != nil {
		return fmt.Errorf("couldn't produce random message: %v", err)
	}
	
	msg := kafka.Message{
		Key:   []byte(key),
		Value: message,
		Time:  time.Now(),
	}
	logger.Log.Info("msg produced")
	err = p.writer.WriteMessages(ctx, msg)

	if err != nil {
		p.metrics.IncErrors()
		return err
	}
	p.metrics.IncProduced()
	p.metrics.AddLatency(time.Since(msg.Time))
	logger.Log.Info("exiting Send Message function")

	return nil
}

func (p *Producer) Close() error {
	err := p.writer.Close()
	if err != nil {
		logger.Log.Error("cannot close producer")
		fmt.Printf("canot close producer: %v", err)
	}
	return err
}
