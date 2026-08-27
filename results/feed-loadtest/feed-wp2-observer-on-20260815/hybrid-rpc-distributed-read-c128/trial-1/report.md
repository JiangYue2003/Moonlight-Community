# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-observer-on-20260815`
- 开始时间：2026-08-15T18:57:17+08:00
- 采样时长：20.0336236s
- 并发：128
- 重复轮次：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 52094 | 52094 | 0 | 0 | 2600.83 | 48.455 | 64.123 | 73.116 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 80 | 0.002 |
| mysql | 589 | 0.011 |
| redis | 156871 | 3.011 |
| relation | 52094 | 1.000 |

- Cold compute：52094（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 52094 | 6.708 |
| counter | 80 | 8.717 |
| hydrate | 52094 | 9.552 |
| inbox | 52094 | 7.607 |
| merge_dedup | 52094 | 0.031 |
| relation | 52094 | 22.759 |
| route | 52094 | 0.029 |
| total | 52094 | 46.735 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker-state:zg-counter | healthy | 1.000 |
| docker-state:zg-counter | restart_count | 0.000 |
| docker-state:zg-counter | running | 1.000 |
| docker-state:zg-gateway | healthy | 1.000 |
| docker-state:zg-gateway | restart_count | 0.000 |
| docker-state:zg-gateway | running | 1.000 |
| docker-state:zg-knowpost | healthy | 1.000 |
| docker-state:zg-knowpost | restart_count | 0.000 |
| docker-state:zg-knowpost | running | 1.000 |
| docker-state:zg-relation | healthy | 1.000 |
| docker-state:zg-relation | restart_count | 0.000 |
| docker-state:zg-relation | running | 1.000 |
| docker:zg-counter | cpu_percent | 7.610 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 17.000 |
| docker:zg-gateway | cpu_percent | 1.500 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 326.550 |
| docker:zg-knowpost | memory_percent | 0.380 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 267.430 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 431773.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 126.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3833205.000 |
| redis | connected_clients | 269.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 17981657.000 |
| redis | keyspace_misses | 2154156.000 |
| redis | net_input_bytes | 834709634.000 |
| redis | net_output_bytes | 3696009303.000 |
| redis | ops_per_sec | 21189.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90592920.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
