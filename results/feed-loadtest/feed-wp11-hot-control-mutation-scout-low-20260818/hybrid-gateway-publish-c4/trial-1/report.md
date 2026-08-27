# Feed 压测报告：hybrid / gateway / publish-c4

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:04:12+08:00
- 采样时长：15.3908823s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 84 | 84 | 0 | 0 | 5.46 | 734.915 | 784.939 | 803.041 | 974.253 | 974.253 |
| publish_draft | 84 | 84 | 0 | 0 | 5.46 | 4.684 | 6.354 | 6.883 | 8.243 | 8.243 |
| publish_metadata | 84 | 84 | 0 | 0 | 5.46 | 249.795 | 274.387 | 306.571 | 431.850 | 431.850 |
| publish_confirm | 84 | 84 | 0 | 0 | 5.46 | 234.651 | 257.918 | 270.346 | 276.173 | 276.173 |
| publish_commit | 84 | 84 | 0 | 0 | 5.46 | 238.955 | 270.003 | 281.635 | 363.593 | 363.593 |

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

- Commands：261247；input：20411950 bytes；output：4677778 bytes
- Hits/Misses：201/1521；run hit rate：11.67%
- Evicted/Rejected：0/0；ops/s max：17901；safety epoch：10566 -> 10650

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.082 |
| client:loadtest | cpu_percent_total | 1.320 |
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
| docker:zg-canal | cpu_percent | 2.860 |
| docker:zg-canal | memory_percent | 4.670 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.720 |
| docker:zg-es | memory_percent | 11.870 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.890 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 221.360 |
| docker:zg-kafka | memory_percent | 7.550 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 45.540 |
| docker:zg-zk | memory_percent | 1.290 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7248.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 7248.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 75006.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 10.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.202 |
| process:counter | cpu_seconds_total | 19.719 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45432832.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 2.326 |
| process:gateway | cpu_seconds_total | 0.422 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 44097536.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 65.210 |
| process:knowpost | cpu_seconds_total | 198.938 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 67678208.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.129 |
| process:relation | cpu_seconds_total | 5.844 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48312320.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.322 |
| process:search | cpu_seconds_total | 3.047 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42647552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.101 |
| process:user-storage | cpu_seconds_total | 0.609 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 39964672.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 25716141.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 10650.000 |
| redis | hit_rate | 0.729 |
| redis | keys | 592735.000 |
| redis | keyspace_hits | 982002.000 |
| redis | keyspace_misses | 366689.000 |
| redis | net_input_bytes | 1918144232.000 |
| redis | net_output_bytes | 560334488.000 |
| redis | ops_per_sec | 17901.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45854.000 |
| redis | used_memory_bytes | 95365544.000 |

## 停止施压后的恢复

- Kafka drain：6.4826568s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：84
- 测量前恢复：complete=true；耗时=4.96581s；删除帖子/Outbox=0/0；safety epoch=10544
- 预热后恢复：complete=true；耗时=7.6831725s；删除帖子/Outbox=20/40；safety epoch=10565
- 测量后恢复：complete=true；耗时=4.8631962s；删除帖子/Outbox=84/168；safety epoch=10651

## 说明

- SLA values are reference lines, not pass/fail gates.
