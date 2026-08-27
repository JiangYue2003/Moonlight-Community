# Feed 压测报告：hybrid / gateway / publish-c2

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-publish-c2`
- 开始时间：2026-08-27T05:48:10+08:00
- 采样时长：1m0.0242272s
- 并发：2
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 378 | 378 | 0 | 0 | 6.30 | 301.922 | 387.695 | 421.279 | 505.002 | 514.139 |
| publish_draft | 378 | 378 | 0 | 0 | 6.30 | 4.438 | 5.749 | 6.306 | 111.761 | 192.684 |
| publish_metadata | 378 | 378 | 0 | 0 | 6.30 | 97.975 | 126.758 | 135.748 | 275.673 | 318.288 |
| publish_confirm | 378 | 378 | 0 | 0 | 6.30 | 96.834 | 124.256 | 130.529 | 179.221 | 270.178 |
| publish_commit | 378 | 378 | 0 | 0 | 6.30 | 98.650 | 128.512 | 141.530 | 234.367 | 239.864 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 0 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 0 | 0.000 |
| relation | 0 | 0.000 |

- Cold compute：0（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：917651；input：72431748 bytes；output：16305562 bytes
- Hits/Misses：1734/6917；run hit rate：20.04%
- Evicted/Rejected：0/0；ops/s max：17155；safety epoch：26524 -> 26902

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.145 |
| client:loadtest | cpu_percent_total | 2.317 |
| client:loadtest | logical_cpus | 16.000 |
| docker-state:zg-canal | health_configured | 1.000 |
| docker-state:zg-canal | healthy | 1.000 |
| docker-state:zg-canal | restart_count | 0.000 |
| docker-state:zg-canal | running | 1.000 |
| docker-state:zg-es | health_configured | 1.000 |
| docker-state:zg-es | healthy | 1.000 |
| docker-state:zg-es | restart_count | 0.000 |
| docker-state:zg-es | running | 1.000 |
| docker-state:zg-etcd | health_configured | 1.000 |
| docker-state:zg-etcd | healthy | 1.000 |
| docker-state:zg-etcd | restart_count | 0.000 |
| docker-state:zg-etcd | running | 1.000 |
| docker-state:zg-kafka | health_configured | 1.000 |
| docker-state:zg-kafka | healthy | 1.000 |
| docker-state:zg-kafka | restart_count | 4.000 |
| docker-state:zg-kafka | running | 1.000 |
| docker-state:zg-zk | health_configured | 1.000 |
| docker-state:zg-zk | healthy | 1.000 |
| docker-state:zg-zk | restart_count | 0.000 |
| docker-state:zg-zk | running | 1.000 |
| docker:zg-canal | cpu_percent | 3.740 |
| docker:zg-canal | memory_percent | 4.220 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.410 |
| docker:zg-es | memory_percent | 12.540 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.590 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 154.800 |
| docker:zg-kafka | memory_percent | 7.470 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 39.260 |
| docker:zg-zk | memory_percent | 1.210 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 19847.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 19847.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 103275.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 6.191 |
| process:counter | cpu_seconds_total | 23.406 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 44453888.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 5.373 |
| process:gateway | cpu_seconds_total | 4.266 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 46166016.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 69.640 |
| process:knowpost | cpu_seconds_total | 270.250 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 63266816.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.096 |
| process:relation | cpu_seconds_total | 18.781 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 54554624.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.874 |
| process:search | cpu_seconds_total | 6.062 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 42553344.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 4.649 |
| process:user-storage | cpu_seconds_total | 3.547 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 47214592.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7031712.000 |
| redis | connected_clients | 76.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 26902.000 |
| redis | hit_rate | 0.558 |
| redis | keys | 571985.000 |
| redis | keyspace_hits | 99748.000 |
| redis | keyspace_misses | 84682.000 |
| redis | net_input_bytes | 547582816.000 |
| redis | net_output_bytes | 124529038.000 |
| redis | ops_per_sec | 17155.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 11508.000 |
| redis | used_memory_bytes | 80422136.000 |

## 停止施压后的恢复

- Kafka drain：5.3345138s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：378
- 测量前恢复：complete=true；耗时=4.9276737s；删除帖子/Outbox=0/0；safety epoch=26452
- 预热后恢复：complete=true；耗时=6.3462828s；删除帖子/Outbox=70/140；safety epoch=26523
- 测量后恢复：complete=true；耗时=4.9682222s；删除帖子/Outbox=378/756；safety epoch=26903

## 说明

- SLA values are reference lines, not pass/fail gates.
