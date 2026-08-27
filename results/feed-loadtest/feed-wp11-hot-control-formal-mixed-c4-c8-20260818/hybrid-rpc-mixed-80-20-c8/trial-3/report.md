# Feed 压测报告：hybrid / rpc / mixed-80-20-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:54:13+08:00
- 采样时长：1m0.1805316s
- 并发：8
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1184 | 1184 | 0 | 0 | 19.67 | 3.180 | 4.871 | 5.413 | 6.885 | 8.508 |
| publish_total | 296 | 296 | 0 | 0 | 4.92 | 1602.920 | 1745.919 | 1774.196 | 1814.785 | 1828.109 |
| publish_draft | 296 | 296 | 0 | 0 | 4.92 | 3.794 | 5.550 | 36.306 | 161.341 | 175.899 |
| publish_metadata | 296 | 296 | 0 | 0 | 4.92 | 513.612 | 628.970 | 648.428 | 677.764 | 694.743 |
| publish_confirm | 296 | 296 | 0 | 0 | 4.92 | 514.038 | 589.098 | 606.271 | 689.244 | 708.842 |
| publish_commit | 296 | 296 | 0 | 0 | 4.92 | 534.102 | 652.625 | 744.912 | 769.416 | 773.180 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 191 | 0.161 |
| mysql | 122 | 0.103 |
| redis | 3674 | 3.103 |
| relation | 1184 | 1.000 |

- Cold compute：1184（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 47360 | 40.000 |
| merge_candidates | 80079 | 67.634 |
| redis_commands | 7104 | 6.000 |
| redis_members | 83752 | 70.736 |
| redis_roundtrips | 2368 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1184 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1184 | 0.524 |
| counter | 191 | 0.804 |
| hydrate | 1184 | 0.697 |
| inbox | 1184 | 0.539 |
| merge_dedup | 1184 | 0.016 |
| relation | 1184 | 1.218 |
| route | 1184 | 0.136 |
| total | 1184 | 3.150 |

## Redis 本轮边界增量

- Commands：978219；input：77722452 bytes；output：28202248 bytes
- Hits/Misses：59143/7805；run hit rate：88.34%
- Evicted/Rejected：0/0；ops/s max：18044；safety epoch：19683 -> 19979

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.138 |
| client:loadtest | cpu_percent_total | 2.207 |
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
| docker:zg-canal | cpu_percent | 3.480 |
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.510 |
| docker:zg-es | memory_percent | 12.250 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.440 |
| docker:zg-etcd | memory_percent | 0.310 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 141.390 |
| docker:zg-kafka | memory_percent | 7.710 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 41.430 |
| docker:zg-zk | memory_percent | 1.560 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 14452.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 14452.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 256575.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 11.000 |
| process:counter | cpu_percent | 8.316 |
| process:counter | cpu_seconds_total | 99.547 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44417024.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 12.172 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 40452096.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 91.399 |
| process:knowpost | cpu_seconds_total | 1417.812 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66117632.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.970 |
| process:relation | cpu_seconds_total | 42.500 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47575040.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 8.316 |
| process:search | cpu_seconds_total | 22.219 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42270720.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 8.609 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35770368.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 56398797.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 19979.000 |
| redis | hit_rate | 0.786 |
| redis | keys | 597667.000 |
| redis | keyspace_hits | 2887516.000 |
| redis | keyspace_misses | 786801.000 |
| redis | net_input_bytes | 4339326132.000 |
| redis | net_output_bytes | 1400442162.000 |
| redis | ops_per_sec | 18044.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48899.000 |
| redis | used_memory_bytes | 96719000.000 |

## 停止施压后的恢复

- Kafka drain：6.5231086s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：296
- 测量前恢复：complete=true；耗时=4.8949069s；删除帖子/Outbox=0/0；safety epoch=19625
- 预热后恢复：complete=true；耗时=7.7124012s；删除帖子/Outbox=56/112；safety epoch=19682
- 测量后恢复：complete=true；耗时=4.9357331s；删除帖子/Outbox=296/592；safety epoch=19980

## 说明

- SLA values are reference lines, not pass/fail gates.
