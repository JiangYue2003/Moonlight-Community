# Feed 压测报告：push / rpc / mixed-90-10-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:27:34+08:00
- 采样时长：3.3150392s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 1617.20 | 50.795 | 112.756 | 122.296 |
| publish_total | 30 | 30 | 0 | 0 | 9.05 | 1172.071 | 1487.335 | 1487.860 |
| publish_draft | 30 | 30 | 0 | 0 | 9.05 | 8.426 | 82.696 | 83.294 |
| publish_metadata | 30 | 30 | 0 | 0 | 9.05 | 379.723 | 555.269 | 558.871 |
| publish_confirm | 30 | 30 | 0 | 0 | 9.05 | 393.608 | 436.500 | 439.159 |
| publish_commit | 30 | 30 | 0 | 0 | 9.05 | 395.584 | 447.672 | 449.466 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.130 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.090 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 112.680 |
| docker:zg-knowpost | memory_percent | 0.440 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 1.010 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 13.000 |
| kafka | lag_total | 13.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 270305.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 88.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2045069.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 536673.000 |
| redis | keyspace_hits | 4664031.000 |
| redis | keyspace_misses | 21164.000 |
| redis | net_input_bytes | 287303047.000 |
| redis | net_output_bytes | 877500084.000 |
| redis | ops_per_sec | 16970.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92402776.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
