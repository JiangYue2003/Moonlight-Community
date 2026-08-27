# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T19:56:02+08:00
- 采样时长：1m0.0244248s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 140040 | 140040 | 0 | 0 | 2333.48 | 27.016 | 33.191 | 35.302 | 40.164 | 75.088 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 788 | 0.006 |
| redis | 420908 | 3.006 |
| relation | 140040 | 1.000 |

- Cold compute：140040（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 140040 | 3.671 |
| counter | 240 | 4.971 |
| hydrate | 140040 | 5.258 |
| inbox | 140040 | 4.147 |
| merge_dedup | 140040 | 0.025 |
| relation | 140040 | 12.511 |
| route | 140040 | 0.018 |
| total | 140040 | 25.675 |

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
| docker:zg-counter | cpu_percent | 4.710 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.320 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 326.430 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 253.100 |
| docker:zg-relation | memory_percent | 0.360 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 4722517.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 102.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 38323990.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 196241799.000 |
| redis | keyspace_misses | 23423981.000 |
| redis | net_input_bytes | 8858011506.000 |
| redis | net_output_bytes | 40595077376.000 |
| redis | ops_per_sec | 19515.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89319744.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
