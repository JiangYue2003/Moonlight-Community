# Feed 压测报告：hybrid / rpc / distributed-read-c512

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:26:33+08:00
- 采样时长：1m0.1942437s
- 并发：512
- 重复轮次：3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 167903 | 167903 | 0 | 0 | 2789.98 | 153.025 | 389.646 | 556.207 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 2425 | 0.014 |
| redis | 506134 | 3.014 |
| relation | 167903 | 1.000 |

- Cold compute：167903（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 167903 | 11.529 |
| counter | 240 | 14.168 |
| hydrate | 167903 | 15.764 |
| inbox | 167903 | 12.323 |
| merge_dedup | 167903 | 0.032 |
| relation | 167903 | 140.845 |
| route | 167903 | 0.053 |
| total | 167903 | 180.598 |

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
| docker:zg-counter | cpu_percent | 6.520 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 3.790 |
| docker:zg-gateway | memory_percent | 0.130 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 343.960 |
| docker:zg-knowpost | memory_percent | 0.650 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 272.440 |
| docker:zg-relation | memory_percent | 0.470 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 3805413.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 30827323.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 157935102.000 |
| redis | keyspace_misses | 18862060.000 |
| redis | net_input_bytes | 7129754678.000 |
| redis | net_output_bytes | 32665684220.000 |
| redis | ops_per_sec | 22982.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 94958392.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
