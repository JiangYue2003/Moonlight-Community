# Feed 压测报告：pull / rpc / distributed-read-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:35:48+08:00
- 采样时长：3.0100209s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1566 | 1566 | 0 | 0 | 520.26 | 15.277 | 18.494 | 19.623 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.510 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.170 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 238.900 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 78.010 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 323406.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 123.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4496291.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537380.000 |
| redis | keyspace_hits | 6276064.000 |
| redis | keyspace_misses | 57737.000 |
| redis | net_input_bytes | 512242407.000 |
| redis | net_output_bytes | 1276971290.000 |
| redis | ops_per_sec | 15274.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91154872.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
