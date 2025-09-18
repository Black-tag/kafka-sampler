package main

import (
	"context"
	"fmt"
	"log"
	"os"

	
	"github.com/Black-tag/kafka-sampler/internal/kafka/publisher"
	"github.com/Black-tag/kafka-sampler/internal/kafka/subscriber"
	"github.com/Black-tag/kafka-sampler/internal/load"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
	"go.yaml.in/yaml/v2"
)


func main() {


	// brokers := []string{"localhost:9092"}
	// topic := "load-test"

	// cfg := load.GeneratorConfig{
	// 	NumMessages: 20,
	// 	Topic: "load-test",
	// 	Key: "test-key",
	// 	EnableMetrics: true,

	// 	Brokers: brokers,
	// 	NumConsumers: 5,
	// 	ConsumerGroup: "load-group",
	// 	Partitions: 1,
	// 	Replication: 1,
		

	// }

	var cfg load.GeneratorConfig
	file, err := os.Open("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatal(err)
	}



	
	m := &metrics.Metrics{}
	producer := publisher.NewProducer(cfg.Brokers, cfg.Topic, m)
	defer producer.Close()


	load.Generate(producer, m, cfg)
	

	ctx := context.Background()
	consumer := subscriber.NewConsumer([]string{"localhost:9092"}, "load-test", "load-group", m)
	defer consumer.Close()

	

	consumer.StartConsuming(ctx, func(key, value string) {
		log.Println("recieved:", key, value)
		
	})
	fmt.Printf("produced=%d consumed=%d, Errors=%d latency=%v", m.Produced, m.Consumed, m.Errors, m.Latencies)

	select {}

	
}