# Feed 压测报告：hybrid / rpc / distributed-read-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:44+08:00
- 采样时长：3.008334s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2204 | 2204 | 0 | 0 | 732.63 | 10.663 | 13.471 | 14.677 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 57.330 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.190 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 193.540 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 102.310 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 387031.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 115.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7359520.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 538713.000 |
| redis | keyspace_hits | 9726675.000 |
| redis | keyspace_misses | 106787.000 |
| redis | net_input_bytes | 783627697.000 |
| redis | net_output_bytes | 2234705801.000 |
| redis | ops_per_sec | 23761.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90027432.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
