# Feed 压测报告：pull / rpc / hot-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:35:16+08:00
- 采样时长：3.003524s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 491 | 491 | 0 | 0 | 163.47 | 5.874 | 7.252 | 8.010 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.060 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.650 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 67.980 |
| docker:zg-knowpost | memory_percent | 0.350 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 22.720 |
| docker:zg-relation | memory_percent | 0.260 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 314114.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4233164.000 |
| redis | connected_clients | 388.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 537415.000 |
| redis | keyspace_hits | 5670019.000 |
| redis | keyspace_misses | 48552.000 |
| redis | net_input_bytes | 481320685.000 |
| redis | net_output_bytes | 1095920592.000 |
| redis | ops_per_sec | 4992.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85577144.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
