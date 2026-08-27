# Feed 压测报告：pull / gateway / publish-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:40:02+08:00
- 采样时长：4.6613784s
- 并发：128
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 512 | 3 | 509 | 0 | 0.64 | 2447.918 | 2459.919 | 2459.919 |
| publish_draft | 512 | 512 | 0 | 0 | 109.84 | 36.771 | 67.549 | 79.588 |
| publish_metadata | 512 | 3 | 509 | 0 | 0.64 | 1992.932 | 1994.301 | 1994.301 |
| publish_confirm | 3 | 3 | 0 | 0 | 0.64 | 211.218 | 216.898 | 216.898 |
| publish_commit | 3 | 3 | 0 | 0 | 0.64 | 213.995 | 218.199 | 218.199 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.300 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 16.410 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 112.110 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 0.290 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 349369.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 103.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5690258.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537459.000 |
| redis | keyspace_hits | 7857903.000 |
| redis | keyspace_misses | 87070.000 |
| redis | net_input_bytes | 644344563.000 |
| redis | net_output_bytes | 1694892052.000 |
| redis | ops_per_sec | 10918.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 88620152.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
