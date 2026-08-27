# Feed 压测报告：hybrid / gateway / distributed-read-c128

- Run ID：`feedopt-20260814-route-pipeline-nolog`
- 开始时间：2026-08-14T20:47:14+08:00
- 采样时长：3.0634206s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4901 | 4901 | 0 | 0 | 1599.85 | 76.315 | 111.198 | 163.988 |

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
| docker:zg-counter | cpu_percent | 2.590 |
| docker:zg-counter | memory_percent | 0.250 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 155.920 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 188.970 |
| docker:zg-knowpost | memory_percent | 0.350 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 151.640 |
| docker:zg-relation | memory_percent | 0.380 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2956.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2956.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 913568.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 95.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 20515355.000 |
| redis | connected_clients | 352.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.982 |
| redis | keys | 541013.000 |
| redis | keyspace_hits | 35670111.000 |
| redis | keyspace_misses | 654107.000 |
| redis | net_input_bytes | 2294608080.000 |
| redis | net_output_bytes | 7378481216.000 |
| redis | ops_per_sec | 12612.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93495192.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
