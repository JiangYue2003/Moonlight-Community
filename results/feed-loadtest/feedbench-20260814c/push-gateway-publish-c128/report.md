# Feed 压测报告：push / gateway / publish-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:30:06+08:00
- 采样时长：4.7005229s
- 并发：128
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 512 | 2 | 510 | 0 | 0.43 | 2452.432 | 2458.191 | 2458.191 |
| publish_draft | 512 | 512 | 0 | 0 | 108.92 | 31.471 | 90.964 | 108.849 |
| publish_metadata | 512 | 2 | 510 | 0 | 0.43 | 1997.008 | 1999.094 | 1999.094 |
| publish_confirm | 2 | 2 | 0 | 0 | 0.43 | 204.423 | 207.066 | 207.066 |
| publish_commit | 2 | 2 | 0 | 0 | 0.43 | 216.457 | 216.981 | 216.981 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.040 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 3.950 |
| docker:zg-gateway | memory_percent | 0.180 |
| docker:zg-gateway | pids | 22.000 |
| docker:zg-knowpost | cpu_percent | 104.980 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 0.980 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 280666.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 98.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3154363.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 537640.000 |
| redis | keyspace_hits | 4724905.000 |
| redis | keyspace_misses | 36021.000 |
| redis | net_input_bytes | 370890551.000 |
| redis | net_output_bytes | 902920640.000 |
| redis | ops_per_sec | 10874.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 123946368.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
