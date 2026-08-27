# Feed 压测报告：hybrid / rpc / distributed-read-c7

- Run ID：`feed-wp6-combined-validation-20260815`
- 开始时间：2026-08-15T22:05:45+08:00
- 采样时长：10.0039766s
- 并发：7
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 34959 | 34959 | 0 | 0 | 3495.15 | 2.089 | 2.418 | 2.644 | 3.148 | 10.487 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 69918 | 2.000 |
| relation | 40 | 0.001 |

- Cold compute：34959（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 34959 | 0.676 |
| counter | 40 | 1.295 |
| hydrate | 34959 | 0.666 |
| inbox | 34959 | 0.675 |
| merge_dedup | 34959 | 0.001 |
| relation | 40 | 3.257 |
| route | 34959 | 0.014 |
| total | 34959 | 1.375 |

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
| docker:zg-counter | cpu_percent | 1.170 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 2.120 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 191.750 |
| docker:zg-knowpost | memory_percent | 0.210 |
| docker:zg-knowpost | pids | 22.000 |
| docker:zg-relation | cpu_percent | 2.330 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2966.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2966.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6507852.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 76048270.000 |
| redis | connected_clients | 156.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.900 |
| redis | keys | 515093.000 |
| redis | keyspace_hits | 405214277.000 |
| redis | keyspace_misses | 45273105.000 |
| redis | net_input_bytes | 18186469437.000 |
| redis | net_output_bytes | 84457235367.000 |
| redis | ops_per_sec | 25157.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 80931704.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
