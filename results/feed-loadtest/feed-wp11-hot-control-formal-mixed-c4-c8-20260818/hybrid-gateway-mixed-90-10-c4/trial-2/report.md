# Feed 压测报告：hybrid / gateway / mixed-90-10-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:57:31+08:00
- 采样时长：1m0.4629224s
- 并发：4
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2484 | 2484 | 0 | 0 | 41.08 | 2.641 | 3.449 | 3.855 | 4.840 | 6.296 |
| publish_total | 276 | 276 | 0 | 0 | 4.56 | 837.699 | 986.873 | 1019.499 | 1081.027 | 1086.715 |
| publish_draft | 276 | 276 | 0 | 0 | 4.56 | 4.028 | 4.690 | 5.050 | 166.852 | 180.173 |
| publish_metadata | 276 | 276 | 0 | 0 | 4.56 | 273.620 | 337.611 | 351.872 | 379.849 | 385.715 |
| publish_confirm | 276 | 276 | 0 | 0 | 4.56 | 271.365 | 345.081 | 366.531 | 461.273 | 476.196 |
| publish_commit | 276 | 276 | 0 | 0 | 4.56 | 273.941 | 343.669 | 397.659 | 451.095 | 480.199 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 230 | 0.093 |
| mysql | 142 | 0.057 |
| redis | 7594 | 3.057 |
| relation | 2484 | 1.000 |

- Cold compute：2484（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 99360 | 40.000 |
| merge_candidates | 168422 | 67.803 |
| redis_commands | 14904 | 6.000 |
| redis_members | 173664 | 69.913 |
| redis_roundtrips | 4968 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2484 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2484 | 0.273 |
| counter | 230 | 0.561 |
| hydrate | 2484 | 0.351 |
| inbox | 2484 | 0.249 |
| merge_dedup | 2484 | 0.011 |
| relation | 2484 | 0.969 |
| route | 2484 | 0.057 |
| total | 2484 | 1.936 |

## Redis 本轮边界增量

- Commands：927593；input：75318401 bytes；output：39019620 bytes
- Hits/Misses：119071/9534；run hit rate：92.59%
- Evicted/Rejected：0/0；ops/s max：16951；safety epoch：20371 -> 20647

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.142 |
| client:loadtest | cpu_percent_total | 2.274 |
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
| docker:zg-canal | cpu_percent | 3.100 |
| docker:zg-canal | memory_percent | 5.000 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.880 |
| docker:zg-es | memory_percent | 12.250 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.100 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 142.720 |
| docker:zg-kafka | memory_percent | 7.700 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 42.660 |
| docker:zg-zk | memory_percent | 1.360 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 14985.000 |
| kafka | lag_max | 5.000 |
| kafka | lag_total | 5.000 |
| kafka | log_end_offset_total | 14985.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 273415.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.205 |
| process:counter | cpu_seconds_total | 105.906 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 43986944.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 10.076 |
| process:gateway | cpu_seconds_total | 17.859 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46346240.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 77.502 |
| process:knowpost | cpu_seconds_total | 1511.547 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66666496.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.453 |
| process:relation | cpu_seconds_total | 47.438 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47865856.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.335 |
| process:search | cpu_seconds_total | 23.469 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42151936.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 7.179 |
| process:user-storage | cpu_seconds_total | 11.375 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40292352.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 58672287.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 20647.000 |
| redis | hit_rate | 0.795 |
| redis | keys | 597677.000 |
| redis | keyspace_hits | 3191078.000 |
| redis | keyspace_misses | 820991.000 |
| redis | net_input_bytes | 4523145311.000 |
| redis | net_output_bytes | 1494435227.000 |
| redis | ops_per_sec | 16951.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49097.000 |
| redis | used_memory_bytes | 96599864.000 |

## 停止施压后的恢复

- Kafka drain：5.3511443s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：276
- 测量前恢复：complete=true；耗时=4.9004545s；删除帖子/Outbox=0/0；safety epoch=20313
- 预热后恢复：complete=true；耗时=6.2892272s；删除帖子/Outbox=56/112；safety epoch=20370
- 测量后恢复：complete=true；耗时=4.9713915s；删除帖子/Outbox=276/552；safety epoch=20648

## 说明

- SLA values are reference lines, not pass/fail gates.
