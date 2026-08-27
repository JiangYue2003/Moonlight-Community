# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:28:00+08:00
- 采样时长：1m0.0330202s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 238777 | 238777 | 0 | 0 | 3978.74 | 15.762 | 20.406 | 22.058 | 26.003 | 56.188 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 1164 | 0.005 |
| redis | 717495 | 3.005 |
| relation | 240 | 0.001 |

- Cold compute：238777（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 238777 | 4.196 |
| counter | 240 | 4.508 |
| hydrate | 238777 | 5.714 |
| inbox | 238777 | 4.595 |
| merge_dedup | 238777 | 0.022 |
| relation | 240 | 10.974 |
| route | 238777 | 0.078 |
| total | 238777 | 14.644 |

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
| docker:zg-counter | cpu_percent | 2.600 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 2.650 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 438.610 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 2.830 |
| docker:zg-relation | memory_percent | 0.350 |
| docker:zg-relation | pids | 29.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6485987.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 78.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 64044442.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.897 |
| redis | keys | 519503.000 |
| redis | keyspace_hits | 336347793.000 |
| redis | keyspace_misses | 38532115.000 |
| redis | net_input_bytes | 15123742439.000 |
| redis | net_output_bytes | 69906173982.000 |
| redis | ops_per_sec | 28955.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85797600.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
