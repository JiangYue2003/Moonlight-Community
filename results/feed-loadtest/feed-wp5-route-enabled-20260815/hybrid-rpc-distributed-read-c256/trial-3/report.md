# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:35:51+08:00
- 采样时长：1m0.0616188s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 231833 | 231833 | 0 | 0 | 3861.21 | 64.065 | 85.061 | 89.979 | 104.378 | 206.395 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 3550 | 0.015 |
| redis | 699049 | 3.015 |
| relation | 240 | 0.001 |

- Cold compute：231833（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 231833 | 20.371 |
| counter | 240 | 12.010 |
| hydrate | 231833 | 23.028 |
| inbox | 231833 | 20.699 |
| merge_dedup | 231833 | 0.020 |
| relation | 240 | 24.998 |
| route | 231833 | 0.564 |
| total | 231833 | 64.719 |

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
| docker:zg-counter | cpu_percent | 2.590 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.300 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 428.770 |
| docker:zg-knowpost | memory_percent | 0.570 |
| docker:zg-knowpost | pids | 34.000 |
| docker:zg-relation | cpu_percent | 0.560 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6507666.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 89.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 75602750.000 |
| redis | connected_clients | 309.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.900 |
| redis | keys | 519499.000 |
| redis | keyspace_hits | 405042675.000 |
| redis | keyspace_misses | 45102421.000 |
| redis | net_input_bytes | 18154243643.000 |
| redis | net_output_bytes | 84436238802.000 |
| redis | ops_per_sec | 30317.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91809352.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
