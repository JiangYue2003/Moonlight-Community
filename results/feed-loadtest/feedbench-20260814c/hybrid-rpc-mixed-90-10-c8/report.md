# Feed 压测报告：hybrid / rpc / mixed-90-10-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:46:56+08:00
- 采样时长：20.716655s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 683.28 | 9.798 | 12.417 | 14.871 |
| publish_total | 30 | 30 | 0 | 0 | 1.45 | 685.655 | 840.583 | 878.433 |
| publish_draft | 30 | 30 | 0 | 0 | 1.45 | 5.332 | 6.544 | 8.507 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.45 | 225.570 | 238.572 | 368.224 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.45 | 222.849 | 237.262 | 244.986 |
| publish_commit | 30 | 30 | 0 | 0 | 1.45 | 227.405 | 240.344 | 407.979 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.970 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 1.630 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 19.840 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 2.090 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 1.000 |
| kafka | lag_total | 1.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 415469.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 124.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8346290.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 538790.000 |
| redis | keyspace_hits | 12673063.000 |
| redis | keyspace_misses | 108845.000 |
| redis | net_input_bytes | 896131512.000 |
| redis | net_output_bytes | 2731663061.000 |
| redis | ops_per_sec | 2507.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89117800.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
