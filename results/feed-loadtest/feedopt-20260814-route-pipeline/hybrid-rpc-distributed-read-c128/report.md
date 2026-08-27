# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feedopt-20260814-route-pipeline`
- 开始时间：2026-08-14T20:30:40+08:00
- 采样时长：3.0332497s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 6661 | 6661 | 0 | 0 | 2196.39 | 55.160 | 72.104 | 194.708 |

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
| docker:zg-counter | cpu_percent | 4.030 |
| docker:zg-counter | memory_percent | 0.250 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.340 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 22.000 |
| docker:zg-knowpost | cpu_percent | 283.770 |
| docker:zg-knowpost | memory_percent | 0.350 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 214.860 |
| docker:zg-relation | memory_percent | 0.390 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2956.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2956.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 675238.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18423930.000 |
| redis | connected_clients | 317.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.984 |
| redis | keys | 541373.000 |
| redis | keyspace_hits | 24864908.000 |
| redis | keyspace_misses | 405350.000 |
| redis | net_input_bytes | 1835689371.000 |
| redis | net_output_bytes | 5004547360.000 |
| redis | ops_per_sec | 18683.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93115952.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
