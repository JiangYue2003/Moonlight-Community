# Feed 压测报告：hybrid / rpc / publish-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:44:10+08:00
- 采样时长：20.4539025s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 30 | 30 | 0 | 0 | 1.47 | 682.762 | 710.227 | 736.669 |
| publish_draft | 30 | 30 | 0 | 0 | 1.47 | 5.390 | 8.481 | 20.545 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.47 | 224.257 | 241.237 | 292.339 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.47 | 223.835 | 234.921 | 240.880 |
| publish_commit | 30 | 30 | 0 | 0 | 1.47 | 224.796 | 235.509 | 236.509 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.960 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.670 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 20.340 |
| docker:zg-knowpost | memory_percent | 0.160 |
| docker:zg-knowpost | pids | 20.000 |
| docker:zg-relation | cpu_percent | 2.140 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 364092.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6479093.000 |
| redis | connected_clients | 219.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 538650.000 |
| redis | keyspace_hits | 8375706.000 |
| redis | keyspace_misses | 103825.000 |
| redis | net_input_bytes | 715103754.000 |
| redis | net_output_bytes | 2019441507.000 |
| redis | ops_per_sec | 2095.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 81830712.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
