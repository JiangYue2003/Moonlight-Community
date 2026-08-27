# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:49+08:00
- 采样时长：3.0193897s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3914 | 3914 | 0 | 0 | 1296.51 | 24.226 | 30.969 | 35.506 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 83.210 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.220 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 260.320 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 148.870 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 390950.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 102.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7489598.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 538711.000 |
| redis | keyspace_hits | 10008487.000 |
| redis | keyspace_misses | 106790.000 |
| redis | net_input_bytes | 793798057.000 |
| redis | net_output_bytes | 2278727333.000 |
| redis | ops_per_sec | 43382.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92641056.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
