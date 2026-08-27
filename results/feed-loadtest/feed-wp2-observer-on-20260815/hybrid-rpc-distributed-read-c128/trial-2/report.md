# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-observer-on-20260815`
- 开始时间：2026-08-15T18:57:51+08:00
- 采样时长：20.0399961s
- 并发：128
- 重复轮次：2
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 52705 | 52705 | 0 | 0 | 2630.54 | 47.998 | 62.454 | 70.572 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 80 | 0.002 |
| mysql | 157 | 0.003 |
| redis | 158272 | 3.003 |
| relation | 52705 | 1.000 |

- Cold compute：52705（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 52705 | 6.677 |
| counter | 80 | 8.610 |
| hydrate | 52705 | 9.338 |
| inbox | 52705 | 7.544 |
| merge_dedup | 52705 | 0.026 |
| relation | 52705 | 22.646 |
| route | 52705 | 0.030 |
| total | 52705 | 46.308 |

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
| docker:zg-counter | cpu_percent | 2.760 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 17.000 |
| docker:zg-gateway | cpu_percent | 2.870 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 324.000 |
| docker:zg-knowpost | memory_percent | 0.400 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 258.780 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 498209.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4367024.000 |
| redis | connected_clients | 270.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 20756815.000 |
| redis | keyspace_misses | 2484403.000 |
| redis | net_input_bytes | 959266523.000 |
| redis | net_output_bytes | 4270422199.000 |
| redis | ops_per_sec | 21443.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90523304.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
