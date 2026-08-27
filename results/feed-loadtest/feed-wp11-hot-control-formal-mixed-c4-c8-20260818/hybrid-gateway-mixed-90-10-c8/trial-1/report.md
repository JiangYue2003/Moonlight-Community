# Feed 压测报告：hybrid / gateway / mixed-90-10-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-19T00:00:48+08:00
- 采样时长：1m1.511162s
- 并发：8
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2520 | 2520 | 0 | 0 | 40.97 | 3.986 | 5.477 | 6.245 | 7.721 | 9.877 |
| publish_total | 280 | 280 | 0 | 0 | 4.55 | 1706.168 | 1842.884 | 1874.412 | 2029.760 | 2055.912 |
| publish_draft | 280 | 280 | 0 | 0 | 4.55 | 4.357 | 5.181 | 5.403 | 6.726 | 24.110 |
| publish_metadata | 280 | 280 | 0 | 0 | 4.55 | 572.334 | 687.509 | 708.517 | 744.742 | 791.663 |
| publish_confirm | 280 | 280 | 0 | 0 | 4.55 | 555.594 | 646.062 | 673.489 | 739.984 | 773.077 |
| publish_commit | 280 | 280 | 0 | 0 | 4.55 | 569.869 | 671.455 | 689.082 | 736.096 | 747.928 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 239 | 0.095 |
| mysql | 148 | 0.059 |
| redis | 7708 | 3.059 |
| relation | 2520 | 1.000 |

- Cold compute：2520（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 100800 | 40.000 |
| merge_candidates | 169960 | 67.444 |
| redis_commands | 15120 | 6.000 |
| redis_members | 175198 | 69.523 |
| redis_roundtrips | 5040 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2520 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2520 | 0.577 |
| counter | 239 | 0.970 |
| hydrate | 2520 | 0.752 |
| inbox | 2520 | 0.602 |
| merge_dedup | 2520 | 0.008 |
| relation | 2520 | 1.332 |
| route | 2520 | 0.097 |
| total | 2520 | 3.396 |

## Redis 本轮边界增量

- Commands：940635；input：76377659 bytes；output：39557021 bytes
- Hits/Misses：120817/9626；run hit rate：92.62%
- Evicted/Rejected：0/0；ops/s max：17665；safety epoch：20991 -> 21271

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.173 |
| client:loadtest | cpu_percent_total | 2.769 |
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
| docker:zg-canal | cpu_percent | 1.900 |
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.080 |
| docker:zg-es | memory_percent | 12.260 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.100 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 167.520 |
| docker:zg-kafka | memory_percent | 7.880 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 41.170 |
| docker:zg-zk | memory_percent | 1.370 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 15479.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 15479.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 289150.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 9.300 |
| process:counter | cpu_seconds_total | 112.438 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44060672.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 9.545 |
| process:gateway | cpu_seconds_total | 23.875 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46907392.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 89.874 |
| process:knowpost | cpu_seconds_total | 1608.438 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 67203072.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.971 |
| process:relation | cpu_seconds_total | 52.469 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48054272.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.325 |
| process:search | cpu_seconds_total | 24.969 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42381312.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 9.545 |
| process:user-storage | cpu_seconds_total | 14.594 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40452096.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 60802103.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 21271.000 |
| redis | hit_rate | 0.803 |
| redis | keys | 597547.000 |
| redis | keyspace_hits | 3476088.000 |
| redis | keyspace_misses | 854018.000 |
| redis | net_input_bytes | 4695217340.000 |
| redis | net_output_bytes | 1582032627.000 |
| redis | ops_per_sec | 17665.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49295.000 |
| redis | used_memory_bytes | 96699376.000 |

## 停止施压后的恢复

- Kafka drain：5.2628262s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：280
- 测量前恢复：complete=true；耗时=4.9883611s；删除帖子/Outbox=0/0；safety epoch=20941
- 预热后恢复：complete=true；耗时=6.4732318s；删除帖子/Outbox=48/96；safety epoch=20990
- 测量后恢复：complete=true；耗时=5.0231405s；删除帖子/Outbox=280/560；safety epoch=21272

## 说明

- SLA values are reference lines, not pass/fail gates.
