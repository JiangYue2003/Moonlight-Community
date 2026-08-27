# Feed 压测报告：push / rpc / distributed-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:59+08:00
- 采样时长：3.0366413s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 8379 | 8379 | 0 | 0 | 2759.77 | 45.645 | 60.037 | 66.048 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.410 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.290 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 276.490 |
| docker:zg-knowpost | memory_percent | 0.440 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 238.330 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 239278.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 100.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1707753.000 |
| redis | connected_clients | 418.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 536257.000 |
| redis | keyspace_hits | 2042728.000 |
| redis | keyspace_misses | 15578.000 |
| redis | net_input_bytes | 174981169.000 |
| redis | net_output_bytes | 381545131.000 |
| redis | ops_per_sec | 8526.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93512536.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
