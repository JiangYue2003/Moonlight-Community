# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:29:19+08:00
- 采样时长：1m0.037458s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 233866 | 233866 | 0 | 0 | 3896.52 | 31.390 | 45.413 | 49.414 | 56.144 | 100.754 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 1723 | 0.007 |
| redis | 703321 | 3.007 |
| relation | 240 | 0.001 |

- Cold compute：233866（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 233866 | 9.562 |
| counter | 240 | 6.077 |
| hydrate | 233866 | 11.454 |
| inbox | 233866 | 10.041 |
| merge_dedup | 233866 | 0.020 |
| relation | 240 | 16.555 |
| route | 233866 | 0.201 |
| total | 233866 | 31.316 |

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
| docker:zg-counter | cpu_percent | 5.150 |
| docker:zg-counter | memory_percent | 0.180 |
| docker:zg-counter | pids | 27.000 |
| docker:zg-gateway | cpu_percent | 1.700 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 410.340 |
| docker:zg-knowpost | memory_percent | 0.460 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 1.690 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6488672.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 88.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 65953157.000 |
| redis | connected_clients | 280.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.898 |
| redis | keys | 519503.000 |
| redis | keyspace_hits | 347693632.000 |
| redis | keyspace_misses | 39616599.000 |
| redis | net_input_bytes | 15624123358.000 |
| redis | net_output_bytes | 72305962812.000 |
| redis | ops_per_sec | 30588.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89853496.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
