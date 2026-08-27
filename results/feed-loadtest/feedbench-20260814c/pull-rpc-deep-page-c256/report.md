# Feed 压测报告：pull / rpc / deep-page-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:36:30+08:00
- 采样时长：3.2302984s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1850 | 1850 | 0 | 0 | 572.70 | 431.451 | 560.720 | 598.776 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.270 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.680 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 288.970 |
| docker:zg-knowpost | memory_percent | 0.640 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 78.460 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 336959.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 90.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4862698.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537399.000 |
| redis | keyspace_hits | 7691490.000 |
| redis | keyspace_misses | 71224.000 |
| redis | net_input_bytes | 575428485.000 |
| redis | net_output_bytes | 1614347398.000 |
| redis | ops_per_sec | 16331.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93278232.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
