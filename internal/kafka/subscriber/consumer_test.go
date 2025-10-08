package subscriber_test

import(
	"github.com/Black-tag/kafka-sampler/internal/kafka/subscriber"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		brokers []string
		topic   string
		groupID string
		m       *metrics.Metrics
		want    *subscriber.Consumer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subscriber.NewConsumer(tt.brokers, tt.topic, tt.groupID, tt.m)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewConsumer() = %v, want %v", got, tt.want)
			}
		})
	}
}
