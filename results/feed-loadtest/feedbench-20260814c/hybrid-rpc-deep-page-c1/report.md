# Feed 压测报告：hybrid / rpc / deep-page-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:46:05+08:00
- 采样时长：3.0030764s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 427 | 427 | 0 | 0 | 142.19 | 6.862 | 8.616 | 9.800 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 13.080 |
| docker:zg-counter | memory_percent | 0.240 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.150 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 41.340 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 20.050 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 401273.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7831471.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 538767.000 |
| redis | keyspace_hits | 10782167.000 |
| redis | keyspace_misses | 106870.000 |
| redis | net_input_bytes | 821729455.000 |
| redis | net_output_bytes | 2400783516.000 |
| redis | ops_per_sec | 5135.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 88904736.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
