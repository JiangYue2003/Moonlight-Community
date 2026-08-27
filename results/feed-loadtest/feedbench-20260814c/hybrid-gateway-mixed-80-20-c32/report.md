# Feed 压测报告：hybrid / gateway / mixed-80-20-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:52:19+08:00
- 采样时长：8.8855478s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 650.88 | 35.447 | 52.443 | 61.873 |
| publish_total | 60 | 60 | 0 | 0 | 6.75 | 979.541 | 1311.359 | 1320.570 |
| publish_draft | 60 | 60 | 0 | 0 | 6.75 | 8.000 | 20.800 | 25.837 |
| publish_metadata | 60 | 60 | 0 | 0 | 6.75 | 306.481 | 633.813 | 640.249 |
| publish_confirm | 60 | 60 | 0 | 0 | 6.75 | 334.773 | 381.498 | 385.021 |
| publish_commit | 60 | 60 | 0 | 0 | 6.75 | 309.449 | 355.600 | 358.810 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 16.360 |
| docker:zg-counter | memory_percent | 0.200 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 5.280 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 85.840 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 2.110 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 6.000 |
| kafka | lag_total | 6.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 444722.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 40.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 10068431.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 540495.000 |
| redis | keyspace_hits | 13806321.000 |
| redis | keyspace_misses | 128176.000 |
| redis | net_input_bytes | 1029842885.000 |
| redis | net_output_bytes | 2999147446.000 |
| redis | ops_per_sec | 10901.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 94732288.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
