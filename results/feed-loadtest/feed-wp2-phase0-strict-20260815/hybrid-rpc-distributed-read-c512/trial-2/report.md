# Feed 压测报告：hybrid / rpc / distributed-read-c512

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T20:06:31+08:00
- 采样时长：1m0.1564666s
- 并发：512
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 167191 | 167191 | 0 | 0 | 2779.88 | 155.411 | 314.994 | 383.916 | 544.696 | 1209.360 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 2432 | 0.015 |
| redis | 504005 | 3.015 |
| relation | 167191 | 1.000 |

- Cold compute：167191（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 167191 | 12.622 |
| counter | 240 | 13.057 |
| hydrate | 167191 | 16.827 |
| inbox | 167191 | 13.479 |
| merge_dedup | 167191 | 0.031 |
| relation | 167191 | 138.057 |
| route | 167191 | 0.050 |
| total | 167191 | 181.121 |

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
| docker:zg-counter | cpu_percent | 5.590 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 2.310 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 354.650 |
| docker:zg-knowpost | memory_percent | 0.640 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 270.110 |
| docker:zg-relation | memory_percent | 0.490 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6275354.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 50686351.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 260457943.000 |
| redis | keyspace_misses | 31088077.000 |
| redis | net_input_bytes | 11744670292.000 |
| redis | net_output_bytes | 53887834820.000 |
| redis | ops_per_sec | 23293.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95696032.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
