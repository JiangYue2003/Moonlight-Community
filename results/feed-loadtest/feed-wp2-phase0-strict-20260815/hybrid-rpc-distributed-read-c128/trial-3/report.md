# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T19:59:58+08:00
- 采样时长：1m0.0528551s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 155523 | 155523 | 0 | 0 | 2590.31 | 48.665 | 60.319 | 64.343 | 73.725 | 134.126 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 1683 | 0.011 |
| redis | 468252 | 3.011 |
| relation | 155523 | 1.000 |

- Cold compute：155523（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 155523 | 6.916 |
| counter | 240 | 8.245 |
| hydrate | 155523 | 9.675 |
| inbox | 155523 | 7.794 |
| merge_dedup | 155523 | 0.028 |
| relation | 155523 | 22.600 |
| route | 155523 | 0.030 |
| total | 155523 | 47.092 |

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
| docker:zg-counter | cpu_percent | 7.940 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 3.350 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 333.870 |
| docker:zg-knowpost | memory_percent | 0.460 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 267.810 |
| docker:zg-relation | memory_percent | 0.350 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 5276391.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 42746205.000 |
| redis | connected_clients | 284.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 219215528.000 |
| redis | keyspace_misses | 26164043.000 |
| redis | net_input_bytes | 9890370388.000 |
| redis | net_output_bytes | 45350483573.000 |
| redis | ops_per_sec | 22050.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91569136.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
