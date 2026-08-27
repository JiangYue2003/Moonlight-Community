# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:26:42+08:00
- 采样时长：1m0.0292876s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 238163 | 238163 | 0 | 0 | 3968.69 | 15.798 | 20.462 | 22.054 | 25.897 | 56.250 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 992 | 0.004 |
| redis | 715481 | 3.004 |
| relation | 240 | 0.001 |

- Cold compute：238163（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 238163 | 4.229 |
| counter | 240 | 4.695 |
| hydrate | 238163 | 5.685 |
| inbox | 238163 | 4.613 |
| merge_dedup | 238163 | 0.021 |
| relation | 240 | 10.556 |
| route | 238163 | 0.080 |
| total | 238163 | 14.667 |

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
| docker:zg-counter | cpu_percent | 4.620 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.410 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 426.430 |
| docker:zg-knowpost | memory_percent | 0.400 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 0.540 |
| docker:zg-relation | memory_percent | 0.350 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6483823.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 63.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 62076123.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.897 |
| redis | keys | 519502.000 |
| redis | keyspace_hits | 324638219.000 |
| redis | keyspace_misses | 37413881.000 |
| redis | net_input_bytes | 14607581420.000 |
| redis | net_output_bytes | 67429486105.000 |
| redis | ops_per_sec | 29335.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85881760.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
