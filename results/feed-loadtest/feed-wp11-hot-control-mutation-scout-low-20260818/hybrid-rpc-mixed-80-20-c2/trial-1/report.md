# Feed 压测报告：hybrid / rpc / mixed-80-20-c2

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:00:08+08:00
- 采样时长：15.0520449s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 288 | 288 | 0 | 0 | 19.13 | 1.584 | 2.237 | 2.707 | 5.871 | 14.445 |
| publish_total | 72 | 72 | 0 | 0 | 4.78 | 403.848 | 472.789 | 523.083 | 535.131 | 535.131 |
| publish_draft | 72 | 72 | 0 | 0 | 4.78 | 3.230 | 3.992 | 4.277 | 4.806 | 4.806 |
| publish_metadata | 72 | 72 | 0 | 0 | 4.78 | 132.479 | 162.553 | 200.228 | 231.411 | 231.411 |
| publish_confirm | 72 | 72 | 0 | 0 | 4.78 | 131.933 | 163.408 | 170.334 | 189.257 | 189.257 |
| publish_commit | 72 | 72 | 0 | 0 | 4.78 | 133.339 | 155.940 | 182.637 | 252.917 | 252.917 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 56 | 0.194 |
| mysql | 34 | 0.118 |
| redis | 898 | 3.118 |
| relation | 288 | 1.000 |

- Cold compute：288（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 11520 | 40.000 |
| merge_candidates | 14202 | 49.312 |
| redis_commands | 1728 | 6.000 |
| redis_members | 14202 | 49.312 |
| redis_roundtrips | 576 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 288 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 288 | 0.097 |
| counter | 56 | 0.499 |
| hydrate | 288 | 0.254 |
| inbox | 288 | 0.144 |
| merge_dedup | 288 | 0.014 |
| relation | 288 | 0.821 |
| route | 288 | 0.099 |
| total | 288 | 1.443 |

## Redis 本轮边界增量

- Commands：227858；input：18094520 bytes；output：6448165 bytes
- Hits/Misses：14371/2067；run hit rate：87.43%
- Evicted/Rejected：0/0；ops/s max：16121；safety epoch：10038 -> 10110

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.175 |
| client:loadtest | cpu_percent_total | 2.803 |
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
| docker:zg-canal | cpu_percent | 1.830 |
| docker:zg-canal | memory_percent | 4.650 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.140 |
| docker:zg-es | memory_percent | 11.860 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.830 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 179.610 |
| docker:zg-kafka | memory_percent | 7.700 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 18.760 |
| docker:zg-zk | memory_percent | 1.480 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 6845.000 |
| kafka | lag_max | 6.000 |
| kafka | lag_total | 6.000 |
| kafka | log_end_offset_total | 6845.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 65061.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.649 |
| process:counter | cpu_seconds_total | 13.422 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44056576.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.794 |
| process:gateway | cpu_seconds_total | 0.141 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 37421056.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 66.605 |
| process:knowpost | cpu_seconds_total | 129.766 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66658304.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.649 |
| process:relation | cpu_seconds_total | 4.219 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47566848.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.734 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42176512.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.203 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35250176.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 23940767.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 10110.000 |
| redis | hit_rate | 0.725 |
| redis | keys | 592451.000 |
| redis | keyspace_hits | 866990.000 |
| redis | keyspace_misses | 328106.000 |
| redis | net_input_bytes | 1780000729.000 |
| redis | net_output_bytes | 516541967.000 |
| redis | ops_per_sec | 16121.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45608.000 |
| redis | used_memory_bytes | 95244896.000 |

## 停止施压后的恢复

- Kafka drain：5.2209629s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：72
- 测量前恢复：complete=true；耗时=4.8204719s；删除帖子/Outbox=0/0；safety epoch=10020
- 预热后恢复：complete=true；耗时=6.2325175s；删除帖子/Outbox=16/32；safety epoch=10037
- 测量后恢复：complete=true；耗时=4.9932838s；删除帖子/Outbox=72/144；safety epoch=10111

## 说明

- SLA values are reference lines, not pass/fail gates.
