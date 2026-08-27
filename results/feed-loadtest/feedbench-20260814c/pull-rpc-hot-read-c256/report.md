# Feed 压测报告：pull / rpc / hot-read-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:35:37+08:00
- 采样时长：3.23361s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1951 | 1951 | 0 | 0 | 603.45 | 415.687 | 512.097 | 553.677 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.280 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.270 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 283.070 |
| docker:zg-knowpost | memory_percent | 0.600 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 85.560 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 321336.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 130.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4436726.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537393.000 |
| redis | keyspace_hits | 6140817.000 |
| redis | keyspace_misses | 55688.000 |
| redis | net_input_bytes | 505283582.000 |
| redis | net_output_bytes | 1236547554.000 |
| redis | ops_per_sec | 17582.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92963544.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
