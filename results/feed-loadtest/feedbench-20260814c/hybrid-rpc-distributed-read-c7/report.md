# Feed 压测报告：hybrid / rpc / distributed-read-c7

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-15T18:45:11+08:00
- 采样时长：10.005779s
- 并发：7
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 10623 | 10623 | 0 | 0 | 1061.74 | 6.396 | 8.024 | 9.345 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.004 |
| mysql | 0 | 0.000 |
| redis | 31869 | 3.000 |
| relation | 10623 | 1.000 |

- Cold compute：10623（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 10623 | 0.740 |
| counter | 40 | 1.663 |
| hydrate | 10623 | 1.060 |
| inbox | 10623 | 0.818 |
| merge_dedup | 10623 | 0.015 |
| relation | 10623 | 3.153 |
| route | 10623 | 0.011 |
| total | 10623 | 5.826 |

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
| docker:zg-counter | cpu_percent | 1.490 |
| docker:zg-counter | memory_percent | 0.090 |
| docker:zg-counter | pids | 16.000 |
| docker:zg-gateway | cpu_percent | 0.220 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 168.850 |
| docker:zg-knowpost | memory_percent | 0.160 |
| docker:zg-knowpost | pids | 22.000 |
| docker:zg-relation | cpu_percent | 143.580 |
| docker:zg-relation | memory_percent | 0.130 |
| docker:zg-relation | pids | 21.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 166714.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 13.000 |
| mysql | threads_running | 3.000 |
| prometheus:knowpost | cold_compute:success | 10635.000 |
| prometheus:knowpost | dependency_calls:counter:batch_follower_counts:success | 41.000 |
| prometheus:knowpost | dependency_calls:mysql:feed_item_db:success | 3.000 |
| prometheus:knowpost | dependency_calls:redis:bigv_pipeline:success | 10631.000 |
| prometheus:knowpost | dependency_calls:redis:feed_item_cache_write:success | 3.000 |
| prometheus:knowpost | dependency_calls:redis:feed_item_mget:success | 10635.000 |
| prometheus:knowpost | dependency_calls:redis:inbox_read:success | 10635.000 |
| prometheus:knowpost | dependency_calls:relation:list_followings:success | 10635.000 |
| prometheus:knowpost | stage_duration_count:bigv_pipeline:success | 10631.000 |
| prometheus:knowpost | stage_duration_count:counter:success | 41.000 |
| prometheus:knowpost | stage_duration_count:hydrate:success | 10635.000 |
| prometheus:knowpost | stage_duration_count:inbox:success | 10635.000 |
| prometheus:knowpost | stage_duration_count:merge_dedup:success | 10635.000 |
| prometheus:knowpost | stage_duration_count:relation:success | 10635.000 |
| prometheus:knowpost | stage_duration_count:route:success | 10635.000 |
| prometheus:knowpost | stage_duration_count:total:success | 10635.000 |
| prometheus:knowpost | stage_duration_sum:bigv_pipeline:success | 7865.204 |
| prometheus:knowpost | stage_duration_sum:counter:success | 67.864 |
| prometheus:knowpost | stage_duration_sum:hydrate:success | 11277.914 |
| prometheus:knowpost | stage_duration_sum:inbox:success | 8699.466 |
| prometheus:knowpost | stage_duration_sum:merge_dedup:success | 157.531 |
| prometheus:knowpost | stage_duration_sum:relation:success | 33531.981 |
| prometheus:knowpost | stage_duration_sum:route:success | 113.487 |
| prometheus:knowpost | stage_duration_sum:total:success | 61960.240 |
| prometheus:knowpost | stage_total:bigv_pipeline:success | 10631.000 |
| prometheus:knowpost | stage_total:counter:success | 41.000 |
| prometheus:knowpost | stage_total:hydrate:success | 10635.000 |
| prometheus:knowpost | stage_total:inbox:success | 10635.000 |
| prometheus:knowpost | stage_total:merge_dedup:success | 10635.000 |
| prometheus:knowpost | stage_total:relation:success | 10635.000 |
| prometheus:knowpost | stage_total:route:success | 10635.000 |
| prometheus:knowpost | stage_total:total:success | 10635.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1654790.000 |
| redis | connected_clients | 31.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519568.000 |
| redis | keyspace_hits | 6978856.000 |
| redis | keyspace_misses | 837963.000 |
| redis | net_input_bytes | 335473533.000 |
| redis | net_output_bytes | 1417341978.000 |
| redis | ops_per_sec | 8578.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 79631640.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
