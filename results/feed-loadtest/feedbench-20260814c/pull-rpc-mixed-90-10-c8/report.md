# Feed 压测报告：pull / rpc / mixed-90-10-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:36:59+08:00
- 采样时长：20.0902766s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 449.35 | 15.324 | 18.876 | 23.909 |
| publish_total | 30 | 30 | 0 | 0 | 1.49 | 657.019 | 703.889 | 1056.186 |
| publish_draft | 30 | 30 | 0 | 0 | 1.49 | 5.352 | 6.275 | 14.296 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.49 | 217.417 | 227.201 | 594.540 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.49 | 212.620 | 220.964 | 222.381 |
| publish_commit | 30 | 30 | 0 | 0 | 1.49 | 218.454 | 236.289 | 248.993 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.640 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.670 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 19.680 |
| docker:zg-knowpost | memory_percent | 0.400 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 1.770 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 338343.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 90.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4952707.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537373.000 |
| redis | keyspace_hits | 7727085.000 |
| redis | keyspace_misses | 72880.000 |
| redis | net_input_bytes | 583141213.000 |
| redis | net_output_bytes | 1626726382.000 |
| redis | ops_per_sec | 1928.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85922888.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
