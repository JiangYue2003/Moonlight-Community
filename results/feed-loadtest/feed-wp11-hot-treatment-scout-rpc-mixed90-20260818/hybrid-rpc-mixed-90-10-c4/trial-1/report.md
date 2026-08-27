# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-treatment-scout-rpc-mixed90-20260818`
- 开始时间：2026-08-18T22:29:20+08:00
- 采样时长：15.2392426s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 756 | 756 | 0 | 0 | 49.61 | 0.524 | 2.089 | 2.597 | 4.222 | 8.621 |
| publish_total | 84 | 84 | 0 | 0 | 5.51 | 696.721 | 814.815 | 838.022 | 848.218 | 848.218 |
| publish_draft | 84 | 84 | 0 | 0 | 5.51 | 3.703 | 5.192 | 5.318 | 6.285 | 6.285 |
| publish_metadata | 84 | 84 | 0 | 0 | 5.51 | 232.720 | 310.266 | 314.097 | 318.268 | 318.268 |
| publish_confirm | 84 | 84 | 0 | 0 | 5.51 | 226.318 | 282.877 | 290.270 | 319.546 | 319.546 |
| publish_commit | 84 | 84 | 0 | 0 | 5.51 | 225.134 | 286.446 | 294.556 | 317.582 | 317.582 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 369 | 0.488 |
| counter | 60 | 0.079 |
| mysql | 48 | 0.063 |
| redis | 1460 | 1.931 |
| relation | 60 | 0.079 |

- Cold compute：353（0.467 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 14120 | 18.677 |
| merge_candidates | 17491 | 23.136 |
| redis_commands | 2118 | 2.802 |
| redis_members | 17491 | 23.136 |
| redis_roundtrips | 353 | 0.467 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 387 |
| l2_fresh | 0 |
| miss | 369 |

- L1+L2 Fresh ratio：51.19%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 353 | 0.202 |
| counter | 60 | 0.533 |
| hydrate | 353 | 0.295 |
| inbox | 353 | 0.200 |
| merge_dedup | 353 | 0.008 |
| relation | 60 | 0.960 |
| route | 353 | 0.261 |
| total | 756 | 0.585 |

## Redis 本轮边界增量

- Commands：262310；input：21894463 bytes；output：7395645 bytes
- Hits/Misses：16409/3648；run hit rate：81.81%
- Evicted/Rejected：0/0；ops/s max：18866；safety epoch：5139 -> 5307

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.147 |
| client:loadtest | cpu_percent_total | 2.358 |
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
| docker:zg-canal | cpu_percent | 5.230 |
| docker:zg-canal | memory_percent | 2.980 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.550 |
| docker:zg-es | memory_percent | 11.570 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.530 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 143.150 |
| docker:zg-kafka | memory_percent | 7.500 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 0.160 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4669.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 4669.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 12979.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 4.650 |
| process:counter | cpu_seconds_total | 473.375 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 42217472.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 4.188 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 37552128.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 83.807 |
| process:knowpost | cpu_seconds_total | 481.188 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74993664.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 13.500 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 44318720.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 11.766 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 41967616.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 10.469 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35053568.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 13979346.000 |
| redis | connected_clients | 68.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 5307.000 |
| redis | hit_rate | 0.708 |
| redis | keys | 591283.000 |
| redis | keyspace_hits | 236243.000 |
| redis | keyspace_misses | 97445.000 |
| redis | net_input_bytes | 997172535.000 |
| redis | net_output_bytes | 274307063.000 |
| redis | ops_per_sec | 18866.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43760.000 |
| redis | used_memory_bytes | 95948096.000 |

## 停止施压后的恢复

- Kafka drain：5.3329523s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：84
- 测量前恢复：complete=true；耗时=4.8638467s；删除帖子/Outbox=0/0；safety epoch=5097
- 预热后恢复：complete=true；耗时=6.383496s；删除帖子/Outbox=20/40；safety epoch=5138
- 测量后恢复：complete=true；耗时=4.899241s；删除帖子/Outbox=84/168；safety epoch=5308

## 说明

- SLA values are reference lines, not pass/fail gates.
