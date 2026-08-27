# Feed 压测报告：push / gateway / mixed-80-20-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:32:20+08:00
- 采样时长：8.7145698s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 927.34 | 22.832 | 57.530 | 63.353 |
| publish_total | 60 | 60 | 0 | 0 | 6.89 | 972.958 | 1210.701 | 1213.577 |
| publish_draft | 60 | 60 | 0 | 0 | 6.89 | 7.975 | 34.970 | 44.831 |
| publish_metadata | 60 | 60 | 0 | 0 | 6.89 | 341.900 | 517.670 | 529.082 |
| publish_confirm | 60 | 60 | 0 | 0 | 6.89 | 298.199 | 335.104 | 339.250 |
| publish_commit | 60 | 60 | 0 | 0 | 6.89 | 321.610 | 382.128 | 396.842 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.160 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 6.300 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 84.450 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 5.980 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 306156.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 90.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3664890.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 538562.000 |
| redis | keyspace_hits | 5626265.000 |
| redis | keyspace_misses | 41363.000 |
| redis | net_input_bytes | 436450240.000 |
| redis | net_output_bytes | 1075032533.000 |
| redis | ops_per_sec | 18037.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 126762432.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
