# Feed 压测报告：pull / gateway / hot-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:40:20+08:00
- 采样时长：3.1590035s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1452 | 1452 | 0 | 0 | 459.64 | 268.601 | 348.206 | 401.617 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.370 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 50.370 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 218.900 |
| docker:zg-knowpost | memory_percent | 0.510 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 53.560 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 352809.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 89.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5782845.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537592.000 |
| redis | keyspace_hits | 8069214.000 |
| redis | keyspace_misses | 90311.000 |
| redis | net_input_bytes | 655185990.000 |
| redis | net_output_bytes | 1822514998.000 |
| redis | ops_per_sec | 13720.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93222080.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
