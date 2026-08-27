# Feed 压测报告：hybrid / rpc / publish-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:01+08:00
- 采样时长：2.8017451s
- 并发：128
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 512 | 1 | 511 | 0 | 0.36 | 721.608 | 721.608 | 721.608 |
| publish_draft | 512 | 512 | 0 | 0 | 182.74 | 28.721 | 70.297 | 71.974 |
| publish_metadata | 512 | 1 | 511 | 0 | 0.36 | 256.794 | 256.794 | 256.794 |
| publish_confirm | 1 | 1 | 0 | 0 | 0.36 | 215.412 | 215.412 | 215.412 |
| publish_commit | 1 | 1 | 0 | 0 | 0.36 | 224.394 | 224.394 | 224.394 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.660 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.180 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 115.770 |
| docker:zg-knowpost | memory_percent | 0.320 |
| docker:zg-knowpost | pids | 23.000 |
| docker:zg-relation | cpu_percent | 2.570 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 367107.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6732070.000 |
| redis | connected_clients | 383.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.987 |
| redis | keys | 538722.000 |
| redis | keyspace_hits | 8376452.000 |
| redis | keyspace_misses | 106728.000 |
| redis | net_input_bytes | 734605453.000 |
| redis | net_output_bytes | 2023672956.000 |
| redis | ops_per_sec | 7988.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86200040.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
