# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp6-combined-enabled-20260815`
- 开始时间：2026-08-15T22:21:14+08:00
- 采样时长：1m0.0652442s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 635415 | 635415 | 0 | 0 | 10587.21 | 23.136 | 31.473 | 33.068 | 35.944 | 132.075 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.000 |
| mysql | 405 | 0.001 |
| redis | 1271235 | 2.001 |
| relation | 240 | 0.000 |

- Cold compute：635415（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 635415 | 11.165 |
| counter | 240 | 6.457 |
| hydrate | 635415 | 11.143 |
| inbox | 635415 | 11.163 |
| merge_dedup | 635415 | 0.002 |
| relation | 240 | 15.279 |
| route | 635415 | 0.147 |
| total | 635415 | 22.474 |

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
| docker:zg-counter | cpu_percent | 2.150 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 3.330 |
| docker:zg-gateway | memory_percent | 0.200 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 476.540 |
| docker:zg-knowpost | memory_percent | 0.710 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 0.540 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2966.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2966.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6516218.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 89.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 133623004.000 |
| redis | connected_clients | 282.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.852 |
| redis | keys | 515088.000 |
| redis | keyspace_hits | 438097257.000 |
| redis | keyspace_misses | 78096809.000 |
| redis | net_input_bytes | 22434369383.000 |
| redis | net_output_bytes | 88019899231.000 |
| redis | ops_per_sec | 81109.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89750168.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
