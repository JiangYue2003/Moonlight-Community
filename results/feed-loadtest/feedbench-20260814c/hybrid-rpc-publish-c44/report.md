# Feed 压测报告：hybrid / rpc / publish-c44

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T16:01:57+08:00
- 采样时长：3.9959715s
- 并发：44
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 176 | 34 | 142 | 0 | 8.51 | 3810.885 | 3861.256 | 3868.764 |
| publish_draft | 176 | 176 | 0 | 0 | 44.04 | 13.447 | 62.897 | 64.529 |
| publish_metadata | 176 | 34 | 142 | 0 | 8.51 | 1446.057 | 1514.213 | 1517.445 |
| publish_confirm | 34 | 34 | 0 | 0 | 8.51 | 1061.886 | 1091.453 | 1094.685 |
| publish_commit | 34 | 34 | 0 | 0 | 8.51 | 1247.714 | 1269.880 | 1276.557 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.690 |
| docker:zg-counter | memory_percent | 0.200 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 1.570 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 134.150 |
| docker:zg-knowpost | memory_percent | 0.200 |
| docker:zg-knowpost | pids | 23.000 |
| docker:zg-relation | cpu_percent | 1.290 |
| docker:zg-relation | memory_percent | 0.260 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 454680.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 47.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 11102132.000 |
| redis | connected_clients | 363.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 541551.000 |
| redis | keyspace_hits | 13826295.000 |
| redis | keyspace_misses | 138909.000 |
| redis | net_input_bytes | 1108645163.000 |
| redis | net_output_bytes | 3019879202.000 |
| redis | ops_per_sec | 12820.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 100428528.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
