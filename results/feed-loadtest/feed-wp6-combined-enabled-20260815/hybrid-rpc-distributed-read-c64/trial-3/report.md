# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp6-combined-enabled-20260815`
- 开始时间：2026-08-15T22:13:23+08:00
- 采样时长：1m0.045385s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 597103 | 597103 | 0 | 0 | 9950.94 | 6.295 | 8.379 | 9.088 | 10.855 | 37.865 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.000 |
| mysql | 123 | 0.000 |
| redis | 1194329 | 2.000 |
| relation | 240 | 0.000 |

- Cold compute：597103（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 597103 | 2.504 |
| counter | 240 | 2.952 |
| hydrate | 597103 | 2.437 |
| inbox | 597103 | 2.502 |
| merge_dedup | 597103 | 0.002 |
| relation | 240 | 6.938 |
| route | 597103 | 0.026 |
| total | 597103 | 4.991 |

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
| docker:zg-counter | cpu_percent | 4.050 |
| docker:zg-counter | memory_percent | 0.160 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 1.590 |
| docker:zg-gateway | memory_percent | 0.210 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 468.630 |
| docker:zg-knowpost | memory_percent | 0.520 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 1.660 |
| docker:zg-relation | memory_percent | 0.340 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2966.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2966.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6511070.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 87.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 102414383.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.878 |
| redis | keys | 515092.000 |
| redis | keyspace_hits | 420272738.000 |
| redis | keyspace_misses | 60300352.000 |
| redis | net_input_bytes | 20131465345.000 |
| redis | net_output_bytes | 86088558533.000 |
| redis | ops_per_sec | 70857.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 83934952.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
