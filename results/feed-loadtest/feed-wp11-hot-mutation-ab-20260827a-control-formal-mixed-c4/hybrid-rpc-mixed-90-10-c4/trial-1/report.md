# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-mixed-c4`
- 开始时间：2026-08-27T05:54:17+08:00
- 采样时长：1m0.2427798s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3537 | 3537 | 0 | 0 | 58.71 | 2.102 | 2.820 | 3.232 | 4.299 | 10.683 |
| publish_total | 393 | 393 | 0 | 0 | 6.52 | 586.018 | 689.107 | 718.609 | 830.921 | 851.520 |
| publish_draft | 393 | 393 | 0 | 0 | 6.52 | 3.658 | 4.072 | 4.263 | 5.804 | 6.552 |
| publish_metadata | 393 | 393 | 0 | 0 | 6.52 | 194.248 | 230.288 | 240.670 | 291.060 | 312.975 |
| publish_confirm | 393 | 393 | 0 | 0 | 6.52 | 190.060 | 234.533 | 247.313 | 316.912 | 330.581 |
| publish_commit | 393 | 393 | 0 | 0 | 6.52 | 191.566 | 237.945 | 268.235 | 309.026 | 315.122 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 237 | 0.067 |
| mysql | 173 | 0.049 |
| redis | 10784 | 3.049 |
| relation | 3537 | 1.000 |

- Cold compute：3537（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 141480 | 40.000 |
| merge_candidates | 266690 | 75.400 |
| redis_commands | 21222 | 6.000 |
| redis_members | 315171 | 89.107 |
| redis_roundtrips | 7074 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 3537 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 3537 | 0.268 |
| counter | 237 | 0.576 |
| hydrate | 3537 | 0.347 |
| inbox | 3537 | 0.261 |
| merge_dedup | 3537 | 0.009 |
| relation | 3537 | 0.954 |
| route | 3537 | 0.043 |
| total | 3537 | 1.911 |

## Redis 本轮边界增量

- Commands：997683；input：82980779 bytes；output：54383058 bytes
- Hits/Misses：173581/6862；run hit rate：96.20%
- Evicted/Rejected：0/0；ops/s max：18965；safety epoch：27537 -> 27930

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.188 |
| client:loadtest | cpu_percent_total | 3.009 |
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
| docker:zg-canal | cpu_percent | 4.540 |
| docker:zg-canal | memory_percent | 4.690 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 9.510 |
| docker:zg-es | memory_percent | 12.840 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.400 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 158.820 |
| docker:zg-kafka | memory_percent | 7.440 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 44.060 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 20661.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 20661.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 128282.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.969 |
| process:counter | cpu_seconds_total | 31.922 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 44687360.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.266 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 43151360.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 82.882 |
| process:knowpost | cpu_seconds_total | 385.641 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 64159744.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 10.081 |
| process:relation | cpu_seconds_total | 26.188 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 54575104.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.437 |
| process:search | cpu_seconds_total | 8.250 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 43159552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 3.672 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 42287104.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9753053.000 |
| redis | connected_clients | 80.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 27930.000 |
| redis | hit_rate | 0.830 |
| redis | keys | 572254.000 |
| redis | keyspace_hits | 579876.000 |
| redis | keyspace_misses | 119085.000 |
| redis | net_input_bytes | 771929598.000 |
| redis | net_output_bytes | 266021516.000 |
| redis | ops_per_sec | 18965.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 11876.000 |
| redis | used_memory_bytes | 80816784.000 |

## 停止施压后的恢复

- Kafka drain：5.6889102s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：393
- 测量前恢复：complete=true；耗时=5.0103396s；删除帖子/Outbox=0/0；safety epoch=27459
- 预热后恢复：complete=true；耗时=6.2922043s；删除帖子/Outbox=76/152；safety epoch=27536
- 测量后恢复：complete=true；耗时=5.0202921s；删除帖子/Outbox=393/786；safety epoch=27931

## 说明

- SLA values are reference lines, not pass/fail gates.
