# Feed 压测报告：hybrid / rpc / hot-read-c16

- Run ID：`feed-wp9-functional-20260816`
- 开始时间：2026-08-16T00:19:59+08:00
- 采样时长：22.5852ms
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 400 | 400 | 0 | 0 | 17710.71 | 0.804 | 1.089 | 1.582 | 3.114 | 3.114 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 19 | 0.048 |
| counter | 1 | 0.003 |
| mysql | 1 | 0.003 |
| redis | 4 | 0.010 |
| relation | 1 | 0.003 |

- Cold compute：1（0.003 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1 | 0.737 |
| counter | 1 | 1.443 |
| hydrate | 1 | 2.290 |
| inbox | 1 | 0.736 |
| merge_dedup | 1 | 0.004 |
| relation | 1 | 3.206 |
| route | 1 | 4.691 |
| total | 400 | 0.081 |

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
| docker:zg-counter | cpu_percent | 0.910 |
| docker:zg-counter | memory_percent | 0.160 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.150 |
| docker:zg-gateway | memory_percent | 0.190 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 1.950 |
| docker:zg-knowpost | memory_percent | 0.210 |
| docker:zg-knowpost | pids | 22.000 |
| docker:zg-relation | cpu_percent | 0.220 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2969.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2969.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6516651.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 7.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 134117525.000 |
| redis | connected_clients | 121.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.849 |
| redis | keys | 515160.000 |
| redis | keyspace_hits | 438107438.000 |
| redis | keyspace_misses | 78104588.000 |
| redis | net_input_bytes | 22468431006.000 |
| redis | net_output_bytes | 88028685098.000 |
| redis | ops_per_sec | 96.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 79944136.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
