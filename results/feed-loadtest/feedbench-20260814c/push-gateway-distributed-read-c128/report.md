# Feed 压测报告：push / gateway / distributed-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:30:41+08:00
- 采样时长：3.0608574s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4601 | 4601 | 0 | 0 | 1503.17 | 82.497 | 118.268 | 133.783 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.370 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 137.500 |
| docker:zg-gateway | memory_percent | 0.220 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 165.880 |
| docker:zg-knowpost | memory_percent | 0.470 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 142.550 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 301228.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 100.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3222485.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 537996.000 |
| redis | keyspace_hits | 5571448.000 |
| redis | keyspace_misses | 36153.000 |
| redis | net_input_bytes | 402119210.000 |
| redis | net_output_bytes | 1058922123.000 |
| redis | ops_per_sec | 4330.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 125619704.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
