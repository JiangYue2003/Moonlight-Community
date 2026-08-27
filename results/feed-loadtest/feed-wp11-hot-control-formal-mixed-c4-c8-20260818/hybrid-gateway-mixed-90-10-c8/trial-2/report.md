# Feed 压测报告：hybrid / gateway / mixed-90-10-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-19T00:02:29+08:00
- 采样时长：1m1.7206582s
- 并发：8
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2304 | 2304 | 0 | 0 | 37.33 | 4.202 | 5.971 | 6.902 | 8.388 | 12.716 |
| publish_total | 256 | 256 | 0 | 0 | 4.15 | 1818.478 | 2190.159 | 2400.878 | 2533.381 | 2560.996 |
| publish_draft | 256 | 256 | 0 | 0 | 4.15 | 4.418 | 5.071 | 5.727 | 6.188 | 7.159 |
| publish_metadata | 256 | 256 | 0 | 0 | 4.15 | 591.421 | 707.842 | 781.203 | 869.976 | 884.803 |
| publish_confirm | 256 | 256 | 0 | 0 | 4.15 | 636.756 | 797.590 | 840.534 | 903.077 | 932.917 |
| publish_commit | 256 | 256 | 0 | 0 | 4.15 | 627.463 | 836.035 | 916.214 | 982.306 | 989.024 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 220 | 0.095 |
| mysql | 138 | 0.060 |
| redis | 7050 | 3.060 |
| relation | 2304 | 1.000 |

- Cold compute：2304（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 92160 | 40.000 |
| merge_candidates | 152294 | 66.100 |
| redis_commands | 13824 | 6.000 |
| redis_members | 155383 | 67.441 |
| redis_roundtrips | 4608 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2304 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2304 | 0.612 |
| counter | 220 | 1.037 |
| hydrate | 2304 | 0.800 |
| inbox | 2304 | 0.633 |
| merge_dedup | 2304 | 0.011 |
| relation | 2304 | 1.406 |
| route | 2304 | 0.103 |
| total | 2304 | 3.588 |

## Redis 本轮边界增量

- Commands：860246；input：69824822 bytes；output：35985597 bytes
- Hits/Misses：110386/9094；run hit rate：92.39%
- Evicted/Rejected：0/0；ops/s max：16722；safety epoch：21331 -> 21587

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.180 |
| client:loadtest | cpu_percent_total | 2.886 |
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
| docker:zg-canal | cpu_percent | 3.370 |
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.290 |
| docker:zg-es | memory_percent | 12.260 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.930 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 169.740 |
| docker:zg-kafka | memory_percent | 7.950 |
| docker:zg-kafka | pids | 151.000 |
| docker:zg-zk | cpu_percent | 37.780 |
| docker:zg-zk | memory_percent | 1.560 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 15729.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 15729.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 297159.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 8.525 |
| process:counter | cpu_seconds_total | 115.625 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44064768.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 10.845 |
| process:gateway | cpu_seconds_total | 26.984 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47771648.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 89.105 |
| process:knowpost | cpu_seconds_total | 1659.266 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66826240.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 10.803 |
| process:relation | cpu_seconds_total | 55.328 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48246784.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.086 |
| process:search | cpu_seconds_total | 25.703 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42659840.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 7.749 |
| process:user-storage | cpu_seconds_total | 16.078 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40640512.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 61878016.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 21587.000 |
| redis | hit_rate | 0.806 |
| redis | keys | 597295.000 |
| redis | keyspace_hits | 3619899.000 |
| redis | keyspace_misses | 870863.000 |
| redis | net_input_bytes | 4782156369.000 |
| redis | net_output_bytes | 1626258826.000 |
| redis | ops_per_sec | 16722.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49397.000 |
| redis | used_memory_bytes | 96612080.000 |

## 停止施压后的恢复

- Kafka drain：5.5986584s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：256
- 测量前恢复：complete=true；耗时=4.982632s；删除帖子/Outbox=0/0；safety epoch=21273
- 预热后恢复：complete=true；耗时=7.7375557s；删除帖子/Outbox=56/112；safety epoch=21330
- 测量后恢复：complete=true；耗时=5.0497365s；删除帖子/Outbox=256/512；safety epoch=21588

## 说明

- SLA values are reference lines, not pass/fail gates.
