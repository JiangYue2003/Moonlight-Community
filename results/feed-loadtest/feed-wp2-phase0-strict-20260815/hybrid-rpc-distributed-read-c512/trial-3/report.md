# Feed 压测报告：hybrid / rpc / distributed-read-c512

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T20:07:49+08:00
- 采样时长：1m0.1726226s
- 并发：512
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 168527 | 168527 | 0 | 0 | 2801.37 | 154.611 | 309.743 | 378.928 | 536.935 | 1602.948 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 2535 | 0.015 |
| redis | 508116 | 3.015 |
| relation | 168527 | 1.000 |

- Cold compute：168527（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 168527 | 13.028 |
| counter | 240 | 13.749 |
| hydrate | 168527 | 17.312 |
| inbox | 168527 | 13.801 |
| merge_dedup | 168527 | 0.030 |
| relation | 168527 | 135.469 |
| route | 168527 | 0.050 |
| total | 168527 | 179.742 |

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
| docker:zg-counter | cpu_percent | 8.560 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.470 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 349.340 |
| docker:zg-knowpost | memory_percent | 0.610 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 267.230 |
| docker:zg-relation | memory_percent | 0.480 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6476176.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 134.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 52282217.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 268745926.000 |
| redis | keyspace_misses | 32077671.000 |
| redis | net_input_bytes | 12117324274.000 |
| redis | net_output_bytes | 55603452543.000 |
| redis | ops_per_sec | 23445.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 96682120.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
