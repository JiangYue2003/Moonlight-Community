# Feed 压测报告：push / rpc / publish-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:24:42+08:00
- 采样时长：3.5670584s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 32 | 32 | 0 | 0 | 8.97 | 851.042 | 1035.575 | 1042.546 |
| publish_draft | 32 | 32 | 0 | 0 | 8.97 | 7.342 | 16.900 | 17.429 |
| publish_metadata | 32 | 32 | 0 | 0 | 8.97 | 283.442 | 325.560 | 326.979 |
| publish_confirm | 32 | 32 | 0 | 0 | 8.97 | 267.695 | 372.383 | 373.854 |
| publish_commit | 32 | 32 | 0 | 0 | 8.97 | 296.877 | 329.879 | 333.575 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 5.220 |
| docker:zg-counter | memory_percent | 0.100 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.150 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 94.310 |
| docker:zg-knowpost | memory_percent | 0.160 |
| docker:zg-knowpost | pids | 21.000 |
| docker:zg-relation | cpu_percent | 4.410 |
| docker:zg-relation | memory_percent | 0.160 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 188469.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1243208.000 |
| redis | connected_clients | 84.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.855 |
| redis | keys | 535917.000 |
| redis | keyspace_hits | 73149.000 |
| redis | keyspace_misses | 12876.000 |
| redis | net_input_bytes | 80129764.000 |
| redis | net_output_bytes | 14848952.000 |
| redis | ops_per_sec | 13889.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 79895288.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
