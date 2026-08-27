# Feed 压测报告：pull / rpc / mixed-80-20-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:39:04+08:00
- 采样时长：6.3223301s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 426.61 | 223.939 | 300.058 | 308.674 |
| publish_total | 60 | 60 | 0 | 0 | 9.49 | 2512.576 | 2879.121 | 2883.635 |
| publish_draft | 60 | 60 | 0 | 0 | 9.49 | 13.430 | 89.082 | 100.616 |
| publish_metadata | 60 | 60 | 0 | 0 | 9.49 | 929.305 | 1313.795 | 1326.727 |
| publish_confirm | 60 | 60 | 0 | 0 | 9.49 | 738.245 | 801.492 | 804.920 |
| publish_commit | 60 | 60 | 0 | 0 | 9.49 | 786.279 | 792.650 | 796.279 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.620 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.180 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 138.050 |
| docker:zg-knowpost | memory_percent | 0.460 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 2.880 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 344937.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 93.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5389611.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 537415.000 |
| redis | keyspace_hits | 7842294.000 |
| redis | keyspace_misses | 81997.000 |
| redis | net_input_bytes | 620069449.000 |
| redis | net_output_bytes | 1681040584.000 |
| redis | ops_per_sec | 12505.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87189992.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
