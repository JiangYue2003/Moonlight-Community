# Feed 压测报告：push / rpc / mixed-90-10-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:27:23+08:00
- 采样时长：5.5483422s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 1562.32 | 15.815 | 30.986 | 35.172 |
| publish_total | 30 | 30 | 0 | 0 | 5.41 | 673.708 | 842.370 | 842.922 |
| publish_draft | 30 | 30 | 0 | 0 | 5.41 | 6.377 | 15.960 | 16.977 |
| publish_metadata | 30 | 30 | 0 | 0 | 5.41 | 221.793 | 356.573 | 358.520 |
| publish_confirm | 30 | 30 | 0 | 0 | 5.41 | 220.492 | 234.499 | 236.065 |
| publish_commit | 30 | 30 | 0 | 0 | 5.41 | 225.278 | 239.394 | 241.172 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.130 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.150 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 57.620 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 17.400 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 269399.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 111.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1987544.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 536614.000 |
| redis | keyspace_hits | 4653816.000 |
| redis | keyspace_misses | 19332.000 |
| redis | net_input_bytes | 282515894.000 |
| redis | net_output_bytes | 874830993.000 |
| redis | ops_per_sec | 11284.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91242272.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
