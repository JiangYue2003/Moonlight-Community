# Feed 压测报告：hybrid / rpc / publish-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:44:36+08:00
- 采样时长：3.9321258s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 32 | 32 | 0 | 0 | 8.14 | 961.474 | 1048.121 | 1060.732 |
| publish_draft | 32 | 32 | 0 | 0 | 8.14 | 7.007 | 15.206 | 15.206 |
| publish_metadata | 32 | 32 | 0 | 0 | 8.14 | 312.160 | 332.469 | 335.490 |
| publish_confirm | 32 | 32 | 0 | 0 | 8.14 | 327.076 | 383.928 | 385.400 |
| publish_commit | 32 | 32 | 0 | 0 | 8.14 | 303.940 | 356.579 | 356.664 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.110 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.570 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 88.970 |
| docker:zg-knowpost | memory_percent | 0.160 |
| docker:zg-knowpost | pids | 21.000 |
| docker:zg-relation | cpu_percent | 2.750 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 364536.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6521538.000 |
| redis | connected_clients | 228.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 538637.000 |
| redis | keyspace_hits | 8375822.000 |
| redis | keyspace_misses | 104322.000 |
| redis | net_input_bytes | 718410438.000 |
| redis | net_output_bytes | 2020177969.000 |
| redis | ops_per_sec | 11451.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 82750912.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
