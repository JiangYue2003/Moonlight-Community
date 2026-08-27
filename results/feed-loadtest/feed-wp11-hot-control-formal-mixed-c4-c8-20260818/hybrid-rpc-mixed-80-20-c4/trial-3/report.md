# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:49:15+08:00
- 采样时长：1m0.6977307s
- 并发：4
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1216 | 1216 | 0 | 0 | 20.03 | 1.634 | 2.672 | 3.174 | 3.837 | 12.053 |
| publish_total | 304 | 304 | 0 | 0 | 5.01 | 770.451 | 895.441 | 931.307 | 985.042 | 997.378 |
| publish_draft | 304 | 304 | 0 | 0 | 5.01 | 3.642 | 4.242 | 4.479 | 5.841 | 11.545 |
| publish_metadata | 304 | 304 | 0 | 0 | 5.01 | 253.564 | 315.224 | 353.821 | 447.589 | 466.007 |
| publish_confirm | 304 | 304 | 0 | 0 | 5.01 | 250.393 | 315.275 | 348.871 | 405.804 | 415.960 |
| publish_commit | 304 | 304 | 0 | 0 | 5.01 | 248.331 | 322.597 | 349.595 | 432.318 | 437.796 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 202 | 0.166 |
| mysql | 149 | 0.123 |
| redis | 3797 | 3.123 |
| relation | 1216 | 1.000 |

- Cold compute：1216（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 48640 | 40.000 |
| merge_candidates | 83370 | 68.561 |
| redis_commands | 7296 | 6.000 |
| redis_members | 88126 | 72.472 |
| redis_roundtrips | 2432 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1216 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1216 | 0.203 |
| counter | 202 | 0.502 |
| hydrate | 1216 | 0.301 |
| inbox | 1216 | 0.197 |
| merge_dedup | 1216 | 0.010 |
| relation | 1216 | 0.864 |
| route | 1216 | 0.088 |
| total | 1216 | 1.688 |

## Redis 本轮边界增量

- Commands：999826；input：79462537 bytes；output：28931113 bytes
- Hits/Misses：60837/8093；run hit rate：88.26%
- Evicted/Rejected：0/0；ops/s max：18370；safety epoch：18603 -> 18907

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.182 |
| client:loadtest | cpu_percent_total | 2.909 |
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
| docker:zg-canal | cpu_percent | 1.560 |
| docker:zg-canal | memory_percent | 5.000 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.560 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.120 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 152.180 |
| docker:zg-kafka | memory_percent | 7.700 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 38.490 |
| docker:zg-zk | memory_percent | 1.350 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 13619.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 13619.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 234914.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 7.740 |
| process:counter | cpu_seconds_total | 90.703 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 43761664.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.100 |
| process:gateway | cpu_seconds_total | 12.172 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 40402944.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 73.537 |
| process:knowpost | cpu_seconds_total | 1271.047 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66109440.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 9.288 |
| process:relation | cpu_seconds_total | 37.781 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47538176.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.251 |
| process:search | cpu_seconds_total | 20.000 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42405888.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.813 |
| process:user-storage | cpu_seconds_total | 8.453 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35688448.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 52820663.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 18907.000 |
| redis | hit_rate | 0.781 |
| redis | keys | 597312.000 |
| redis | keyspace_hits | 2646182.000 |
| redis | keyspace_misses | 742265.000 |
| redis | net_input_bytes | 4055993314.000 |
| redis | net_output_bytes | 1297936483.000 |
| redis | ops_per_sec | 18370.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48601.000 |
| redis | used_memory_bytes | 96810200.000 |

## 停止施压后的恢复

- Kafka drain：5.3689173s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：304
- 测量前恢复：complete=true；耗时=4.9181996s；删除帖子/Outbox=0/0；safety epoch=18541
- 预热后恢复：complete=true；耗时=6.3359551s；删除帖子/Outbox=60/120；safety epoch=18602
- 测量后恢复：complete=true；耗时=4.9909894s；删除帖子/Outbox=304/608；safety epoch=18908

## 说明

- SLA values are reference lines, not pass/fail gates.
