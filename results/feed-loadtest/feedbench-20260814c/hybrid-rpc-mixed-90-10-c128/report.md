# Feed 压测报告：hybrid / rpc / mixed-90-10-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:47:31+08:00
- 采样时长：3.5724996s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 1083.29 | 88.598 | 133.619 | 138.995 |
| publish_total | 30 | 30 | 0 | 0 | 8.40 | 1254.840 | 1614.963 | 1616.592 |
| publish_draft | 30 | 30 | 0 | 0 | 8.40 | 8.022 | 59.924 | 59.924 |
| publish_metadata | 30 | 30 | 0 | 0 | 8.40 | 411.094 | 689.993 | 692.726 |
| publish_confirm | 30 | 30 | 0 | 0 | 8.40 | 424.601 | 445.595 | 446.113 |
| publish_commit | 30 | 30 | 0 | 0 | 8.40 | 418.267 | 438.697 | 442.898 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 12.040 |
| docker:zg-counter | memory_percent | 0.200 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.180 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 116.990 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 0.920 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 417072.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 80.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8448798.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 538871.000 |
| redis | keyspace_hits | 12711145.000 |
| redis | keyspace_misses | 110833.000 |
| redis | net_input_bytes | 904216872.000 |
| redis | net_output_bytes | 2739402830.000 |
| redis | ops_per_sec | 13089.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89747184.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
