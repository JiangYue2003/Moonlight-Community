# Feed 压测报告：pull / rpc / mixed-80-20-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:39:14+08:00
- 采样时长：5.8382791s
- 并发：256
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 445.39 | 428.774 | 457.638 | 478.685 |
| publish_total | 60 | 1 | 59 | 0 | 0.17 | 837.847 | 837.847 | 837.847 |
| publish_draft | 60 | 60 | 0 | 0 | 10.28 | 45.621 | 103.787 | 117.875 |
| publish_metadata | 60 | 59 | 1 | 0 | 10.11 | 1925.705 | 1987.803 | 1993.026 |
| publish_confirm | 59 | 59 | 0 | 0 | 10.11 | 1446.379 | 1517.799 | 1543.820 |
| publish_commit | 59 | 1 | 58 | 0 | 0.17 | 228.307 | 228.307 | 228.307 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.500 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.710 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 133.900 |
| docker:zg-knowpost | memory_percent | 0.540 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 3.200 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 345970.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 118.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5459643.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 537439.000 |
| redis | keyspace_hits | 7857265.000 |
| redis | keyspace_misses | 83972.000 |
| redis | net_input_bytes | 626041490.000 |
| redis | net_output_bytes | 1690566871.000 |
| redis | ops_per_sec | 13867.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 88130624.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
