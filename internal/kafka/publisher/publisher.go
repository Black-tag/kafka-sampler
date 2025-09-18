package publisher

import (
	"context"
	
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
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
	return &Producer{writer: w, metrics: m}
}


func (p *Producer) SendMessage(key,value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg := kafka.Message{
		Key: []byte(key),
		Value: []byte(value),
		Time: time.Now(),
	}
	err := p.writer.WriteMessages(ctx, msg)

	if err != nil {
		p.metrics.IncErrors()
		return err
	}
	p.metrics.IncProduced()
	p.metrics.AddLatency(time.Since(msg.Time))
	return nil
}


func (p *Producer) Close() error {
	return p.writer.Close()
}