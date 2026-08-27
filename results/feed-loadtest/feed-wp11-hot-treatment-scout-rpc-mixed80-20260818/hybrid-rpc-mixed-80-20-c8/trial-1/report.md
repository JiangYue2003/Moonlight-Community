# Feed 压测报告：hybrid / rpc / mixed-80-20-c8

- Run ID：`feed-wp11-hot-treatment-scout-rpc-mixed80-20260818`
- 开始时间：2026-08-18T22:34:44+08:00
- 采样时长：15.1165966s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 320 | 320 | 0 | 0 | 21.17 | 1.636 | 4.784 | 5.789 | 7.955 | 9.764 |
| publish_total | 80 | 80 | 0 | 0 | 5.29 | 1471.344 | 1642.840 | 1647.405 | 1670.912 | 1670.912 |
| publish_draft | 80 | 80 | 0 | 0 | 5.29 | 3.922 | 4.833 | 5.344 | 6.354 | 6.354 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.29 | 482.179 | 561.620 | 580.625 | 593.194 | 593.194 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.29 | 487.151 | 550.555 | 559.353 | 580.192 | 580.192 |
| publish_commit | 80 | 80 | 0 | 0 | 5.29 | 516.372 | 581.893 | 630.429 | 653.057 | 653.057 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 166 | 0.519 |
| counter | 55 | 0.172 |
| mysql | 36 | 0.113 |
| redis | 668 | 2.087 |
| relation | 55 | 0.172 |

- Cold compute：158（0.494 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6320 | 19.750 |
| merge_candidates | 7679 | 23.997 |
| redis_commands | 948 | 2.962 |
| redis_members | 7679 | 23.997 |
| redis_roundtrips | 158 | 0.494 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 154 |
| l2_fresh | 0 |
| miss | 166 |

- L1+L2 Fresh ratio：48.12%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 158 | 0.536 |
| counter | 55 | 0.800 |
| hydrate | 158 | 0.881 |
| inbox | 158 | 0.536 |
| merge_dedup | 158 | 0.004 |
| relation | 55 | 1.260 |
| route | 158 | 0.721 |
| total | 320 | 1.858 |

## Redis 本轮边界增量

- Commands：260238；input：20962648 bytes；output：5844458 bytes
- Hits/Misses：8004/2666；run hit rate：75.01%
- Evicted/Rejected：0/0；ops/s max：19151；safety epoch：6305 -> 6465

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.103 |
| client:loadtest | cpu_percent_total | 1.654 |
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
| docker-state:zg-kafka | restart_count | 0.000 |
| docker-state:zg-kafka | running | 1.000 |
| docker-state:zg-zk | health_configured | 1.000 |
| docker-state:zg-zk | healthy | 1.000 |
| docker-state:zg-zk | restart_count | 0.000 |
| docker-state:zg-zk | running | 1.000 |
| docker:zg-canal | cpu_percent | 2.030 |
| docker:zg-canal | memory_percent | 3.500 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.020 |
| docker:zg-es | memory_percent | 11.650 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.980 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 150.700 |
| docker:zg-kafka | memory_percent | 7.510 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 22.730 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5118.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 5118.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23047.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 17.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.881 |
| process:counter | cpu_seconds_total | 481.438 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 43491328.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.234 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 37527552.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 74.352 |
| process:knowpost | cpu_seconds_total | 559.422 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74698752.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.323 |
| process:relation | cpu_seconds_total | 14.859 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 46534656.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.776 |
| process:search | cpu_seconds_total | 12.953 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 42340352.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 10.531 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35155968.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 16000804.000 |
| redis | connected_clients | 89.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 6465.000 |
| redis | hit_rate | 0.708 |
| redis | keys | 595874.000 |
| redis | keyspace_hits | 357140.000 |
| redis | keyspace_misses | 147612.000 |
| redis | net_input_bytes | 1157405583.000 |
| redis | net_output_bytes | 321008877.000 |
| redis | ops_per_sec | 19151.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44085.000 |
| redis | used_memory_bytes | 98366216.000 |

## 停止施压后的恢复

- Kafka drain：5.1745706s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.9103479s；删除帖子/Outbox=0/0；safety epoch=6255
- 预热后恢复：complete=true；耗时=6.2474602s；删除帖子/Outbox=24/48；safety epoch=6304
- 测量后恢复：complete=true；耗时=4.9954657s；删除帖子/Outbox=80/160；safety epoch=6466

## 说明

- SLA values are reference lines, not pass/fail gates.
