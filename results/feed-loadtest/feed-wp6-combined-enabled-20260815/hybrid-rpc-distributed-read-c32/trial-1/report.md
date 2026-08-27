# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp6-combined-enabled-20260815`
- 开始时间：2026-08-15T22:06:51+08:00
- 采样时长：1m0.0362736s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 478851 | 478851 | 0 | 0 | 7980.49 | 3.843 | 5.117 | 5.530 | 6.787 | 25.533 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 32 | 0.000 |
| redis | 957734 | 2.000 |
| relation | 240 | 0.001 |

- Cold compute：478851（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 478851 | 1.449 |
| counter | 240 | 2.275 |
| hydrate | 478851 | 1.396 |
| inbox | 478851 | 1.447 |
| merge_dedup | 478851 | 0.002 |
| relation | 240 | 5.913 |
| route | 478851 | 0.018 |
| total | 478851 | 2.887 |

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
| docker:zg-counter | cpu_percent | 1.610 |
| docker:zg-counter | memory_percent | 0.170 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 1.030 |
| docker:zg-gateway | memory_percent | 0.220 |
| docker:zg-gateway | pids | 26.000 |
| docker:zg-knowpost | cpu_percent | 422.720 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 2.480 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2966.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2966.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6508334.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 46.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 79981072.000 |
| redis | connected_clients | 178.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.899 |
| redis | keys | 515093.000 |
| redis | keyspace_hits | 407459107.000 |
| redis | keyspace_misses | 47512671.000 |
| redis | net_input_bytes | 18476533370.000 |
| redis | net_output_bytes | 84700419089.000 |
| redis | ops_per_sec | 56577.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 82251160.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
