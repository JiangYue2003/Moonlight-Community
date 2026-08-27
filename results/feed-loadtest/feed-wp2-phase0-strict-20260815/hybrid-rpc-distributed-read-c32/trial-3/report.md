# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T19:52:07+08:00
- 采样时长：1m0.02023s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 118304 | 118304 | 0 | 0 | 1971.36 | 15.982 | 19.374 | 20.547 | 23.044 | 37.331 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 392 | 0.003 |
| redis | 355304 | 3.003 |
| relation | 118304 | 1.000 |

- Cold compute：118304（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 118304 | 2.090 |
| counter | 240 | 3.056 |
| hydrate | 118304 | 2.943 |
| inbox | 118304 | 2.327 |
| merge_dedup | 118304 | 0.022 |
| relation | 118304 | 7.453 |
| route | 118304 | 0.013 |
| total | 118304 | 14.889 |

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
| docker:zg-counter | cpu_percent | 2.420 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 2.750 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 297.070 |
| docker:zg-knowpost | memory_percent | 0.400 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 239.710 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 4229455.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 58.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 34364160.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 175673953.000 |
| redis | keyspace_misses | 20974078.000 |
| redis | net_input_bytes | 7934363771.000 |
| redis | net_output_bytes | 36338172056.000 |
| redis | ops_per_sec | 16376.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87642312.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
