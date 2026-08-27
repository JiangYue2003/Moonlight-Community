# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-observer-off-20260815`
- 开始时间：2026-08-15T18:55:57+08:00
- 采样时长：20.0323s
- 并发：128
- 重复轮次：3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 52782 | 52782 | 0 | 0 | 2635.39 | 47.784 | 63.564 | 73.547 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker-state:zg-counter | healthy | 1.000 |
| docker-state:zg-counter | restart_count | 0.000 |
| docker-state:zg-counter | running | 1.000 |
| docker-state:zg-gateway | healthy | 1.000 |
| docker-state:zg-gateway | restart_count | 0.000 |
| docker-state:zg-gateway | running | 1.000 |
| docker-state:zg-knowpost | healthy | 1.000 |
| docker-state:zg-knowpost | restart_count | 0.000 |
| docker-state:zg-knowpost | running | 1.000 |
| docker-state:zg-relation | healthy | 1.000 |
| docker-state:zg-relation | restart_count | 0.000 |
| docker-state:zg-relation | running | 1.000 |
| docker:zg-counter | cpu_percent | 5.010 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 17.000 |
| docker:zg-gateway | cpu_percent | 0.450 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 327.860 |
| docker:zg-knowpost | memory_percent | 0.380 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 266.620 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 365909.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3305553.000 |
| redis | connected_clients | 272.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 15274312.000 |
| redis | keyspace_misses | 1829768.000 |
| redis | net_input_bytes | 712358637.000 |
| redis | net_output_bytes | 3135432583.000 |
| redis | ops_per_sec | 21771.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90693984.000 |

## 缺失指标

- feed_metrics

## 说明

- SLA values are reference lines, not pass/fail gates.
