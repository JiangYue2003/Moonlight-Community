# Feed 压测报告：hybrid / rpc / deep-page-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:46:16+08:00
- 采样时长：3.025822s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2954 | 2954 | 0 | 0 | 976.26 | 32.158 | 41.969 | 49.040 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 67.630 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.530 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 263.450 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 115.720 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 406374.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 81.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7996532.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 538764.000 |
| redis | keyspace_hits | 11532576.000 |
| redis | keyspace_misses | 107050.000 |
| redis | net_input_bytes | 848488663.000 |
| redis | net_output_bytes | 2531743771.000 |
| redis | ops_per_sec | 33043.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92585152.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
