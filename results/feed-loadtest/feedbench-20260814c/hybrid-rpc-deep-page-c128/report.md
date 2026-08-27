# Feed 压测报告：hybrid / rpc / deep-page-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:46:22+08:00
- 采样时长：3.0711794s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3604 | 3604 | 0 | 0 | 1173.70 | 107.360 | 136.211 | 151.540 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 72.590 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 3.350 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 290.930 |
| docker:zg-knowpost | memory_percent | 0.500 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 128.670 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 410071.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 81.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8116392.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 538760.000 |
| redis | keyspace_hits | 12080180.000 |
| redis | keyspace_misses | 107258.000 |
| redis | net_input_bytes | 867987593.000 |
| redis | net_output_bytes | 2627295396.000 |
| redis | ops_per_sec | 39631.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 97603632.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
