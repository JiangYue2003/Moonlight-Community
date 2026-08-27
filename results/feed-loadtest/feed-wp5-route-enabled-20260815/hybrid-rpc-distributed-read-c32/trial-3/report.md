# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:24:05+08:00
- 采样时长：1m0.0237235s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 215470 | 215470 | 0 | 0 | 3590.73 | 8.708 | 11.086 | 11.932 | 13.989 | 30.987 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 580 | 0.003 |
| redis | 646990 | 3.003 |
| relation | 240 | 0.001 |

- Cold compute：215470（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 215470 | 2.195 |
| counter | 240 | 3.023 |
| hydrate | 215470 | 3.023 |
| inbox | 215470 | 2.391 |
| merge_dedup | 215470 | 0.019 |
| relation | 240 | 6.844 |
| route | 215470 | 0.042 |
| total | 215470 | 7.707 |

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
| docker:zg-counter | cpu_percent | 4.720 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.320 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 394.580 |
| docker:zg-knowpost | memory_percent | 0.330 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 0.720 |
| docker:zg-relation | memory_percent | 0.360 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6480046.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 55.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 58154734.000 |
| redis | connected_clients | 180.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.895 |
| redis | keys | 519508.000 |
| redis | keyspace_hits | 301307445.000 |
| redis | keyspace_misses | 35186328.000 |
| redis | net_input_bytes | 13579268254.000 |
| redis | net_output_bytes | 62494832751.000 |
| redis | ops_per_sec | 25500.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 83749784.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
