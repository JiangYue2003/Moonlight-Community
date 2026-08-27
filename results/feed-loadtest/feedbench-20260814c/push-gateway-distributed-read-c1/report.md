# Feed 压测报告：push / gateway / distributed-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:30:30+08:00
- 采样时长：3.0052527s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 549 | 549 | 0 | 0 | 182.68 | 5.302 | 6.880 | 7.514 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.130 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 24.590 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 27.850 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 24.450 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 292146.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 100.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3193541.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 537890.000 |
| redis | keyspace_hits | 5194394.000 |
| redis | keyspace_misses | 36120.000 |
| redis | net_input_bytes | 388309182.000 |
| redis | net_output_bytes | 989473471.000 |
| redis | ops_per_sec | 801.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 119818160.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
