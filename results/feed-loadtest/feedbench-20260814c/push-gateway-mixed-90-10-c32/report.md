# Feed 压测报告：push / gateway / mixed-90-10-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:31:14+08:00
- 采样时长：6.2264438s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 914.25 | 26.799 | 53.178 | 57.883 |
| publish_total | 30 | 30 | 0 | 0 | 4.82 | 732.087 | 1063.727 | 1065.630 |
| publish_draft | 30 | 30 | 0 | 0 | 4.82 | 7.784 | 35.886 | 35.886 |
| publish_metadata | 30 | 30 | 0 | 0 | 4.82 | 252.627 | 509.148 | 519.976 |
| publish_confirm | 30 | 30 | 0 | 0 | 4.82 | 254.781 | 279.414 | 281.900 |
| publish_commit | 30 | 30 | 0 | 0 | 4.82 | 244.693 | 280.272 | 280.844 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.150 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 25.390 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 56.790 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 25.440 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 4.000 |
| kafka | lag_total | 4.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 302825.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 94.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3352078.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 538157.000 |
| redis | keyspace_hits | 5594027.000 |
| redis | keyspace_misses | 37959.000 |
| redis | net_input_bytes | 412398111.000 |
| redis | net_output_bytes | 1064819731.000 |
| redis | ops_per_sec | 11182.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 122101216.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
