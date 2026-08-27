# Feed 压测报告：hybrid / gateway / publish-c2

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-publish-c2`
- 开始时间：2026-08-27T05:44:53+08:00
- 采样时长：1m0.0854363s
- 并发：2
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 378 | 378 | 0 | 0 | 6.29 | 303.624 | 384.253 | 428.420 | 511.465 | 648.633 |
| publish_draft | 378 | 378 | 0 | 0 | 6.29 | 4.447 | 5.777 | 6.434 | 52.070 | 219.442 |
| publish_metadata | 378 | 378 | 0 | 0 | 6.29 | 97.348 | 126.842 | 135.449 | 182.395 | 356.042 |
| publish_confirm | 378 | 378 | 0 | 0 | 6.29 | 97.709 | 129.035 | 137.339 | 239.984 | 271.362 |
| publish_commit | 378 | 378 | 0 | 0 | 6.29 | 98.074 | 133.551 | 157.219 | 213.641 | 234.323 |

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

- Commands：896021；input：70702233 bytes；output：15895279 bytes
- Hits/Misses：1736/6935；run hit rate：20.02%
- Evicted/Rejected：0/0；ops/s max：17676；safety epoch：25616 -> 25994

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.088 |
| client:loadtest | cpu_percent_total | 1.404 |
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
| docker:zg-canal | cpu_percent | 16.830 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.050 |
| docker:zg-es | memory_percent | 12.530 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 4.530 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 153.870 |
| docker:zg-kafka | memory_percent | 7.670 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 16.710 |
| docker:zg-zk | memory_percent | 1.130 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 19148.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 19148.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 89728.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 5.422 |
| process:counter | cpu_seconds_total | 18.531 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 43986944.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 6.359 |
| process:gateway | cpu_seconds_total | 1.734 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 44503040.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 79.057 |
| process:knowpost | cpu_seconds_total | 185.234 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 63258624.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.385 |
| process:relation | cpu_seconds_total | 17.703 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 54755328.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.058 |
| process:search | cpu_seconds_total | 3.938 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 42659840.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.325 |
| process:user-storage | cpu_seconds_total | 1.766 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 46866432.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4811461.000 |
| redis | connected_clients | 75.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 25994.000 |
| redis | hit_rate | 0.586 |
| redis | keys | 570127.000 |
| redis | keyspace_hits | 74735.000 |
| redis | keyspace_misses | 58493.000 |
| redis | net_input_bytes | 372911453.000 |
| redis | net_output_bytes | 84581146.000 |
| redis | ops_per_sec | 17676.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 11313.000 |
| redis | used_memory_bytes | 80212616.000 |

## 停止施压后的恢复

- Kafka drain：6.8651391s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：378
- 测量前恢复：complete=true；耗时=4.9598325s；删除帖子/Outbox=0/0；safety epoch=25540
- 预热后恢复：complete=true；耗时=6.381233s；删除帖子/Outbox=74/148；safety epoch=25615
- 测量后恢复：complete=true；耗时=4.9575067s；删除帖子/Outbox=378/756；safety epoch=25995

## 说明

- SLA values are reference lines, not pass/fail gates.
