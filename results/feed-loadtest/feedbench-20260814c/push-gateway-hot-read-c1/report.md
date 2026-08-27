# Feed 压测报告：push / gateway / hot-read-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:30:14+08:00
- 采样时长：3.002259s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 553 | 553 | 0 | 0 | 184.19 | 5.293 | 6.418 | 7.011 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.230 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 23.350 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 22.000 |
| docker:zg-knowpost | cpu_percent | 27.730 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 24.850 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 281299.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 61.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3158139.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 537726.000 |
| redis | keyspace_hits | 4748194.000 |
| redis | keyspace_misses | 36072.000 |
| redis | net_input_bytes | 371885501.000 |
| redis | net_output_bytes | 907254225.000 |
| redis | ops_per_sec | 791.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 119746704.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
