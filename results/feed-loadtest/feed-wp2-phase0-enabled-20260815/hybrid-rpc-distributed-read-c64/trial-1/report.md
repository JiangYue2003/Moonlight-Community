# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:04:21+08:00
- 采样时长：1m0.0307172s
- 并发：64
- 重复轮次：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 137575 | 137575 | 0 | 0 | 2292.16 | 27.478 | 36.185 | 40.861 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 241 | 0.002 |
| mysql | 626 | 0.005 |
| redis | 413351 | 3.005 |
| relation | 137575 | 1.000 |

- Cold compute：137575（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 137575 | 3.705 |
| counter | 241 | 5.305 |
| hydrate | 137575 | 5.260 |
| inbox | 137575 | 4.203 |
| merge_dedup | 137575 | 0.025 |
| relation | 137575 | 12.843 |
| route | 137575 | 0.020 |
| total | 137575 | 26.102 |

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
| docker:zg-counter | cpu_percent | 7.200 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 3.210 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 317.060 |
| docker:zg-knowpost | memory_percent | 0.360 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 256.840 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 1140475.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 103.000 |
| mysql | threads_running | 7.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9543507.000 |
| redis | connected_clients | 270.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 47566582.000 |
| redis | keyspace_misses | 5678534.000 |
| redis | net_input_bytes | 2164202007.000 |
| redis | net_output_bytes | 9818882243.000 |
| redis | ops_per_sec | 18740.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 88277376.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
