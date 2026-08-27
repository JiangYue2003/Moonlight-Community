# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T19:58:40+08:00
- 采样时长：1m0.0447278s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 156623 | 156623 | 0 | 0 | 2609.04 | 48.344 | 59.816 | 63.659 | 72.582 | 117.097 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 1543 | 0.010 |
| redis | 471412 | 3.010 |
| relation | 156623 | 1.000 |

- Cold compute：156623（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 156623 | 6.898 |
| counter | 240 | 8.478 |
| hydrate | 156623 | 9.533 |
| inbox | 156623 | 7.730 |
| merge_dedup | 156623 | 0.029 |
| relation | 156623 | 22.431 |
| route | 156623 | 0.030 |
| total | 156623 | 46.701 |

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
| docker:zg-counter | cpu_percent | 4.970 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 2.900 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 342.140 |
| docker:zg-knowpost | memory_percent | 0.450 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 268.210 |
| docker:zg-relation | memory_percent | 0.360 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 5091787.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 132.000 |
| mysql | threads_running | 7.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 41272850.000 |
| redis | connected_clients | 295.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 211560547.000 |
| redis | keyspace_misses | 25251131.000 |
| redis | net_input_bytes | 9546406528.000 |
| redis | net_output_bytes | 43765954527.000 |
| redis | ops_per_sec | 21365.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91472088.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
