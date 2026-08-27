# Feed 压测报告：pull / gateway / mixed-80-20-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:41:24+08:00
- 采样时长：40.8273545s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 114.85 | 8.500 | 10.588 | 11.648 |
| publish_total | 60 | 60 | 0 | 0 | 1.47 | 670.376 | 731.552 | 878.530 |
| publish_draft | 60 | 60 | 0 | 0 | 1.47 | 6.835 | 7.514 | 9.598 |
| publish_metadata | 60 | 60 | 0 | 0 | 1.47 | 218.495 | 242.817 | 289.009 |
| publish_confirm | 60 | 60 | 0 | 0 | 1.47 | 217.691 | 242.269 | 293.064 |
| publish_commit | 60 | 60 | 0 | 0 | 1.47 | 226.276 | 248.589 | 334.345 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.200 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 18.920 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 72.390 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 19.140 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 359510.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 52.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6086924.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537838.000 |
| redis | keyspace_hits | 8343907.000 |
| redis | keyspace_misses | 97698.000 |
| redis | net_input_bytes | 683911907.000 |
| redis | net_output_bytes | 1992633092.000 |
| redis | ops_per_sec | 3103.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85890624.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
