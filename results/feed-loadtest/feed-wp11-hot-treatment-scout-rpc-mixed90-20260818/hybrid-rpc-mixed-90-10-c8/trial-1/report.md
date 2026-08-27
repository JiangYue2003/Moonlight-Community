# Feed 压测报告：hybrid / rpc / mixed-90-10-c8

- Run ID：`feed-wp11-hot-treatment-scout-rpc-mixed90-20260818`
- 开始时间：2026-08-18T22:30:05+08:00
- 采样时长：15.5933193s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 792 | 792 | 0 | 0 | 50.79 | 0.519 | 2.918 | 3.682 | 6.017 | 8.117 |
| publish_total | 88 | 88 | 0 | 0 | 5.64 | 1421.232 | 1467.229 | 1471.844 | 1483.301 | 1483.301 |
| publish_draft | 88 | 88 | 0 | 0 | 5.64 | 4.212 | 6.289 | 6.359 | 7.066 | 7.066 |
| publish_metadata | 88 | 88 | 0 | 0 | 5.64 | 472.243 | 494.697 | 496.219 | 502.957 | 502.957 |
| publish_confirm | 88 | 88 | 0 | 0 | 5.64 | 479.388 | 542.486 | 549.970 | 558.905 | 558.905 |
| publish_commit | 88 | 88 | 0 | 0 | 5.64 | 471.756 | 549.316 | 555.984 | 565.975 | 565.975 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 242 | 0.306 |
| counter | 60 | 0.076 |
| mysql | 34 | 0.043 |
| redis | 878 | 1.109 |
| relation | 60 | 0.076 |

- Cold compute：211（0.266 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 8440 | 10.657 |
| merge_candidates | 10329 | 13.042 |
| redis_commands | 1266 | 1.598 |
| redis_members | 10329 | 13.042 |
| redis_roundtrips | 211 | 0.266 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 550 |
| l2_fresh | 0 |
| miss | 242 |

- L1+L2 Fresh ratio：69.44%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 211 | 0.376 |
| counter | 60 | 0.611 |
| hydrate | 211 | 0.623 |
| inbox | 211 | 0.375 |
| merge_dedup | 211 | 0.007 |
| relation | 60 | 1.152 |
| route | 211 | 0.509 |
| total | 792 | 0.766 |

## Redis 本轮边界增量

- Commands：276456；input：22411003 bytes；output：6541188 bytes
- Hits/Misses：10484/3012；run hit rate：77.68%
- Evicted/Rejected：0/0；ops/s max：19717；safety epoch：5359 -> 5535

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.200 |
| client:loadtest | cpu_percent_total | 3.207 |
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
| docker:zg-canal | cpu_percent | 1.930 |
| docker:zg-canal | memory_percent | 3.130 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 74.160 |
| docker:zg-es | memory_percent | 11.620 |
| docker:zg-es | pids | 154.000 |
| docker:zg-etcd | cpu_percent | 4.070 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 166.820 |
| docker:zg-kafka | memory_percent | 7.380 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 41.770 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4756.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 4756.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 14921.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.968 |
| process:counter | cpu_seconds_total | 474.562 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 42094592.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.203 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 37552128.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 86.744 |
| process:knowpost | cpu_seconds_total | 496.453 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74969088.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.099 |
| process:relation | cpu_seconds_total | 13.750 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 44109824.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.591 |
| process:search | cpu_seconds_total | 11.906 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 41943040.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 10.469 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35053568.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 14357124.000 |
| redis | connected_clients | 73.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 5535.000 |
| redis | hit_rate | 0.710 |
| redis | keys | 592359.000 |
| redis | keyspace_hits | 259801.000 |
| redis | keyspace_misses | 106334.000 |
| redis | net_input_bytes | 1027350689.000 |
| redis | net_output_bytes | 283376543.000 |
| redis | ops_per_sec | 19717.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43806.000 |
| redis | used_memory_bytes | 96789192.000 |

## 停止施压后的恢复

- Kafka drain：5.3536478s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：88
- 测量前恢复：complete=true；耗时=4.9153924s；删除帖子/Outbox=0/0；safety epoch=5309
- 预热后恢复：complete=true；耗时=6.3504765s；删除帖子/Outbox=24/48；safety epoch=5358
- 测量后恢复：complete=true；耗时=4.8932466s；删除帖子/Outbox=88/176；safety epoch=5536

## 说明

- SLA values are reference lines, not pass/fail gates.
