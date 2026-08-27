# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-treatment-scout-rpc-mixed80-20260818`
- 开始时间：2026-08-18T22:33:59+08:00
- 采样时长：15.5421189s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 320 | 320 | 0 | 0 | 20.59 | 1.053 | 2.654 | 3.264 | 6.279 | 6.279 |
| publish_total | 80 | 80 | 0 | 0 | 5.15 | 758.843 | 870.221 | 894.713 | 933.641 | 933.641 |
| publish_draft | 80 | 80 | 0 | 0 | 5.15 | 3.735 | 5.434 | 5.806 | 6.399 | 6.399 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.15 | 241.806 | 299.470 | 311.206 | 332.714 | 332.714 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.15 | 241.565 | 312.816 | 314.878 | 434.879 | 434.879 |
| publish_commit | 80 | 80 | 0 | 0 | 5.15 | 246.816 | 326.241 | 338.176 | 399.113 | 399.113 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 234 | 0.731 |
| counter | 58 | 0.181 |
| mysql | 46 | 0.144 |
| redis | 962 | 3.006 |
| relation | 58 | 0.181 |

- Cold compute：229（0.716 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9160 | 28.625 |
| merge_candidates | 11362 | 35.506 |
| redis_commands | 1374 | 4.294 |
| redis_members | 11362 | 35.506 |
| redis_roundtrips | 229 | 0.716 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 86 |
| l2_fresh | 0 |
| miss | 234 |

- L1+L2 Fresh ratio：26.88%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 229 | 0.223 |
| counter | 58 | 0.511 |
| hydrate | 229 | 0.415 |
| inbox | 229 | 0.220 |
| merge_dedup | 229 | 0.008 |
| relation | 58 | 1.037 |
| route | 229 | 0.406 |
| total | 320 | 1.087 |

## Redis 本轮边界增量

- Commands：259863；input：21247515 bytes；output：6357942 bytes
- Hits/Misses：10869/3182；run hit rate：77.35%
- Evicted/Rejected：0/0；ops/s max：17513；safety epoch：6093 -> 6253

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.151 |
| client:loadtest | cpu_percent_total | 2.413 |
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
| docker:zg-canal | cpu_percent | 3.340 |
| docker:zg-canal | memory_percent | 3.400 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 16.980 |
| docker:zg-es | memory_percent | 11.650 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.760 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 230.800 |
| docker:zg-kafka | memory_percent | 7.840 |
| docker:zg-kafka | pids | 149.000 |
| docker:zg-zk | cpu_percent | 41.140 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5036.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 5036.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 21222.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 10.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 5.423 |
| process:counter | cpu_seconds_total | 480.328 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 43466752.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.234 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 37548032.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 65.321 |
| process:knowpost | cpu_seconds_total | 545.656 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74727424.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.323 |
| process:relation | cpu_seconds_total | 14.672 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 46174208.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.593 |
| process:search | cpu_seconds_total | 12.641 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 43003904.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 10.516 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35155968.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 15635348.000 |
| redis | connected_clients | 89.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 6253.000 |
| redis | hit_rate | 0.708 |
| redis | keys | 595339.000 |
| redis | keyspace_hits | 336467.000 |
| redis | keyspace_misses | 139069.000 |
| redis | net_input_bytes | 1128417927.000 |
| redis | net_output_bytes | 312647667.000 |
| redis | ops_per_sec | 17513.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44039.000 |
| redis | used_memory_bytes | 97929392.000 |

## 停止施压后的恢复

- Kafka drain：5.2056648s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.8901243s；删除帖子/Outbox=0/0；safety epoch=6051
- 预热后恢复：complete=true；耗时=6.2618477s；删除帖子/Outbox=20/40；safety epoch=6092
- 测量后恢复：complete=true；耗时=4.9052514s；删除帖子/Outbox=80/160；safety epoch=6254

## 说明

- SLA values are reference lines, not pass/fail gates.
