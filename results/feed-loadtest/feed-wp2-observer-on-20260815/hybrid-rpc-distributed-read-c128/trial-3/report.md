# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-observer-on-20260815`
- 开始时间：2026-08-15T18:58:24+08:00
- 采样时长：20.037633s
- 并发：128
- 重复轮次：3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 52144 | 52144 | 0 | 0 | 2602.77 | 48.376 | 64.066 | 74.058 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 80 | 0.002 |
| mysql | 1206 | 0.023 |
| redis | 157638 | 3.023 |
| relation | 52144 | 1.000 |

- Cold compute：52144（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 52144 | 6.719 |
| counter | 80 | 8.687 |
| hydrate | 52144 | 9.896 |
| inbox | 52144 | 7.579 |
| merge_dedup | 52144 | 0.028 |
| relation | 52144 | 22.506 |
| route | 52144 | 0.028 |
| total | 52144 | 46.805 |

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
| docker:zg-counter | cpu_percent | 3.020 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 17.000 |
| docker:zg-gateway | cpu_percent | 0.400 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 330.850 |
| docker:zg-knowpost | memory_percent | 0.390 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 261.220 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 564963.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4896938.000 |
| redis | connected_clients | 270.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 23501589.000 |
| redis | keyspace_misses | 2812728.000 |
| redis | net_input_bytes | 1082867641.000 |
| redis | net_output_bytes | 4838611088.000 |
| redis | ops_per_sec | 21218.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90815000.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
