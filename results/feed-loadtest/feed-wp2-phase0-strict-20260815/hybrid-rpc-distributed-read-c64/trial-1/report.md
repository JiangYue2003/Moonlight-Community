# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T19:53:26+08:00
- 采样时长：1m0.0276534s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 139660 | 139660 | 0 | 0 | 2327.03 | 27.091 | 33.279 | 35.336 | 39.950 | 63.239 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 759 | 0.005 |
| redis | 419739 | 3.005 |
| relation | 139660 | 1.000 |

- Cold compute：139660（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 139660 | 3.658 |
| counter | 240 | 4.850 |
| hydrate | 139660 | 5.292 |
| inbox | 139660 | 4.176 |
| merge_dedup | 139660 | 0.024 |
| relation | 139660 | 12.569 |
| route | 139660 | 0.018 |
| total | 139660 | 25.782 |

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
| docker:zg-counter | cpu_percent | 2.890 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.620 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 316.950 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 251.880 |
| docker:zg-relation | memory_percent | 0.400 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 4393986.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 102.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 35685118.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 182535035.000 |
| redis | keyspace_misses | 21791367.000 |
| redis | net_input_bytes | 8242486106.000 |
| redis | net_output_bytes | 37758208094.000 |
| redis | ops_per_sec | 19253.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89851776.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
