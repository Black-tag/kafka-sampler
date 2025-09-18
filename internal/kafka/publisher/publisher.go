package publisher

import (
	"context"

	"time"

	logger "github.com/Black-tag/kafka-sampler/internal/logging"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
	"github.com/segmentio/kafka-go"
)





type Producer struct {
	writer *kafka.Writer
	metrics *metrics.Metrics
}



func NewProducer(brokers []string, topic string, m *metrics.Metrics) *Producer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic: topic,
	})
	logger.Log.Info("produced a new produce" )
	return &Producer{writer: w, metrics: m}
}


func (p *Producer) SendMessage(key,value string) error {
	logger.Log.Info("entered Sendmessege function")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg := kafka.Message{
		Key: []byte(key),
		Value: []byte(value),
		Time: time.Now(),
	}
	logger.Log.Info("msg produced")
	err := p.writer.WriteMessages(ctx, msg)

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
	return p.writer.Close()
}