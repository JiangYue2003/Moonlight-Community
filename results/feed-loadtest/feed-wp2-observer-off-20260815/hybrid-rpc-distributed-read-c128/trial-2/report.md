# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-observer-off-20260815`
- 开始时间：2026-08-15T18:55:24+08:00
- 采样时长：20.0350252s
- 并发：128
- 重复轮次：2
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 52804 | 52804 | 0 | 0 | 2636.07 | 47.815 | 62.935 | 72.339 |

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
| docker:zg-counter | cpu_percent | 3.550 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 17.000 |
| docker:zg-gateway | cpu_percent | 0.310 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 327.600 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 267.500 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 299160.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 8.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2771640.000 |
| redis | connected_clients | 272.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519486.000 |
| redis | keyspace_hits | 12511721.000 |
| redis | keyspace_misses | 1498749.000 |
| redis | net_input_bytes | 587819760.000 |
| redis | net_output_bytes | 2563506404.000 |
| redis | ops_per_sec | 21776.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90740416.000 |

## 缺失指标

- feed_metrics

## 说明

- SLA values are reference lines, not pass/fail gates.
