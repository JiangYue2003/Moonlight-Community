# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:15:35+08:00
- 采样时长：1m0.0733794s
- 并发：256
- 重复轮次：3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 169942 | 169942 | 0 | 0 | 2829.55 | 88.207 | 123.626 | 147.420 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 2785 | 0.016 |
| redis | 512611 | 3.016 |
| relation | 169942 | 1.000 |

- Cold compute：169942（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 169942 | 11.479 |
| counter | 240 | 13.381 |
| hydrate | 169942 | 15.609 |
| inbox | 169942 | 12.250 |
| merge_dedup | 169942 | 0.033 |
| relation | 169942 | 48.055 |
| route | 169942 | 0.047 |
| total | 169942 | 87.529 |

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
| docker:zg-counter | cpu_percent | 7.820 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 2.610 |
| docker:zg-gateway | memory_percent | 0.120 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 345.960 |
| docker:zg-knowpost | memory_percent | 0.520 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 269.290 |
| docker:zg-relation | memory_percent | 0.410 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 2610267.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 134.000 |
| mysql | threads_running | 8.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 21286091.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 108525162.000 |
| redis | keyspace_misses | 12952604.000 |
| redis | net_input_bytes | 4904622778.000 |
| redis | net_output_bytes | 22437042473.000 |
| redis | ops_per_sec | 23336.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95494968.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
