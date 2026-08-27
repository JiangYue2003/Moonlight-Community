# Feed 压测报告：hybrid / gateway / distributed-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:50:30+08:00
- 采样时长：3.0035205s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 385 | 385 | 0 | 0 | 128.18 | 7.485 | 9.557 | 10.588 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 11.520 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 17.340 |
| docker:zg-gateway | memory_percent | 0.190 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 37.620 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 18.880 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 434316.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 106.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9519667.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539935.000 |
| redis | keyspace_hits | 13311321.000 |
| redis | keyspace_misses | 123776.000 |
| redis | net_input_bytes | 987211456.000 |
| redis | net_output_bytes | 2882772026.000 |
| redis | ops_per_sec | 4505.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92192488.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
