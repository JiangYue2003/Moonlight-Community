# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:09:57+08:00
- 采样时长：1m0.0382073s
- 并发：128
- 重复轮次：2
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 153313 | 153313 | 0 | 0 | 2554.12 | 49.365 | 65.044 | 74.852 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 1543 | 0.010 |
| redis | 461482 | 3.010 |
| relation | 153313 | 1.000 |

- Cold compute：153313（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 153313 | 6.802 |
| counter | 240 | 8.417 |
| hydrate | 153313 | 9.729 |
| inbox | 153313 | 7.761 |
| merge_dedup | 153313 | 0.030 |
| relation | 153313 | 23.262 |
| route | 153313 | 0.031 |
| total | 153313 | 47.663 |

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
| docker:zg-counter | cpu_percent | 5.850 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 3.640 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 328.640 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 259.590 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 1823033.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 15014202.000 |
| redis | connected_clients | 275.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 75962113.000 |
| redis | keyspace_misses | 9064329.000 |
| redis | net_input_bytes | 3440274073.000 |
| redis | net_output_bytes | 15696240436.000 |
| redis | ops_per_sec | 21100.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91078888.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
