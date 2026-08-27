# Feed 压测报告：push / rpc / deep-page-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:26:21+08:00
- 采样时长：3.0144545s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4085 | 4085 | 0 | 0 | 1355.38 | 22.906 | 31.813 | 37.983 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.680 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.590 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 246.730 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 156.110 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 255724.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 80.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1760507.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 536372.000 |
| redis | keyspace_hits | 3311625.000 |
| redis | keyspace_misses | 15795.000 |
| redis | net_input_bytes | 220576174.000 |
| redis | net_output_bytes | 619269756.000 |
| redis | ops_per_sec | 4330.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90090776.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
