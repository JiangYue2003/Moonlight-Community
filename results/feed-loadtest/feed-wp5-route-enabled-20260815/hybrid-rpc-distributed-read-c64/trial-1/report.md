# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:25:23+08:00
- 采样时长：1m0.0282954s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 238948 | 238948 | 0 | 0 | 3981.83 | 15.752 | 20.384 | 22.002 | 25.971 | 49.960 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 1064 | 0.004 |
| redis | 717908 | 3.004 |
| relation | 240 | 0.001 |

- Cold compute：238948（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 238948 | 4.219 |
| counter | 240 | 4.334 |
| hydrate | 238948 | 5.685 |
| inbox | 238948 | 4.603 |
| merge_dedup | 238948 | 0.021 |
| relation | 240 | 10.398 |
| route | 238948 | 0.075 |
| total | 238948 | 14.641 |

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
| docker:zg-counter | cpu_percent | 2.740 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 3.200 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 432.380 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 3.960 |
| docker:zg-relation | memory_percent | 0.350 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6482035.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 67.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 60113191.000 |
| redis | connected_clients | 213.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.896 |
| redis | keys | 519502.000 |
| redis | keyspace_hits | 312958755.000 |
| redis | keyspace_misses | 36298795.000 |
| redis | net_input_bytes | 14092818814.000 |
| redis | net_output_bytes | 64959186523.000 |
| redis | ops_per_sec | 29011.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85576760.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
