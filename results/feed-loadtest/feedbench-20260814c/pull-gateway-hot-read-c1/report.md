# Feed 压测报告：pull / gateway / hot-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:40:10+08:00
- 采样时长：3.0011652s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 404 | 404 | 0 | 0 | 134.61 | 7.351 | 8.717 | 9.507 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.410 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 17.740 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 66.550 |
| docker:zg-knowpost | memory_percent | 0.400 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 18.500 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 349824.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 51.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5702870.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537517.000 |
| redis | keyspace_hits | 7884623.000 |
| redis | keyspace_misses | 87514.000 |
| redis | net_input_bytes | 645782755.000 |
| redis | net_output_bytes | 1711046332.000 |
| redis | ops_per_sec | 4049.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86295768.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
