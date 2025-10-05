package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)
var MessageProduced = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "produced_msgs_total",
		Help: "messages total produced by producer",
	})

func Init() {
	prometheus.MustRegister(MessageProduced)
}
