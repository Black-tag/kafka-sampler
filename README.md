Kafka Server Performance Testing

Apache Kafka is a distributed streaming platform that must handle high throughput, low latency, and fault tolerance across producers, brokers, and consumers. Performance testing ensures Kafka can sustain required workloads under real-world conditions and helps identify bottlenecks in hardware, configuration, or application design.

Key Components to Test

Producers

Message publish rate (records/sec, MB/sec).

Latency in sending records.

Impact of batching and compression (e.g., Snappy, LZ4, ZSTD).

Acknowledgment levels (acks=0,1,all).

Brokers

Broker throughput and disk I/O performance.

Partition distribution and replication overhead.

Network bandwidth utilization.

Page cache and OS-level tuning.

Consumers

Message consumption rate (records/sec, MB/sec).

Consumer group rebalancing impact.

Commit latency (auto vs manual).

Effect of fetch size and max.poll settings.

ZooKeeper / KRaft (if using newer versions)

Metadata update latency.

Controller failover performance.

Cluster Infrastructure

Disk types (SSD vs HDD).

Network (1GbE vs 10GbE vs RDMA).

CPU & memory utilization.

Performance Testing Techniques

Throughput Testing

Measure maximum records/sec and MB/sec producers can push and consumers can pull without errors.

Use multiple producers/consumers to simulate realistic load.

Latency Testing

End-to-end latency (producer → broker → consumer).

Publish acknowledgment latency at different acks settings.

Consumer fetch latency under varying loads.

Stress Testing

Push the system beyond capacity (e.g., millions of messages/sec).

Identify breaking points (CPU saturation, disk I/O bottleneck, network congestion).

Scalability Testing

Increase partitions, topics, or nodes to measure horizontal scaling.

Evaluate how leader reassignments and replication affect performance.

Durability & Reliability Testing

Simulate broker failures, disk crashes, and network partitions.

Measure recovery time and data loss (if any).

Soak (Endurance) Testing

Run sustained workloads (hours/days).

Look for memory leaks, disk growth, and broker stability over time.

Backpressure & Lag Testing

Simulate slow consumers and observe lag buildup.

Verify Kafka’s ability to handle backlog replay.







Critical Metrics to Monitor

Producer Metrics

Records per second.

Record send rate vs error rate.

Batch size, retries, request latency.

Broker Metrics (JMX / Kafka Metrics Reporter)

Bytes in/out per second.

Request handler time (produce, fetch).

Disk I/O wait time.

Under-replicated partitions.

Controller latency.

Consumer Metrics

Records consumed per second.

Commit latency.

Lag per partition.

System Metrics

CPU utilization.

Network throughput.

Disk read/write latency.

Garbage collection (GC) pauses.