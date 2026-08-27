# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:35:57+08:00
- 采样时长：1m0.2676177s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2484 | 2484 | 0 | 0 | 41.22 | 2.096 | 2.704 | 3.161 | 3.847 | 15.193 |
| publish_total | 276 | 276 | 0 | 0 | 4.58 | 849.760 | 974.072 | 1009.855 | 1116.169 | 1130.105 |
| publish_draft | 276 | 276 | 0 | 0 | 4.58 | 3.706 | 4.755 | 10.221 | 18.953 | 23.598 |
| publish_metadata | 276 | 276 | 0 | 0 | 4.58 | 279.675 | 336.909 | 368.865 | 426.585 | 465.952 |
| publish_confirm | 276 | 276 | 0 | 0 | 4.58 | 283.734 | 346.873 | 368.403 | 438.524 | 454.362 |
| publish_commit | 276 | 276 | 0 | 0 | 4.58 | 270.069 | 343.459 | 355.234 | 374.935 | 379.244 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 221 | 0.089 |
| mysql | 131 | 0.053 |
| redis | 7583 | 3.053 |
| relation | 2484 | 1.000 |

- Cold compute：2484（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 99360 | 40.000 |
| merge_candidates | 168574 | 67.864 |
| redis_commands | 14904 | 6.000 |
| redis_members | 173880 | 70.000 |
| redis_roundtrips | 4968 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2484 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2484 | 0.235 |
| counter | 221 | 0.537 |
| hydrate | 2484 | 0.316 |
| inbox | 2484 | 0.224 |
| merge_dedup | 2484 | 0.013 |
| relation | 2484 | 0.931 |
| route | 2484 | 0.055 |
| total | 2484 | 1.799 |

## Redis 本轮边界增量

- Commands：910123；input：73969667 bytes；output：38708763 bytes
- Hits/Misses：118841/9096；run hit rate：92.89%
- Evicted/Rejected：0/0；ops/s max：17040；safety epoch：15817 -> 16093

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.269 |
| client:loadtest | cpu_percent_total | 4.304 |
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
| docker:zg-canal | cpu_percent | 3.020 |
| docker:zg-canal | memory_percent | 4.950 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 10.250 |
| docker:zg-es | memory_percent | 12.190 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.220 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 241.730 |
| docker:zg-kafka | memory_percent | 7.650 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 8.950 |
| docker:zg-zk | memory_percent | 1.480 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 11398.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 11398.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 169449.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.969 |
| process:counter | cpu_seconds_total | 69.375 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46264320.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 12.078 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47321088.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 78.263 |
| process:knowpost | cpu_seconds_total | 895.047 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68644864.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.745 |
| process:relation | cpu_seconds_total | 21.094 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49532928.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.320 |
| process:search | cpu_seconds_total | 14.594 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43454464.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.324 |
| process:user-storage | cpu_seconds_total | 8.172 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40968192.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 43417026.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 16093.000 |
| redis | hit_rate | 0.727 |
| redis | keys | 595608.000 |
| redis | keyspace_hits | 1621986.000 |
| redis | keyspace_misses | 609592.000 |
| redis | net_input_bytes | 3301300729.000 |
| redis | net_output_bytes | 952803070.000 |
| redis | ops_per_sec | 17040.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47803.000 |
| redis | used_memory_bytes | 96297896.000 |

## 停止施压后的恢复

- Kafka drain：5.4322103s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：276
- 测量前恢复：complete=true；耗时=4.9708881s；删除帖子/Outbox=0/0；safety epoch=15763
- 预热后恢复：complete=true；耗时=6.2826698s；删除帖子/Outbox=52/104；safety epoch=15816
- 测量后恢复：complete=true；耗时=5.1005451s；删除帖子/Outbox=276/552；safety epoch=16094

## 说明

- SLA values are reference lines, not pass/fail gates.
