# Feed 压测报告：hybrid / gateway / publish-c16

- Run ID：`feed-wp11-hot-treatment-scout-gateway-publish-20260818`
- 开始时间：2026-08-18T22:41:47+08:00
- 采样时长：17.3362177s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 96 | 96 | 0 | 0 | 5.54 | 2849.361 | 3038.808 | 3085.794 | 3150.447 | 3150.447 |
| publish_draft | 96 | 96 | 0 | 0 | 5.54 | 5.391 | 8.932 | 8.932 | 12.593 | 12.593 |
| publish_metadata | 96 | 96 | 0 | 0 | 5.54 | 940.313 | 1011.053 | 1024.119 | 1124.282 | 1124.282 |
| publish_confirm | 96 | 96 | 0 | 0 | 5.54 | 946.832 | 1025.195 | 1040.561 | 1097.226 | 1097.226 |
| publish_commit | 96 | 96 | 0 | 0 | 5.54 | 960.206 | 1146.861 | 1161.438 | 1185.881 | 1185.881 |

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
| bypass | 0 |
| l1_fresh | 0 |
| l2_fresh | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：297139；input：23221356 bytes；output：5313298 bytes
- Hits/Misses：227/1740；run hit rate：11.54%
- Evicted/Rejected：0/0；ops/s max：18616；safety epoch：7391 -> 7583

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.056 |
| client:loadtest | cpu_percent_total | 0.901 |
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
| docker:zg-canal | cpu_percent | 2.090 |
| docker:zg-canal | memory_percent | 3.890 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 11.210 |
| docker:zg-es | memory_percent | 11.690 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.930 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 169.590 |
| docker:zg-kafka | memory_percent | 7.510 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 44.720 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5538.000 |
| kafka | lag_max | 14.000 |
| kafka | lag_total | 14.000 |
| kafka | log_end_offset_total | 5538.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 32107.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.039 |
| process:counter | cpu_seconds_total | 490.438 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 43433984.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.913 |
| process:gateway | cpu_seconds_total | 5.703 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 44470272.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 95.014 |
| process:knowpost | cpu_seconds_total | 639.516 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 76328960.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.913 |
| process:relation | cpu_seconds_total | 15.750 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 46891008.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.324 |
| process:search | cpu_seconds_total | 14.156 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 43126784.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.194 |
| process:user-storage | cpu_seconds_total | 11.281 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 40853504.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 17941126.000 |
| redis | connected_clients | 90.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 7583.000 |
| redis | hit_rate | 0.696 |
| redis | keys | 592359.000 |
| redis | keyspace_hits | 417382.000 |
| redis | keyspace_misses | 183726.000 |
| redis | net_input_bytes | 1306970880.000 |
| redis | net_output_bytes | 358256231.000 |
| redis | ops_per_sec | 18616.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44509.000 |
| redis | used_memory_bytes | 96633760.000 |

## 停止施压后的恢复

- Kafka drain：5.3022967s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：96
- 测量前恢复：complete=true；耗时=4.8587671s；删除帖子/Outbox=0/0；safety epoch=7325
- 预热后恢复：complete=true；耗时=6.2575578s；删除帖子/Outbox=32/64；safety epoch=7390
- 测量后恢复：complete=true；耗时=5.0018065s；删除帖子/Outbox=96/192；safety epoch=7584

## 说明

- SLA values are reference lines, not pass/fail gates.
