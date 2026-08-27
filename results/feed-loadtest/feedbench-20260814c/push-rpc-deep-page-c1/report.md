# Feed 压测报告：push / rpc / deep-page-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:26:10+08:00
- 采样时长：3.0060573s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 569 | 569 | 0 | 0 | 189.28 | 5.247 | 6.466 | 7.422 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.310 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.580 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 33.950 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 25.640 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 248816.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1738310.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 536358.000 |
| redis | keyspace_hits | 2486633.000 |
| redis | keyspace_misses | 15676.000 |
| redis | net_input_bytes | 191179240.000 |
| redis | net_output_bytes | 463623295.000 |
| redis | ops_per_sec | 800.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 88590392.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
