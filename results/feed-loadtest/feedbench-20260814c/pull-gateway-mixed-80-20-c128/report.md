# Feed 压测报告：pull / gateway / mixed-80-20-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:42:20+08:00
- 采样时长：6.7138954s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 332.60 | 279.015 | 364.520 | 381.596 |
| publish_total | 60 | 60 | 0 | 0 | 8.94 | 2571.664 | 3147.683 | 3158.224 |
| publish_draft | 60 | 60 | 0 | 0 | 8.94 | 12.415 | 97.628 | 131.794 |
| publish_metadata | 60 | 60 | 0 | 0 | 8.94 | 883.779 | 1544.708 | 1552.267 |
| publish_confirm | 60 | 60 | 0 | 0 | 8.94 | 736.260 | 794.302 | 803.235 |
| publish_commit | 60 | 60 | 0 | 0 | 8.94 | 821.774 | 892.588 | 898.239 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.970 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 3.850 |
| docker:zg-gateway | memory_percent | 0.190 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 136.980 |
| docker:zg-knowpost | memory_percent | 0.460 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 0.920 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 361824.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 87.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6240396.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 537883.000 |
| redis | keyspace_hits | 8374710.000 |
| redis | keyspace_misses | 101301.000 |
| redis | net_input_bytes | 696821488.000 |
| redis | net_output_bytes | 2015464840.000 |
| redis | ops_per_sec | 12806.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87169208.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
