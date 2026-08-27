# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:00:03+08:00
- 采样时长：1m0.0186081s
- 并发：32
- 重复轮次：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 117898 | 117898 | 0 | 0 | 1964.65 | 16.009 | 20.724 | 23.718 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 250 | 0.002 |
| redis | 353944 | 3.002 |
| relation | 117898 | 1.000 |

- Cold compute：117898（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 117898 | 2.086 |
| counter | 240 | 3.191 |
| hydrate | 117898 | 2.931 |
| inbox | 117898 | 2.323 |
| merge_dedup | 117898 | 0.022 |
| relation | 117898 | 7.501 |
| route | 117898 | 0.013 |
| total | 117898 | 14.917 |

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
| docker:zg-counter | cpu_percent | 2.620 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 17.000 |
| docker:zg-gateway | cpu_percent | 3.480 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 300.770 |
| docker:zg-knowpost | memory_percent | 0.340 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 237.130 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 703373.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6019539.000 |
| redis | connected_clients | 270.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 29292240.000 |
| redis | keyspace_misses | 3502714.000 |
| redis | net_input_bytes | 1343428941.000 |
| redis | net_output_bytes | 6037030649.000 |
| redis | ops_per_sec | 16366.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86357792.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
