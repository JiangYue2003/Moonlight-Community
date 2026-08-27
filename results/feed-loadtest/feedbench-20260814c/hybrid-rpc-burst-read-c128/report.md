# Feed 压测报告：hybrid / rpc / burst-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T17:02:20+08:00
- 采样时长：10.0564792s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 15526 | 15526 | 0 | 0 | 1544.12 | 80.490 | 106.303 | 125.845 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 85.130 |
| docker:zg-counter | memory_percent | 0.240 |
| docker:zg-counter | pids | 28.000 |
| docker:zg-gateway | cpu_percent | 2.920 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 272.810 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 162.370 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2852.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2852.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 615336.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 69.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 14202133.000 |
| redis | connected_clients | 413.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.984 |
| redis | keys | 538655.000 |
| redis | keyspace_hits | 21147389.000 |
| redis | keyspace_misses | 338277.000 |
| redis | net_input_bytes | 1515036362.000 |
| redis | net_output_bytes | 4410270588.000 |
| redis | ops_per_sec | 55203.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95537472.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
