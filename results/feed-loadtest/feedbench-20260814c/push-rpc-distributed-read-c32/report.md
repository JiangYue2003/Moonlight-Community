# Feed 压测报告：push / rpc / distributed-read-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:25:54+08:00
- 采样时长：3.0112277s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 6280 | 6280 | 0 | 0 | 2085.53 | 14.949 | 20.361 | 23.655 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.380 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.280 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 211.510 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 197.030 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 230890.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 103.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1681888.000 |
| redis | connected_clients | 418.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 536249.000 |
| redis | keyspace_hits | 1690805.000 |
| redis | keyspace_misses | 15572.000 |
| redis | net_input_bytes | 162175470.000 |
| redis | net_output_bytes | 316754980.000 |
| redis | ops_per_sec | 6694.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89495536.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
