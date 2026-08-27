# Feed 压测报告：hybrid / rpc / deep-page-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:46:27+08:00
- 采样时长：3.1687701s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3647 | 3647 | 0 | 0 | 1150.92 | 216.209 | 276.740 | 305.585 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 70.500 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.240 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 282.750 |
| docker:zg-knowpost | memory_percent | 0.580 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 128.700 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 414051.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 124.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8238003.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 538754.000 |
| redis | keyspace_hits | 12634124.000 |
| redis | keyspace_misses | 107662.000 |
| redis | net_input_bytes | 887770089.000 |
| redis | net_output_bytes | 2723962209.000 |
| redis | ops_per_sec | 35875.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 97567480.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
