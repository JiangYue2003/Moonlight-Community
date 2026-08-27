# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:37:36+08:00
- 采样时长：1m0.2186923s
- 并发：4
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2484 | 2484 | 0 | 0 | 41.25 | 2.094 | 2.726 | 3.196 | 3.835 | 12.488 |
| publish_total | 276 | 276 | 0 | 0 | 4.58 | 840.369 | 991.727 | 1032.673 | 1065.733 | 1078.668 |
| publish_draft | 276 | 276 | 0 | 0 | 4.58 | 3.718 | 8.031 | 13.073 | 163.768 | 173.942 |
| publish_metadata | 276 | 276 | 0 | 0 | 4.58 | 273.752 | 352.560 | 385.925 | 483.015 | 518.880 |
| publish_confirm | 276 | 276 | 0 | 0 | 4.58 | 278.754 | 352.103 | 378.090 | 459.055 | 474.108 |
| publish_commit | 276 | 276 | 0 | 0 | 4.58 | 271.253 | 331.208 | 344.804 | 380.843 | 387.439 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 226 | 0.091 |
| mysql | 145 | 0.058 |
| redis | 7597 | 3.058 |
| relation | 2484 | 1.000 |

- Cold compute：2484（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 99360 | 40.000 |
| merge_candidates | 168624 | 67.884 |
| redis_commands | 14904 | 6.000 |
| redis_members | 173917 | 70.015 |
| redis_roundtrips | 4968 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2484 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2484 | 0.232 |
| counter | 226 | 0.507 |
| hydrate | 2484 | 0.335 |
| inbox | 2484 | 0.231 |
| merge_dedup | 2484 | 0.010 |
| relation | 2484 | 0.922 |
| route | 2484 | 0.054 |
| total | 2484 | 1.806 |

## Redis 本轮边界增量

- Commands：912227；input：74129205 bytes；output：38730631 bytes
- Hits/Misses：118905/9484；run hit rate：92.61%
- Evicted/Rejected：0/0；ops/s max：17618；safety epoch：16145 -> 16421

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.190 |
| client:loadtest | cpu_percent_total | 3.036 |
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
| docker:zg-canal | cpu_percent | 1.440 |
| docker:zg-canal | memory_percent | 4.950 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.720 |
| docker:zg-es | memory_percent | 12.190 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.130 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 181.920 |
| docker:zg-kafka | memory_percent | 7.690 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 41.020 |
| docker:zg-zk | memory_percent | 1.610 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 11660.000 |
| kafka | lag_max | 4.000 |
| kafka | lag_total | 4.000 |
| kafka | log_end_offset_total | 11660.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 177714.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.970 |
| process:counter | cpu_seconds_total | 71.750 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46542848.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 12.094 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46178304.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 78.355 |
| process:knowpost | cpu_seconds_total | 940.562 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68689920.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 8.528 |
| process:relation | cpu_seconds_total | 23.578 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49725440.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.326 |
| process:search | cpu_seconds_total | 15.297 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43278336.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 8.203 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 38838272.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 44516084.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 16421.000 |
| redis | hit_rate | 0.739 |
| redis | keys | 595848.000 |
| redis | keyspace_hits | 1771145.000 |
| redis | keyspace_misses | 626497.000 |
| redis | net_input_bytes | 3390206455.000 |
| redis | net_output_bytes | 998658777.000 |
| redis | ops_per_sec | 17618.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47902.000 |
| redis | used_memory_bytes | 96384176.000 |

## 停止施压后的恢复

- Kafka drain：5.2951629s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：276
- 测量前恢复：complete=true；耗时=5.0317524s；删除帖子/Outbox=0/0；safety epoch=16095
- 预热后恢复：complete=true；耗时=7.7568242s；删除帖子/Outbox=48/96；safety epoch=16144
- 测量后恢复：complete=true；耗时=4.9721193s；删除帖子/Outbox=276/552；safety epoch=16422

## 说明

- SLA values are reference lines, not pass/fail gates.
