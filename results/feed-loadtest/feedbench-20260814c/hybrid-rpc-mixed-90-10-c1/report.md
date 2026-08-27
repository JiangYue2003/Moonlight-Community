# Feed 压测报告：hybrid / rpc / mixed-90-10-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:46:33+08:00
- 采样时长：20.5148008s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 151.93 | 6.399 | 7.905 | 8.715 |
| publish_total | 30 | 30 | 0 | 0 | 1.46 | 680.140 | 727.028 | 756.386 |
| publish_draft | 30 | 30 | 0 | 0 | 1.46 | 5.394 | 7.885 | 8.076 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.46 | 222.762 | 252.107 | 253.082 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.46 | 222.540 | 242.697 | 250.122 |
| publish_commit | 30 | 30 | 0 | 0 | 1.46 | 225.754 | 240.964 | 253.264 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 16.440 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 1.670 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 55.250 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 23.290 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 414755.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 124.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8291732.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 538752.000 |
| redis | keyspace_hits | 12653705.000 |
| redis | keyspace_misses | 108139.000 |
| redis | net_input_bytes | 891906502.000 |
| redis | net_output_bytes | 2727785174.000 |
| redis | ops_per_sec | 4694.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 88955064.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
