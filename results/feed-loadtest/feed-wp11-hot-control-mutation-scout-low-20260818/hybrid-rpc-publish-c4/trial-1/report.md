# Feed 压测报告：hybrid / rpc / publish-c4

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T22:54:28+08:00
- 采样时长：15.6136934s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 84 | 84 | 0 | 0 | 5.38 | 706.838 | 890.372 | 904.687 | 907.843 | 907.843 |
| publish_draft | 84 | 84 | 0 | 0 | 5.38 | 4.455 | 6.321 | 6.428 | 6.955 | 6.955 |
| publish_metadata | 84 | 84 | 0 | 0 | 5.38 | 237.820 | 290.395 | 314.437 | 360.710 | 360.710 |
| publish_confirm | 84 | 84 | 0 | 0 | 5.38 | 229.171 | 338.201 | 351.076 | 388.696 | 388.696 |
| publish_commit | 84 | 84 | 0 | 0 | 5.38 | 235.578 | 283.352 | 305.118 | 326.116 | 326.116 |

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

- Commands：259232；input：20263299 bytes；output：4645954 bytes
- Hits/Misses：200/1509；run hit rate：11.70%
- Evicted/Rejected：0/0；ops/s max：17771；safety epoch：9267 -> 9351

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.125 |
| client:loadtest | cpu_percent_total | 2.001 |
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
| docker:zg-canal | cpu_percent | 2.680 |
| docker:zg-canal | memory_percent | 4.530 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.910 |
| docker:zg-es | memory_percent | 11.840 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.930 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 165.900 |
| docker:zg-kafka | memory_percent | 7.380 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 39.190 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 6283.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 6283.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 48902.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 9.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.199 |
| process:counter | cpu_seconds_total | 4.453 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 39944192.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.062 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 36990976.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 73.609 |
| process:knowpost | cpu_seconds_total | 31.094 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 61005824.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.550 |
| process:relation | cpu_seconds_total | 0.547 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 43229184.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.549 |
| process:search | cpu_seconds_total | 0.359 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 40345600.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.141 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 34754560.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 21437613.000 |
| redis | connected_clients | 24.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 9351.000 |
| redis | hit_rate | 0.696 |
| redis | keys | 592231.000 |
| redis | keyspace_hits | 608743.000 |
| redis | keyspace_misses | 267177.000 |
| redis | net_input_bytes | 1582572480.000 |
| redis | net_output_bytes | 437018476.000 |
| redis | ops_per_sec | 17771.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45269.000 |
| redis | used_memory_bytes | 94714456.000 |

## 停止施压后的恢复

- Kafka drain：5.1773165s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：84
- 测量前恢复：complete=true；耗时=4.9046068s；删除帖子/Outbox=0/0；safety epoch=9245
- 预热后恢复：complete=true；耗时=6.1413577s；删除帖子/Outbox=20/40；safety epoch=9266
- 测量后恢复：complete=true；耗时=5.0001352s；删除帖子/Outbox=84/168；safety epoch=9352

## 说明

- SLA values are reference lines, not pass/fail gates.
