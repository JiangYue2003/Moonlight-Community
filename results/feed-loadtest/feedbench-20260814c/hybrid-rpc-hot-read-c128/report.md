# Feed 压测报告：hybrid / rpc / hot-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:28+08:00
- 采样时长：3.0501991s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4687 | 4687 | 0 | 0 | 1536.62 | 81.479 | 103.197 | 119.803 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 91.800 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 22.000 |
| docker:zg-gateway | cpu_percent | 0.220 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 272.040 |
| docker:zg-knowpost | memory_percent | 0.450 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 156.250 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 379583.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 131.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7112023.000 |
| redis | connected_clients | 460.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 538736.000 |
| redis | keyspace_hits | 9193642.000 |
| redis | keyspace_misses | 106787.000 |
| redis | net_input_bytes | 764290969.000 |
| redis | net_output_bytes | 2151389634.000 |
| redis | ops_per_sec | 51057.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 97170392.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
