# Feed 压测报告：hybrid / rpc / hot-read-c16

- Run ID：`feed-wp9-final-functional-20260816`
- 开始时间：2026-08-16T01:09:34+08:00
- 采样时长：37.4007ms
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 400 | 400 | 0 | 0 | 10694.99 | 0.677 | 1.193 | 1.351 | 17.662 | 17.662 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 16 | 0.040 |
| counter | 1 | 0.003 |
| mysql | 1 | 0.003 |
| redis | 5 | 0.013 |
| relation | 1 | 0.003 |

- Cold compute：1（0.003 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1 | 0.548 |
| counter | 1 | 1.244 |
| hydrate | 1 | 5.704 |
| inbox | 1 | 0.545 |
| merge_dedup | 1 | 0.005 |
| relation | 1 | 6.297 |
| route | 1 | 7.611 |
| total | 400 | 0.657 |

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
| docker:zg-counter | cpu_percent | 0.850 |
| docker:zg-counter | memory_percent | 0.160 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.160 |
| docker:zg-gateway | memory_percent | 0.200 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 0.290 |
| docker:zg-knowpost | memory_percent | 0.150 |
| docker:zg-knowpost | pids | 20.000 |
| docker:zg-relation | cpu_percent | 0.190 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2969.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2969.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6516699.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 134318099.000 |
| redis | connected_clients | 116.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.849 |
| redis | keys | 515136.000 |
| redis | keyspace_hits | 438110483.000 |
| redis | keyspace_misses | 78104714.000 |
| redis | net_input_bytes | 22482247100.000 |
| redis | net_output_bytes | 88032256403.000 |
| redis | ops_per_sec | 78.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 79832744.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
