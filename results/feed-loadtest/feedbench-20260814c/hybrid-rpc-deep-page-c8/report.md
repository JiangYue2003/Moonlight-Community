# Feed 压测报告：hybrid / rpc / deep-page-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:46:11+08:00
- 采样时长：3.010566s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1984 | 1984 | 0 | 0 | 659.01 | 11.966 | 14.762 | 16.422 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 56.220 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.750 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 209.240 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 93.520 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 403264.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7897878.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 538766.000 |
| redis | keyspace_hits | 11083739.000 |
| redis | keyspace_misses | 106871.000 |
| redis | net_input_bytes | 832476275.000 |
| redis | net_output_bytes | 2453417404.000 |
| redis | ops_per_sec | 22639.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90046768.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
