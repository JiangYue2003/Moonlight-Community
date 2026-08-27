# Feed 压测报告：push / rpc / mixed-90-10-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:27:40+08:00
- 采样时长：3.0830491s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 1722.27 | 112.909 | 143.922 | 148.836 |
| publish_total | 30 | 30 | 0 | 0 | 9.73 | 2430.748 | 2442.066 | 2443.650 |
| publish_draft | 30 | 30 | 0 | 0 | 9.73 | 34.492 | 63.657 | 68.352 |
| publish_metadata | 30 | 30 | 0 | 0 | 9.73 | 926.733 | 934.951 | 937.016 |
| publish_confirm | 30 | 30 | 0 | 0 | 9.73 | 777.163 | 786.961 | 790.861 |
| publish_commit | 30 | 30 | 0 | 0 | 9.73 | 684.625 | 694.136 | 696.839 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.130 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.150 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 143.060 |
| docker:zg-knowpost | memory_percent | 0.480 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 5.770 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 271065.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 112.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2101287.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 536727.000 |
| redis | keyspace_hits | 4674029.000 |
| redis | keyspace_misses | 23200.000 |
| redis | net_input_bytes | 292032884.000 |
| redis | net_output_bytes | 880123069.000 |
| redis | ops_per_sec | 11046.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 96753888.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
