# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-treatment-formal-mixed-c4-rpc`
- 开始时间：2026-08-27T06:19:16+08:00
- 采样时长：1m0.0531046s
- 并发：4
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1424 | 1424 | 0 | 0 | 23.71 | 1.050 | 2.665 | 3.137 | 4.337 | 10.812 |
| publish_total | 356 | 356 | 0 | 0 | 5.93 | 660.129 | 771.486 | 828.745 | 858.716 | 862.173 |
| publish_draft | 356 | 356 | 0 | 0 | 5.93 | 3.771 | 5.435 | 6.360 | 12.072 | 14.221 |
| publish_metadata | 356 | 356 | 0 | 0 | 5.93 | 213.055 | 274.599 | 302.442 | 376.272 | 385.101 |
| publish_confirm | 356 | 356 | 0 | 0 | 5.93 | 211.006 | 272.818 | 297.051 | 370.076 | 383.848 |
| publish_commit | 356 | 356 | 0 | 0 | 5.93 | 215.099 | 254.947 | 274.538 | 287.676 | 297.851 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1038 | 0.729 |
| counter | 211 | 0.148 |
| mysql | 197 | 0.138 |
| redis | 4194 | 2.945 |
| relation | 211 | 0.148 |

- Cold compute：999（0.702 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 39960 | 28.062 |
| merge_candidates | 74107 | 52.041 |
| redis_commands | 5994 | 4.209 |
| redis_members | 85600 | 60.112 |
| redis_roundtrips | 999 | 0.702 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 386 |
| l2_fresh | 1 |
| miss | 1037 |

- L1+L2 Fresh ratio：27.18%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 999 | 0.221 |
| counter | 211 | 0.417 |
| hydrate | 999 | 0.398 |
| inbox | 999 | 0.219 |
| merge_dedup | 999 | 0.012 |
| relation | 211 | 1.062 |
| route | 999 | 0.322 |
| total | 1424 | 1.003 |

## Redis 本轮边界增量

- Commands：1019893；input：85122432 bytes；output：27681391 bytes
- Hits/Misses：48884/12373；run hit rate：79.80%
- Evicted/Rejected：0/0；ops/s max：19346；safety epoch：43013 -> 43725

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.171 |
| client:loadtest | cpu_percent_total | 2.732 |
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
| docker:zg-canal | cpu_percent | 3.150 |
| docker:zg-canal | memory_percent | 4.980 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.590 |
| docker:zg-es | memory_percent | 13.030 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 4.210 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 23.000 |
| docker:zg-kafka | cpu_percent | 262.760 |
| docker:zg-kafka | memory_percent | 8.100 |
| docker:zg-kafka | pids | 152.000 |
| docker:zg-zk | cpu_percent | 43.670 |
| docker:zg-zk | memory_percent | 1.180 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 25366.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 25366.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 236177.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 14.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 8.414 |
| process:counter | cpu_seconds_total | 23.828 |
| process:counter | pid | 20740.000 |
| process:counter | process_start_ms | 1787781931297.000 |
| process:counter | rss_bytes | 44503040.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.784 |
| process:gateway | cpu_seconds_total | 0.234 |
| process:gateway | pid | 39628.000 |
| process:gateway | process_start_ms | 1787781953868.000 |
| process:gateway | rss_bytes | 36683776.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 86.243 |
| process:knowpost | cpu_seconds_total | 371.516 |
| process:knowpost | pid | 39528.000 |
| process:knowpost | process_start_ms | 1787781942481.000 |
| process:knowpost | rss_bytes | 70483968.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 10.065 |
| process:relation | cpu_seconds_total | 5.672 |
| process:relation | pid | 21680.000 |
| process:relation | process_start_ms | 1787781936365.000 |
| process:relation | rss_bytes | 45764608.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.874 |
| process:search | cpu_seconds_total | 7.078 |
| process:search | pid | 27196.000 |
| process:search | process_start_ms | 1787781948810.000 |
| process:search | rss_bytes | 43118592.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.784 |
| process:user-storage | cpu_seconds_total | 0.516 |
| process:user-storage | pid | 34404.000 |
| process:user-storage | process_start_ms | 1787781926436.000 |
| process:user-storage | rss_bytes | 34930688.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 25759906.000 |
| redis | connected_clients | 30.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 43725.000 |
| redis | hit_rate | 0.846 |
| redis | keys | 584175.000 |
| redis | keyspace_hits | 1831535.000 |
| redis | keyspace_misses | 336634.000 |
| redis | net_input_bytes | 2083324131.000 |
| redis | net_output_bytes | 776391774.000 |
| redis | ops_per_sec | 19346.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 13376.000 |
| redis | used_memory_bytes | 86249272.000 |

## 停止施压后的恢复

- Kafka drain：6.6844436s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：356
- 测量前恢复：complete=true；耗时=4.9469443s；删除帖子/Outbox=0/0；safety epoch=42875
- 预热后恢复：complete=true；耗时=6.2848231s；删除帖子/Outbox=68/136；safety epoch=43012
- 测量后恢复：complete=true；耗时=4.9593293s；删除帖子/Outbox=356/712；safety epoch=43726

## 说明

- SLA values are reference lines, not pass/fail gates.
