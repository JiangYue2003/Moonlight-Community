# Feed 压测报告：hybrid / rpc / distributed-read-c512

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T20:05:12+08:00
- 采样时长：1m0.1578207s
- 并发：512
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 168718 | 168718 | 0 | 0 | 2805.25 | 153.780 | 312.925 | 382.092 | 544.774 | 1230.326 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 2239 | 0.013 |
| redis | 508393 | 3.013 |
| relation | 168718 | 1.000 |

- Cold compute：168718（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 168718 | 12.385 |
| counter | 240 | 13.833 |
| hydrate | 168718 | 16.479 |
| inbox | 168718 | 13.177 |
| merge_dedup | 168718 | 0.031 |
| relation | 168718 | 137.361 |
| route | 168718 | 0.052 |
| total | 168718 | 179.537 |

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
| docker:zg-counter | cpu_percent | 8.350 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 3.230 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 352.130 |
| docker:zg-knowpost | memory_percent | 0.660 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 271.880 |
| docker:zg-relation | memory_percent | 0.490 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6075812.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 7.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 49102021.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 252231264.000 |
| redis | keyspace_misses | 30105565.000 |
| redis | net_input_bytes | 11374718325.000 |
| redis | net_output_bytes | 52184903313.000 |
| redis | ops_per_sec | 23316.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 96613832.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
