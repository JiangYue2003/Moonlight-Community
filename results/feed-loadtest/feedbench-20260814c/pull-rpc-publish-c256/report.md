# Feed 压测报告：pull / rpc / publish-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:35:12+08:00
- 采样时长：943.1809ms
- 并发：256
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 1024 | 4 | 1020 | 0 | 4.24 | 908.374 | 915.893 | 915.893 |
| publish_draft | 1024 | 1024 | 0 | 0 | 1085.69 | 43.236 | 112.250 | 141.577 |
| publish_metadata | 1024 | 4 | 1020 | 0 | 4.24 | 411.843 | 418.317 | 418.317 |
| publish_confirm | 4 | 4 | 0 | 0 | 4.24 | 240.551 | 243.906 | 243.906 |
| publish_commit | 4 | 4 | 0 | 0 | 4.24 | 236.260 | 239.744 | 239.744 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 1.880 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.110 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 1.660 |
| docker:zg-knowpost | memory_percent | 0.330 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 0.180 |
| docker:zg-relation | memory_percent | 0.260 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 313615.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 67.000 |
| mysql | threads_running | 9.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4218376.000 |
| redis | connected_clients | 388.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 537385.000 |
| redis | keyspace_hits | 5637647.000 |
| redis | keyspace_misses | 48021.000 |
| redis | net_input_bytes | 479611637.000 |
| redis | net_output_bytes | 1086225020.000 |
| redis | ops_per_sec | 3367.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85896592.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
