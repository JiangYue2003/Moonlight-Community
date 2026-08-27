# Feed 压测报告：pull / rpc / publish-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:34:44+08:00
- 采样时长：3.6580194s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 32 | 32 | 0 | 0 | 8.75 | 908.837 | 978.816 | 979.390 |
| publish_draft | 32 | 32 | 0 | 0 | 8.75 | 6.442 | 20.351 | 20.351 |
| publish_metadata | 32 | 32 | 0 | 0 | 8.75 | 291.365 | 322.275 | 323.403 |
| publish_confirm | 32 | 32 | 0 | 0 | 8.75 | 292.287 | 322.087 | 322.087 |
| publish_commit | 32 | 32 | 0 | 0 | 8.75 | 292.458 | 330.500 | 332.759 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.350 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.170 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 88.230 |
| docker:zg-knowpost | memory_percent | 0.160 |
| docker:zg-knowpost | pids | 20.000 |
| docker:zg-relation | cpu_percent | 2.260 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 309981.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4041032.000 |
| redis | connected_clients | 235.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 537418.000 |
| redis | keyspace_hits | 5637354.000 |
| redis | keyspace_misses | 45604.000 |
| redis | net_input_bytes | 465535496.000 |
| redis | net_output_bytes | 1082961495.000 |
| redis | ops_per_sec | 10541.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 82246232.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
