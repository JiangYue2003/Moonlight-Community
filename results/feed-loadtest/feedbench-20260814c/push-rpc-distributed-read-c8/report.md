# Feed 压测报告：push / rpc / distributed-read-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:49+08:00
- 采样时长：3.0060573s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3459 | 3459 | 0 | 0 | 1150.88 | 6.824 | 9.000 | 10.910 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 5.290 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.300 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 156.800 |
| docker:zg-knowpost | memory_percent | 0.400 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 140.510 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 224600.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 124.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1661994.000 |
| redis | connected_clients | 418.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 536240.000 |
| redis | keyspace_hits | 1427040.000 |
| redis | keyspace_misses | 15563.000 |
| redis | net_input_bytes | 152540491.000 |
| redis | net_output_bytes | 268184029.000 |
| redis | ops_per_sec | 3685.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 88304968.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
