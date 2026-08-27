# Feed 压测报告：push / gateway / publish-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:29:26+08:00
- 采样时长：19.8865567s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 30 | 30 | 0 | 0 | 1.51 | 660.828 | 703.468 | 706.456 |
| publish_draft | 30 | 30 | 0 | 0 | 1.51 | 6.401 | 6.989 | 7.907 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.51 | 217.548 | 238.082 | 239.059 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.51 | 215.673 | 225.579 | 225.660 |
| publish_commit | 30 | 30 | 0 | 0 | 1.51 | 220.411 | 244.477 | 251.938 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.770 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.200 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 19.990 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 2.310 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 277497.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 132.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2839508.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 537432.000 |
| redis | keyspace_hits | 4723508.000 |
| redis | keyspace_misses | 33176.000 |
| redis | net_input_bytes | 347668005.000 |
| redis | net_output_bytes | 898444081.000 |
| redis | ops_per_sec | 6114.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 114933256.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
