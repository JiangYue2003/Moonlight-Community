# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:06:58+08:00
- 采样时长：1m0.0327978s
- 并发：64
- 重复轮次：3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 133886 | 133886 | 0 | 0 | 2230.58 | 28.195 | 37.341 | 42.600 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 824 | 0.006 |
| redis | 402482 | 3.006 |
| relation | 133886 | 1.000 |

- Cold compute：133886（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 133886 | 3.773 |
| counter | 240 | 5.352 |
| hydrate | 133886 | 5.427 |
| inbox | 133886 | 4.295 |
| merge_dedup | 133886 | 0.025 |
| relation | 133886 | 13.256 |
| route | 133886 | 0.020 |
| total | 133886 | 26.843 |

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
| docker:zg-counter | cpu_percent | 5.090 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 2.920 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 325.210 |
| docker:zg-knowpost | memory_percent | 0.370 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 255.300 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 1460483.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 119.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12112903.000 |
| redis | connected_clients | 270.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 60907586.000 |
| redis | keyspace_misses | 7267842.000 |
| redis | net_input_bytes | 2763392524.000 |
| redis | net_output_bytes | 12580011363.000 |
| redis | ops_per_sec | 18751.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87906056.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
