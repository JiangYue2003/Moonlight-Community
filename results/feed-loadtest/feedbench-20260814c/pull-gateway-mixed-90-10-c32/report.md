# Feed 压测报告：pull / gateway / mixed-90-10-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:41:07+08:00
- 采样时长：6.4542189s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 343.73 | 80.564 | 98.224 | 99.990 |
| publish_total | 30 | 30 | 0 | 0 | 4.65 | 724.103 | 1460.029 | 1460.546 |
| publish_draft | 30 | 30 | 0 | 0 | 4.65 | 7.456 | 41.208 | 51.386 |
| publish_metadata | 30 | 30 | 0 | 0 | 4.65 | 239.369 | 964.640 | 970.148 |
| publish_confirm | 30 | 30 | 0 | 0 | 4.65 | 232.146 | 245.239 | 245.239 |
| publish_commit | 30 | 30 | 0 | 0 | 4.65 | 229.020 | 291.104 | 293.249 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.830 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 32.010 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 144.480 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 32.450 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 357565.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 75.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5962093.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537716.000 |
| redis | keyspace_hits | 8310378.000 |
| redis | keyspace_misses | 95290.000 |
| redis | net_input_bytes | 673418819.000 |
| redis | net_output_bytes | 1969707994.000 |
| redis | ops_per_sec | 6500.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85985928.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
