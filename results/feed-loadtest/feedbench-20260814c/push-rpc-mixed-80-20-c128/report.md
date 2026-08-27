# Feed 压测报告：push / rpc / mixed-80-20-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:29:05+08:00
- 采样时长：5.9329559s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 1552.23 | 48.020 | 97.922 | 105.400 |
| publish_total | 60 | 60 | 0 | 0 | 10.11 | 2448.743 | 2471.319 | 2474.483 |
| publish_draft | 60 | 60 | 0 | 0 | 10.11 | 10.562 | 76.002 | 93.685 |
| publish_metadata | 60 | 60 | 0 | 0 | 10.11 | 761.250 | 936.840 | 945.582 |
| publish_confirm | 60 | 60 | 0 | 0 | 10.11 | 742.731 | 810.367 | 811.398 |
| publish_commit | 60 | 60 | 0 | 0 | 10.11 | 729.680 | 890.765 | 892.874 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.580 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.380 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 132.700 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 11.370 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 275762.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 107.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2596955.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 537190.000 |
| redis | keyspace_hits | 4714961.000 |
| redis | keyspace_misses | 28775.000 |
| redis | net_input_bytes | 329681724.000 |
| redis | net_output_bytes | 894022677.000 |
| redis | ops_per_sec | 23667.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 110386184.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
