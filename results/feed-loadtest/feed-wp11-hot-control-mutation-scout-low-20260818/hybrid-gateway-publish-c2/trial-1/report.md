# Feed 压测报告：hybrid / gateway / publish-c2

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:03:24+08:00
- 采样时长：15.0126867s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 76 | 76 | 0 | 0 | 5.06 | 399.423 | 443.800 | 471.221 | 498.744 | 498.744 |
| publish_draft | 76 | 76 | 0 | 0 | 5.06 | 4.290 | 5.687 | 6.055 | 6.365 | 6.365 |
| publish_metadata | 76 | 76 | 0 | 0 | 5.06 | 127.454 | 162.852 | 165.101 | 171.087 | 171.087 |
| publish_confirm | 76 | 76 | 0 | 0 | 5.06 | 128.976 | 151.604 | 158.489 | 164.396 | 164.396 |
| publish_commit | 76 | 76 | 0 | 0 | 5.06 | 124.978 | 151.049 | 155.119 | 179.600 | 179.600 |

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

- Commands：236448；input：18468050 bytes；output：4231920 bytes
- Hits/Misses：183/1410；run hit rate：11.49%
- Evicted/Rejected：0/0；ops/s max：17056；safety epoch：10466 -> 10542

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.039 |
| client:loadtest | cpu_percent_total | 0.624 |
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
| docker:zg-canal | cpu_percent | 2.980 |
| docker:zg-canal | memory_percent | 4.660 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.110 |
| docker:zg-es | memory_percent | 11.870 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.050 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 258.020 |
| docker:zg-kafka | memory_percent | 7.990 |
| docker:zg-kafka | pids | 146.000 |
| docker:zg-zk | cpu_percent | 0.170 |
| docker:zg-zk | memory_percent | 1.290 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7170.000 |
| kafka | lag_max | 6.000 |
| kafka | lag_total | 6.000 |
| kafka | log_end_offset_total | 7170.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 73324.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.648 |
| process:counter | cpu_seconds_total | 18.344 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45408256.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.100 |
| process:gateway | cpu_seconds_total | 0.297 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 42274816.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 64.290 |
| process:knowpost | cpu_seconds_total | 186.531 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 70639616.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.100 |
| process:relation | cpu_seconds_total | 5.641 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48136192.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.550 |
| process:search | cpu_seconds_total | 2.703 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42369024.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.587 |
| process:user-storage | cpu_seconds_total | 0.375 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 38658048.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 25364147.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 10542.000 |
| redis | hit_rate | 0.730 |
| redis | keys | 592689.000 |
| redis | keyspace_hits | 971357.000 |
| redis | keyspace_misses | 359896.000 |
| redis | net_input_bytes | 1891032847.000 |
| redis | net_output_bytes | 553724522.000 |
| redis | ops_per_sec | 17056.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45804.000 |
| redis | used_memory_bytes | 95197456.000 |

## 停止施压后的恢复

- Kafka drain：5.2608017s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：76
- 测量前恢复：complete=true；耗时=4.8584985s；删除帖子/Outbox=0/0；safety epoch=10446
- 预热后恢复：complete=true；耗时=6.2115407s；删除帖子/Outbox=18/36；safety epoch=10465
- 测量后恢复：complete=true；耗时=4.9315611s；删除帖子/Outbox=76/152；safety epoch=10543

## 说明

- SLA values are reference lines, not pass/fail gates.
