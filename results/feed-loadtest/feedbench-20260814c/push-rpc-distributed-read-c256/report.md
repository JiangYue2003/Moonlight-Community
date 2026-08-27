# Feed 压测报告：push / rpc / distributed-read-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:26:05+08:00
- 采样时长：3.0593086s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 8918 | 8918 | 0 | 0 | 2915.53 | 82.873 | 131.227 | 167.052 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.180 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.440 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 269.590 |
| docker:zg-knowpost | memory_percent | 0.510 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 229.370 |
| docker:zg-relation | memory_percent | 0.360 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 248237.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 96.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1735341.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 536267.000 |
| redis | keyspace_hits | 2417288.000 |
| redis | keyspace_misses | 15587.000 |
| redis | net_input_bytes | 188616083.000 |
| redis | net_output_bytes | 450509790.000 |
| redis | ops_per_sec | 8808.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 94543728.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
