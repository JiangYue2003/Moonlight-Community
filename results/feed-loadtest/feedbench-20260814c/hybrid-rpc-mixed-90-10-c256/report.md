# Feed 压测报告：hybrid / rpc / mixed-90-10-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:47:37+08:00
- 采样时长：3.3553888s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 1190.18 | 157.460 | 205.489 | 214.862 |
| publish_total | 30 | 30 | 0 | 0 | 8.94 | 2587.842 | 2605.246 | 2607.372 |
| publish_draft | 30 | 30 | 0 | 0 | 8.94 | 39.489 | 100.199 | 100.199 |
| publish_metadata | 30 | 30 | 0 | 0 | 8.94 | 1024.146 | 1050.728 | 1063.143 |
| publish_confirm | 30 | 30 | 0 | 0 | 8.94 | 748.368 | 755.069 | 755.568 |
| publish_commit | 30 | 30 | 0 | 0 | 8.94 | 756.210 | 762.601 | 766.046 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.450 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.340 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 125.620 |
| docker:zg-knowpost | memory_percent | 0.490 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 3.950 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 417789.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 93.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8497345.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 538910.000 |
| redis | keyspace_hits | 12729904.000 |
| redis | keyspace_misses | 112077.000 |
| redis | net_input_bytes | 908122521.000 |
| redis | net_output_bytes | 2743327808.000 |
| redis | ops_per_sec | 13224.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90363280.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
