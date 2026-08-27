# Feed 压测报告：push / rpc / distributed-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:44+08:00
- 采样时长：3.0005047s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 685 | 685 | 0 | 0 | 228.29 | 4.236 | 5.299 | 6.017 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.370 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.170 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 31.510 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 31.720 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 221113.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 124.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1650728.000 |
| redis | connected_clients | 418.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 536234.000 |
| redis | keyspace_hits | 1281757.000 |
| redis | keyspace_misses | 15557.000 |
| redis | net_input_bytes | 147211849.000 |
| redis | net_output_bytes | 241418795.000 |
| redis | ops_per_sec | 920.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87893424.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
