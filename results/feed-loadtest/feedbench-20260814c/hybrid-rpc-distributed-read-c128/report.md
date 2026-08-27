# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-15T18:06:49+08:00
- 采样时长：30.0373186s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 79185 | 79185 | 0 | 0 | 2636.79 | 47.900 | 62.435 | 69.700 |

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
| docker:zg-counter | cpu_percent | 5.700 |
| docker:zg-counter | memory_percent | 0.100 |
| docker:zg-counter | pids | 19.000 |
| docker:zg-gateway | cpu_percent | 2.620 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 320.800 |
| docker:zg-knowpost | memory_percent | 0.400 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 258.620 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 30.000 |
| kafka | current_offset_total | 2959.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2959.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 155801.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 7.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1354440.000 |
| redis | connected_clients | 257.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519476.000 |
| redis | keyspace_hits | 6521395.000 |
| redis | keyspace_misses | 782459.000 |
| redis | net_input_bytes | 300518901.000 |
| redis | net_output_bytes | 1320656687.000 |
| redis | ops_per_sec | 22159.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90190056.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
