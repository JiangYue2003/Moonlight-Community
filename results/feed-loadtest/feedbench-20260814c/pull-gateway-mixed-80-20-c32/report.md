# Feed 压测报告：pull / gateway / mixed-80-20-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:42:08+08:00
- 采样时长：8.4868731s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 339.15 | 70.975 | 92.499 | 99.667 |
| publish_total | 60 | 60 | 0 | 0 | 7.07 | 847.151 | 1561.544 | 1569.504 |
| publish_draft | 60 | 60 | 0 | 0 | 7.07 | 8.536 | 27.279 | 28.888 |
| publish_metadata | 60 | 60 | 0 | 0 | 7.07 | 277.776 | 955.705 | 972.492 |
| publish_confirm | 60 | 60 | 0 | 0 | 7.07 | 279.798 | 317.300 | 321.438 |
| publish_commit | 60 | 60 | 0 | 0 | 7.07 | 280.331 | 379.599 | 386.347 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.180 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 6.110 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 137.990 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 26.500 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 360628.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 55.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6164087.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 537827.000 |
| redis | keyspace_hits | 8358969.000 |
| redis | keyspace_misses | 99838.000 |
| redis | net_input_bytes | 690445531.000 |
| redis | net_output_bytes | 2003929338.000 |
| redis | ops_per_sec | 9746.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86294504.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
