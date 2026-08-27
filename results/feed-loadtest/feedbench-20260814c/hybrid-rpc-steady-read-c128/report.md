# Feed 压测报告：hybrid / rpc / steady-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T17:02:33+08:00
- 采样时长：30.0519142s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 47777 | 47777 | 0 | 0 | 1590.01 | 79.360 | 103.534 | 115.576 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 90.610 |
| docker:zg-counter | memory_percent | 0.260 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 3.030 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 30.000 |
| docker:zg-knowpost | cpu_percent | 274.960 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 164.320 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2852.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2852.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 663157.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 69.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 15782886.000 |
| redis | connected_clients | 413.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 538665.000 |
| redis | keyspace_hits | 24539578.000 |
| redis | keyspace_misses | 386099.000 |
| redis | net_input_bytes | 1638699479.000 |
| redis | net_output_bytes | 4890026006.000 |
| redis | ops_per_sec | 57294.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 96230648.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
