# Feed 压测报告：hybrid / rpc / hot-read-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:17+08:00
- 采样时长：3.0104989s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2334 | 2334 | 0 | 0 | 775.44 | 10.026 | 12.910 | 17.264 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 60.190 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.650 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 188.610 |
| docker:zg-knowpost | memory_percent | 0.350 |
| docker:zg-knowpost | pids | 23.000 |
| docker:zg-relation | cpu_percent | 100.380 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 370981.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 75.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6829756.000 |
| redis | connected_clients | 398.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 538745.000 |
| redis | keyspace_hits | 8579113.000 |
| redis | keyspace_misses | 106784.000 |
| redis | net_input_bytes | 742209991.000 |
| redis | net_output_bytes | 2055409958.000 |
| redis | ops_per_sec | 26418.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87793256.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
