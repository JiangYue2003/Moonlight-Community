# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-treatment-formal-mixed-c4-rpc`
- 开始时间：2026-08-27T06:12:42+08:00
- 采样时长：1m0.4101665s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3564 | 3564 | 0 | 0 | 59.00 | 0.527 | 2.093 | 2.647 | 3.805 | 9.153 |
| publish_total | 396 | 396 | 0 | 0 | 6.56 | 592.271 | 713.240 | 744.209 | 773.412 | 788.151 |
| publish_draft | 396 | 396 | 0 | 0 | 6.56 | 3.689 | 4.751 | 5.253 | 5.520 | 6.559 |
| publish_metadata | 396 | 396 | 0 | 0 | 6.56 | 192.546 | 238.981 | 269.579 | 299.954 | 320.577 |
| publish_confirm | 396 | 396 | 0 | 0 | 6.56 | 190.346 | 245.546 | 286.084 | 392.036 | 393.674 |
| publish_commit | 396 | 396 | 0 | 0 | 6.56 | 190.263 | 238.790 | 253.982 | 415.318 | 424.527 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1746 | 0.490 |
| counter | 229 | 0.064 |
| mysql | 225 | 0.063 |
| redis | 6841 | 1.919 |
| relation | 229 | 0.064 |

- Cold compute：1654（0.464 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 66160 | 18.563 |
| merge_candidates | 124753 | 35.004 |
| redis_commands | 9924 | 2.785 |
| redis_members | 147772 | 41.462 |
| redis_roundtrips | 1654 | 0.464 |

| page cache source | requests |
|---|---:|
| l1_fresh | 1818 |
| l2_fresh | 0 |
| miss | 1746 |

- L1+L2 Fresh ratio：51.01%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1654 | 0.216 |
| counter | 229 | 0.460 |
| hydrate | 1654 | 0.355 |
| inbox | 1654 | 0.214 |
| merge_dedup | 1654 | 0.010 |
| relation | 229 | 1.033 |
| route | 1654 | 0.214 |
| total | 3564 | 0.606 |

## Redis 本轮边界增量

- Commands：1057908；input：90696216 bytes；output：35160675 bytes
- Hits/Misses：79759/13508；run hit rate：85.52%
- Evicted/Rejected：0/0；ops/s max：19885；safety epoch：39453 -> 40245

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.230 |
| client:loadtest | cpu_percent_total | 3.673 |
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
| docker:zg-canal | cpu_percent | 3.670 |
| docker:zg-canal | memory_percent | 4.920 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.420 |
| docker:zg-es | memory_percent | 13.020 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 4.420 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 131.390 |
| docker:zg-kafka | memory_percent | 7.550 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 42.120 |
| docker:zg-zk | memory_percent | 0.970 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 23993.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 23993.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 207914.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 6.197 |
| process:counter | cpu_seconds_total | 12.125 |
| process:counter | pid | 20740.000 |
| process:counter | process_start_ms | 1787781931297.000 |
| process:counter | rss_bytes | 43798528.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.094 |
| process:gateway | pid | 39628.000 |
| process:gateway | process_start_ms | 1787781953868.000 |
| process:gateway | rss_bytes | 37515264.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 82.096 |
| process:knowpost | cpu_seconds_total | 182.891 |
| process:knowpost | pid | 39528.000 |
| process:knowpost | process_start_ms | 1787781942481.000 |
| process:knowpost | rss_bytes | 70701056.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.876 |
| process:relation | cpu_seconds_total | 2.531 |
| process:relation | pid | 21680.000 |
| process:relation | process_start_ms | 1787781936365.000 |
| process:relation | rss_bytes | 45465600.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.097 |
| process:search | cpu_seconds_total | 3.547 |
| process:search | pid | 27196.000 |
| process:search | process_start_ms | 1787781948810.000 |
| process:search | rss_bytes | 42577920.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 0.359 |
| process:user-storage | pid | 34404.000 |
| process:user-storage | process_start_ms | 1787781926436.000 |
| process:user-storage | rss_bytes | 35049472.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 20748196.000 |
| redis | connected_clients | 29.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 40245.000 |
| redis | hit_rate | 0.855 |
| redis | keys | 579536.000 |
| redis | keyspace_hits | 1495821.000 |
| redis | keyspace_misses | 254591.000 |
| redis | net_input_bytes | 1661430021.000 |
| redis | net_output_bytes | 628633661.000 |
| redis | ops_per_sec | 19885.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 12981.000 |
| redis | used_memory_bytes | 85178264.000 |

## 停止施压后的恢复

- Kafka drain：5.2599839s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：396
- 测量前恢复：complete=true；耗时=4.9687295s；删除帖子/Outbox=0/0；safety epoch=39291
- 预热后恢复：complete=true；耗时=6.2575082s；删除帖子/Outbox=80/160；safety epoch=39452
- 测量后恢复：complete=true；耗时=5.0386157s；删除帖子/Outbox=396/792；safety epoch=40246

## 说明

- SLA values are reference lines, not pass/fail gates.
