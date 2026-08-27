# Feed 压测报告：pull / rpc / burst-publish-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T17:00:22+08:00
- 采样时长：12.0459265s
- 并发：32
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 128 | 127 | 1 | 0 | 10.54 | 2989.727 | 3081.573 | 3095.037 |
| publish_draft | 128 | 128 | 0 | 0 | 10.63 | 9.807 | 29.146 | 33.432 |
| publish_metadata | 128 | 127 | 1 | 0 | 10.54 | 1042.176 | 1114.351 | 1117.887 |
| publish_confirm | 127 | 127 | 0 | 0 | 10.54 | 920.460 | 1133.811 | 1137.990 |
| publish_commit | 127 | 127 | 0 | 0 | 10.54 | 983.429 | 1090.606 | 1092.758 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 8.310 |
| docker:zg-counter | memory_percent | 0.200 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 1.920 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 145.830 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 2.360 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2751.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2751.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 597743.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 99.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 13491159.000 |
| redis | connected_clients | 437.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.984 |
| redis | keys | 537627.000 |
| redis | keyspace_hits | 20044401.000 |
| redis | keyspace_misses | 317969.000 |
| redis | net_input_bytes | 1459709808.000 |
| redis | net_output_bytes | 4251218089.000 |
| redis | ops_per_sec | 13155.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86822816.000 |

## 停止施压后的恢复

- Kafka drain：1.0782838s
- 完成：true
- 最终 lag：0

## 说明

- SLA values are reference lines, not pass/fail gates.
