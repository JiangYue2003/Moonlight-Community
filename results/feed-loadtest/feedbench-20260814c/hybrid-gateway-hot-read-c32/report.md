# Feed 压测报告：hybrid / gateway / hot-read-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:50:19+08:00
- 采样时长：3.0369732s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2996 | 2996 | 0 | 0 | 986.51 | 31.782 | 41.432 | 48.175 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 62.110 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 104.930 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 209.290 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 109.270 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 430696.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 74.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9404082.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539867.000 |
| redis | keyspace_hits | 13060137.000 |
| redis | keyspace_misses | 123776.000 |
| redis | net_input_bytes | 978172672.000 |
| redis | net_output_bytes | 2826522715.000 |
| redis | ops_per_sec | 33258.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95535952.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
