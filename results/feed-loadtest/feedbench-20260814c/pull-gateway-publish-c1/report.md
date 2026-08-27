# Feed 压测报告：pull / gateway / publish-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:39:23+08:00
- 采样时长：19.8027975s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 30 | 30 | 0 | 0 | 1.51 | 662.302 | 676.615 | 695.874 |
| publish_draft | 30 | 30 | 0 | 0 | 1.51 | 6.385 | 7.330 | 7.416 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.51 | 215.048 | 223.576 | 231.059 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.51 | 214.075 | 221.978 | 222.314 |
| publish_commit | 30 | 30 | 0 | 0 | 1.51 | 221.862 | 233.917 | 243.686 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.580 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.330 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 18.960 |
| docker:zg-knowpost | memory_percent | 0.440 |
| docker:zg-knowpost | pids | 32.000 |
| docker:zg-relation | cpu_percent | 1.800 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 346431.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 119.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5497159.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537468.000 |
| redis | keyspace_hits | 7857438.000 |
| redis | keyspace_misses | 84425.000 |
| redis | net_input_bytes | 628992277.000 |
| redis | net_output_bytes | 1691313817.000 |
| redis | ops_per_sec | 1923.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85762224.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
