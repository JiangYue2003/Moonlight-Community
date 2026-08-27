# Feed 压测报告：hybrid / rpc / distributed-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:39+08:00
- 采样时长：3.0012518s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 476 | 476 | 0 | 0 | 158.60 | 6.297 | 7.390 | 8.230 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 14.270 |
| docker:zg-counter | memory_percent | 0.230 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.170 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 41.710 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 22.110 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 384798.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 125.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7285854.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 538715.000 |
| redis | keyspace_hits | 9567982.000 |
| redis | keyspace_misses | 106787.000 |
| redis | net_input_bytes | 777872192.000 |
| redis | net_output_bytes | 2209902347.000 |
| redis | ops_per_sec | 5605.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 88887912.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
