# Feed 压测报告：pull / rpc / burst-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T16:59:37+08:00
- 采样时长：10.144571s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 6291 | 6291 | 0 | 0 | 620.17 | 195.627 | 265.135 | 353.563 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.250 |
| docker:zg-counter | memory_percent | 0.190 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 3.480 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 290.820 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 100.000 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2751.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2751.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 575793.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 69.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12776928.000 |
| redis | connected_clients | 419.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.986 |
| redis | keys | 537644.000 |
| redis | keyspace_hits | 18731562.000 |
| redis | keyspace_misses | 275601.000 |
| redis | net_input_bytes | 1380793433.000 |
| redis | net_output_bytes | 4008079550.000 |
| redis | ops_per_sec | 20331.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91998304.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
