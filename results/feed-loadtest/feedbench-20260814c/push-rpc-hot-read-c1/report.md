# Feed 压测报告：push / rpc / hot-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:17+08:00
- 采样时长：3.0003346s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 693 | 693 | 0 | 0 | 230.97 | 4.230 | 5.174 | 5.833 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.350 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 26.000 |
| docker:zg-gateway | cpu_percent | 1.360 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 32.750 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 29.980 |
| docker:zg-relation | memory_percent | 0.160 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 192956.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1561808.000 |
| redis | connected_clients | 290.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.870 |
| redis | keys | 536187.000 |
| redis | keyspace_hits | 103632.000 |
| redis | keyspace_misses | 15518.000 |
| redis | net_input_bytes | 104172108.000 |
| redis | net_output_bytes | 24430437.000 |
| redis | ops_per_sec | 931.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85205144.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
