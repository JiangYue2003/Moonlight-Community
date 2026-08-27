# Feed 压测报告：pull / rpc / publish-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:35:06+08:00
- 采样时长：2.7717224s
- 并发：128
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 512 | 2 | 510 | 0 | 0.72 | 740.326 | 745.965 | 745.965 |
| publish_draft | 512 | 512 | 0 | 0 | 184.72 | 25.909 | 68.799 | 82.031 |
| publish_metadata | 512 | 2 | 510 | 0 | 0.72 | 275.991 | 285.571 | 285.571 |
| publish_confirm | 2 | 2 | 0 | 0 | 0.72 | 210.871 | 211.394 | 211.394 |
| publish_commit | 2 | 2 | 0 | 0 | 0.72 | 219.353 | 219.353 | 219.353 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 1.920 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.570 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 109.320 |
| docker:zg-knowpost | memory_percent | 0.300 |
| docker:zg-knowpost | pids | 23.000 |
| docker:zg-relation | cpu_percent | 0.300 |
| docker:zg-relation | memory_percent | 0.260 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 312538.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 67.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4212125.000 |
| redis | connected_clients | 388.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 537388.000 |
| redis | keyspace_hits | 5637635.000 |
| redis | keyspace_misses | 47958.000 |
| redis | net_input_bytes | 479143521.000 |
| redis | net_output_bytes | 1086118873.000 |
| redis | ops_per_sec | 8878.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86880208.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
