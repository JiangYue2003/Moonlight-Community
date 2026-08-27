# Feed 压测报告：push / rpc / publish-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:24:49+08:00
- 采样时长：11.2455537s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 128 | 128 | 0 | 0 | 11.38 | 2711.766 | 2985.753 | 2990.257 |
| publish_draft | 128 | 128 | 0 | 0 | 11.38 | 10.077 | 50.084 | 51.161 |
| publish_metadata | 128 | 128 | 0 | 0 | 11.38 | 886.672 | 966.555 | 974.515 |
| publish_confirm | 128 | 128 | 0 | 0 | 11.38 | 911.463 | 1005.607 | 1009.545 |
| publish_commit | 128 | 128 | 0 | 0 | 11.38 | 895.601 | 1095.058 | 1105.087 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 5.660 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.190 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 150.630 |
| docker:zg-knowpost | memory_percent | 0.230 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 12.250 |
| docker:zg-relation | memory_percent | 0.170 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 32.000 |
| kafka | lag_total | 32.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 190299.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 36.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1496563.000 |
| redis | connected_clients | 150.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.850 |
| redis | keys | 536065.000 |
| redis | keyspace_hits | 74300.000 |
| redis | keyspace_misses | 14940.000 |
| redis | net_input_bytes | 98588964.000 |
| redis | net_output_bytes | 18225345.000 |
| redis | ops_per_sec | 28000.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 83643528.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
