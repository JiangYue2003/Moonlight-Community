# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:21:28+08:00
- 采样时长：1m0.0232684s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 213333 | 213333 | 0 | 0 | 3555.13 | 8.848 | 11.171 | 12.082 | 14.171 | 30.754 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 517 | 0.002 |
| redis | 640516 | 3.002 |
| relation | 240 | 0.001 |

- Cold compute：213333（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 213333 | 2.227 |
| counter | 240 | 2.997 |
| hydrate | 213333 | 3.043 |
| inbox | 213333 | 2.423 |
| merge_dedup | 213333 | 0.019 |
| relation | 240 | 7.248 |
| route | 213333 | 0.043 |
| total | 213333 | 7.792 |

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
| docker:zg-counter | cpu_percent | 4.390 |
| docker:zg-counter | memory_percent | 0.160 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 1.670 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 396.430 |
| docker:zg-knowpost | memory_percent | 0.340 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 0.660 |
| docker:zg-relation | memory_percent | 0.370 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6477522.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 44.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 54610925.000 |
| redis | connected_clients | 201.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.894 |
| redis | keys | 519543.000 |
| redis | keyspace_hits | 280231856.000 |
| redis | keyspace_misses | 33175033.000 |
| redis | net_input_bytes | 12650556670.000 |
| redis | net_output_bytes | 58037458691.000 |
| redis | ops_per_sec | 25650.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 84021664.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
