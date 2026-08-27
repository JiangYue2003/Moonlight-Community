# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-mixed-smoke-v2-20260818`
- 开始时间：2026-08-18T15:45:10+08:00
- 采样时长：3.5946251s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 175 | 175 | 0 | 0 | 48.68 | 0.528 | 2.141 | 2.642 | 5.933 | 9.931 |
| publish_total | 19 | 19 | 0 | 0 | 5.29 | 751.247 | 768.253 | 772.504 | 772.504 | 772.504 |
| publish_draft | 19 | 19 | 0 | 0 | 5.29 | 3.723 | 4.778 | 4.889 | 4.889 | 4.889 |
| publish_metadata | 19 | 19 | 0 | 0 | 5.29 | 222.967 | 305.852 | 306.832 | 306.832 | 306.832 |
| publish_confirm | 19 | 19 | 0 | 0 | 5.29 | 249.776 | 273.513 | 280.810 | 280.810 | 280.810 |
| publish_commit | 19 | 19 | 0 | 0 | 5.29 | 241.433 | 263.396 | 264.578 | 264.578 | 264.578 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 83 | 0.474 |
| counter | 20 | 0.114 |
| mysql | 17 | 0.097 |
| redis | 334 | 1.909 |
| relation | 20 | 0.114 |

- Cold compute：79（0.451 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3160 | 18.057 |
| merge_candidates | 3421 | 19.549 |
| redis_commands | 474 | 2.709 |
| redis_members | 3421 | 19.549 |
| redis_roundtrips | 79 | 0.451 |

| page cache source | requests |
|---|---:|
| l1_fresh | 92 |
| l2_fresh | 1 |
| miss | 82 |

- L1+L2 Fresh ratio：53.14%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 79 | 0.161 |
| counter | 20 | 0.667 |
| hydrate | 79 | 0.375 |
| inbox | 79 | 0.161 |
| merge_dedup | 79 | 0.026 |
| relation | 20 | 1.053 |
| route | 79 | 0.455 |
| total | 175 | 0.645 |

## Redis 本轮边界增量

- Commands：62594；input：5239738 bytes；output：1718298 bytes
- Hits/Misses：3548/1136；run hit rate：75.75%
- Evicted/Rejected：0/0；ops/s max：16441；safety epoch：3917 -> 3955

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.109 |
| client:loadtest | cpu_percent_total | 1.739 |
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
| docker:zg-canal | cpu_percent | 1.850 |
| docker:zg-canal | memory_percent | 2.050 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 9.530 |
| docker:zg-es | memory_percent | 10.970 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.910 |
| docker:zg-etcd | memory_percent | 0.250 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 134.940 |
| docker:zg-kafka | memory_percent | 7.440 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 0.150 |
| docker:zg-zk | memory_percent | 0.970 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4170.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4170.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 1639.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.136 |
| process:counter | cpu_seconds_total | 39.625 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 38928384.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.531 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 36724736.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 79.507 |
| process:knowpost | cpu_seconds_total | 42.000 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 63410176.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.352 |
| process:relation | cpu_seconds_total | 1.422 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 42274816.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.125 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 39092224.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.062 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35540992.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1345130.000 |
| redis | connected_clients | 29.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3955.000 |
| redis | hit_rate | 0.699 |
| redis | keys | 596254.000 |
| redis | keyspace_hits | 58760.000 |
| redis | keyspace_misses | 25340.000 |
| redis | net_input_bytes | 96017772.000 |
| redis | net_output_bytes | 27858855.000 |
| redis | ops_per_sec | 16441.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19496.000 |
| redis | used_memory_bytes | 94564696.000 |

## 停止施压后的恢复

- Kafka drain：2.4579764s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：19
- 测量前恢复：complete=true；耗时=4.9836288s；删除帖子/Outbox=0/0；safety epoch=3899
- 预热后恢复：complete=true；耗时=6.2718595s；删除帖子/Outbox=8/16；safety epoch=3916
- 测量后恢复：complete=true；耗时=4.9519642s；删除帖子/Outbox=19/38；safety epoch=3956

## 说明

- SLA values are reference lines, not pass/fail gates.
