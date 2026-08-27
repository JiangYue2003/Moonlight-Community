# Feed 压测报告：push / rpc / mixed-80-20-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:28:30+08:00
- 采样时长：19.8371408s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 899.92 | 6.396 | 7.948 | 12.434 |
| publish_total | 60 | 60 | 0 | 0 | 3.02 | 651.175 | 717.845 | 776.175 |
| publish_draft | 60 | 60 | 0 | 0 | 3.02 | 5.374 | 7.532 | 9.287 |
| publish_metadata | 60 | 60 | 0 | 0 | 3.02 | 213.632 | 264.616 | 307.882 |
| publish_confirm | 60 | 60 | 0 | 0 | 3.02 | 212.040 | 229.804 | 233.058 |
| publish_commit | 60 | 60 | 0 | 0 | 3.02 | 218.680 | 234.042 | 239.477 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.760 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.790 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 32.110 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 2.300 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 4.000 |
| kafka | lag_total | 4.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 273468.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 111.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2374089.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 537003.000 |
| redis | keyspace_hits | 4695224.000 |
| redis | keyspace_misses | 25521.000 |
| redis | net_input_bytes | 312416421.000 |
| redis | net_output_bytes | 887415826.000 |
| redis | ops_per_sec | 10445.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 103197184.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
