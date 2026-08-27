# Feed 压测报告：hybrid / rpc / mixed-90-10-c2

- Run ID：`feed-wp11-hot-mutation-script-smoke-20260818`
- 开始时间：2026-08-18T16:25:31+08:00
- 采样时长：2.3888762s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 108 | 108 | 0 | 0 | 45.21 | 0.529 | 2.091 | 2.243 | 5.918 | 5.918 |
| publish_total | 12 | 12 | 0 | 0 | 5.02 | 384.224 | 423.072 | 425.741 | 425.741 | 425.741 |
| publish_draft | 12 | 12 | 0 | 0 | 5.02 | 3.405 | 4.265 | 4.678 | 4.678 | 4.678 |
| publish_metadata | 12 | 12 | 0 | 0 | 5.02 | 129.076 | 160.592 | 161.676 | 161.676 | 161.676 |
| publish_confirm | 12 | 12 | 0 | 0 | 5.02 | 122.936 | 154.862 | 157.741 | 157.741 | 157.741 |
| publish_commit | 12 | 12 | 0 | 0 | 5.02 | 111.192 | 141.341 | 141.368 | 141.368 | 141.368 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 76 | 0.704 |
| counter | 20 | 0.185 |
| mysql | 8 | 0.074 |
| redis | 308 | 2.852 |
| relation | 20 | 0.185 |

- Cold compute：75（0.694 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3000 | 27.778 |
| merge_candidates | 3200 | 29.630 |
| redis_commands | 450 | 4.167 |
| redis_members | 3200 | 29.630 |
| redis_roundtrips | 75 | 0.694 |

| page cache source | requests |
|---|---:|
| l1_fresh | 32 |
| l2_fresh | 0 |
| miss | 76 |

- L1+L2 Fresh ratio：29.63%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 75 | 0.077 |
| counter | 20 | 0.529 |
| hydrate | 75 | 0.319 |
| inbox | 75 | 0.077 |
| merge_dedup | 75 | 0.000 |
| relation | 20 | 1.061 |
| route | 75 | 0.431 |
| total | 108 | 0.752 |

## Redis 本轮边界增量

- Commands：38946；input：3337220 bytes；output：1270109 bytes
- Hits/Misses：3456/892；run hit rate：79.48%
- Evicted/Rejected：0/0；ops/s max：14758；safety epoch：4060 -> 4084

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.245 |
| client:loadtest | cpu_percent_total | 3.924 |
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
| docker:zg-canal | cpu_percent | 0.110 |
| docker:zg-canal | memory_percent | 2.160 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.520 |
| docker:zg-es | memory_percent | 11.330 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 0.530 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 166.430 |
| docker:zg-kafka | memory_percent | 7.450 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 0.180 |
| docker:zg-zk | memory_percent | 0.980 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4212.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4212.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 2902.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 2.966 |
| process:counter | cpu_seconds_total | 82.688 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 38768640.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1.000 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 38555648.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 56.091 |
| process:knowpost | cpu_seconds_total | 87.406 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 66883584.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.580 |
| process:relation | cpu_seconds_total | 2.875 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 41738240.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.790 |
| process:search | cpu_seconds_total | 2.094 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 38473728.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.559 |
| process:user-storage | cpu_seconds_total | 2.297 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 36282368.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2674404.000 |
| redis | connected_clients | 29.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 4084.000 |
| redis | hit_rate | 0.706 |
| redis | keys | 588636.000 |
| redis | keyspace_hits | 110634.000 |
| redis | keyspace_misses | 46028.000 |
| redis | net_input_bytes | 190988524.000 |
| redis | net_output_bytes | 55999622.000 |
| redis | ops_per_sec | 14758.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21919.000 |
| redis | used_memory_bytes | 93825408.000 |

## 停止施压后的恢复

- Kafka drain：5.193472s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：12
- 测量前恢复：complete=true；耗时=4.9581345s；删除帖子/Outbox=0/0；safety epoch=4046
- 预热后恢复：complete=true；耗时=7.6072714s；删除帖子/Outbox=6/12；safety epoch=4059
- 测量后恢复：complete=true；耗时=5.0349004s；删除帖子/Outbox=12/24；safety epoch=4085

## 说明

- SLA values are reference lines, not pass/fail gates.
