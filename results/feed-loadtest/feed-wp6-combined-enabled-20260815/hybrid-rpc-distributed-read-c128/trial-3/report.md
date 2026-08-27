# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp6-combined-enabled-20260815`
- 开始时间：2026-08-15T22:17:18+08:00
- 采样时长：1m0.059631s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 649686 | 649686 | 0 | 0 | 10825.93 | 11.425 | 16.036 | 17.462 | 20.443 | 60.281 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.000 |
| mysql | 244 | 0.000 |
| redis | 1299616 | 2.000 |
| relation | 240 | 0.000 |

- Cold compute：649686（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 649686 | 5.082 |
| counter | 240 | 4.152 |
| hydrate | 649686 | 5.010 |
| inbox | 649686 | 5.080 |
| merge_dedup | 649686 | 0.002 |
| relation | 240 | 9.548 |
| route | 649686 | 0.057 |
| total | 649686 | 10.171 |

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
| docker:zg-counter | cpu_percent | 2.340 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.460 |
| docker:zg-gateway | memory_percent | 0.200 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 496.040 |
| docker:zg-knowpost | memory_percent | 0.610 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 0.540 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2966.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2966.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6513174.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 89.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 118135656.000 |
| redis | connected_clients | 250.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.865 |
| redis | keys | 515088.000 |
| redis | keyspace_hits | 429252467.000 |
| redis | keyspace_misses | 69265243.000 |
| redis | net_input_bytes | 21291493753.000 |
| redis | net_output_bytes | 87061569919.000 |
| redis | ops_per_sec | 80510.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87810520.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
