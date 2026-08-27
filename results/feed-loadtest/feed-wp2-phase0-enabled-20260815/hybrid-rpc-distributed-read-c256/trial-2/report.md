# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:14:17+08:00
- 采样时长：1m0.0764599s
- 并发：256
- 重复轮次：2
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 169373 | 169373 | 0 | 0 | 2820.01 | 88.152 | 125.638 | 150.554 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 2400 | 0.014 |
| redis | 510519 | 3.014 |
| relation | 169373 | 1.000 |

- Cold compute：169373（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 169373 | 11.327 |
| counter | 240 | 12.801 |
| hydrate | 169373 | 15.385 |
| inbox | 169373 | 12.192 |
| merge_dedup | 169373 | 0.030 |
| relation | 169373 | 48.793 |
| route | 169373 | 0.044 |
| total | 169373 | 87.826 |

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
| docker:zg-counter | cpu_percent | 5.410 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 19.000 |
| docker:zg-gateway | cpu_percent | 2.910 |
| docker:zg-gateway | memory_percent | 0.130 |
| docker:zg-gateway | pids | 25.000 |
| docker:zg-knowpost | cpu_percent | 354.690 |
| docker:zg-knowpost | memory_percent | 0.530 |
| docker:zg-knowpost | pids | 34.000 |
| docker:zg-relation | cpu_percent | 277.660 |
| docker:zg-relation | memory_percent | 0.390 |
| docker:zg-relation | pids | 34.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 2407791.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19679646.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 100182000.000 |
| redis | keyspace_misses | 11955920.000 |
| redis | net_input_bytes | 4529380601.000 |
| redis | net_output_bytes | 20709884997.000 |
| redis | ops_per_sec | 23116.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 94936648.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
