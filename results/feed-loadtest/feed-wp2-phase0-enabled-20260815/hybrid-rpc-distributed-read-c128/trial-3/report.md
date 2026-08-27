# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:11:16+08:00
- 采样时长：1m0.0466789s
- 并发：128
- 重复轮次：3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 154520 | 154520 | 0 | 0 | 2573.84 | 48.873 | 64.815 | 75.885 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 241 | 0.002 |
| mysql | 1513 | 0.010 |
| redis | 465073 | 3.010 |
| relation | 154520 | 1.000 |

- Cold compute：154520（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 154520 | 6.804 |
| counter | 241 | 8.431 |
| hydrate | 154520 | 9.705 |
| inbox | 154520 | 7.726 |
| merge_dedup | 154520 | 0.027 |
| relation | 154520 | 22.936 |
| route | 154520 | 0.029 |
| total | 154520 | 47.276 |

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
| docker:zg-counter | cpu_percent | 5.560 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 2.880 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 336.640 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 259.690 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 2005974.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 16476547.000 |
| redis | connected_clients | 275.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 83559889.000 |
| redis | keyspace_misses | 9970317.000 |
| redis | net_input_bytes | 3781646400.000 |
| redis | net_output_bytes | 17268904061.000 |
| redis | ops_per_sec | 21345.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90919552.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
