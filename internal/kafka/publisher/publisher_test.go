package publisher_test

import(
	"github.com/Black-tag/kafka-sampler/internal/kafka/publisher"
	"github.com/Black-tag/kafka-sampler/internal/metrics"
	"testing"
)

func TestProducer_SendMessage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		brokers []string
		topic   string
		m       *metrics.Metrics
		// Named input parameters for target function.
		key     string
		value   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := publisher.NewProducer(tt.brokers, tt.topic, tt.m)
			gotErr := p.SendMessage(tt.key, tt.value)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SendMessage() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SendMessage() succeeded unexpectedly")
			}
		})
	}
}
