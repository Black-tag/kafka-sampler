package load

import (
	"fmt"
	"time"

	"github.com/Black-tag/kafka-sampler/internal/kafka/publisher"
	logger "github.com/Black-tag/kafka-sampler/internal/logging"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
)


type GeneratorConfig struct {
	NumMessages int    `yaml:"num_messages"`
	Topic       string `yaml:"topic"`
	Key         string `yaml:"key"`
	EnableMetrics bool `yaml:"enable_metrics"`


	Brokers    []string `yaml:"brokers"`
	NumConsumers int    `yaml:"num_consumers"`
	ConsumerGroup string `yaml:"consumer_group"`
	Partitions    int    `yaml:"partitions"`
	Replication   int    `yaml:"replication"`


}


func Generate(producer *publisher.Producer, m *metrics.Metrics, cfg GeneratorConfig) {
	logger.Log.Info("entered generate function")
	
    for i := 0; i < cfg.NumMessages; i++ {
		logger.Log.Info("starting to produce msg")
        msg := fmt.Sprintf("message-%d", i)
		err := producer.SendMessage(cfg.Key, msg)
		if err != nil {
			logger.Log.Error("error producing msgs")
			fmt.Println("error in producing message", err)
		} else {
			fmt.Println("produced:", msg)
		}
        
    }
	if cfg.EnableMetrics  {
		logger.Log.Info("metrics is enabled")
		go func() {
		ticker := time.NewTicker(5*time.Second)
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

