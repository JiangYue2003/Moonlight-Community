# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:18:13+08:00
- 采样时长：1m0.1943913s
- 并发：2
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 258 | 258 | 0 | 0 | 4.29 | 447.941 | 593.300 | 693.404 | 744.339 | 748.859 |
| publish_draft | 258 | 258 | 0 | 0 | 4.29 | 4.050 | 5.622 | 6.620 | 169.906 | 253.162 |
| publish_metadata | 258 | 258 | 0 | 0 | 4.29 | 146.540 | 188.769 | 236.702 | 274.640 | 281.663 |
| publish_confirm | 258 | 258 | 0 | 0 | 4.29 | 140.945 | 196.079 | 232.248 | 250.929 | 264.399 |
| publish_commit | 258 | 258 | 0 | 0 | 4.29 | 142.560 | 205.952 | 225.427 | 261.552 | 366.823 |

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

- Commands：804384；input：62924469 bytes；output：14756941 bytes
- Hits/Misses：612/5010；run hit rate：10.89%
- Evicted/Rejected：0/0；ops/s max：16320；safety epoch：12467 -> 12713

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.130 |
| client:loadtest | cpu_percent_total | 2.077 |
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
| docker:zg-canal | cpu_percent | 4.640 |
| docker:zg-canal | memory_percent | 4.740 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.600 |
| docker:zg-es | memory_percent | 12.140 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.570 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 136.590 |
| docker:zg-kafka | memory_percent | 7.490 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 42.250 |
| docker:zg-zk | memory_percent | 1.520 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 8808.000 |
| kafka | lag_max | 9.000 |
| kafka | lag_total | 9.000 |
| kafka | log_end_offset_total | 8808.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 113099.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 12.404 |
| process:counter | cpu_seconds_total | 40.984 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46223360.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 6.359 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 44797952.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 64.309 |
| process:knowpost | cpu_seconds_total | 468.391 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68571136.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 10.067 |
| process:relation | cpu_seconds_total | 12.766 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49229824.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.103 |
| process:search | cpu_seconds_total | 7.141 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43479040.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.548 |
| process:user-storage | cpu_seconds_total | 3.984 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 38129664.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 32484173.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 12713.000 |
| redis | hit_rate | 0.739 |
| redis | keys | 593338.000 |
| redis | keyspace_hits | 1367267.000 |
| redis | keyspace_misses | 487390.000 |
| redis | net_input_bytes | 2447417925.000 |
| redis | net_output_bytes | 728198361.000 |
| redis | ops_per_sec | 16320.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46738.000 |
| redis | used_memory_bytes | 95773960.000 |

## 停止施压后的恢复

- Kafka drain：5.7049614s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：258
- 测量前恢复：complete=true；耗时=5.0111624s；删除帖子/Outbox=0/0；safety epoch=12415
- 预热后恢复：complete=true；耗时=6.4576925s；删除帖子/Outbox=50/100；safety epoch=12466
- 测量后恢复：complete=true；耗时=5.0579489s；删除帖子/Outbox=258/516；safety epoch=12714

## 说明

- SLA values are reference lines, not pass/fail gates.
