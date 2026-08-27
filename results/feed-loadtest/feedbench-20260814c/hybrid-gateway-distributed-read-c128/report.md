# Feed 压测报告：hybrid / gateway / distributed-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:50:40+08:00
- 采样时长：3.1010617s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2695 | 2695 | 0 | 0 | 869.06 | 143.972 | 180.420 | 200.639 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 56.980 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 93.370 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 203.850 |
| docker:zg-knowpost | memory_percent | 0.450 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 101.990 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 439991.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 106.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9705494.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 540013.000 |
| redis | keyspace_hits | 13712980.000 |
| redis | keyspace_misses | 123779.000 |
| redis | net_input_bytes | 1001735680.000 |
| redis | net_output_bytes | 2972714853.000 |
| redis | ops_per_sec | 28738.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 101675064.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
