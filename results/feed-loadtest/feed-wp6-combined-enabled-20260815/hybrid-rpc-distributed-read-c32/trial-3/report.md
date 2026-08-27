# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp6-combined-enabled-20260815`
- 开始时间：2026-08-15T22:09:27+08:00
- 采样时长：1m0.0342023s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 479856 | 479856 | 0 | 0 | 7997.20 | 3.844 | 5.085 | 5.487 | 6.644 | 20.518 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 64 | 0.000 |
| redis | 959776 | 2.000 |
| relation | 240 | 0.001 |

- Cold compute：479856（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 479856 | 1.442 |
| counter | 240 | 2.282 |
| hydrate | 479856 | 1.395 |
| inbox | 479856 | 1.440 |
| merge_dedup | 479856 | 0.002 |
| relation | 240 | 4.934 |
| route | 479856 | 0.018 |
| total | 479856 | 2.878 |

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
| docker:zg-counter | cpu_percent | 3.610 |
| docker:zg-counter | memory_percent | 0.160 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.310 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 420.930 |
| docker:zg-knowpost | memory_percent | 0.440 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 0.450 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2966.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2966.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6509316.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 55.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 87872541.000 |
| redis | connected_clients | 155.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.891 |
| redis | keys | 515091.000 |
| redis | keyspace_hits | 411966479.000 |
| redis | keyspace_misses | 52009559.000 |
| redis | net_input_bytes | 19058595597.000 |
| redis | net_output_bytes | 85188612331.000 |
| redis | ops_per_sec | 56177.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 82015872.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
