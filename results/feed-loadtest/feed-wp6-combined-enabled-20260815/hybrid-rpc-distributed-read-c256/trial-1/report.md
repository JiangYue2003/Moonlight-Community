# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp6-combined-enabled-20260815`
- 开始时间：2026-08-15T22:18:37+08:00
- 采样时长：1m0.0663168s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 629556 | 629556 | 0 | 0 | 10489.76 | 23.379 | 31.933 | 33.749 | 36.574 | 133.309 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.000 |
| mysql | 400 | 0.001 |
| redis | 1259512 | 2.001 |
| relation | 240 | 0.000 |

- Cold compute：629556（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 629556 | 11.285 |
| counter | 240 | 6.175 |
| hydrate | 629556 | 11.259 |
| inbox | 629556 | 11.284 |
| merge_dedup | 629556 | 0.002 |
| relation | 240 | 14.230 |
| route | 629556 | 0.156 |
| total | 629556 | 22.719 |

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
| docker:zg-counter | cpu_percent | 4.660 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 3.260 |
| docker:zg-gateway | memory_percent | 0.200 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 476.330 |
| docker:zg-knowpost | memory_percent | 0.930 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 4.120 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2966.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2966.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6514050.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 89.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 123223770.000 |
| redis | connected_clients | 282.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.861 |
| redis | keys | 515088.000 |
| redis | keyspace_hits | 432158366.000 |
| redis | keyspace_misses | 72166543.000 |
| redis | net_input_bytes | 21666947235.000 |
| redis | net_output_bytes | 87376425874.000 |
| redis | ops_per_sec | 80665.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89913448.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
