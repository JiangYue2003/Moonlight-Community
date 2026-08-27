# Feed 压测报告：hybrid / gateway / publish-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:50:06+08:00
- 采样时长：4.2808309s
- 并发：128
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 512 | 0 | 512 | 0 | 0.00 | 0.000 | 0.000 | 0.000 |
| publish_draft | 512 | 512 | 0 | 0 | 119.60 | 32.758 | 84.351 | 91.672 |
| publish_metadata | 512 | 0 | 512 | 0 | 0.00 | 0.000 | 0.000 | 0.000 |
| publish_confirm | 0 | 0 | 0 | 0 | 0.00 | 0.000 | 0.000 | 0.000 |
| publish_commit | 0 | 0 | 0 | 0 | 0.00 | 0.000 | 0.000 | 0.000 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.050 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 1.120 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 126.010 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 1.390 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 427141.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 67.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9289225.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 539755.000 |
| redis | keyspace_hits | 12815257.000 |
| redis | keyspace_misses | 123736.000 |
| redis | net_input_bytes | 969203870.000 |
| redis | net_output_bytes | 2771645894.000 |
| redis | ops_per_sec | 11811.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95805568.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
