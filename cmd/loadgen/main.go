package main

import (
	"context"
	"log"

	"github.com/Black-tag/kafka-sampler/internal/kafka/publisher"
	"github.com/Black-tag/kafka-sampler/internal/kafka/subscriber"
	"github.com/Black-tag/kafka-sampler/internal/load"
)


func main() {

	brokers := []string{"localhost:9092"}
	topic := "load-test"

	producer := publisher.NewProducer(brokers, topic)
	defer producer.Close()


	load.Generate(producer)
	

	ctx := context.Background()
	consumer := subscriber.NewConsumer([]string{"localhost:9092"}, "load-test", "load-group")
	defer consumer.Close()

	consumer.StartConsuming(ctx, func(key, value string) {
		log.Println("recieved:", key, value)
		
	})

	select {}
}