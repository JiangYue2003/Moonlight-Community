# Feed 压测报告：pull / rpc / deep-page-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:36:25+08:00
- 采样时长：3.1656026s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1740 | 1740 | 0 | 0 | 549.66 | 225.508 | 297.496 | 318.212 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.180 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.410 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 272.980 |
| docker:zg-knowpost | memory_percent | 0.530 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 75.110 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 334833.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 78.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4809810.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537402.000 |
| redis | keyspace_hits | 7421590.000 |
| redis | keyspace_misses | 69166.000 |
| redis | net_input_bytes | 563992675.000 |
| redis | net_output_bytes | 1555857427.000 |
| redis | ops_per_sec | 15707.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93096968.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
