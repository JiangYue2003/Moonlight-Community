# Feed 压测报告：hybrid / gateway / hot-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:50:14+08:00
- 采样时长：3.0019816s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 403 | 403 | 0 | 0 | 134.24 | 7.343 | 8.840 | 10.011 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 13.020 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 18.000 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 38.460 |
| docker:zg-knowpost | memory_percent | 0.400 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 18.880 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 427620.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 55.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9304214.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 539824.000 |
| redis | keyspace_hits | 12844331.000 |
| redis | keyspace_misses | 123776.000 |
| redis | net_input_bytes | 970367569.000 |
| redis | net_output_bytes | 2778198283.000 |
| redis | ops_per_sec | 4819.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92172288.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
