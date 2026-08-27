# Feed 压测报告：pull / rpc / deep-page-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:36:09+08:00
- 采样时长：3.0014088s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 447 | 447 | 0 | 0 | 148.93 | 6.504 | 7.880 | 9.015 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.040 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.620 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 67.860 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 21.230 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 329593.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4671177.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537426.000 |
| redis | keyspace_hits | 6714893.000 |
| redis | keyspace_misses | 63928.000 |
| redis | net_input_bytes | 534065869.000 |
| redis | net_output_bytes | 1402713437.000 |
| redis | ops_per_sec | 4464.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86276040.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
