# Feed 压测报告：push / rpc / steady-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T16:57:19+08:00
- 采样时长：30.034806s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 82405 | 82405 | 0 | 0 | 2744.23 | 46.119 | 61.066 | 68.412 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.700 |
| docker:zg-counter | memory_percent | 0.190 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.470 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 289.790 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 243.700 |
| docker:zg-relation | memory_percent | 0.360 |
| docker:zg-relation | pids | 26.000 |
| kafka | current_offset_total | 2729.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2729.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 567122.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12385429.000 |
| redis | connected_clients | 389.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 539035.000 |
| redis | keyspace_hits | 18322261.000 |
| redis | keyspace_misses | 257950.000 |
| redis | net_input_bytes | 1343201962.000 |
| redis | net_output_bytes | 3929576818.000 |
| redis | ops_per_sec | 8835.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91563128.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
