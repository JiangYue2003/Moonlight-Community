# Feed 压测报告：hybrid / rpc / publish-c40

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T16:01:25+08:00
- 采样时长：16.9031517s
- 并发：40
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 160 | 160 | 0 | 0 | 9.47 | 4231.339 | 4299.724 | 4318.382 |
| publish_draft | 160 | 160 | 0 | 0 | 9.47 | 10.244 | 49.808 | 54.158 |
| publish_metadata | 160 | 160 | 0 | 0 | 9.47 | 1400.296 | 1462.412 | 1472.784 |
| publish_confirm | 160 | 160 | 0 | 0 | 9.47 | 1364.322 | 1462.101 | 1475.127 |
| publish_commit | 160 | 160 | 0 | 0 | 9.47 | 1450.216 | 1500.882 | 1510.861 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 10.850 |
| docker:zg-counter | memory_percent | 0.200 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.000 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 137.610 |
| docker:zg-knowpost | memory_percent | 0.230 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 7.620 |
| docker:zg-relation | memory_percent | 0.260 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 35.000 |
| kafka | lag_total | 35.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 454012.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 44.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 11042596.000 |
| redis | connected_clients | 356.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 541515.000 |
| redis | keyspace_hits | 13826108.000 |
| redis | keyspace_misses | 138330.000 |
| redis | net_input_bytes | 1104108639.000 |
| redis | net_output_bytes | 3018881412.000 |
| redis | ops_per_sec | 18482.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 100178656.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
