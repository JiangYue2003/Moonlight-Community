# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:30:37+08:00
- 采样时长：1m0.0375117s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 231771 | 231771 | 0 | 0 | 3861.64 | 31.784 | 45.602 | 49.301 | 55.978 | 100.694 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 1862 | 0.008 |
| redis | 697175 | 3.008 |
| relation | 240 | 0.001 |

- Cold compute：231771（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 231771 | 9.674 |
| counter | 240 | 6.781 |
| hydrate | 231771 | 11.522 |
| inbox | 231771 | 10.147 |
| merge_dedup | 231771 | 0.022 |
| relation | 240 | 15.933 |
| route | 231771 | 0.204 |
| total | 231771 | 31.607 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker-state:zg-counter | health_configured | 1.000 |
| docker-state:zg-counter | healthy | 1.000 |
| docker-state:zg-counter | restart_count | 0.000 |
| docker-state:zg-counter | running | 1.000 |
| docker-state:zg-gateway | health_configured | 1.000 |
| docker-state:zg-gateway | healthy | 1.000 |
| docker-state:zg-gateway | restart_count | 0.000 |
| docker-state:zg-gateway | running | 1.000 |
| docker-state:zg-knowpost | health_configured | 1.000 |
| docker-state:zg-knowpost | healthy | 1.000 |
| docker-state:zg-knowpost | restart_count | 0.000 |
| docker-state:zg-knowpost | running | 1.000 |
| docker-state:zg-relation | health_configured | 1.000 |
| docker-state:zg-relation | healthy | 1.000 |
| docker-state:zg-relation | restart_count | 0.000 |
| docker-state:zg-relation | running | 1.000 |
| docker:zg-counter | cpu_percent | 2.780 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.290 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 432.210 |
| docker:zg-knowpost | memory_percent | 0.480 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 1.080 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6491516.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 89.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 67871243.000 |
| redis | connected_clients | 276.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.898 |
| redis | keys | 519500.000 |
| redis | keyspace_hits | 359097494.000 |
| redis | keyspace_misses | 40706401.000 |
| redis | net_input_bytes | 16126997443.000 |
| redis | net_output_bytes | 74718006007.000 |
| redis | ops_per_sec | 29910.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89681800.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
