# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:16:33+08:00
- 采样时长：1m0.0977633s
- 并发：2
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 264 | 264 | 0 | 0 | 4.39 | 436.658 | 559.502 | 606.509 | 661.415 | 671.539 |
| publish_draft | 264 | 264 | 0 | 0 | 4.39 | 3.872 | 5.421 | 5.822 | 8.613 | 243.184 |
| publish_metadata | 264 | 264 | 0 | 0 | 4.39 | 142.178 | 194.667 | 214.087 | 321.642 | 331.906 |
| publish_confirm | 264 | 264 | 0 | 0 | 4.39 | 138.092 | 178.644 | 190.159 | 241.958 | 281.583 |
| publish_commit | 264 | 264 | 0 | 0 | 4.39 | 142.726 | 193.847 | 237.197 | 292.094 | 402.913 |

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

- Commands：820906；input：64119213 bytes；output：14668085 bytes
- Hits/Misses：628/5205；run hit rate：10.77%
- Evicted/Rejected：0/0；ops/s max：16319；safety epoch：12149 -> 12413

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.119 |
| client:loadtest | cpu_percent_total | 1.898 |
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
| docker:zg-canal | cpu_percent | 2.000 |
| docker:zg-canal | memory_percent | 4.720 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 11.900 |
| docker:zg-es | memory_percent | 12.140 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.450 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 141.740 |
| docker:zg-kafka | memory_percent | 7.580 |
| docker:zg-kafka | pids | 136.000 |
| docker:zg-zk | cpu_percent | 46.730 |
| docker:zg-zk | memory_percent | 1.310 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 8572.000 |
| kafka | lag_max | 5.000 |
| kafka | lag_total | 5.000 |
| kafka | log_end_offset_total | 8572.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 108182.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 7.746 |
| process:counter | cpu_seconds_total | 38.250 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45826048.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 6.344 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46690304.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 66.762 |
| process:knowpost | cpu_seconds_total | 428.141 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 69341184.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.874 |
| process:relation | cpu_seconds_total | 11.891 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49057792.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.326 |
| process:search | cpu_seconds_total | 6.219 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43642880.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.325 |
| process:user-storage | cpu_seconds_total | 3.875 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 38125568.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 31495715.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 12413.000 |
| redis | hit_rate | 0.742 |
| redis | keys | 592912.000 |
| redis | keyspace_hits | 1356161.000 |
| redis | keyspace_misses | 476394.000 |
| redis | net_input_bytes | 2370455214.000 |
| redis | net_output_bytes | 709884487.000 |
| redis | ops_per_sec | 16319.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46640.000 |
| redis | used_memory_bytes | 95795632.000 |

## 停止施压后的恢复

- Kafka drain：6.7855724s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：264
- 测量前恢复：complete=true；耗时=4.8784188s；删除帖子/Outbox=0/0；safety epoch=12093
- 预热后恢复：complete=true；耗时=6.3145779s；删除帖子/Outbox=54/108；safety epoch=12148
- 测量后恢复：complete=true；耗时=5.0513159s；删除帖子/Outbox=264/528；safety epoch=12414

## 说明

- SLA values are reference lines, not pass/fail gates.
