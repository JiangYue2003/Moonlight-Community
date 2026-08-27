# Feed 压测报告：pull / rpc / deep-page-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:36:14+08:00
- 采样时长：3.0120543s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1434 | 1434 | 0 | 0 | 476.09 | 16.570 | 20.689 | 22.803 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.300 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.730 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 230.210 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 72.730 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 331044.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 71.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4712239.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537421.000 |
| redis | keyspace_hits | 6924256.000 |
| redis | keyspace_misses | 65368.000 |
| redis | net_input_bytes | 542911635.000 |
| redis | net_output_bytes | 1448079075.000 |
| redis | ops_per_sec | 13911.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89198736.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
