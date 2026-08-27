# Feed 压测报告：hybrid / rpc / publish-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:09+08:00
- 采样时长：906.1749ms
- 并发：256
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 1024 | 1 | 1023 | 0 | 1.10 | 904.588 | 904.588 | 904.588 |
| publish_draft | 1024 | 1024 | 0 | 0 | 1130.02 | 50.173 | 116.234 | 157.261 |
| publish_metadata | 1024 | 1 | 1023 | 0 | 1.10 | 433.388 | 433.388 | 433.388 |
| publish_confirm | 1 | 1 | 0 | 0 | 1.10 | 225.152 | 225.152 | 225.152 |
| publish_commit | 1 | 1 | 0 | 0 | 1.10 | 236.982 | 236.982 | 236.982 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.030 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.130 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 0.410 |
| docker:zg-knowpost | memory_percent | 0.350 |
| docker:zg-knowpost | pids | 23.000 |
| docker:zg-relation | cpu_percent | 0.250 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 368148.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6735688.000 |
| redis | connected_clients | 398.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.987 |
| redis | keys | 538719.000 |
| redis | keyspace_hits | 8376463.000 |
| redis | keyspace_misses | 106744.000 |
| redis | net_input_bytes | 734854259.000 |
| redis | net_output_bytes | 2023731523.000 |
| redis | ops_per_sec | 1558.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87212032.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
