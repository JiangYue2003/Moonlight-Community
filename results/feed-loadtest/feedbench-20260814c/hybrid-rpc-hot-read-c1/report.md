# Feed 压测报告：hybrid / rpc / hot-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:13+08:00
- 采样时长：3.0036508s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 481 | 481 | 0 | 0 | 160.14 | 6.251 | 7.390 | 8.070 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 14.340 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.430 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 40.430 |
| docker:zg-knowpost | memory_percent | 0.340 |
| docker:zg-knowpost | pids | 23.000 |
| docker:zg-relation | cpu_percent | 21.500 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 368634.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6752616.000 |
| redis | connected_clients | 398.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.987 |
| redis | keys | 538749.000 |
| redis | keyspace_hits | 8411060.000 |
| redis | keyspace_misses | 106784.000 |
| redis | net_input_bytes | 736175302.000 |
| redis | net_output_bytes | 2029158919.000 |
| redis | ops_per_sec | 5586.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86442880.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
