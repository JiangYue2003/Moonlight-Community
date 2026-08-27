# Feed 压测报告：hybrid / rpc / hot-read-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:23+08:00
- 采样时长：3.0168035s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3848 | 3848 | 0 | 0 | 1275.77 | 24.530 | 31.508 | 39.364 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 76.510 |
| docker:zg-counter | memory_percent | 0.140 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.280 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 238.020 |
| docker:zg-knowpost | memory_percent | 0.380 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 135.130 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 374859.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 99.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6956776.000 |
| redis | connected_clients | 404.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 538739.000 |
| redis | keyspace_hits | 8856175.000 |
| redis | keyspace_misses | 106784.000 |
| redis | net_input_bytes | 752148005.000 |
| redis | net_output_bytes | 2098674219.000 |
| redis | ops_per_sec | 43032.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89762640.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
