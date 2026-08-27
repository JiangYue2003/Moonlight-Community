# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T20:01:17+08:00
- 采样时长：1m0.0745986s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 167432 | 167432 | 0 | 0 | 2787.75 | 89.870 | 114.180 | 123.603 | 145.376 | 233.852 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 2743 | 0.016 |
| redis | 505039 | 3.016 |
| relation | 167432 | 1.000 |

- Cold compute：167432（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 167432 | 12.515 |
| counter | 240 | 13.921 |
| hydrate | 167432 | 16.700 |
| inbox | 167432 | 13.279 |
| merge_dedup | 167432 | 0.033 |
| relation | 167432 | 46.319 |
| route | 167432 | 0.052 |
| total | 167432 | 88.954 |

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
| docker:zg-counter | cpu_percent | 4.940 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 3.140 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 341.990 |
| docker:zg-knowpost | memory_percent | 0.530 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 273.670 |
| docker:zg-relation | memory_percent | 0.420 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 5475229.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 9.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 44325959.000 |
| redis | connected_clients | 339.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 227420583.000 |
| redis | keyspace_misses | 27143852.000 |
| redis | net_input_bytes | 10259327973.000 |
| redis | net_output_bytes | 47049019352.000 |
| redis | ops_per_sec | 23102.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95285672.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
