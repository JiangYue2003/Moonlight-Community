# Feed 压测报告：hybrid / rpc / burst-publish-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T17:03:05+08:00
- 采样时长：12.6407151s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 128 | 128 | 0 | 0 | 10.13 | 3151.892 | 3310.403 | 3333.011 |
| publish_draft | 128 | 128 | 0 | 0 | 10.13 | 9.627 | 27.845 | 31.789 |
| publish_metadata | 128 | 128 | 0 | 0 | 10.13 | 1059.306 | 1106.643 | 1109.952 |
| publish_confirm | 128 | 128 | 0 | 0 | 10.13 | 1008.354 | 1140.281 | 1150.321 |
| publish_commit | 128 | 128 | 0 | 0 | 10.13 | 1011.489 | 1217.865 | 1229.473 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 7.030 |
| docker:zg-counter | memory_percent | 0.250 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.590 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 142.040 |
| docker:zg-knowpost | memory_percent | 0.380 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 9.800 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2926.000 |
| kafka | lag_max | 28.000 |
| kafka | lag_total | 28.000 |
| kafka | log_end_offset_total | 2954.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 665025.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 99.000 |
| mysql | threads_running | 7.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 15954483.000 |
| redis | connected_clients | 428.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 538753.000 |
| redis | keyspace_hits | 24539851.000 |
| redis | keyspace_misses | 388270.000 |
| redis | net_input_bytes | 1652032045.000 |
| redis | net_output_bytes | 4892849791.000 |
| redis | ops_per_sec | 17527.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87445720.000 |

## 停止施压后的恢复

- Kafka drain：1.1390657s
- 完成：true
- 最终 lag：0

## 说明

- SLA values are reference lines, not pass/fail gates.
