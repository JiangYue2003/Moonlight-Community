# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:47:35+08:00
- 采样时长：1m0.5532742s
- 并发：4
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1216 | 1216 | 0 | 0 | 20.08 | 1.655 | 2.689 | 3.171 | 4.331 | 5.416 |
| publish_total | 304 | 304 | 0 | 0 | 5.02 | 783.007 | 886.317 | 920.613 | 964.899 | 1007.228 |
| publish_draft | 304 | 304 | 0 | 0 | 5.02 | 3.648 | 4.255 | 4.928 | 155.807 | 164.771 |
| publish_metadata | 304 | 304 | 0 | 0 | 5.02 | 254.052 | 314.411 | 337.957 | 362.119 | 392.166 |
| publish_confirm | 304 | 304 | 0 | 0 | 5.02 | 253.921 | 312.275 | 337.881 | 393.137 | 438.583 |
| publish_commit | 304 | 304 | 0 | 0 | 5.02 | 253.317 | 322.242 | 335.936 | 376.075 | 387.551 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 204 | 0.168 |
| mysql | 160 | 0.132 |
| redis | 3808 | 3.132 |
| relation | 1216 | 1.000 |

- Cold compute：1216（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 48640 | 40.000 |
| merge_candidates | 83342 | 68.538 |
| redis_commands | 7296 | 6.000 |
| redis_members | 88052 | 72.411 |
| redis_roundtrips | 2432 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1216 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1216 | 0.217 |
| counter | 204 | 0.479 |
| hydrate | 1216 | 0.318 |
| inbox | 1216 | 0.199 |
| merge_dedup | 1216 | 0.010 |
| relation | 1216 | 0.871 |
| route | 1216 | 0.086 |
| total | 1216 | 1.725 |

## Redis 本轮边界增量

- Commands：997907；input：79301595 bytes；output：28906858 bytes
- Hits/Misses：60937/8021；run hit rate：88.37%
- Evicted/Rejected：0/0；ops/s max：18396；safety epoch：18235 -> 18539

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.211 |
| client:loadtest | cpu_percent_total | 3.380 |
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
| docker:zg-canal | cpu_percent | 3.300 |
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.070 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.550 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 165.920 |
| docker:zg-kafka | memory_percent | 7.630 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 38.270 |
| docker:zg-zk | memory_percent | 1.610 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 13333.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 13333.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 227478.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 7.743 |
| process:counter | cpu_seconds_total | 87.688 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44666880.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 12.094 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 40529920.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 79.834 |
| process:knowpost | cpu_seconds_total | 1225.828 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66265088.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 10.851 |
| process:relation | cpu_seconds_total | 36.484 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47652864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.328 |
| process:search | cpu_seconds_total | 19.312 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42283008.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 8.422 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 36343808.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 51597442.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 18539.000 |
| redis | hit_rate | 0.779 |
| redis | keys | 597038.000 |
| redis | keyspace_hits | 2563472.000 |
| redis | keyspace_misses | 727249.000 |
| redis | net_input_bytes | 3959098341.000 |
| redis | net_output_bytes | 1262784229.000 |
| redis | ops_per_sec | 18396.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48502.000 |
| redis | used_memory_bytes | 96681232.000 |

## 停止施压后的恢复

- Kafka drain：6.6604928s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：304
- 测量前恢复：complete=true；耗时=4.9401321s；删除帖子/Outbox=0/0；safety epoch=18177
- 预热后恢复：complete=true；耗时=7.5581667s；删除帖子/Outbox=56/112；safety epoch=18234
- 测量后恢复：complete=true；耗时=4.9357901s；删除帖子/Outbox=304/608；safety epoch=18540

## 说明

- SLA values are reference lines, not pass/fail gates.
