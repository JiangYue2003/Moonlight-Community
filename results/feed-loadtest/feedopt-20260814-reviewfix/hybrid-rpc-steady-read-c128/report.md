# Feed 压测报告：hybrid / rpc / steady-read-c128

- Run ID：`feedopt-20260814-reviewfix`
- 开始时间：2026-08-14T21:14:08+08:00
- 采样时长：30.0421341s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 71968 | 71968 | 0 | 0 | 2396.03 | 52.618 | 68.551 | 77.851 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker-state:zg-counter | healthy | 1.000 |
| docker-state:zg-counter | restart_count | 0.000 |
| docker-state:zg-counter | running | 1.000 |
| docker-state:zg-gateway | healthy | 1.000 |
| docker-state:zg-gateway | restart_count | 0.000 |
| docker-state:zg-gateway | running | 1.000 |
| docker-state:zg-knowpost | healthy | 1.000 |
| docker-state:zg-knowpost | restart_count | 0.000 |
| docker-state:zg-knowpost | running | 1.000 |
| docker-state:zg-relation | healthy | 1.000 |
| docker-state:zg-relation | restart_count | 0.000 |
| docker-state:zg-relation | running | 1.000 |
| docker:zg-counter | cpu_percent | 6.630 |
| docker:zg-counter | memory_percent | 0.250 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.660 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 298.590 |
| docker:zg-knowpost | memory_percent | 0.380 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 237.500 |
| docker:zg-relation | memory_percent | 0.370 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2958.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2958.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 1058826.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 22016714.000 |
| redis | connected_clients | 304.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.982 |
| redis | keys | 542691.000 |
| redis | keyspace_hits | 42310922.000 |
| redis | keyspace_misses | 801807.000 |
| redis | net_input_bytes | 2590686143.000 |
| redis | net_output_bytes | 8856404762.000 |
| redis | ops_per_sec | 19981.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93483032.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
