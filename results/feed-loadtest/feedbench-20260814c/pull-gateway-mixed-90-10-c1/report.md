# Feed 压测报告：pull / gateway / mixed-90-10-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:40:43+08:00
- 采样时长：20.8096449s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 120.38 | 8.063 | 10.155 | 11.708 |
| publish_total | 30 | 30 | 0 | 0 | 1.44 | 673.745 | 865.465 | 932.359 |
| publish_draft | 30 | 30 | 0 | 0 | 1.44 | 6.376 | 8.654 | 8.849 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.44 | 217.999 | 264.356 | 267.969 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.44 | 219.840 | 314.989 | 383.145 |
| publish_commit | 30 | 30 | 0 | 0 | 1.44 | 223.054 | 276.203 | 435.754 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.840 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 19.760 |
| docker:zg-gateway | memory_percent | 0.180 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 72.280 |
| docker:zg-knowpost | memory_percent | 0.440 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 18.830 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 356825.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 90.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5918755.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537715.000 |
| redis | keyspace_hits | 8292818.000 |
| redis | keyspace_misses | 94147.000 |
| redis | net_input_bytes | 669629605.000 |
| redis | net_output_bytes | 1958208253.000 |
| redis | ops_per_sec | 4843.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86361896.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
