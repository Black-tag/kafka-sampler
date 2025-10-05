package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/Black-tag/kafka-sampler/internal/kafka/publisher"
	"github.com/Black-tag/kafka-sampler/internal/kafka/subscriber"
	"github.com/Black-tag/kafka-sampler/internal/load"
	logger "github.com/Black-tag/kafka-sampler/internal/logging"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.yaml.in/yaml/v2"
)

func main() {

	metrics.RecordMetrics()
	logger.Log.Info("starting application")

	var cfg load.GeneratorConfig
	file, err := os.Open("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
		logger.Log.DPanic("cannot open config")

	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			logger.Log.Error("cannot close file")
			fmt.Println(cerr)

		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		logger.Log.Error("could not decode config into struct")
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	m := &metrics.Metrics{}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":2112", nil)
		if err != nil {
			logger.Log.Error("cannot start the prometheus endpoint")
			fmt.Println("error in loading metrics")
		}

	}()

	// producer := publisher.NewProducer(cfg.Brokers, cfg.Topic, m)
	producer := make([]*publisher.Producer, cfg.NumProducer)
	for i := 0; i < cfg.NumProducer; i++ {
		producer[i] = publisher.NewProducer(cfg.Brokers, cfg.Topic, m)
		defer func() {
			if cerr := producer[i].Close(); cerr != nil {
				logger.Log.Error("cannot close producer")
				fmt.Printf("cannot close producer: %v", cerr)
			}

		}()
	}

	load.Generate(producer, m, cfg)

	ctx := context.Background()
	consumer := subscriber.NewConsumer(cfg.Brokers, cfg.Topic, cfg.ConsumerGroup, m)
	defer consumer.Close()

	for i := 0; i < cfg.NumConsumers; i++ {
		wg.Add(1)
		go consumer.StartConsuming(ctx, func(key, value string) {
			fmt.Printf("Consumer %d consumed key=%s, value=%s | produced=%d, consumed=%d, errors=%d\n",
				i, key, value, m.Produced, m.Consumed, m.Errors)
		}, i, &wg)
	}

	select {}

}
