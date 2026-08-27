# Feed 压测报告：hybrid / rpc / distributed-read-c1

- Run ID：`feed-wp7-l1l2-validation-20260815`
- 开始时间：2026-08-15T23:06:49+08:00
- 采样时长：133.1761ms
- 并发：1
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 40 | 40 | 0 | 0 | 300.35 | 0.651 | 6.900 | 7.045 | 8.869 | 8.869 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 18 | 0.450 |
| counter | 18 | 0.450 |
| mysql | 0 | 0.000 |
| redis | 72 | 1.800 |
| relation | 18 | 0.450 |

- Cold compute：18（0.450 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 18 | 0.532 |
| counter | 18 | 0.998 |
| hydrate | 18 | 0.518 |
| inbox | 18 | 0.531 |
| merge_dedup | 18 | 0.001 |
| relation | 18 | 2.449 |
| route | 18 | 3.499 |
| total | 40 | 2.797 |

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
| docker:zg-counter | cpu_percent | 0.900 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 1.860 |
| docker:zg-gateway | memory_percent | 0.200 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 0.280 |
| docker:zg-knowpost | memory_percent | 0.140 |
| docker:zg-knowpost | pids | 21.000 |
| docker:zg-relation | cpu_percent | 1.760 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2969.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2969.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6516476.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 7.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 133819356.000 |
| redis | connected_clients | 114.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.849 |
| redis | keys | 515124.000 |
| redis | keyspace_hits | 438102227.000 |
| redis | keyspace_misses | 78104088.000 |
| redis | net_input_bytes | 22447842293.000 |
| redis | net_output_bytes | 88023321584.000 |
| redis | ops_per_sec | 530.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 79797160.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
