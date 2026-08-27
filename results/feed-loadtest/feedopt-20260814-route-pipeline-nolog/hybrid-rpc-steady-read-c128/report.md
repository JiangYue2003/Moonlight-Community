# Feed 压测报告：hybrid / rpc / steady-read-c128

- Run ID：`feedopt-20260814-route-pipeline-nolog`
- 开始时间：2026-08-14T20:39:09+08:00
- 采样时长：30.0499442s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 71145 | 71145 | 0 | 0 | 2368.02 | 53.186 | 69.479 | 79.677 |

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
| docker:zg-counter | cpu_percent | 8.170 |
| docker:zg-counter | memory_percent | 0.250 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 3.190 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 22.000 |
| docker:zg-knowpost | cpu_percent | 299.360 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 229.520 |
| docker:zg-relation | memory_percent | 0.410 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2956.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2956.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 896141.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 87.000 |
| mysql | threads_running | 8.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 20280765.000 |
| redis | connected_clients | 353.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.983 |
| redis | keys | 540952.000 |
| redis | keyspace_hits | 34911050.000 |
| redis | keyspace_misses | 628424.000 |
| redis | net_input_bytes | 2254966536.000 |
| redis | net_output_bytes | 7210014535.000 |
| redis | ops_per_sec | 19731.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93894864.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
