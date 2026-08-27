# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:01:21+08:00
- 采样时长：1m0.0184231s
- 并发：32
- 重复轮次：2
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 117822 | 117822 | 0 | 0 | 1963.39 | 16.032 | 20.738 | 23.413 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 295 | 0.003 |
| redis | 353761 | 3.003 |
| relation | 117822 | 1.000 |

- Cold compute：117822（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 117822 | 2.088 |
| counter | 240 | 3.171 |
| hydrate | 117822 | 2.938 |
| inbox | 117822 | 2.328 |
| merge_dedup | 117822 | 0.022 |
| relation | 117822 | 7.511 |
| route | 117822 | 0.013 |
| total | 117822 | 14.943 |

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
| docker:zg-counter | cpu_percent | 4.810 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 17.000 |
| docker:zg-gateway | cpu_percent | 1.510 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 298.440 |
| docker:zg-knowpost | memory_percent | 0.340 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 240.600 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 841526.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 64.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7134088.000 |
| redis | connected_clients | 270.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 35072253.000 |
| redis | keyspace_misses | 4190869.000 |
| redis | net_input_bytes | 1603004423.000 |
| redis | net_output_bytes | 7233130840.000 |
| redis | ops_per_sec | 16155.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86274856.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
