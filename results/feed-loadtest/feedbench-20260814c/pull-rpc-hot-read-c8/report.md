# Feed 压测报告：pull / rpc / hot-read-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:35:21+08:00
- 采样时长：3.0101814s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1551 | 1551 | 0 | 0 | 515.34 | 15.483 | 18.473 | 21.043 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.310 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 1.790 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 225.930 |
| docker:zg-knowpost | memory_percent | 0.350 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 74.370 |
| docker:zg-relation | memory_percent | 0.270 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 315678.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 75.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4277518.000 |
| redis | connected_clients | 388.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 537413.000 |
| redis | keyspace_hits | 5772390.000 |
| redis | keyspace_misses | 50103.000 |
| redis | net_input_bytes | 486537342.000 |
| redis | net_output_bytes | 1126500320.000 |
| redis | ops_per_sec | 14958.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90407576.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
