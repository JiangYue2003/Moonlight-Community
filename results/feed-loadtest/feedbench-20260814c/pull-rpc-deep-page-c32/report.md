# Feed 压测报告：pull / rpc / deep-page-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:36:19+08:00
- 采样时长：3.0346286s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1669 | 1669 | 0 | 0 | 549.98 | 56.675 | 75.564 | 82.649 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.250 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.780 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 242.220 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 79.310 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 332791.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 71.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4759809.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537408.000 |
| redis | keyspace_hits | 7167785.000 |
| redis | keyspace_misses | 67185.000 |
| redis | net_input_bytes | 553212561.000 |
| redis | net_output_bytes | 1500847318.000 |
| redis | ops_per_sec | 14689.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92693256.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
