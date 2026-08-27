# Feed 压测报告：hybrid / rpc / distributed-read-c512

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:23:56+08:00
- 采样时长：1m0.1577626s
- 并发：512
- 重复轮次：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 168940 | 168940 | 0 | 0 | 2808.89 | 151.154 | 390.378 | 558.838 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 1487 | 0.009 |
| redis | 508307 | 3.009 |
| relation | 168940 | 1.000 |

- Cold compute：168940（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 168940 | 11.116 |
| counter | 240 | 12.678 |
| hydrate | 168940 | 14.903 |
| inbox | 168940 | 11.935 |
| merge_dedup | 168940 | 0.029 |
| relation | 168940 | 141.271 |
| route | 168940 | 0.046 |
| total | 168940 | 179.352 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker-state:zg-counter | healthy | 1.000 |
| docker-state:zg-counter | restart_count | 0.000 |
| docker-state:zg-counter | running | 1.000 |
| docker-state:zg-gateway | healthy | 1.000 |
| docker-state:zg-gateway | restart_count | 0.000 |
| docker-state:zg-gateway | running | 1.000 |
| docker-state:zg-knowpost | healthy | 1.000 |
| docker-state:zg-knowpost | restart_count | 0.000 |
| docker-state:zg-knowpost | running | 1.000 |
| docker-state:zg-relation | healthy | 1.000 |
| docker-state:zg-relation | restart_count | 0.000 |
| docker-state:zg-relation | running | 1.000 |
| docker:zg-counter | cpu_percent | 6.390 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 3.340 |
| docker:zg-gateway | memory_percent | 0.130 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 345.030 |
| docker:zg-knowpost | memory_percent | 0.620 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 271.070 |
| docker:zg-relation | memory_percent | 0.450 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 3405875.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 7.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 27651688.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 141448413.000 |
| redis | keyspace_misses | 16892485.000 |
| redis | net_input_bytes | 6388213741.000 |
| redis | net_output_bytes | 29252882062.000 |
| redis | ops_per_sec | 23418.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95002888.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
