# Feed 压测报告：push / rpc / publish-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:24:18+08:00
- 采样时长：18.8698516s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 30 | 30 | 0 | 0 | 1.59 | 624.827 | 664.068 | 669.526 |
| publish_draft | 30 | 30 | 0 | 0 | 1.59 | 5.337 | 12.052 | 19.650 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.59 | 204.649 | 226.488 | 235.381 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.59 | 203.476 | 217.675 | 235.667 |
| publish_commit | 30 | 30 | 0 | 0 | 1.59 | 209.222 | 221.078 | 236.409 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.560 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.930 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 20.610 |
| docker:zg-knowpost | memory_percent | 0.160 |
| docker:zg-knowpost | pids | 21.000 |
| docker:zg-relation | cpu_percent | 2.960 |
| docker:zg-relation | memory_percent | 0.160 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 188027.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 6.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1201891.000 |
| redis | connected_clients | 74.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.860 |
| redis | keys | 535892.000 |
| redis | keyspace_hits | 73027.000 |
| redis | keyspace_misses | 12374.000 |
| redis | net_input_bytes | 76930131.000 |
| redis | net_output_bytes | 14152858.000 |
| redis | ops_per_sec | 8301.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 79334776.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
