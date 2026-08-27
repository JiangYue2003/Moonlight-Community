# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-mixed-smoke-20260818`
- 开始时间：2026-08-18T15:41:46+08:00
- 采样时长：3.1572608s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 180 | 180 | 0 | 0 | 57.01 | 0.520 | 2.120 | 5.378 | 12.119 | 13.678 |
| publish_total | 16 | 16 | 0 | 0 | 5.07 | 692.968 | 981.216 | 984.957 | 984.957 | 984.957 |
| publish_draft | 16 | 16 | 0 | 0 | 5.07 | 4.273 | 4.833 | 5.260 | 5.260 | 5.260 |
| publish_metadata | 16 | 16 | 0 | 0 | 5.07 | 220.848 | 331.934 | 340.536 | 340.536 | 340.536 |
| publish_confirm | 16 | 16 | 0 | 0 | 5.07 | 223.565 | 388.755 | 394.585 | 394.585 | 394.585 |
| publish_commit | 16 | 16 | 0 | 0 | 5.07 | 246.933 | 273.530 | 280.110 | 280.110 | 280.110 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 42 | 0.233 |
| counter | 20 | 0.111 |
| mysql | 4 | 0.022 |
| redis | 160 | 0.889 |
| relation | 20 | 0.111 |

- Cold compute：39（0.217 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1560 | 8.667 |
| merge_candidates | 1599 | 8.883 |
| redis_commands | 234 | 1.300 |
| redis_members | 1599 | 8.883 |
| redis_roundtrips | 39 | 0.217 |

| page cache source | requests |
|---|---:|
| l1_fresh | 138 |
| miss | 42 |

- L1+L2 Fresh ratio：76.67%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 39 | 0.088 |
| counter | 20 | 0.886 |
| hydrate | 39 | 0.417 |
| inbox | 39 | 0.088 |
| merge_dedup | 39 | 0.013 |
| relation | 20 | 0.953 |
| route | 39 | 0.957 |
| total | 180 | 0.445 |

## Redis 本轮边界增量

- Commands：51999；input：4218869 bytes；output：1234034 bytes
- Hits/Misses：2022/707；run hit rate：74.09%
- Evicted/Rejected：0/0；ops/s max：17017；safety epoch：3865 -> 3897

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.155 |
| client:loadtest | cpu_percent_total | 2.474 |
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
| docker:zg-canal | cpu_percent | 2.590 |
| docker:zg-canal | memory_percent | 2.020 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 8.700 |
| docker:zg-es | memory_percent | 10.880 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 0.720 |
| docker:zg-etcd | memory_percent | 0.250 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 164.710 |
| docker:zg-kafka | memory_percent | 7.430 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 0.120 |
| docker:zg-zk | memory_percent | 0.980 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4153.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4153.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 1096.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.192 |
| process:counter | cpu_seconds_total | 35.781 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 39403520.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.484 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 37351424.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 60.649 |
| process:knowpost | cpu_seconds_total | 34.469 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 62439424.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.596 |
| process:relation | cpu_seconds_total | 1.156 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 42008576.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.798 |
| process:search | cpu_seconds_total | 1.078 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 38961152.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.969 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35545088.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1144722.000 |
| redis | connected_clients | 29.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3897.000 |
| redis | hit_rate | 0.696 |
| redis | keys | 595919.000 |
| redis | keyspace_hits | 43145.000 |
| redis | keyspace_misses | 18942.000 |
| redis | net_input_bytes | 80971926.000 |
| redis | net_output_bytes | 23057528.000 |
| redis | ops_per_sec | 17017.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19290.000 |
| redis | used_memory_bytes | 94423552.000 |

## 停止施压后的恢复

- Kafka drain：1.1044734s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：16
- 测量前恢复：complete=true；耗时=5.0516173s；删除帖子/Outbox=0/0；safety epoch=3847
- 预热后恢复：complete=true；耗时=6.304534s；删除帖子/Outbox=8/16；safety epoch=3864
- 测量后恢复：complete=true；耗时=4.905097s；删除帖子/Outbox=16/32；safety epoch=3898

## 说明

- SLA values are reference lines, not pass/fail gates.
