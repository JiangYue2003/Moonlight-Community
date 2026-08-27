# Feed 压测报告：hybrid / rpc / steady-read-c128

- Run ID：`feedopt-20260814-route-pipeline`
- 开始时间：2026-08-14T20:31:25+08:00
- 采样时长：30.0371552s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 69675 | 69675 | 0 | 0 | 2320.08 | 54.047 | 72.870 | 86.214 |

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
| docker:zg-counter | cpu_percent | 7.030 |
| docker:zg-counter | memory_percent | 0.250 |
| docker:zg-counter | pids | 27.000 |
| docker:zg-gateway | cpu_percent | 2.240 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 22.000 |
| docker:zg-knowpost | cpu_percent | 294.930 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 223.800 |
| docker:zg-relation | memory_percent | 0.420 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2956.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2956.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 754250.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 131.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19056720.000 |
| redis | connected_clients | 385.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.984 |
| redis | keys | 541329.000 |
| redis | keyspace_hits | 28409509.000 |
| redis | keyspace_misses | 484506.000 |
| redis | net_input_bytes | 1982150716.000 |
| redis | net_output_bytes | 5782337861.000 |
| redis | ops_per_sec | 19679.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 94406792.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
