# Feed 压测报告：pull / rpc / publish-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:34:51+08:00
- 采样时长：11.8967714s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 128 | 128 | 0 | 0 | 10.76 | 2956.678 | 3029.476 | 3041.171 |
| publish_draft | 128 | 128 | 0 | 0 | 10.76 | 9.312 | 58.270 | 62.253 |
| publish_metadata | 128 | 128 | 0 | 0 | 10.76 | 1012.859 | 1118.732 | 1124.580 |
| publish_confirm | 128 | 128 | 0 | 0 | 10.76 | 945.015 | 1005.970 | 1011.858 |
| publish_commit | 128 | 128 | 0 | 0 | 10.76 | 982.785 | 1017.649 | 1024.062 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.520 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.910 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 145.520 |
| docker:zg-knowpost | memory_percent | 0.220 |
| docker:zg-knowpost | pids | 23.000 |
| docker:zg-relation | cpu_percent | 2.260 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 311683.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 35.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4185360.000 |
| redis | connected_clients | 262.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 537411.000 |
| redis | keyspace_hits | 5637625.000 |
| redis | keyspace_misses | 47524.000 |
| redis | net_input_bytes | 477034674.000 |
| redis | net_output_bytes | 1085605133.000 |
| redis | ops_per_sec | 13214.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 83873728.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
