# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp5-route-enabled-20260815`
- 开始时间：2026-08-15T21:34:32+08:00
- 采样时长：1m0.0727846s
- 并发：256
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 233437 | 233437 | 0 | 0 | 3887.13 | 63.761 | 84.013 | 88.936 | 102.037 | 190.773 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 3235 | 0.014 |
| redis | 703546 | 3.014 |
| relation | 240 | 0.001 |

- Cold compute：233437（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 233437 | 20.271 |
| counter | 240 | 12.035 |
| hydrate | 233437 | 22.818 |
| inbox | 233437 | 20.591 |
| merge_dedup | 233437 | 0.020 |
| relation | 240 | 24.444 |
| route | 233437 | 0.532 |
| total | 233437 | 64.270 |

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
| docker:zg-counter | cpu_percent | 5.460 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 2.970 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 445.470 |
| docker:zg-knowpost | memory_percent | 0.570 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 3.230 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6503244.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 89.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 73691775.000 |
| redis | connected_clients | 308.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.899 |
| redis | keys | 519499.000 |
| redis | keyspace_hits | 393690968.000 |
| redis | keyspace_misses | 44015597.000 |
| redis | net_input_bytes | 17653207897.000 |
| redis | net_output_bytes | 82035124456.000 |
| redis | ops_per_sec | 30487.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91911672.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
