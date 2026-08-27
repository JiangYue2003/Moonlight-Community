# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:08:39+08:00
- 采样时长：1m0.0391277s
- 并发：128
- 重复轮次：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 153622 | 153622 | 0 | 0 | 2559.22 | 49.346 | 64.609 | 73.778 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 1439 | 0.009 |
| redis | 462305 | 3.009 |
| relation | 153622 | 1.000 |

- Cold compute：153622（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 153622 | 6.849 |
| counter | 240 | 8.166 |
| hydrate | 153622 | 9.681 |
| inbox | 153622 | 7.791 |
| merge_dedup | 153622 | 0.029 |
| relation | 153622 | 23.064 |
| route | 153622 | 0.029 |
| total | 153622 | 47.493 |

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
| docker:zg-counter | cpu_percent | 7.500 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 3.670 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 329.210 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 262.170 |
| docker:zg-relation | memory_percent | 0.350 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 1641893.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 13566643.000 |
| redis | connected_clients | 275.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 68442485.000 |
| redis | keyspace_misses | 8167614.000 |
| redis | net_input_bytes | 3102391931.000 |
| redis | net_output_bytes | 14139759128.000 |
| redis | ops_per_sec | 21031.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90759952.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
