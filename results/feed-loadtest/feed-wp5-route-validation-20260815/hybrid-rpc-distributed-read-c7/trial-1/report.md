# Feed 压测报告：hybrid / rpc / distributed-read-c7

- Run ID：`feed-wp5-route-validation-20260815`
- 开始时间：2026-08-15T21:20:29+08:00
- 采样时长：10.0047228s
- 并发：7
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 20251 | 20251 | 0 | 0 | 2024.46 | 3.274 | 4.187 | 4.490 | 5.351 | 12.383 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 55 | 0.003 |
| redis | 60808 | 3.003 |
| relation | 40 | 0.002 |

- Cold compute：20251（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 20251 | 0.764 |
| counter | 40 | 1.445 |
| hydrate | 20251 | 1.059 |
| inbox | 20251 | 0.828 |
| merge_dedup | 20251 | 0.013 |
| relation | 40 | 3.256 |
| route | 20251 | 0.023 |
| total | 20251 | 2.714 |

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
| docker:zg-counter | cpu_percent | 1.520 |
| docker:zg-counter | memory_percent | 0.160 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.190 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 203.680 |
| docker:zg-knowpost | memory_percent | 0.200 |
| docker:zg-knowpost | pids | 22.000 |
| docker:zg-relation | cpu_percent | 0.320 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2965.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2965.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6476622.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 52848680.000 |
| redis | connected_clients | 180.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519582.000 |
| redis | keyspace_hits | 269769618.000 |
| redis | keyspace_misses | 32176597.000 |
| redis | net_input_bytes | 12189324087.000 |
| redis | net_output_bytes | 55824709218.000 |
| redis | ops_per_sec | 14029.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 82519528.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
