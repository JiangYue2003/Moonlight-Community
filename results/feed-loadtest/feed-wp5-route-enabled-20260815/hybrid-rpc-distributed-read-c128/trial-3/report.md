# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:31:55+08:00
- 采样时长：1m0.0400246s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 239419 | 239419 | 0 | 0 | 3989.05 | 30.915 | 43.554 | 47.190 | 53.937 | 98.391 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 2283 | 0.010 |
| redis | 720540 | 3.010 |
| relation | 240 | 0.001 |

- Cold compute：239419（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 239419 | 9.275 |
| counter | 240 | 6.131 |
| hydrate | 239419 | 11.257 |
| inbox | 239419 | 9.736 |
| merge_dedup | 239419 | 0.021 |
| relation | 240 | 17.176 |
| route | 239419 | 0.187 |
| total | 239419 | 30.515 |

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
| docker:zg-counter | cpu_percent | 5.020 |
| docker:zg-counter | memory_percent | 0.180 |
| docker:zg-counter | pids | 26.000 |
| docker:zg-gateway | cpu_percent | 3.490 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 419.510 |
| docker:zg-knowpost | memory_percent | 0.470 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 2.750 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6494670.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 89.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 69842393.000 |
| redis | connected_clients | 276.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.899 |
| redis | keys | 519499.000 |
| redis | keyspace_hits | 370817022.000 |
| redis | keyspace_misses | 41826734.000 |
| redis | net_input_bytes | 16643873765.000 |
| redis | net_output_bytes | 77196867308.000 |
| redis | ops_per_sec | 30815.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89814976.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
