# Feed 压测报告：hybrid / rpc / publish-c8

- Run ID：`feed-wp11-hot-treatment-scout-rpc-publish-low-20260818`
- 开始时间：2026-08-18T22:25:22+08:00
- 采样时长：15.9806289s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 88 | 88 | 0 | 0 | 5.51 | 1432.138 | 1571.396 | 1580.532 | 1610.945 | 1610.945 |
| publish_draft | 88 | 88 | 0 | 0 | 5.51 | 4.168 | 6.450 | 6.996 | 7.330 | 7.330 |
| publish_metadata | 88 | 88 | 0 | 0 | 5.51 | 503.232 | 562.081 | 603.500 | 626.155 | 626.155 |
| publish_confirm | 88 | 88 | 0 | 0 | 5.51 | 467.038 | 540.632 | 543.134 | 551.876 | 551.876 |
| publish_commit | 88 | 88 | 0 | 0 | 5.51 | 478.302 | 532.431 | 543.575 | 560.191 | 560.191 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 0 | 0.000 |
| counter | 0 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 0 | 0.000 |
| relation | 0 | 0.000 |

- Cold compute：0（0.000 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 0 |
| l2_fresh | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：262969；input：20553752 bytes；output：4704923 bytes
- Hits/Misses：216/1582；run hit rate：12.01%
- Evicted/Rejected：0/0；ops/s max：17430；safety epoch：4659 -> 4835

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.098 |
| client:loadtest | cpu_percent_total | 1.564 |
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
| docker:zg-canal | cpu_percent | 6.960 |
| docker:zg-canal | memory_percent | 2.760 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.910 |
| docker:zg-es | memory_percent | 11.490 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.030 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 208.520 |
| docker:zg-kafka | memory_percent | 7.470 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 0.160 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4491.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4491.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 9048.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.769 |
| process:counter | cpu_seconds_total | 468.391 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 42299392.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 4.141 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 38637568.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 75.924 |
| process:knowpost | cpu_seconds_total | 446.719 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 72052736.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.549 |
| process:relation | cpu_seconds_total | 13.078 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 44617728.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.326 |
| process:search | cpu_seconds_total | 11.297 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 42897408.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 10.375 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 36515840.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 13154550.000 |
| redis | connected_clients | 60.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 4835.000 |
| redis | hit_rate | 0.712 |
| redis | keys | 589293.000 |
| redis | keyspace_hits | 194637.000 |
| redis | keyspace_misses | 80264.000 |
| redis | net_input_bytes | 932173572.000 |
| redis | net_output_bytes | 255539610.000 |
| redis | ops_per_sec | 17430.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43523.000 |
| redis | used_memory_bytes | 94738912.000 |

## 停止施压后的恢复

- Kafka drain：5.2560974s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：88
- 测量前恢复：complete=true；耗时=5.0134177s；删除帖子/Outbox=0/0；safety epoch=4609
- 预热后恢复：complete=true；耗时=6.3185868s；删除帖子/Outbox=24/48；safety epoch=4658
- 测量后恢复：complete=true；耗时=4.9427067s；删除帖子/Outbox=88/176；safety epoch=4836

## 说明

- SLA values are reference lines, not pass/fail gates.
