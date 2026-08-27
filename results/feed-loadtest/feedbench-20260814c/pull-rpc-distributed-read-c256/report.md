# Feed 压测报告：pull / rpc / distributed-read-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:36:04+08:00
- 采样时长：3.25126s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2022 | 2022 | 0 | 0 | 621.91 | 406.841 | 471.097 | 513.875 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.010 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.790 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 276.170 |
| docker:zg-knowpost | memory_percent | 0.610 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 72.720 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 329135.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 100.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4657521.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537351.000 |
| redis | keyspace_hits | 6649705.000 |
| redis | keyspace_misses | 63401.000 |
| redis | net_input_bytes | 531237692.000 |
| redis | net_output_bytes | 1388561165.000 |
| redis | ops_per_sec | 17516.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93444032.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
