# Feed 压测报告：push / rpc / hot-read-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:28+08:00
- 采样时长：3.0110588s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 6368 | 6368 | 0 | 0 | 2114.87 | 14.826 | 19.653 | 23.080 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 5.350 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 3.190 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 231.040 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 206.980 |
| docker:zg-relation | memory_percent | 0.180 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 202964.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 92.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1593541.000 |
| redis | connected_clients | 285.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.971 |
| redis | keys | 536205.000 |
| redis | keyspace_hits | 522256.000 |
| redis | keyspace_misses | 15533.000 |
| redis | net_input_bytes | 119475146.000 |
| redis | net_output_bytes | 101532067.000 |
| redis | ops_per_sec | 6656.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86835120.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
