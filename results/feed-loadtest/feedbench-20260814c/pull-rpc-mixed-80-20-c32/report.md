# Feed 压测报告：pull / rpc / mixed-80-20-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:38:52+08:00
- 采样时长：8.1769419s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 420.40 | 58.475 | 75.017 | 82.060 |
| publish_total | 60 | 60 | 0 | 0 | 7.34 | 859.765 | 1410.408 | 1422.177 |
| publish_draft | 60 | 60 | 0 | 0 | 7.34 | 6.468 | 23.000 | 31.015 |
| publish_metadata | 60 | 60 | 0 | 0 | 7.34 | 272.118 | 810.425 | 815.234 |
| publish_confirm | 60 | 60 | 0 | 0 | 7.34 | 267.434 | 327.216 | 328.255 |
| publish_commit | 60 | 60 | 0 | 0 | 7.34 | 278.341 | 359.873 | 367.513 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.940 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.170 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 88.720 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 2.300 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 343840.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 35.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5314323.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 537396.000 |
| redis | keyspace_hits | 7826724.000 |
| redis | keyspace_misses | 80454.000 |
| redis | net_input_bytes | 613760496.000 |
| redis | net_output_bytes | 1671853083.000 |
| redis | ops_per_sec | 9908.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86628936.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
