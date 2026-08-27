# Feed 压测报告：hybrid / rpc / burst-read-c128

- Run ID：`feedopt-20260814-route-pipeline`
- 开始时间：2026-08-14T20:35:45+08:00
- 采样时长：30.0385539s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 70202 | 70202 | 0 | 0 | 2337.52 | 53.807 | 70.665 | 81.084 |

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
| docker:zg-counter | cpu_percent | 9.390 |
| docker:zg-counter | memory_percent | 0.250 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 4.690 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 22.000 |
| docker:zg-knowpost | cpu_percent | 294.540 |
| docker:zg-knowpost | memory_percent | 0.380 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 222.760 |
| docker:zg-relation | memory_percent | 0.430 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2956.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2956.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 824756.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 98.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19670670.000 |
| redis | connected_clients | 348.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.983 |
| redis | keys | 541029.000 |
| redis | keyspace_hits | 31638342.000 |
| redis | keyspace_misses | 556298.000 |
| redis | net_input_bytes | 2118112205.000 |
| redis | net_output_bytes | 6491475828.000 |
| redis | ops_per_sec | 19690.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93512288.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
