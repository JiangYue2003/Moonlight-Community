# Feed 压测报告：hybrid / rpc / hot-read-c32

- Run ID：`feed-wp10-functional-20260816`
- 开始时间：2026-08-16T01:48:23+08:00
- 采样时长：5.0062711s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 162586 | 162586 | 0 | 0 | 32513.61 | 1.038 | 1.549 | 1.618 | 2.173 | 6.861 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 200 | 0.001 |
| counter | 1 | 0.000 |
| mysql | 2 | 0.000 |
| redis | 13 | 0.000 |
| relation | 1 | 0.000 |

- Cold compute：2（0.000 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 162129 |
| l1_stale | 263 |
| l2_fresh | 162 |
| l2_stale | 32 |
| miss | 0 |

- L1+L2 Fresh ratio：99.82%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2 | 0.785 |
| counter | 1 | 1.445 |
| hydrate | 2 | 2.935 |
| inbox | 2 | 0.783 |
| merge_dedup | 2 | 0.005 |
| relation | 1 | 3.782 |
| route | 2 | 2.632 |
| total | 162586 | 0.010 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 13.674 |
| client:loadtest | cpu_percent_total | 218.788 |
| client:loadtest | logical_cpus | 16.000 |
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
| docker:zg-counter | cpu_percent | 1.120 |
| docker:zg-counter | memory_percent | 0.160 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.160 |
| docker:zg-gateway | memory_percent | 0.190 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 404.400 |
| docker:zg-knowpost | memory_percent | 0.300 |
| docker:zg-knowpost | pids | 21.000 |
| docker:zg-relation | cpu_percent | 0.310 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2969.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2969.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6516731.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 134474715.000 |
| redis | connected_clients | 111.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 33.000 |
| redis | hit_rate | 0.849 |
| redis | keys | 515119.000 |
| redis | keyspace_hits | 438112973.000 |
| redis | keyspace_misses | 78104802.000 |
| redis | net_input_bytes | 22493019478.000 |
| redis | net_output_bytes | 88035072113.000 |
| redis | ops_per_sec | 99.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 79643296.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
