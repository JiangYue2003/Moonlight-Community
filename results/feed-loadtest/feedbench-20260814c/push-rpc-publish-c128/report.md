# Feed 压测报告：push / rpc / publish-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:04+08:00
- 采样时长：2.7604299s
- 并发：128
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 512 | 3 | 509 | 0 | 1.09 | 698.054 | 705.127 | 705.127 |
| publish_draft | 512 | 512 | 0 | 0 | 185.51 | 29.908 | 62.698 | 69.396 |
| publish_metadata | 512 | 3 | 509 | 0 | 1.09 | 259.854 | 265.117 | 265.117 |
| publish_confirm | 3 | 3 | 0 | 0 | 1.09 | 198.930 | 200.195 | 200.195 |
| publish_commit | 3 | 3 | 0 | 0 | 1.09 | 208.885 | 209.401 | 209.401 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 1.860 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.130 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 103.150 |
| docker:zg-knowpost | memory_percent | 0.310 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 2.010 |
| docker:zg-relation | memory_percent | 0.160 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 191179.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1548159.000 |
| redis | connected_clients | 290.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.833 |
| redis | keys | 536116.000 |
| redis | keyspace_hits | 74508.000 |
| redis | keyspace_misses | 15406.000 |
| redis | net_input_bytes | 102303019.000 |
| redis | net_output_bytes | 18918836.000 |
| redis | ops_per_sec | 12983.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87908264.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
