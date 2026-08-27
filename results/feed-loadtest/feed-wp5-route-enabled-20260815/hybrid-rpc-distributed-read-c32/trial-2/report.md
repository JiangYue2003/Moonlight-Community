# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:22:46+08:00
- 采样时长：1m0.0217942s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 214380 | 214380 | 0 | 0 | 3572.68 | 8.755 | 11.148 | 12.050 | 14.230 | 32.584 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 452 | 0.002 |
| redis | 643592 | 3.002 |
| relation | 240 | 0.001 |

- Cold compute：214380（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 214380 | 2.201 |
| counter | 240 | 3.020 |
| hydrate | 214380 | 3.041 |
| inbox | 214380 | 2.408 |
| merge_dedup | 214380 | 0.019 |
| relation | 240 | 7.016 |
| route | 214380 | 0.042 |
| total | 214380 | 7.748 |

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
| docker:zg-counter | cpu_percent | 2.700 |
| docker:zg-counter | memory_percent | 0.160 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 3.260 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 393.440 |
| docker:zg-knowpost | memory_percent | 0.350 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 3.400 |
| docker:zg-relation | memory_percent | 0.370 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6478798.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 53.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 56377391.000 |
| redis | connected_clients | 182.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.895 |
| redis | keys | 519516.000 |
| redis | keyspace_hits | 290736974.000 |
| redis | keyspace_misses | 34177569.000 |
| redis | net_input_bytes | 13113474870.000 |
| redis | net_output_bytes | 60259232768.000 |
| redis | ops_per_sec | 25328.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 83694832.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
