# Feed 压测报告：hybrid / gateway / mixed-80-20-c4

- Run ID：`feed-wp11-hot-treatment-scout-gateway-mixed80-20260818`
- 开始时间：2026-08-18T22:47:35+08:00
- 采样时长：15.6471685s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 336 | 336 | 0 | 0 | 21.47 | 1.567 | 3.114 | 3.756 | 6.302 | 6.822 |
| publish_total | 84 | 84 | 0 | 0 | 5.37 | 729.867 | 788.113 | 857.098 | 928.922 | 928.922 |
| publish_draft | 84 | 84 | 0 | 0 | 5.37 | 3.946 | 5.296 | 5.738 | 6.414 | 6.414 |
| publish_metadata | 84 | 84 | 0 | 0 | 5.37 | 241.168 | 289.706 | 295.157 | 405.150 | 405.150 |
| publish_confirm | 84 | 84 | 0 | 0 | 5.37 | 230.270 | 284.058 | 287.760 | 295.779 | 295.779 |
| publish_commit | 84 | 84 | 0 | 0 | 5.37 | 235.550 | 280.626 | 282.648 | 287.038 | 287.038 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 244 | 0.726 |
| counter | 59 | 0.176 |
| mysql | 46 | 0.137 |
| redis | 1006 | 2.994 |
| relation | 59 | 0.176 |

- Cold compute：240（0.714 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9600 | 28.571 |
| merge_candidates | 12028 | 35.798 |
| redis_commands | 1440 | 4.286 |
| redis_members | 12028 | 35.798 |
| redis_roundtrips | 240 | 0.714 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 92 |
| l2_fresh | 0 |
| miss | 244 |

- L1+L2 Fresh ratio：27.38%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 240 | 0.162 |
| counter | 59 | 0.516 |
| hydrate | 240 | 0.387 |
| inbox | 240 | 0.159 |
| merge_dedup | 240 | 0.007 |
| relation | 59 | 0.963 |
| route | 240 | 0.370 |
| total | 336 | 0.950 |

## Redis 本轮边界增量

- Commands：275563；input：22547891 bytes；output：6693479 bytes
- Hits/Misses：11199/3504；run hit rate：76.17%
- Evicted/Rejected：0/0；ops/s max：18800；safety epoch：8535 -> 8703

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.156 |
| client:loadtest | cpu_percent_total | 2.496 |
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
| docker:zg-canal | cpu_percent | 2.970 |
| docker:zg-canal | memory_percent | 4.300 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 6.760 |
| docker:zg-es | memory_percent | 11.770 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.120 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 261.730 |
| docker:zg-kafka | memory_percent | 7.540 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 21.520 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5963.000 |
| kafka | lag_max | 5.000 |
| kafka | lag_total | 5.000 |
| kafka | log_end_offset_total | 5963.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 41841.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 10.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.852 |
| process:counter | cpu_seconds_total | 498.344 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 44380160.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 5.424 |
| process:gateway | cpu_seconds_total | 9.594 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 46628864.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 75.152 |
| process:knowpost | cpu_seconds_total | 720.250 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74391552.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.325 |
| process:relation | cpu_seconds_total | 16.875 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 47464448.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.582 |
| process:search | cpu_seconds_total | 15.188 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 42532864.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.100 |
| process:user-storage | cpu_seconds_total | 12.984 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 40280064.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19910365.000 |
| redis | connected_clients | 95.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 8703.000 |
| redis | hit_rate | 0.699 |
| redis | keys | 596241.000 |
| redis | keyspace_hits | 545899.000 |
| redis | keyspace_misses | 235231.000 |
| redis | net_input_bytes | 1464329967.000 |
| redis | net_output_bytes | 405835759.000 |
| redis | ops_per_sec | 18800.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44857.000 |
| redis | used_memory_bytes | 98433728.000 |

## 停止施压后的恢复

- Kafka drain：6.5564185s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：84
- 测量前恢复：complete=true；耗时=4.9205407s；删除帖子/Outbox=0/0；safety epoch=8493
- 预热后恢复：complete=true；耗时=6.2668774s；删除帖子/Outbox=20/40；safety epoch=8534
- 测量后恢复：complete=true；耗时=4.9122948s；删除帖子/Outbox=84/168；safety epoch=8704

## 说明

- SLA values are reference lines, not pass/fail gates.
