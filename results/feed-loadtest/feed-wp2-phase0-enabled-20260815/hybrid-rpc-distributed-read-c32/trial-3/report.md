# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:02:39+08:00
- 采样时长：1m0.0207061s
- 并发：32
- 重复轮次：3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 116862 | 116862 | 0 | 0 | 1947.33 | 16.162 | 20.894 | 23.714 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 425 | 0.004 |
| redis | 351011 | 3.004 |
| relation | 116862 | 1.000 |

- Cold compute：116862（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 116862 | 2.112 |
| counter | 240 | 3.050 |
| hydrate | 116862 | 2.970 |
| inbox | 116862 | 2.348 |
| merge_dedup | 116862 | 0.022 |
| relation | 116862 | 7.556 |
| route | 116862 | 0.013 |
| total | 116862 | 15.063 |

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
| docker:zg-counter | cpu_percent | 2.750 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 3.250 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 298.330 |
| docker:zg-knowpost | memory_percent | 0.340 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 235.660 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 978781.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 60.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8240966.000 |
| redis | connected_clients | 270.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 40811420.000 |
| redis | keyspace_misses | 4874186.000 |
| redis | net_input_bytes | 1860760145.000 |
| redis | net_output_bytes | 8420775112.000 |
| redis | ops_per_sec | 16263.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86234936.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
