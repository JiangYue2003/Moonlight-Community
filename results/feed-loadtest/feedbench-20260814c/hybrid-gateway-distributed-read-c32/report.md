# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:50:35+08:00
- 采样时长：3.02302s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2881 | 2881 | 0 | 0 | 953.19 | 33.184 | 42.968 | 48.129 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 59.890 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 98.960 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 201.010 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 101.620 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 437247.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 106.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9615775.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539976.000 |
| redis | keyspace_hits | 13518847.000 |
| redis | keyspace_misses | 123779.000 |
| redis | net_input_bytes | 994722817.000 |
| redis | net_output_bytes | 2929244141.000 |
| redis | ops_per_sec | 31871.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95260960.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
