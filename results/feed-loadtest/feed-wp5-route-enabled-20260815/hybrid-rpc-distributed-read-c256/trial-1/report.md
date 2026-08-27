# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:33:14+08:00
- 采样时长：1m0.0567704s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 233194 | 233194 | 0 | 0 | 3884.14 | 63.985 | 82.928 | 87.092 | 102.477 | 186.461 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 3335 | 0.014 |
| redis | 702917 | 3.014 |
| relation | 240 | 0.001 |

- Cold compute：233194（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 233194 | 20.306 |
| counter | 240 | 12.408 |
| hydrate | 233194 | 22.844 |
| inbox | 233194 | 20.615 |
| merge_dedup | 233194 | 0.019 |
| relation | 240 | 23.128 |
| route | 233194 | 0.510 |
| total | 233194 | 64.334 |

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
| docker:zg-counter | cpu_percent | 2.560 |
| docker:zg-counter | memory_percent | 0.180 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.740 |
| docker:zg-gateway | memory_percent | 0.220 |
| docker:zg-gateway | pids | 26.000 |
| docker:zg-knowpost | cpu_percent | 443.710 |
| docker:zg-knowpost | memory_percent | 0.560 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 0.520 |
| docker:zg-relation | memory_percent | 0.350 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6499145.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 87.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 71767837.000 |
| redis | connected_clients | 308.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.899 |
| redis | keys | 519499.000 |
| redis | keyspace_hits | 382257349.000 |
| redis | keyspace_misses | 42921630.000 |
| redis | net_input_bytes | 17148727412.000 |
| redis | net_output_bytes | 79616717140.000 |
| redis | ops_per_sec | 30464.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91706144.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
