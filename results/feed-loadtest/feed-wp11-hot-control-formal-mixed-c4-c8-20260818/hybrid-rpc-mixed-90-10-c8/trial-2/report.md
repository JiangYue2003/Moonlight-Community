# Feed 压测报告：hybrid / rpc / mixed-90-10-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:42:36+08:00
- 采样时长：1m0.8545801s
- 并发：8
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2664 | 2664 | 0 | 0 | 43.78 | 3.156 | 4.709 | 5.285 | 6.414 | 10.584 |
| publish_total | 296 | 296 | 0 | 0 | 4.86 | 1587.270 | 1783.547 | 1845.810 | 1913.899 | 1939.185 |
| publish_draft | 296 | 296 | 0 | 0 | 4.86 | 3.847 | 4.824 | 63.171 | 235.924 | 267.048 |
| publish_metadata | 296 | 296 | 0 | 0 | 4.86 | 519.934 | 647.728 | 693.412 | 823.440 | 874.291 |
| publish_confirm | 296 | 296 | 0 | 0 | 4.86 | 511.551 | 608.510 | 635.869 | 706.083 | 715.395 |
| publish_commit | 296 | 296 | 0 | 0 | 4.86 | 533.884 | 622.644 | 666.505 | 697.758 | 707.489 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 218 | 0.082 |
| mysql | 131 | 0.049 |
| redis | 8123 | 3.049 |
| relation | 2664 | 1.000 |

- Cold compute：2664（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 106560 | 40.000 |
| merge_candidates | 181862 | 68.267 |
| redis_commands | 15984 | 6.000 |
| redis_members | 189013 | 70.951 |
| redis_roundtrips | 5328 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2664 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2664 | 0.533 |
| counter | 218 | 0.888 |
| hydrate | 2664 | 0.658 |
| inbox | 2664 | 0.556 |
| merge_dedup | 2664 | 0.011 |
| relation | 2664 | 1.215 |
| route | 2664 | 0.077 |
| total | 2664 | 3.072 |

## Redis 本轮边界增量

- Commands：981629；input：79792824 bytes；output：41715107 bytes
- Hits/Misses：126895/10225；run hit rate：92.54%
- Evicted/Rejected：0/0；ops/s max：18300；safety epoch：17162 -> 17458

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.186 |
| client:loadtest | cpu_percent_total | 2.978 |
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
| docker:zg-canal | cpu_percent | 3.070 |
| docker:zg-canal | memory_percent | 4.950 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.950 |
| docker:zg-es | memory_percent | 12.200 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.150 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 19.000 |
| docker:zg-kafka | cpu_percent | 157.780 |
| docker:zg-kafka | memory_percent | 7.720 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 46.900 |
| docker:zg-zk | memory_percent | 1.340 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 12487.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 12487.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 203851.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 11.000 |
| process:counter | cpu_percent | 6.201 |
| process:counter | cpu_seconds_total | 79.688 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46923776.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 12.094 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 43241472.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 90.612 |
| process:knowpost | cpu_seconds_total | 1084.109 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68820992.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 8.530 |
| process:relation | cpu_seconds_total | 31.188 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49004544.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.101 |
| process:search | cpu_seconds_total | 17.109 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43098112.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 8.328 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 37736448.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 47997144.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 17458.000 |
| redis | hit_rate | 0.767 |
| redis | keys | 596543.000 |
| redis | keyspace_hits | 2239902.000 |
| redis | keyspace_misses | 679666.000 |
| redis | net_input_bytes | 3671933000.000 |
| redis | net_output_bytes | 1143822115.000 |
| redis | ops_per_sec | 18300.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48203.000 |
| redis | used_memory_bytes | 96806792.000 |

## 停止施压后的恢复

- Kafka drain：6.6577139s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：296
- 测量前恢复：complete=true；耗时=4.9608226s；删除帖子/Outbox=0/0；safety epoch=17104
- 预热后恢复：complete=true；耗时=7.5681597s；删除帖子/Outbox=56/112；safety epoch=17161
- 测量后恢复：complete=true；耗时=4.9635573s；删除帖子/Outbox=296/592；safety epoch=17459

## 说明

- SLA values are reference lines, not pass/fail gates.
