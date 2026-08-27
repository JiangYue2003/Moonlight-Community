# Feed 压测报告：hybrid / rpc / mixed-80-20-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:49:13+08:00
- 采样时长：5.8846392s
- 并发：256
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 989.05 | 170.315 | 208.419 | 214.934 |
| publish_total | 60 | 1 | 59 | 0 | 0.17 | 883.741 | 883.741 | 883.741 |
| publish_draft | 60 | 60 | 0 | 0 | 10.20 | 41.270 | 102.397 | 121.463 |
| publish_metadata | 60 | 60 | 0 | 0 | 10.20 | 1781.995 | 1836.718 | 1852.176 |
| publish_confirm | 60 | 60 | 0 | 0 | 10.20 | 1717.352 | 1774.416 | 1801.751 |
| publish_commit | 60 | 1 | 59 | 0 | 0.17 | 239.359 | 239.359 | 239.359 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.910 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 24.000 |
| docker:zg-gateway | cpu_percent | 0.200 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 137.820 |
| docker:zg-knowpost | memory_percent | 0.480 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 18.710 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 31.000 |
| kafka | lag_max | 43.000 |
| kafka | lag_total | 43.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 423458.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 118.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8963450.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539208.000 |
| redis | keyspace_hits | 12813891.000 |
| redis | keyspace_misses | 120396.000 |
| redis | net_input_bytes | 944601346.000 |
| redis | net_output_bytes | 2766497307.000 |
| redis | ops_per_sec | 12617.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 94072808.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
