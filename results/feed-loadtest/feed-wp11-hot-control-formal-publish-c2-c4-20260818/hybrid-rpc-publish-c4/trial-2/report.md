# Feed 压测报告：hybrid / rpc / publish-c4

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:21:44+08:00
- 采样时长：1m1.7177137s
- 并发：4
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 232 | 232 | 0 | 0 | 3.76 | 893.119 | 1728.105 | 1795.656 | 1945.970 | 1962.320 |
| publish_draft | 232 | 232 | 0 | 0 | 3.76 | 4.706 | 10.566 | 12.222 | 25.804 | 26.325 |
| publish_metadata | 232 | 232 | 0 | 0 | 3.76 | 299.408 | 554.465 | 606.912 | 631.848 | 645.835 |
| publish_confirm | 232 | 232 | 0 | 0 | 3.76 | 313.601 | 594.626 | 634.088 | 687.973 | 694.729 |
| publish_commit | 232 | 232 | 0 | 0 | 3.76 | 295.686 | 589.381 | 649.476 | 798.439 | 819.268 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 0 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 0 | 0.000 |
| relation | 0 | 0.000 |

- Cold compute：0（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：733635；input：57277495 bytes；output：13143926 bytes
- Hits/Misses：568/4641；run hit rate：10.90%
- Evicted/Rejected：0/0；ops/s max：16430；safety epoch：13116 -> 13348

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.176 |
| client:loadtest | cpu_percent_total | 2.810 |
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
| docker:zg-canal | cpu_percent | 2.020 |
| docker:zg-canal | memory_percent | 4.860 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.990 |
| docker:zg-es | memory_percent | 12.150 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 9.350 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 245.520 |
| docker:zg-kafka | memory_percent | 7.580 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 43.250 |
| docker:zg-zk | memory_percent | 1.480 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 9293.000 |
| kafka | lag_max | 4.000 |
| kafka | lag_total | 4.000 |
| kafka | log_end_offset_total | 9293.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 123134.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 10.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 7.048 |
| process:counter | cpu_seconds_total | 47.062 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45576192.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 6.375 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 43036672.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 81.381 |
| process:knowpost | cpu_seconds_total | 553.922 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 69918720.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.871 |
| process:relation | cpu_seconds_total | 14.031 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48222208.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.132 |
| process:search | cpu_seconds_total | 8.844 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43241472.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.550 |
| process:user-storage | cpu_seconds_total | 4.125 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 37695488.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 34515552.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 13348.000 |
| redis | hit_rate | 0.733 |
| redis | keys | 593907.000 |
| redis | keyspace_hits | 1389525.000 |
| redis | keyspace_misses | 509640.000 |
| redis | net_input_bytes | 2605369052.000 |
| redis | net_output_bytes | 765064133.000 |
| redis | ops_per_sec | 16430.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46953.000 |
| redis | used_memory_bytes | 95967048.000 |

## 停止施压后的恢复

- Kafka drain：7.2977322s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：232
- 测量前恢复：complete=true；耗时=4.9374788s；删除帖子/Outbox=0/0；safety epoch=13058
- 预热后恢复：complete=true；耗时=6.3961547s；删除帖子/Outbox=56/112；safety epoch=13115
- 测量后恢复：complete=true；耗时=5.1275827s；删除帖子/Outbox=232/464；safety epoch=13349

## 说明

- SLA values are reference lines, not pass/fail gates.
