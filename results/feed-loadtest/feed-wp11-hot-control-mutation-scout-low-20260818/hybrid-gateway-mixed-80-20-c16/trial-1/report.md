# Feed 压测报告：hybrid / gateway / mixed-80-20-c16

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:12:20+08:00
- 采样时长：17.5938184s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 384 | 384 | 0 | 0 | 21.83 | 6.785 | 9.606 | 11.132 | 14.059 | 16.415 |
| publish_total | 95 | 95 | 0 | 0 | 5.40 | 2888.187 | 3130.911 | 3146.976 | 3186.670 | 3186.670 |
| publish_draft | 95 | 95 | 0 | 0 | 5.40 | 5.230 | 6.353 | 7.432 | 8.956 | 8.956 |
| publish_metadata | 95 | 95 | 0 | 0 | 5.40 | 958.572 | 1163.266 | 1183.787 | 1223.890 | 1223.890 |
| publish_confirm | 95 | 95 | 0 | 0 | 5.40 | 964.438 | 1033.876 | 1068.436 | 1079.025 | 1079.025 |
| publish_commit | 95 | 95 | 0 | 0 | 5.40 | 949.774 | 1066.056 | 1075.102 | 1128.750 | 1128.750 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 60 | 0.156 |
| mysql | 33 | 0.086 |
| redis | 1185 | 3.086 |
| relation | 384 | 1.000 |

- Cold compute：384（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 15360 | 40.000 |
| merge_candidates | 18939 | 49.320 |
| redis_commands | 2304 | 6.000 |
| redis_members | 18939 | 49.320 |
| redis_roundtrips | 768 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 384 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 384 | 1.214 |
| counter | 60 | 1.353 |
| hydrate | 384 | 1.485 |
| inbox | 384 | 1.217 |
| merge_dedup | 384 | 0.014 |
| relation | 384 | 2.006 |
| route | 384 | 0.225 |
| total | 384 | 6.180 |

## Redis 本轮边界增量

- Commands：302256；input：24020591 bytes；output：8547600 bytes
- Hits/Misses：18721/2801；run hit rate：86.99%
- Evicted/Rejected：0/0；ops/s max：18580；safety epoch：11634 -> 11729

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.089 |
| client:loadtest | cpu_percent_total | 1.421 |
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
| docker:zg-canal | cpu_percent | 1.780 |
| docker:zg-canal | memory_percent | 4.700 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.930 |
| docker:zg-es | memory_percent | 11.890 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.760 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 203.950 |
| docker:zg-kafka | memory_percent | 7.910 |
| docker:zg-kafka | pids | 148.000 |
| docker:zg-zk | cpu_percent | 0.160 |
| docker:zg-zk | memory_percent | 1.300 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 8050.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 8050.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 97575.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.562 |
| process:counter | cpu_seconds_total | 31.469 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46485504.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 4.651 |
| process:gateway | cpu_seconds_total | 6.328 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 48361472.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 82.912 |
| process:knowpost | cpu_seconds_total | 345.406 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 70348800.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.384 |
| process:relation | cpu_seconds_total | 10.875 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48766976.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.776 |
| process:search | cpu_seconds_total | 5.219 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42885120.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.874 |
| process:user-storage | cpu_seconds_total | 3.734 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41037824.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 29302295.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 11729.000 |
| redis | hit_rate | 0.746 |
| redis | keys | 593247.000 |
| redis | keyspace_hits | 1330276.000 |
| redis | keyspace_misses | 452185.000 |
| redis | net_input_bytes | 2200316270.000 |
| redis | net_output_bytes | 669989285.000 |
| redis | ops_per_sec | 18580.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46343.000 |
| redis | used_memory_bytes | 95921480.000 |

## 停止施压后的恢复

- Kafka drain：5.2045151s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：95
- 测量前恢复：complete=true；耗时=4.8474437s；删除帖子/Outbox=0/0；safety epoch=11600
- 预热后恢复：complete=true；耗时=6.2020285s；删除帖子/Outbox=32/64；safety epoch=11633
- 测量后恢复：complete=true；耗时=4.915444s；删除帖子/Outbox=95/190；safety epoch=11730

## 说明

- SLA values are reference lines, not pass/fail gates.
