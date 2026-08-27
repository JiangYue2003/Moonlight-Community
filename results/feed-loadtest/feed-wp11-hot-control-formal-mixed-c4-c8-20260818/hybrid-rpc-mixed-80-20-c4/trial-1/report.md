# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:45:56+08:00
- 采样时长：1m0.243533s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1216 | 1216 | 0 | 0 | 20.18 | 1.604 | 2.648 | 2.797 | 3.690 | 6.308 |
| publish_total | 304 | 304 | 0 | 0 | 5.05 | 779.282 | 875.915 | 1009.308 | 1099.251 | 1115.863 |
| publish_draft | 304 | 304 | 0 | 0 | 5.05 | 3.666 | 5.403 | 6.900 | 10.656 | 12.782 |
| publish_metadata | 304 | 304 | 0 | 0 | 5.05 | 254.885 | 304.657 | 332.307 | 382.927 | 389.699 |
| publish_confirm | 304 | 304 | 0 | 0 | 5.05 | 250.801 | 310.553 | 365.135 | 561.815 | 581.317 |
| publish_commit | 304 | 304 | 0 | 0 | 5.05 | 250.193 | 292.112 | 335.038 | 353.242 | 376.285 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 205 | 0.169 |
| mysql | 156 | 0.128 |
| redis | 3804 | 3.128 |
| relation | 1216 | 1.000 |

- Cold compute：1216（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 48640 | 40.000 |
| merge_candidates | 83358 | 68.551 |
| redis_commands | 7296 | 6.000 |
| redis_members | 88136 | 72.480 |
| redis_roundtrips | 2432 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1216 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1216 | 0.183 |
| counter | 205 | 0.419 |
| hydrate | 1216 | 0.298 |
| inbox | 1216 | 0.179 |
| merge_dedup | 1216 | 0.013 |
| relation | 1216 | 0.846 |
| route | 1216 | 0.078 |
| total | 1216 | 1.619 |

## Redis 本轮边界增量

- Commands：994875；input：79067465 bytes；output：28836265 bytes
- Hits/Misses：60902/8058；run hit rate：88.31%
- Evicted/Rejected：0/0；ops/s max：18979；safety epoch：17871 -> 18175

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.169 |
| client:loadtest | cpu_percent_total | 2.697 |
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
| docker:zg-canal | cpu_percent | 3.240 |
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 8.000 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.300 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 184.130 |
| docker:zg-kafka | memory_percent | 7.860 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 50.180 |
| docker:zg-zk | memory_percent | 1.350 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 13050.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 13050.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 220116.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 6.204 |
| process:counter | cpu_seconds_total | 84.859 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44445696.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 12.094 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 40570880.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 85.312 |
| process:knowpost | cpu_seconds_total | 1180.250 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66248704.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.956 |
| process:relation | cpu_seconds_total | 34.938 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47951872.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.327 |
| process:search | cpu_seconds_total | 18.594 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 41893888.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 8.375 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 36343808.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 50388255.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 18175.000 |
| redis | hit_rate | 0.777 |
| redis | keys | 596737.000 |
| redis | keyspace_hits | 2481288.000 |
| redis | keyspace_misses | 712466.000 |
| redis | net_input_bytes | 3863363663.000 |
| redis | net_output_bytes | 1227994925.000 |
| redis | ops_per_sec | 18979.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48402.000 |
| redis | used_memory_bytes | 96771072.000 |

## 停止施压后的恢复

- Kafka drain：5.2941053s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：304
- 测量前恢复：complete=true；耗时=4.8880153s；删除帖子/Outbox=0/0；safety epoch=17813
- 预热后恢复：complete=true；耗时=6.3575515s；删除帖子/Outbox=56/112；safety epoch=17870
- 测量后恢复：complete=true；耗时=5.0049423s；删除帖子/Outbox=304/608；safety epoch=18176

## 说明

- SLA values are reference lines, not pass/fail gates.
