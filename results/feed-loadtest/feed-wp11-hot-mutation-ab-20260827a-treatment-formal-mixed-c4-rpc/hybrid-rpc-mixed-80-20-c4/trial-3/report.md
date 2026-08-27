# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-treatment-formal-mixed-c4-rpc`
- 开始时间：2026-08-27T06:20:55+08:00
- 采样时长：1m0.5563738s
- 并发：4
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1456 | 1456 | 0 | 0 | 24.04 | 1.052 | 2.684 | 3.217 | 4.416 | 6.315 |
| publish_total | 364 | 364 | 0 | 0 | 6.01 | 654.250 | 739.130 | 749.094 | 813.547 | 820.985 |
| publish_draft | 364 | 364 | 0 | 0 | 6.01 | 3.780 | 5.410 | 6.356 | 9.466 | 16.256 |
| publish_metadata | 364 | 364 | 0 | 0 | 6.01 | 214.575 | 257.588 | 274.292 | 355.536 | 361.501 |
| publish_confirm | 364 | 364 | 0 | 0 | 6.01 | 210.202 | 264.474 | 279.197 | 377.776 | 394.466 |
| publish_commit | 364 | 364 | 0 | 0 | 6.01 | 213.256 | 260.729 | 275.738 | 314.530 | 325.099 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1047 | 0.719 |
| counter | 216 | 0.148 |
| mysql | 198 | 0.136 |
| redis | 4282 | 2.941 |
| relation | 216 | 0.148 |

- Cold compute：1021（0.701 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 40840 | 28.049 |
| merge_candidates | 75891 | 52.123 |
| redis_commands | 6126 | 4.207 |
| redis_members | 88360 | 60.687 |
| redis_roundtrips | 1021 | 0.701 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 409 |
| l2_fresh | 0 |
| miss | 1047 |

- L1+L2 Fresh ratio：28.09%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1021 | 0.222 |
| counter | 216 | 0.492 |
| hydrate | 1021 | 0.415 |
| inbox | 1021 | 0.221 |
| merge_dedup | 1021 | 0.013 |
| relation | 216 | 1.048 |
| route | 1021 | 0.335 |
| total | 1456 | 1.045 |

## Redis 本轮边界增量

- Commands：1038119；input：86678378 bytes；output：28220851 bytes
- Hits/Misses：49880/12703；run hit rate：79.70%
- Evicted/Rejected：0/0；ops/s max：18646；safety epoch：43857 -> 44585

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.179 |
| client:loadtest | cpu_percent_total | 2.864 |
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
| docker-state:zg-kafka | restart_count | 4.000 |
| docker-state:zg-kafka | running | 1.000 |
| docker-state:zg-zk | health_configured | 1.000 |
| docker-state:zg-zk | healthy | 1.000 |
| docker-state:zg-zk | restart_count | 0.000 |
| docker-state:zg-zk | running | 1.000 |
| docker:zg-canal | cpu_percent | 3.220 |
| docker:zg-canal | memory_percent | 4.980 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.380 |
| docker:zg-es | memory_percent | 13.030 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.330 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 215.080 |
| docker:zg-kafka | memory_percent | 7.950 |
| docker:zg-kafka | pids | 154.000 |
| docker:zg-zk | cpu_percent | 49.000 |
| docker:zg-zk | memory_percent | 1.230 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 25702.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 25702.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 243123.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 6.337 |
| process:counter | cpu_seconds_total | 26.688 |
| process:counter | pid | 20740.000 |
| process:counter | process_start_ms | 1787781931297.000 |
| process:counter | rss_bytes | 43778048.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.234 |
| process:gateway | pid | 39628.000 |
| process:gateway | process_start_ms | 1787781953868.000 |
| process:gateway | rss_bytes | 37294080.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 82.947 |
| process:knowpost | cpu_seconds_total | 418.172 |
| process:knowpost | pid | 39528.000 |
| process:knowpost | process_start_ms | 1787781942481.000 |
| process:knowpost | rss_bytes | 70479872.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.872 |
| process:relation | cpu_seconds_total | 6.516 |
| process:relation | pid | 21680.000 |
| process:relation | process_start_ms | 1787781936365.000 |
| process:relation | rss_bytes | 46288896.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.327 |
| process:search | cpu_seconds_total | 8.078 |
| process:search | pid | 27196.000 |
| process:search | process_start_ms | 1787781948810.000 |
| process:search | rss_bytes | 43036672.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 0.578 |
| process:user-storage | pid | 34404.000 |
| process:user-storage | process_start_ms | 1787781926436.000 |
| process:user-storage | rss_bytes | 35008512.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 27006889.000 |
| redis | connected_clients | 30.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 44585.000 |
| redis | hit_rate | 0.843 |
| redis | keys | 584032.000 |
| redis | keyspace_hits | 1900939.000 |
| redis | keyspace_misses | 356397.000 |
| redis | net_input_bytes | 2187043990.000 |
| redis | net_output_bytes | 810142457.000 |
| redis | ops_per_sec | 18646.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 13474.000 |
| redis | used_memory_bytes | 86058224.000 |

## 停止施压后的恢复

- Kafka drain：5.297874s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：364
- 测量前恢复：complete=true；耗时=4.8864941s；删除帖子/Outbox=0/0；safety epoch=43727
- 预热后恢复：complete=true；耗时=6.306122s；删除帖子/Outbox=64/128；safety epoch=43856
- 测量后恢复：complete=true；耗时=4.9701292s；删除帖子/Outbox=364/728；safety epoch=44586

## 说明

- SLA values are reference lines, not pass/fail gates.
