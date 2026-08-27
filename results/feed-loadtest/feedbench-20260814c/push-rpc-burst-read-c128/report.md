# Feed 压测报告：push / rpc / burst-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T16:57:07+08:00
- 采样时长：10.026563s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 27020 | 27020 | 0 | 0 | 2695.27 | 46.271 | 62.438 | 73.917 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 7.140 |
| docker:zg-counter | memory_percent | 0.190 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 1.930 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 278.760 |
| docker:zg-knowpost | memory_percent | 0.370 |
| docker:zg-knowpost | pids | 31.000 |
| docker:zg-relation | cpu_percent | 247.670 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 25.000 |
| kafka | current_offset_total | 2729.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2729.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 484674.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 69.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12132223.000 |
| redis | connected_clients | 388.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 539066.000 |
| redis | keyspace_hits | 14943625.000 |
| redis | keyspace_misses | 175500.000 |
| redis | net_input_bytes | 1217348517.000 |
| redis | net_output_bytes | 3256418120.000 |
| redis | ops_per_sec | 8551.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91393336.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
