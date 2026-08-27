# Feed 压测报告：push / gateway / mixed-90-10-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:30:47+08:00
- 采样时长：21.160509s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 171.55 | 5.770 | 7.353 | 8.112 |
| publish_total | 30 | 30 | 0 | 0 | 1.42 | 698.855 | 787.377 | 873.434 |
| publish_draft | 30 | 30 | 0 | 0 | 1.42 | 6.850 | 8.185 | 183.827 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.42 | 226.102 | 249.798 | 260.855 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.42 | 226.529 | 245.671 | 251.457 |
| publish_commit | 30 | 30 | 0 | 0 | 1.42 | 232.996 | 262.667 | 272.666 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.990 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 24.720 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 42.870 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 26.590 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 302009.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 100.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3294469.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 538054.000 |
| redis | keyspace_hits | 5583208.000 |
| redis | keyspace_misses | 36637.000 |
| redis | net_input_bytes | 407664834.000 |
| redis | net_output_bytes | 1062017285.000 |
| redis | ops_per_sec | 6201.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 121258016.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
