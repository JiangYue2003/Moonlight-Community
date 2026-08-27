# Feed 压测报告：hybrid / rpc / mixed-80-20-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:48:27+08:00
- 采样时长：21.018856s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 581.72 | 10.097 | 12.350 | 14.689 |
| publish_total | 60 | 60 | 0 | 0 | 2.85 | 694.736 | 752.834 | 884.788 |
| publish_draft | 60 | 60 | 0 | 0 | 2.85 | 5.789 | 7.870 | 8.875 |
| publish_metadata | 60 | 60 | 0 | 0 | 2.85 | 228.167 | 249.783 | 390.432 |
| publish_confirm | 60 | 60 | 0 | 0 | 2.85 | 225.078 | 242.881 | 246.012 |
| publish_commit | 60 | 60 | 0 | 0 | 2.85 | 229.993 | 247.517 | 252.707 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.830 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.050 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 29.880 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 2.230 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 420112.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 93.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8699145.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539023.000 |
| redis | keyspace_hits | 12764847.000 |
| redis | keyspace_misses | 114268.000 |
| redis | net_input_bytes | 923575140.000 |
| redis | net_output_bytes | 2752618330.000 |
| redis | ops_per_sec | 4736.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89664256.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
