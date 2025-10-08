package load

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Black-tag/kafka-sampler/internal/kafka/publisher"
	logger "github.com/Black-tag/kafka-sampler/internal/logging"
	"github.com/Black-tag/kafka-sampler/internal/messages"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
	"go.uber.org/zap"
)

type GeneratorConfig struct {
	NumMessages   int    `yaml:"num_messages"`
	Topic         string `yaml:"topic"`
	Key           string `yaml:"key"`
	EnableMetrics bool   `yaml:"enable_metrics"`

	Brokers       []string `yaml:"brokers"`
	NumProducer   int      `yaml:"num_producers"`
	NumConsumers  int      `yaml:"num_consumers"`
	ConsumerGroup string   `yaml:"consumer_group"`
	Partitions    int      `yaml:"partitions"`
	Replication   int      `yaml:"replication"`
	PayloadFormat    string   `yaml:"payload_format"` 
	MinPayloadSize int     `yaml:"min_payload_size"`
	MaxPayloadSize int     `yaml:"max_payload_size"`
	MessageRate    int     `yaml:"message_rate"`
	WebhookEndpoint string `yaml:"webhook_endpoint"`
	EnableWebhook   bool   `yaml:"enable_webhook"`

}

func Generate(producer []*publisher.Producer, m *metrics.Metrics, cfg GeneratorConfig) {
	logger.Log.Info("entered generate function")

	var wg sync.WaitGroup

	for i := 0; i < cfg.NumMessages; i++ {
		wg.Add(1)


		go func(i int) {
			defer wg.Done()

			logger.Log.Info("starting to produce msg")
			msg, payload, err := messages.GenerateRandomMessage()
			
			logger.Log.Info("current value of i", zap.Int("i", i))
			if len(producer) == 0 {
				panic("no producers availible")

			}
			p := producer[i%len(producer)]
			logger.Log.Info("current producer", zap.Int("producer", cfg.NumProducer))
			err = p.SendMessage(cfg.Key, string(msg))
			if err != nil {
				logger.Log.Error("error producing msgs")
				fmt.Println("error in producing message", err)
			} else {
				fmt.Println("produced:", msg)
			}
			metrics.MessageProduced.Inc()

			if cfg.EnableWebhook && cfg.WebhookEndpoint != "" {
				go func() {
					http.Post(cfg.WebhookEndpoint, "application/octet-stream", bytes.NewReader([]byte(payload)))
				}()
			}
		}(i)

	}

	wg.Wait()
	if cfg.EnableMetrics {
		logger.Log.Info("metrics is enabled")
		go func() {
			ticker := time.NewTicker(5 * time.Second)
			logger.Log.Info("ticker")
			defer ticker.Stop()
			for range ticker.C {
				snap := m.TakeSnapshot()
				fmt.Printf("produced=%d, consumed=%d, Errors=%d, latency=%v\n", snap.Produced, snap.Consumed, snap.Errors, snap.Latencies)

			}

		}()
		logger.Log.Info("completed logging of msgs")
	}
	logger.Log.Info("exiting generate function")

}
