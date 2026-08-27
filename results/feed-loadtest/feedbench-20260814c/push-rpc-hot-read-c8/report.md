# Feed 压测报告：push / rpc / hot-read-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:22+08:00
- 采样时长：3.0050794s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3599 | 3599 | 0 | 0 | 1197.64 | 6.432 | 8.125 | 9.583 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.820 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.120 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 157.530 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 142.450 |
| docker:zg-relation | memory_percent | 0.160 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 196570.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 75.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1573508.000 |
| redis | connected_clients | 276.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.943 |
| redis | keys | 536196.000 |
| redis | keyspace_hits | 254795.000 |
| redis | keyspace_misses | 15524.000 |
| redis | net_input_bytes | 109714874.000 |
| redis | net_output_bytes | 52279521.000 |
| redis | ops_per_sec | 3882.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85286216.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
