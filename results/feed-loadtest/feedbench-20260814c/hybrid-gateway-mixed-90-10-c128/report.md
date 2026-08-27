# Feed 压测报告：hybrid / gateway / mixed-90-10-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:51:24+08:00
- 采样时长：4.1330311s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 670.63 | 148.139 | 216.553 | 223.962 |
| publish_total | 30 | 30 | 0 | 0 | 7.26 | 1417.449 | 1928.023 | 1933.123 |
| publish_draft | 30 | 30 | 0 | 0 | 7.26 | 10.028 | 81.671 | 90.679 |
| publish_metadata | 30 | 30 | 0 | 0 | 7.26 | 503.591 | 861.038 | 865.554 |
| publish_confirm | 30 | 30 | 0 | 0 | 7.26 | 409.641 | 525.594 | 529.838 |
| publish_commit | 30 | 30 | 0 | 0 | 7.26 | 492.646 | 501.283 | 503.802 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.720 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 6.560 |
| docker:zg-gateway | memory_percent | 0.200 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 98.390 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 19.340 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 442356.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 40.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9866834.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 540210.000 |
| redis | keyspace_hits | 13771131.000 |
| redis | keyspace_misses | 126039.000 |
| redis | net_input_bytes | 1014318876.000 |
| redis | net_output_bytes | 2988112300.000 |
| redis | ops_per_sec | 11424.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93742792.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
