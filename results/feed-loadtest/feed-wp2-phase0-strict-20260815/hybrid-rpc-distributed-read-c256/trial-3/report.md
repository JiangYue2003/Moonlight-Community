# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp2-phase0-strict-20260815`
- 开始时间：2026-08-15T20:03:53+08:00
- 采样时长：1m0.0754353s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 167782 | 167782 | 0 | 0 | 2793.47 | 89.378 | 115.254 | 125.255 | 148.977 | 351.275 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 2482 | 0.015 |
| redis | 505828 | 3.015 |
| relation | 167782 | 1.000 |

- Cold compute：167782（1.000 / successful read）

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 167782 | 12.039 |
| counter | 240 | 13.724 |
| hydrate | 167782 | 16.168 |
| inbox | 167782 | 12.822 |
| merge_dedup | 167782 | 0.032 |
| relation | 167782 | 47.582 |
| route | 167782 | 0.045 |
| total | 167782 | 88.741 |

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
| docker:zg-counter | cpu_percent | 6.000 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 21.000 |
| docker:zg-gateway | cpu_percent | 1.900 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 343.210 |
| docker:zg-knowpost | memory_percent | 0.520 |
| docker:zg-knowpost | pids | 28.000 |
| docker:zg-relation | cpu_percent | 272.650 |
| docker:zg-relation | memory_percent | 0.400 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 5875062.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 133.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 47504296.000 |
| redis | connected_clients | 341.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 243932375.000 |
| redis | keyspace_misses | 29114913.000 |
| redis | net_input_bytes | 11001631022.000 |
| redis | net_output_bytes | 50467045180.000 |
| redis | ops_per_sec | 23394.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95167872.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
