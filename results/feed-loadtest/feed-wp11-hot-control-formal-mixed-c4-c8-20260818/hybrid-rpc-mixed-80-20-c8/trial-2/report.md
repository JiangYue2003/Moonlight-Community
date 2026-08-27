# Feed 压测报告：hybrid / rpc / mixed-80-20-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:52:33+08:00
- 采样时长：1m0.9429994s
- 并发：8
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1200 | 1200 | 0 | 0 | 19.69 | 3.215 | 4.805 | 5.384 | 6.895 | 7.530 |
| publish_total | 300 | 300 | 0 | 0 | 4.92 | 1606.121 | 1709.048 | 1772.457 | 1972.885 | 1994.160 |
| publish_draft | 300 | 300 | 0 | 0 | 4.92 | 3.854 | 5.139 | 36.923 | 87.321 | 135.044 |
| publish_metadata | 300 | 300 | 0 | 0 | 4.92 | 525.482 | 607.474 | 653.155 | 834.461 | 857.979 |
| publish_confirm | 300 | 300 | 0 | 0 | 4.92 | 519.005 | 603.476 | 617.727 | 663.825 | 678.098 |
| publish_commit | 300 | 300 | 0 | 0 | 4.92 | 522.812 | 616.671 | 646.773 | 692.199 | 712.361 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 196 | 0.163 |
| mysql | 130 | 0.108 |
| redis | 3730 | 3.108 |
| relation | 1200 | 1.000 |

- Cold compute：1200（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 48000 | 40.000 |
| merge_candidates | 81282 | 67.735 |
| redis_commands | 7200 | 6.000 |
| redis_members | 85320 | 71.100 |
| redis_roundtrips | 2400 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1200 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1200 | 0.531 |
| counter | 196 | 0.883 |
| hydrate | 1200 | 0.710 |
| inbox | 1200 | 0.543 |
| merge_dedup | 1200 | 0.014 |
| relation | 1200 | 1.240 |
| route | 1200 | 0.150 |
| total | 1200 | 3.210 |

## Redis 本轮边界增量

- Commands：988254；input：78539845 bytes；output：28546800 bytes
- Hits/Misses：59968/7939；run hit rate：88.31%
- Evicted/Rejected：0/0；ops/s max：17606；safety epoch：19323 -> 19623

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.143 |
| client:loadtest | cpu_percent_total | 2.282 |
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
| docker:zg-canal | cpu_percent | 1.870 |
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.150 |
| docker:zg-es | memory_percent | 12.250 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.200 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 149.990 |
| docker:zg-kafka | memory_percent | 7.700 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 42.210 |
| docker:zg-zk | memory_percent | 1.360 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 14175.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 14175.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 249375.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 17.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 7.704 |
| process:counter | cpu_seconds_total | 96.469 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44212224.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 12.172 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 40448000.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 87.056 |
| process:knowpost | cpu_seconds_total | 1369.828 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 65998848.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 8.525 |
| process:relation | cpu_seconds_total | 41.031 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47853568.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.116 |
| process:search | cpu_seconds_total | 21.453 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42536960.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.549 |
| process:user-storage | cpu_seconds_total | 8.562 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35758080.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 55209135.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 19623.000 |
| redis | hit_rate | 0.784 |
| redis | keys | 597497.000 |
| redis | keyspace_hits | 2807447.000 |
| redis | keyspace_misses | 771909.000 |
| redis | net_input_bytes | 4245117456.000 |
| redis | net_output_bytes | 1366392827.000 |
| redis | ops_per_sec | 17606.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48799.000 |
| redis | used_memory_bytes | 96814144.000 |

## 停止施压后的恢复

- Kafka drain：5.2744417s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：300
- 测量前恢复：complete=true；耗时=4.9057919s；删除帖子/Outbox=0/0；safety epoch=19265
- 预热后恢复：complete=true；耗时=7.6196937s；删除帖子/Outbox=56/112；safety epoch=19322
- 测量后恢复：complete=true；耗时=5.0748047s；删除帖子/Outbox=300/600；safety epoch=19624

## 说明

- SLA values are reference lines, not pass/fail gates.
