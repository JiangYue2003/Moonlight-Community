# Feed 压测报告：push / rpc / deep-page-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:26:26+08:00
- 采样时长：3.0534993s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 5204 | 5204 | 0 | 0 | 1704.57 | 71.783 | 108.155 | 130.474 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 1.920 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.230 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 296.480 |
| docker:zg-knowpost | memory_percent | 0.510 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 173.860 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 261246.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 120.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1777401.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 536383.000 |
| redis | keyspace_hits | 3945919.000 |
| redis | keyspace_misses | 16401.000 |
| redis | net_input_bytes | 243258931.000 |
| redis | net_output_bytes | 738944807.000 |
| redis | ops_per_sec | 5483.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93888312.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
