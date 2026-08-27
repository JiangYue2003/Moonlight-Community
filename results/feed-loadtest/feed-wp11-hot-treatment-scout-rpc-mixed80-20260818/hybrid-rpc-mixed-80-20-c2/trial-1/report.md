# Feed 压测报告：hybrid / rpc / mixed-80-20-c2

- Run ID：`feed-wp11-hot-treatment-scout-rpc-mixed80-20260818`
- 开始时间：2026-08-18T22:33:12+08:00
- 采样时长：15.0737125s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 296 | 296 | 0 | 0 | 19.64 | 1.043 | 2.157 | 2.668 | 4.789 | 5.795 |
| publish_total | 74 | 74 | 0 | 0 | 4.91 | 396.126 | 463.783 | 482.579 | 625.813 | 625.813 |
| publish_draft | 74 | 74 | 0 | 0 | 4.91 | 3.276 | 3.746 | 4.011 | 5.327 | 5.327 |
| publish_metadata | 74 | 74 | 0 | 0 | 4.91 | 129.067 | 160.360 | 172.575 | 190.825 | 190.825 |
| publish_confirm | 74 | 74 | 0 | 0 | 4.91 | 133.941 | 151.196 | 164.567 | 317.046 | 317.046 |
| publish_commit | 74 | 74 | 0 | 0 | 4.91 | 125.796 | 156.082 | 174.706 | 176.580 | 176.580 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 247 | 0.834 |
| counter | 56 | 0.189 |
| mysql | 58 | 0.196 |
| redis | 1046 | 3.534 |
| relation | 56 | 0.189 |

- Cold compute：247（0.834 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9880 | 33.378 |
| merge_candidates | 12237 | 41.341 |
| redis_commands | 1482 | 5.007 |
| redis_members | 12237 | 41.341 |
| redis_roundtrips | 247 | 0.834 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 49 |
| l2_fresh | 0 |
| miss | 247 |

- L1+L2 Fresh ratio：16.55%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 247 | 0.133 |
| counter | 56 | 0.367 |
| hydrate | 247 | 0.370 |
| inbox | 247 | 0.133 |
| merge_dedup | 247 | 0.013 |
| relation | 56 | 0.957 |
| route | 247 | 0.313 |
| total | 296 | 0.919 |

## Redis 本轮边界增量

- Commands：238916；input：19812351 bytes；output：5991386 bytes
- Hits/Misses：10774/3953；run hit rate：73.16%
- Evicted/Rejected：0/0；ops/s max：17186；safety epoch：5901 -> 6049

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.052 |
| client:loadtest | cpu_percent_total | 0.829 |
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
| docker:zg-canal | cpu_percent | 2.290 |
| docker:zg-canal | memory_percent | 3.320 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.870 |
| docker:zg-es | memory_percent | 11.640 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.030 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 156.440 |
| docker:zg-kafka | memory_percent | 7.730 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 3.020 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4958.000 |
| kafka | lag_max | 6.000 |
| kafka | lag_total | 6.000 |
| kafka | log_end_offset_total | 4958.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 19461.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.883 |
| process:counter | cpu_seconds_total | 479.047 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 44113920.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.234 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 37552128.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 76.713 |
| process:knowpost | cpu_seconds_total | 532.625 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74870784.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.875 |
| process:relation | cpu_seconds_total | 14.453 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 45412352.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 12.453 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 42631168.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 10.500 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35155968.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 15284015.000 |
| redis | connected_clients | 89.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 6049.000 |
| redis | hit_rate | 0.706 |
| redis | keys | 594156.000 |
| redis | keyspace_hits | 312494.000 |
| redis | keyspace_misses | 129972.000 |
| redis | net_input_bytes | 1100144554.000 |
| redis | net_output_bytes | 303925071.000 |
| redis | ops_per_sec | 17186.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43994.000 |
| redis | used_memory_bytes | 97253168.000 |

## 停止施压后的恢复

- Kafka drain：6.69003s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：74
- 测量前恢复：complete=true；耗时=4.9258288s；删除帖子/Outbox=0/0；safety epoch=5863
- 预热后恢复：complete=true；耗时=6.3248289s；删除帖子/Outbox=18/36；safety epoch=5900
- 测量后恢复：complete=true；耗时=4.9261869s；删除帖子/Outbox=74/148；safety epoch=6050

## 说明

- SLA values are reference lines, not pass/fail gates.
