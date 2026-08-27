# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp6-combined-enabled-20260815`
- 开始时间：2026-08-15T22:14:42+08:00
- 采样时长：1m0.0540329s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 638407 | 638407 | 0 | 0 | 10638.67 | 11.534 | 16.567 | 18.267 | 21.498 | 63.114 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.000 |
| mysql | 207 | 0.000 |
| redis | 1277021 | 2.000 |
| relation | 240 | 0.000 |

- Cold compute：638407（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 638407 | 5.188 |
| counter | 240 | 4.614 |
| hydrate | 638407 | 5.108 |
| inbox | 638407 | 5.187 |
| merge_dedup | 638407 | 0.002 |
| relation | 240 | 10.361 |
| route | 638407 | 0.060 |
| total | 638407 | 10.379 |

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
| docker:zg-counter | cpu_percent | 1.980 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 0.540 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 533.180 |
| docker:zg-knowpost | memory_percent | 0.670 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 0.730 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2966.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2966.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6511746.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 87.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 107584896.000 |
| redis | connected_clients | 254.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.874 |
| redis | keys | 515092.000 |
| redis | keyspace_hits | 423225986.000 |
| redis | keyspace_misses | 63248616.000 |
| redis | net_input_bytes | 20512974786.000 |
| redis | net_output_bytes | 86408568034.000 |
| redis | ops_per_sec | 79643.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87812080.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
