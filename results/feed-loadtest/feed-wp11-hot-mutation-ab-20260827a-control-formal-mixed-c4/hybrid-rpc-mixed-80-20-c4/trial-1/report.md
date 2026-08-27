# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-mixed-c4`
- 开始时间：2026-08-27T05:59:34+08:00
- 采样时长：1m0.1233155s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1648 | 1648 | 0 | 0 | 27.41 | 1.988 | 2.686 | 3.001 | 3.702 | 5.899 |
| publish_total | 412 | 412 | 0 | 0 | 6.85 | 560.929 | 678.666 | 716.333 | 760.685 | 789.412 |
| publish_draft | 412 | 412 | 0 | 0 | 6.85 | 3.473 | 4.126 | 4.285 | 4.838 | 9.054 |
| publish_metadata | 412 | 412 | 0 | 0 | 6.85 | 182.197 | 227.739 | 252.649 | 288.962 | 340.603 |
| publish_confirm | 412 | 412 | 0 | 0 | 6.85 | 182.261 | 223.232 | 250.319 | 327.555 | 344.156 |
| publish_commit | 412 | 412 | 0 | 0 | 6.85 | 182.312 | 231.044 | 249.544 | 274.131 | 287.407 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 215 | 0.130 |
| mysql | 185 | 0.112 |
| redis | 5129 | 3.112 |
| relation | 1648 | 1.000 |

- Cold compute：1648（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 65920 | 40.000 |
| merge_candidates | 124219 | 75.376 |
| redis_commands | 9888 | 6.000 |
| redis_members | 151658 | 92.025 |
| redis_roundtrips | 3296 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1648 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1648 | 0.214 |
| counter | 215 | 0.486 |
| hydrate | 1648 | 0.334 |
| inbox | 1648 | 0.208 |
| merge_dedup | 1648 | 0.010 |
| relation | 1648 | 0.892 |
| route | 1648 | 0.070 |
| total | 1648 | 1.755 |

## Redis 本轮边界增量

- Commands：1036828；input：83750468 bytes；output：35636986 bytes
- Hits/Misses：84159/7826；run hit rate：91.49%
- Evicted/Rejected：0/0；ops/s max：19203；safety epoch：28958 -> 29370

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.172 |
| client:loadtest | cpu_percent_total | 2.755 |
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
| docker:zg-canal | cpu_percent | 3.790 |
| docker:zg-canal | memory_percent | 4.890 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 11.420 |
| docker:zg-es | memory_percent | 12.980 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 4.390 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 141.860 |
| docker:zg-kafka | memory_percent | 7.500 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 44.840 |
| docker:zg-zk | memory_percent | 1.200 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 21798.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 21798.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 160809.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 8.521 |
| process:counter | cpu_seconds_total | 39.703 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 45051904.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 4.344 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 41345024.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 89.745 |
| process:knowpost | cpu_seconds_total | 532.625 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 62971904.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.197 |
| process:relation | cpu_seconds_total | 34.156 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 53796864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.902 |
| process:search | cpu_seconds_total | 11.016 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 43225088.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.780 |
| process:user-storage | cpu_seconds_total | 3.750 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 42115072.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 13459079.000 |
| redis | connected_clients | 80.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 29370.000 |
| redis | hit_rate | 0.874 |
| redis | keys | 573119.000 |
| redis | keyspace_hits | 1124185.000 |
| redis | keyspace_misses | 161566.000 |
| redis | net_input_bytes | 1075903232.000 |
| redis | net_output_bytes | 438116972.000 |
| redis | ops_per_sec | 19203.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 12192.000 |
| redis | used_memory_bytes | 81121904.000 |

## 停止施压后的恢复

- Kafka drain：5.3048717s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：412
- 测量前恢复：complete=true；耗时=4.9601551s；删除帖子/Outbox=0/0；safety epoch=28876
- 预热后恢复：complete=true；耗时=6.394119s；删除帖子/Outbox=80/160；safety epoch=28957
- 测量后恢复：complete=true；耗时=5.0710657s；删除帖子/Outbox=412/824；safety epoch=29371

## 说明

- SLA values are reference lines, not pass/fail gates.
