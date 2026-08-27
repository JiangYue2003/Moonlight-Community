# Feed 压测报告：hybrid / gateway / hot-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:50:25+08:00
- 采样时长：3.0846912s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3101 | 3101 | 0 | 0 | 1005.29 | 124.321 | 160.612 | 190.657 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 59.930 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 108.740 |
| docker:zg-gateway | memory_percent | 0.220 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 205.980 |
| docker:zg-knowpost | memory_percent | 0.460 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 103.160 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 433880.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 106.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9506758.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539902.000 |
| redis | keyspace_hits | 13283506.000 |
| redis | keyspace_misses | 123776.000 |
| redis | net_input_bytes | 986203329.000 |
| redis | net_output_bytes | 2876527308.000 |
| redis | ops_per_sec | 34568.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 101404920.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
