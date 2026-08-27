# Feed 压测报告：push / gateway / hot-read-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:30:19+08:00
- 采样时长：3.0151455s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4873 | 4873 | 0 | 0 | 1616.45 | 19.342 | 26.135 | 30.763 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.130 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 143.760 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 22.000 |
| docker:zg-knowpost | cpu_percent | 175.130 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 153.510 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 286257.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3173791.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 537781.000 |
| redis | keyspace_hits | 4952954.000 |
| redis | keyspace_misses | 36087.000 |
| redis | net_input_bytes | 379379613.000 |
| redis | net_output_bytes | 944965817.000 |
| redis | ops_per_sec | 5044.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 121498864.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
