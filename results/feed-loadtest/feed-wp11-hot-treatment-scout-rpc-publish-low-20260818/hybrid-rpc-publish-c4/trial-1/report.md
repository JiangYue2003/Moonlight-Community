# Feed 压测报告：hybrid / rpc / publish-c4

- Run ID：`feed-wp11-hot-treatment-scout-rpc-publish-low-20260818`
- 开始时间：2026-08-18T22:24:34+08:00
- 采样时长：15.4943453s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 88 | 88 | 0 | 0 | 5.68 | 684.769 | 788.821 | 828.976 | 841.711 | 841.711 |
| publish_draft | 88 | 88 | 0 | 0 | 5.68 | 3.783 | 5.883 | 6.282 | 8.465 | 8.465 |
| publish_metadata | 88 | 88 | 0 | 0 | 5.68 | 225.525 | 277.611 | 286.311 | 295.911 | 295.911 |
| publish_confirm | 88 | 88 | 0 | 0 | 5.68 | 236.847 | 279.209 | 304.412 | 333.744 | 333.744 |
| publish_commit | 88 | 88 | 0 | 0 | 5.68 | 223.650 | 264.407 | 289.954 | 313.825 | 313.825 |

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

- Commands：262226；input：20481663 bytes；output：4684548 bytes
- Hits/Misses：215/1568；run hit rate：12.06%
- Evicted/Rejected：0/0；ops/s max：17881；safety epoch：4431 -> 4607

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.189 |
| client:loadtest | cpu_percent_total | 3.025 |
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
| docker:zg-canal | cpu_percent | 21.090 |
| docker:zg-canal | memory_percent | 2.700 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.210 |
| docker:zg-es | memory_percent | 11.490 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.150 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 283.710 |
| docker:zg-kafka | memory_percent | 7.490 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 39.000 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4407.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 4407.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 7232.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 9.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.874 |
| process:counter | cpu_seconds_total | 467.359 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 41480192.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.125 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 38637568.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 77.677 |
| process:knowpost | cpu_seconds_total | 432.016 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 73363456.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.877 |
| process:relation | cpu_seconds_total | 12.953 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 44888064.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.548 |
| process:search | cpu_seconds_total | 11.078 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 42201088.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 10.375 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 36515840.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12792760.000 |
| redis | connected_clients | 60.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 4607.000 |
| redis | hit_rate | 0.719 |
| redis | keys | 588850.000 |
| redis | keyspace_hits | 183984.000 |
| redis | keyspace_misses | 73259.000 |
| redis | net_input_bytes | 904271974.000 |
| redis | net_output_bytes | 248774025.000 |
| redis | ops_per_sec | 17881.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43476.000 |
| redis | used_memory_bytes | 94494256.000 |

## 停止施压后的恢复

- Kafka drain：6.6652062s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：88
- 测量前恢复：complete=true；耗时=4.9244901s；删除帖子/Outbox=0/0；safety epoch=4389
- 预热后恢复：complete=true；耗时=7.5422329s；删除帖子/Outbox=20/40；safety epoch=4430
- 测量后恢复：complete=true；耗时=5.0512802s；删除帖子/Outbox=88/176；safety epoch=4608

## 说明

- SLA values are reference lines, not pass/fail gates.
