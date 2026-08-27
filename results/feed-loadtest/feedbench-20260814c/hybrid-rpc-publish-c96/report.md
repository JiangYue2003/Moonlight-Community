# Feed 压测报告：hybrid / rpc / publish-c96

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:58:36+08:00
- 采样时长：1.0568791s
- 并发：96
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 384 | 6 | 378 | 0 | 5.68 | 1028.561 | 1051.131 | 1051.131 |
| publish_draft | 384 | 384 | 0 | 0 | 363.33 | 24.410 | 42.944 | 46.684 |
| publish_metadata | 384 | 6 | 378 | 0 | 5.68 | 368.276 | 405.821 | 405.821 |
| publish_confirm | 6 | 6 | 0 | 0 | 5.68 | 309.740 | 312.856 | 312.856 |
| publish_commit | 6 | 6 | 0 | 0 | 5.68 | 318.727 | 323.774 | 323.774 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.100 |
| docker:zg-counter | memory_percent | 0.190 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.140 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 1.290 |
| docker:zg-knowpost | memory_percent | 0.260 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 1.650 |
| docker:zg-relation | memory_percent | 0.260 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 449570.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 10568371.000 |
| redis | connected_clients | 412.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 541049.000 |
| redis | keyspace_hits | 13824667.000 |
| redis | keyspace_misses | 133269.000 |
| redis | net_input_bytes | 1067550527.000 |
| redis | net_output_bytes | 3010887570.000 |
| redis | ops_per_sec | 4927.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 96889272.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
