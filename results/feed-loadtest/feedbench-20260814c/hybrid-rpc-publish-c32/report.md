# Feed 压测报告：hybrid / rpc / publish-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:44:43+08:00
- 采样时长：12.6875508s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 128 | 128 | 0 | 0 | 10.09 | 3161.405 | 3242.933 | 3247.612 |
| publish_draft | 128 | 128 | 0 | 0 | 10.09 | 10.613 | 49.184 | 52.183 |
| publish_metadata | 128 | 128 | 0 | 0 | 10.09 | 1075.131 | 1151.758 | 1156.956 |
| publish_confirm | 128 | 128 | 0 | 0 | 10.09 | 975.669 | 1092.126 | 1100.509 |
| publish_commit | 128 | 128 | 0 | 0 | 10.09 | 1040.973 | 1157.903 | 1163.556 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.940 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.170 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 147.160 |
| docker:zg-knowpost | memory_percent | 0.240 |
| docker:zg-knowpost | pids | 23.000 |
| docker:zg-relation | cpu_percent | 6.460 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 366329.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 36.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6697608.000 |
| redis | connected_clients | 277.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 538697.000 |
| redis | keyspace_hits | 8376360.000 |
| redis | keyspace_misses | 106331.000 |
| redis | net_input_bytes | 732023639.000 |
| redis | net_output_bytes | 2023114617.000 |
| redis | ops_per_sec | 16927.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85117312.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
