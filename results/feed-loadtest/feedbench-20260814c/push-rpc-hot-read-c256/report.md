# Feed 压测报告：push / rpc / hot-read-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:38+08:00
- 采样时长：3.0623646s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 9153 | 9153 | 0 | 0 | 2989.43 | 80.722 | 127.779 | 157.689 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 5.350 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.340 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 249.700 |
| docker:zg-knowpost | memory_percent | 0.500 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 219.540 |
| docker:zg-relation | memory_percent | 0.350 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 220413.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 130.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1647631.000 |
| redis | connected_clients | 418.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 536225.000 |
| redis | keyspace_hits | 1252981.000 |
| redis | keyspace_misses | 15548.000 |
| redis | net_input_bytes | 146093447.000 |
| redis | net_output_bytes | 236093837.000 |
| redis | ops_per_sec | 9373.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 94687808.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
