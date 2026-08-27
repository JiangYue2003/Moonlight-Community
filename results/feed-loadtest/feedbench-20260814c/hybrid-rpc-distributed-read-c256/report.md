# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:46:00+08:00
- 采样时长：3.11338s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4980 | 4980 | 0 | 0 | 1599.55 | 157.074 | 193.743 | 216.622 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 89.240 |
| docker:zg-counter | memory_percent | 0.230 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.160 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 280.650 |
| docker:zg-knowpost | memory_percent | 0.540 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 160.160 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 400833.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 102.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7816074.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 538701.000 |
| redis | keyspace_hits | 10717336.000 |
| redis | keyspace_misses | 106790.000 |
| redis | net_input_bytes | 819328209.000 |
| redis | net_output_bytes | 2389439062.000 |
| redis | ops_per_sec | 52178.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 99125496.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
