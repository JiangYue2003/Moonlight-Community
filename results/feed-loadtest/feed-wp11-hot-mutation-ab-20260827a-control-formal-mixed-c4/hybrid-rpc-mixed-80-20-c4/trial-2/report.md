# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-mixed-c4`
- 开始时间：2026-08-27T06:01:10+08:00
- 采样时长：1m0.3993229s
- 并发：4
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1680 | 1680 | 0 | 0 | 27.82 | 2.072 | 2.752 | 3.215 | 4.050 | 8.419 |
| publish_total | 420 | 420 | 0 | 0 | 6.95 | 551.357 | 665.631 | 712.669 | 807.461 | 828.561 |
| publish_draft | 420 | 420 | 0 | 0 | 6.95 | 3.406 | 4.173 | 4.444 | 5.279 | 6.381 |
| publish_metadata | 420 | 420 | 0 | 0 | 6.95 | 176.780 | 222.153 | 237.472 | 256.847 | 272.062 |
| publish_confirm | 420 | 420 | 0 | 0 | 6.95 | 179.244 | 230.069 | 281.693 | 359.050 | 450.066 |
| publish_commit | 420 | 420 | 0 | 0 | 6.95 | 178.370 | 234.367 | 256.793 | 321.032 | 326.630 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 221 | 0.132 |
| mysql | 200 | 0.119 |
| redis | 5240 | 3.119 |
| relation | 1680 | 1.000 |

- Cold compute：1680（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 67200 | 40.000 |
| merge_candidates | 126866 | 75.515 |
| redis_commands | 10080 | 6.000 |
| redis_members | 156006 | 92.861 |
| redis_roundtrips | 3360 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1680 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1680 | 0.224 |
| counter | 221 | 0.488 |
| hydrate | 1680 | 0.327 |
| inbox | 1680 | 0.209 |
| merge_dedup | 1680 | 0.013 |
| relation | 1680 | 0.893 |
| route | 1680 | 0.070 |
| total | 1680 | 1.763 |

## Redis 本轮边界增量

- Commands：1055007；input：85236561 bytes；output：36350272 bytes
- Hits/Misses：85802/8193；run hit rate：91.28%
- Evicted/Rejected：0/0；ops/s max：20115；safety epoch：29450 -> 29870

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.173 |
| client:loadtest | cpu_percent_total | 2.768 |
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
| docker:zg-canal | cpu_percent | 3.700 |
| docker:zg-canal | memory_percent | 4.880 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.300 |
| docker:zg-es | memory_percent | 12.990 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 4.180 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 143.400 |
| docker:zg-kafka | memory_percent | 7.490 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 41.010 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 22187.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 22187.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 170438.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 4.654 |
| process:counter | cpu_seconds_total | 42.000 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 45654016.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.344 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 41680896.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 90.633 |
| process:knowpost | cpu_seconds_total | 582.062 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 64483328.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 14.714 |
| process:relation | cpu_seconds_total | 36.406 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 53870592.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.511 |
| process:search | cpu_seconds_total | 11.766 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 42926080.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 3.781 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 42115072.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 14728966.000 |
| redis | connected_clients | 80.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 29870.000 |
| redis | hit_rate | 0.875 |
| redis | keys | 573227.000 |
| redis | keyspace_hits | 1235863.000 |
| redis | keyspace_misses | 176090.000 |
| redis | net_input_bytes | 1178190162.000 |
| redis | net_output_bytes | 481286129.000 |
| redis | ops_per_sec | 20115.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 12289.000 |
| redis | used_memory_bytes | 81125120.000 |

## 停止施压后的恢复

- Kafka drain：5.3325642s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：420
- 测量前恢复：complete=true；耗时=4.891817s；删除帖子/Outbox=0/0；safety epoch=29372
- 预热后恢复：complete=true；耗时=6.3511081s；删除帖子/Outbox=76/152；safety epoch=29449
- 测量后恢复：complete=true；耗时=4.9959905s；删除帖子/Outbox=420/840；safety epoch=29871

## 说明

- SLA values are reference lines, not pass/fail gates.
