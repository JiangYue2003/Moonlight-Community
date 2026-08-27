# Feed 压测报告：hybrid / gateway / mixed-90-10-c8

- Run ID：`feed-wp11-hot-treatment-scout-gateway-mixed90-20260818`
- 开始时间：2026-08-18T22:43:55+08:00
- 采样时长：16.2739772s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 792 | 792 | 0 | 0 | 48.67 | 0.556 | 3.738 | 5.253 | 6.562 | 7.967 |
| publish_total | 88 | 88 | 0 | 0 | 5.41 | 1459.018 | 1537.121 | 1579.360 | 1603.106 | 1603.106 |
| publish_draft | 88 | 88 | 0 | 0 | 5.41 | 4.153 | 4.767 | 5.216 | 6.735 | 6.735 |
| publish_metadata | 88 | 88 | 0 | 0 | 5.41 | 467.463 | 572.369 | 644.747 | 653.253 | 653.253 |
| publish_confirm | 88 | 88 | 0 | 0 | 5.41 | 474.915 | 584.244 | 595.495 | 613.635 | 613.635 |
| publish_commit | 88 | 88 | 0 | 0 | 5.41 | 488.268 | 553.034 | 576.724 | 594.402 | 594.402 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 235 | 0.297 |
| counter | 60 | 0.076 |
| mysql | 42 | 0.053 |
| redis | 886 | 1.119 |
| relation | 60 | 0.076 |

- Cold compute：211（0.266 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 8440 | 10.657 |
| merge_candidates | 10344 | 13.061 |
| redis_commands | 1266 | 1.598 |
| redis_members | 10344 | 13.061 |
| redis_roundtrips | 211 | 0.266 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 557 |
| l2_fresh | 0 |
| miss | 235 |

- L1+L2 Fresh ratio：70.33%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 211 | 0.414 |
| counter | 60 | 0.755 |
| hydrate | 211 | 0.693 |
| inbox | 211 | 0.408 |
| merge_dedup | 211 | 0.013 |
| relation | 60 | 1.290 |
| route | 211 | 0.591 |
| total | 792 | 0.854 |

## Redis 本轮边界增量

- Commands：282529；input：22888203 bytes；output：6653366 bytes
- Hits/Misses：10484/3010；run hit rate：77.69%
- Evicted/Rejected：0/0；ops/s max：18780；safety epoch：7845 -> 8021

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.090 |
| client:loadtest | cpu_percent_total | 1.440 |
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
| docker:zg-canal | cpu_percent | 1.470 |
| docker:zg-canal | memory_percent | 4.050 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.120 |
| docker:zg-es | memory_percent | 11.750 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.240 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 165.720 |
| docker:zg-kafka | memory_percent | 7.520 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 40.430 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5703.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 5703.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 35901.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 8.000 |
| process:counter | cpu_percent | 5.426 |
| process:counter | cpu_seconds_total | 493.562 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 44068864.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 9.287 |
| process:gateway | cpu_seconds_total | 7.422 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 45916160.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 98.439 |
| process:knowpost | cpu_seconds_total | 671.406 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74354688.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.101 |
| process:relation | cpu_seconds_total | 16.125 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 47362048.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.799 |
| process:search | cpu_seconds_total | 14.516 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 43421696.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 6.192 |
| process:user-storage | cpu_seconds_total | 12.109 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 40538112.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18698124.000 |
| redis | connected_clients | 90.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 8021.000 |
| redis | hit_rate | 0.699 |
| redis | keys | 594382.000 |
| redis | keyspace_hits | 471237.000 |
| redis | keyspace_misses | 202484.000 |
| redis | net_input_bytes | 1368059401.000 |
| redis | net_output_bytes | 377721248.000 |
| redis | ops_per_sec | 18780.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44636.000 |
| redis | used_memory_bytes | 97741512.000 |

## 停止施压后的恢复

- Kafka drain：5.277477s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：88
- 测量前恢复：complete=true；耗时=4.9375716s；删除帖子/Outbox=0/0；safety epoch=7795
- 预热后恢复：complete=true；耗时=6.2658919s；删除帖子/Outbox=24/48；safety epoch=7844
- 测量后恢复：complete=true；耗时=4.9312598s；删除帖子/Outbox=88/176；safety epoch=8022

## 说明

- SLA values are reference lines, not pass/fail gates.
