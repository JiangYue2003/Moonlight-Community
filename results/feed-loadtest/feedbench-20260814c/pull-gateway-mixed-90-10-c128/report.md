# Feed 压测报告：pull / gateway / mixed-90-10-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:41:17+08:00
- 采样时长：4.0554025s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 361.68 | 296.449 | 426.195 | 439.816 |
| publish_total | 30 | 30 | 0 | 0 | 7.40 | 1220.781 | 1980.571 | 1980.571 |
| publish_draft | 30 | 30 | 0 | 0 | 7.40 | 10.200 | 60.253 | 71.249 |
| publish_metadata | 30 | 30 | 0 | 0 | 7.40 | 407.123 | 1162.436 | 1164.481 |
| publish_confirm | 30 | 30 | 0 | 0 | 7.40 | 381.741 | 410.284 | 412.550 |
| publish_commit | 30 | 30 | 0 | 0 | 7.40 | 388.160 | 427.433 | 427.434 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.200 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 6.860 |
| docker:zg-gateway | memory_percent | 0.180 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 118.320 |
| docker:zg-knowpost | memory_percent | 0.500 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 0.670 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 358369.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 85.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6004973.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537770.000 |
| redis | keyspace_hits | 8327845.000 |
| redis | keyspace_misses | 96520.000 |
| redis | net_input_bytes | 677189237.000 |
| redis | net_output_bytes | 1981281503.000 |
| redis | ops_per_sec | 12313.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86289840.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
