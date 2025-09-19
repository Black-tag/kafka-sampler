package metrics

import (
	"sync"
	"time"
)

type Metrics struct {
	Produced  int64
	Consumed  int64
	Errors    int64
	Latencies []time.Duration

	mu sync.Mutex
}

type Snapshot struct {
	Produced  int64
	Consumed  int64
	Errors    int64
	Latencies []time.Duration
}

func (m *Metrics) TakeSnapshot() Snapshot {

	m.mu.Lock()
	defer m.mu.Unlock()
	latenciesCopy := make([]time.Duration, len(m.Latencies))
	copy(latenciesCopy, m.Latencies)
	return Snapshot{
		Produced:  m.Produced,
		Consumed:  m.Consumed,
		Errors:    m.Errors,
		Latencies: latenciesCopy,
	}
}

func (m *Metrics) IncProduced() {
	m.mu.Lock()
	m.Produced++
	m.mu.Unlock()
}

func (m *Metrics) IncConsumed() {
	m.mu.Lock()
	m.Consumed++
	m.mu.Unlock()
}

func (m *Metrics) IncErrors() {
	m.mu.Lock()
	m.Errors++
	m.mu.Unlock()
}

func (m *Metrics) AddLatency(d time.Duration) {
	m.mu.Lock()
	m.Latencies = append(m.Latencies, d)
	m.mu.Unlock()
}
