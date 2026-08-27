# Feed 压测报告：hybrid / gateway / mixed-90-10-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:51:12+08:00
- 采样时长：6.3966353s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 711.26 | 37.963 | 48.349 | 51.563 |
| publish_total | 30 | 30 | 0 | 0 | 4.69 | 751.760 | 1148.126 | 1148.126 |
| publish_draft | 30 | 30 | 0 | 0 | 4.69 | 7.431 | 30.817 | 33.501 |
| publish_metadata | 30 | 30 | 0 | 0 | 4.69 | 247.229 | 589.264 | 590.291 |
| publish_confirm | 30 | 30 | 0 | 0 | 4.69 | 248.501 | 270.541 | 270.946 |
| publish_commit | 30 | 30 | 0 | 0 | 4.69 | 256.450 | 268.162 | 270.307 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.380 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.580 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 58.320 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 27.770 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 3.000 |
| kafka | lag_total | 3.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 441524.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 71.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9814588.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 540122.000 |
| redis | keyspace_hits | 13752173.000 |
| redis | keyspace_misses | 124874.000 |
| redis | net_input_bytes | 1010162573.000 |
| redis | net_output_bytes | 2983051363.000 |
| redis | ops_per_sec | 7383.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93124528.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
