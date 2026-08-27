# Feed 压测报告：pull / gateway / distributed-read-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:40:31+08:00
- 采样时长：3.0402965s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1351 | 1351 | 0 | 0 | 444.36 | 70.779 | 94.320 | 101.973 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.290 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 53.780 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 212.870 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 56.890 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 354651.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 90.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5833715.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537657.000 |
| redis | keyspace_hits | 8184178.000 |
| redis | keyspace_misses | 92050.000 |
| redis | net_input_bytes | 661113907.000 |
| redis | net_output_bytes | 1891942598.000 |
| redis | ops_per_sec | 13490.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92614976.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
