# Feed 压测报告：push / rpc / burst-publish-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T16:57:51+08:00
- 采样时长：2.4587462s
- 并发：32
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 128 | 22 | 106 | 0 | 8.95 | 2318.578 | 2336.120 | 2336.120 |
| publish_draft | 128 | 128 | 0 | 0 | 52.06 | 11.941 | 27.675 | 28.225 |
| publish_metadata | 128 | 22 | 106 | 0 | 8.95 | 986.757 | 1003.414 | 1006.120 |
| publish_confirm | 22 | 22 | 0 | 0 | 8.95 | 661.295 | 668.295 | 672.674 |
| publish_commit | 22 | 22 | 0 | 0 | 8.95 | 642.616 | 652.354 | 653.172 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.120 |
| docker:zg-counter | memory_percent | 0.190 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.010 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 126.880 |
| docker:zg-knowpost | memory_percent | 0.360 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 7.720 |
| docker:zg-relation | memory_percent | 0.360 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2729.000 |
| kafka | lag_max | 22.000 |
| kafka | lag_total | 22.000 |
| kafka | log_end_offset_total | 2751.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 567555.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 99.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12414808.000 |
| redis | connected_clients | 392.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.986 |
| redis | keys | 538939.000 |
| redis | keyspace_hits | 18322312.000 |
| redis | keyspace_misses | 258324.000 |
| redis | net_input_bytes | 1345528956.000 |
| redis | net_output_bytes | 3930123702.000 |
| redis | ops_per_sec | 12550.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86012096.000 |

## 停止施压后的恢复

- Kafka drain：1.1571923s
- 完成：true
- 最终 lag：0

## 说明

- SLA values are reference lines, not pass/fail gates.
