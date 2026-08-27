# Feed 压测报告：push / rpc / hot-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:33+08:00
- 采样时长：3.0307137s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 8245 | 8245 | 0 | 0 | 2720.97 | 45.629 | 62.708 | 91.483 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 5.870 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.420 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 246.260 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 221.200 |
| docker:zg-relation | memory_percent | 0.270 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 211253.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 130.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1619309.000 |
| redis | connected_clients | 376.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.982 |
| redis | keys | 536214.000 |
| redis | keyspace_hits | 868550.000 |
| redis | keyspace_misses | 15539.000 |
| redis | net_input_bytes | 132098864.000 |
| redis | net_output_bytes | 165309981.000 |
| redis | ops_per_sec | 8441.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92696432.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
