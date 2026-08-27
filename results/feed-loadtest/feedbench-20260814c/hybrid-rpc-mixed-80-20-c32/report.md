# Feed 压测报告：hybrid / rpc / mixed-80-20-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:48:52+08:00
- 采样时长：8.4010949s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 850.17 | 26.612 | 49.707 | 51.267 |
| publish_total | 60 | 60 | 0 | 0 | 7.14 | 905.098 | 1246.946 | 1249.645 |
| publish_draft | 60 | 60 | 0 | 0 | 7.14 | 6.368 | 17.761 | 21.560 |
| publish_metadata | 60 | 60 | 0 | 0 | 7.14 | 309.516 | 536.791 | 548.380 |
| publish_confirm | 60 | 60 | 0 | 0 | 7.14 | 296.120 | 348.225 | 352.242 |
| publish_commit | 60 | 60 | 0 | 0 | 7.14 | 309.513 | 342.086 | 346.060 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 16.490 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.060 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 88.230 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 20.700 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 421221.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 35.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8791784.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539097.000 |
| redis | keyspace_hits | 12782148.000 |
| redis | keyspace_misses | 115454.000 |
| redis | net_input_bytes | 930768633.000 |
| redis | net_output_bytes | 2757300747.000 |
| redis | ops_per_sec | 11633.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90452368.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
