# Feed 压测报告：pull / gateway / hot-read-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:40:15+08:00
- 采样时长：3.0459773s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1342 | 1342 | 0 | 0 | 440.58 | 71.870 | 94.896 | 105.336 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.210 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 53.660 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 210.210 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 57.750 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 351271.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 57.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5741280.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537556.000 |
| redis | keyspace_hits | 7973288.000 |
| redis | keyspace_misses | 88856.000 |
| redis | net_input_bytes | 650298724.000 |
| redis | net_output_bytes | 1764587906.000 |
| redis | ops_per_sec | 13340.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92689216.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
