# Feed 压测报告：pull / rpc / mixed-90-10-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:37:31+08:00
- 采样时长：3.7886767s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 477.20 | 229.436 | 270.700 | 282.411 |
| publish_total | 30 | 30 | 0 | 0 | 7.92 | 1162.932 | 1907.713 | 1907.713 |
| publish_draft | 30 | 30 | 0 | 0 | 7.92 | 8.491 | 92.195 | 93.761 |
| publish_metadata | 30 | 30 | 0 | 0 | 7.92 | 397.849 | 985.184 | 993.189 |
| publish_confirm | 30 | 30 | 0 | 0 | 7.92 | 373.144 | 465.804 | 465.971 |
| publish_commit | 30 | 30 | 0 | 0 | 7.92 | 388.343 | 398.494 | 399.014 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.230 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.480 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 116.820 |
| docker:zg-knowpost | memory_percent | 0.450 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 2.280 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 339849.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 60.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5037856.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537400.000 |
| redis | keyspace_hits | 7761900.000 |
| redis | keyspace_misses | 75283.000 |
| redis | net_input_bytes | 590629263.000 |
| redis | net_output_bytes | 1640183004.000 |
| redis | ops_per_sec | 12443.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 88627768.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
