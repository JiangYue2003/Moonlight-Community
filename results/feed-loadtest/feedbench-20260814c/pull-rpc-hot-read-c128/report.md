# Feed 压测报告：pull / rpc / hot-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:35:32+08:00
- 采样时长：3.1127065s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1862 | 1862 | 0 | 0 | 598.19 | 210.455 | 247.701 | 275.320 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.230 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.310 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 268.680 |
| docker:zg-knowpost | memory_percent | 0.460 |
| docker:zg-knowpost | pids | 30.000 |
| docker:zg-relation | cpu_percent | 82.520 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 319378.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 130.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4381109.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537398.000 |
| redis | keyspace_hits | 6012046.000 |
| redis | keyspace_misses | 53734.000 |
| redis | net_input_bytes | 498732999.000 |
| redis | net_output_bytes | 1198087517.000 |
| redis | ops_per_sec | 17437.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92631736.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
