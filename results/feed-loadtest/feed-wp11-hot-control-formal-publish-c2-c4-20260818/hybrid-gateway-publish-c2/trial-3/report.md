# Feed 压测报告：hybrid / gateway / publish-c2

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:28:23+08:00
- 采样时长：1m0.141983s
- 并发：2
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 282 | 282 | 0 | 0 | 4.69 | 412.470 | 483.613 | 525.875 | 595.509 | 623.758 |
| publish_draft | 282 | 282 | 0 | 0 | 4.69 | 4.361 | 5.757 | 6.138 | 37.750 | 51.453 |
| publish_metadata | 282 | 282 | 0 | 0 | 4.69 | 135.868 | 164.455 | 175.769 | 188.219 | 205.727 |
| publish_confirm | 282 | 282 | 0 | 0 | 4.69 | 133.775 | 164.676 | 178.537 | 237.928 | 270.844 |
| publish_commit | 282 | 282 | 0 | 0 | 4.69 | 135.810 | 167.878 | 183.066 | 323.751 | 324.022 |

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

- Commands：895071；input：69943542 bytes；output：16010511 bytes
- Hits/Misses：664/5493；run hit rate：10.78%
- Evicted/Rejected：0/0；ops/s max：16443；safety epoch：14422 -> 14704

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.112 |
| client:loadtest | cpu_percent_total | 1.793 |
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
| docker:zg-canal | cpu_percent | 42.820 |
| docker:zg-canal | memory_percent | 4.950 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 9.070 |
| docker:zg-es | memory_percent | 12.160 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.970 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 155.170 |
| docker:zg-kafka | memory_percent | 7.640 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 38.120 |
| docker:zg-zk | memory_percent | 1.560 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 10327.000 |
| kafka | lag_max | 5.000 |
| kafka | lag_total | 5.000 |
| kafka | log_end_offset_total | 10327.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 144419.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 7.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 6.195 |
| process:counter | cpu_seconds_total | 57.172 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45948928.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.122 |
| process:gateway | cpu_seconds_total | 9.141 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47464448.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 65.452 |
| process:knowpost | cpu_seconds_total | 718.219 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 71077888.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 16.812 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48660480.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.092 |
| process:search | cpu_seconds_total | 11.375 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43515904.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.865 |
| process:user-storage | cpu_seconds_total | 6.125 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41025536.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 38862634.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 14704.000 |
| redis | hit_rate | 0.723 |
| redis | keys | 595168.000 |
| redis | keyspace_hits | 1434182.000 |
| redis | keyspace_misses | 555562.000 |
| redis | net_input_bytes | 2943699091.000 |
| redis | net_output_bytes | 843895181.000 |
| redis | ops_per_sec | 16443.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47350.000 |
| redis | used_memory_bytes | 96067880.000 |

## 停止施压后的恢复

- Kafka drain：6.6156459s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：282
- 测量前恢复：complete=true；耗时=4.9122253s；删除帖子/Outbox=0/0；safety epoch=14370
- 预热后恢复：complete=true；耗时=6.4212433s；删除帖子/Outbox=50/100；safety epoch=14421
- 测量后恢复：complete=true；耗时=4.9945343s；删除帖子/Outbox=282/564；safety epoch=14705

## 说明

- SLA values are reference lines, not pass/fail gates.
