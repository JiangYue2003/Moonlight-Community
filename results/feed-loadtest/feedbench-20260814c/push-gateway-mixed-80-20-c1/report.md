# Feed 压测报告：push / gateway / mixed-80-20-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:31:34+08:00
- 采样时长：42.3990328s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 151.58 | 6.452 | 7.953 | 9.025 |
| publish_total | 60 | 60 | 0 | 0 | 1.42 | 697.369 | 780.361 | 879.609 |
| publish_draft | 60 | 60 | 0 | 0 | 1.42 | 6.503 | 7.908 | 9.500 |
| publish_metadata | 60 | 60 | 0 | 0 | 1.42 | 229.249 | 244.895 | 265.223 |
| publish_confirm | 60 | 60 | 0 | 0 | 1.42 | 227.650 | 253.850 | 350.467 |
| publish_commit | 60 | 60 | 0 | 0 | 1.42 | 232.542 | 259.208 | 421.497 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.770 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 12.630 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 31.040 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 25.340 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 304921.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 95.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3555192.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 538430.000 |
| redis | keyspace_hits | 5616477.000 |
| redis | keyspace_misses | 39585.000 |
| redis | net_input_bytes | 427864983.000 |
| redis | net_output_bytes | 1071696957.000 |
| redis | ops_per_sec | 6071.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 125209440.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
