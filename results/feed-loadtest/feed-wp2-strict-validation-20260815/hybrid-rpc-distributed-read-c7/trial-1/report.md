# Feed 压测报告：hybrid / rpc / distributed-read-c7

- Run ID：`feed-wp2-strict-validation-20260815`
- 开始时间：2026-08-15T19:48:09+08:00
- 采样时长：10.006198s
- 并发：7
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 10579 | 10579 | 0 | 0 | 1057.30 | 6.453 | 7.505 | 7.959 | 9.007 | 12.983 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.004 |
| mysql | 0 | 0.000 |
| redis | 31737 | 3.000 |
| relation | 10579 | 1.000 |

- Cold compute：10579（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 10579 | 0.746 |
| counter | 40 | 1.627 |
| hydrate | 10579 | 1.072 |
| inbox | 10579 | 0.823 |
| merge_dedup | 10579 | 0.015 |
| relation | 10579 | 3.151 |
| route | 10579 | 0.011 |
| total | 10579 | 5.847 |

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
| docker:zg-counter | cpu_percent | 1.640 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 1.590 |
| docker:zg-gateway | memory_percent | 0.130 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 159.890 |
| docker:zg-knowpost | memory_percent | 0.350 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 135.660 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 3818274.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 31041730.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519484.000 |
| redis | keyspace_hits | 158471137.000 |
| redis | keyspace_misses | 18925860.000 |
| redis | net_input_bytes | 7161434644.000 |
| redis | net_output_bytes | 32778215620.000 |
| redis | ops_per_sec | 8636.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86210464.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
