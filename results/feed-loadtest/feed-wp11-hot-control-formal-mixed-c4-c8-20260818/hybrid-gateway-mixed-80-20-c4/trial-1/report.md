# Feed 压测报告：hybrid / gateway / mixed-80-20-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-19T00:05:51+08:00
- 采样时长：1m0.4536431s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1072 | 1072 | 0 | 0 | 17.73 | 2.668 | 3.783 | 4.288 | 5.277 | 6.881 |
| publish_total | 268 | 268 | 0 | 0 | 4.43 | 870.072 | 1069.248 | 1103.999 | 1284.871 | 1286.430 |
| publish_draft | 268 | 268 | 0 | 0 | 4.43 | 4.198 | 5.342 | 14.043 | 124.830 | 134.928 |
| publish_metadata | 268 | 268 | 0 | 0 | 4.43 | 282.463 | 368.126 | 427.791 | 475.055 | 491.379 |
| publish_confirm | 268 | 268 | 0 | 0 | 4.43 | 282.048 | 375.834 | 400.456 | 425.368 | 697.489 |
| publish_commit | 268 | 268 | 0 | 0 | 4.43 | 283.753 | 382.486 | 407.769 | 583.109 | 663.341 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 200 | 0.187 |
| mysql | 124 | 0.116 |
| redis | 3340 | 3.116 |
| relation | 1072 | 1.000 |

- Cold compute：1072（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 42880 | 40.000 |
| merge_candidates | 71494 | 66.692 |
| redis_commands | 6432 | 6.000 |
| redis_members | 73649 | 68.702 |
| redis_roundtrips | 2144 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1072 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1072 | 0.241 |
| counter | 200 | 0.548 |
| hydrate | 1072 | 0.389 |
| inbox | 1072 | 0.251 |
| merge_dedup | 1072 | 0.011 |
| relation | 1072 | 1.002 |
| route | 1072 | 0.112 |
| total | 1072 | 2.029 |

## Redis 本轮边界增量

- Commands：884115；input：70204227 bytes；output：25424279 bytes
- Hits/Misses：54229/7070；run hit rate：88.47%
- Evicted/Rejected：0/0；ops/s max：17878；safety epoch：21964 -> 22232

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.134 |
| client:loadtest | cpu_percent_total | 2.145 |
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
| docker:zg-canal | cpu_percent | 3.180 |
| docker:zg-canal | memory_percent | 5.010 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.540 |
| docker:zg-es | memory_percent | 12.270 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.340 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 20.000 |
| docker:zg-kafka | cpu_percent | 142.250 |
| docker:zg-kafka | memory_percent | 7.800 |
| docker:zg-kafka | pids | 128.000 |
| docker:zg-zk | cpu_percent | 44.970 |
| docker:zg-zk | memory_percent | 1.380 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 16237.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 16237.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 311811.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 7.740 |
| process:counter | cpu_seconds_total | 122.203 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 43991040.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 9.296 |
| process:gateway | cpu_seconds_total | 31.828 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47726592.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 79.726 |
| process:knowpost | cpu_seconds_total | 1753.172 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 67383296.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.978 |
| process:relation | cpu_seconds_total | 59.484 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47624192.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.873 |
| process:search | cpu_seconds_total | 27.047 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42708992.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 5.422 |
| process:user-storage | cpu_seconds_total | 19.172 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40550400.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 64052315.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 22232.000 |
| redis | hit_rate | 0.810 |
| redis | keys | 596983.000 |
| redis | keyspace_hits | 3840590.000 |
| redis | keyspace_misses | 901766.000 |
| redis | net_input_bytes | 4956035738.000 |
| redis | net_output_bytes | 1702247840.000 |
| redis | ops_per_sec | 17878.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49597.000 |
| redis | used_memory_bytes | 96512032.000 |

## 停止施压后的恢复

- Kafka drain：5.3436602s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：268
- 测量前恢复：complete=true；耗时=4.9001422s；删除帖子/Outbox=0/0；safety epoch=21910
- 预热后恢复：complete=true；耗时=6.3212949s；删除帖子/Outbox=52/104；safety epoch=21963
- 测量后恢复：complete=true；耗时=4.9372056s；删除帖子/Outbox=268/536；safety epoch=22233

## 说明

- SLA values are reference lines, not pass/fail gates.
