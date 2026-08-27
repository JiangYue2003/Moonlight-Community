# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-mixed-c4`
- 开始时间：2026-08-27T05:55:55+08:00
- 采样时长：1m0.140849s
- 并发：4
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3528 | 3528 | 0 | 0 | 58.66 | 2.090 | 2.785 | 3.251 | 4.156 | 15.090 |
| publish_total | 392 | 392 | 0 | 0 | 6.52 | 579.284 | 725.080 | 766.426 | 932.375 | 952.010 |
| publish_draft | 392 | 392 | 0 | 0 | 6.52 | 3.561 | 4.180 | 4.421 | 158.120 | 197.327 |
| publish_metadata | 392 | 392 | 0 | 0 | 6.52 | 184.436 | 238.712 | 278.242 | 328.396 | 334.327 |
| publish_confirm | 392 | 392 | 0 | 0 | 6.52 | 191.797 | 239.643 | 262.505 | 337.121 | 344.015 |
| publish_commit | 392 | 392 | 0 | 0 | 6.52 | 185.375 | 245.200 | 293.818 | 367.148 | 410.779 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 234 | 0.066 |
| mysql | 194 | 0.055 |
| redis | 10778 | 3.055 |
| relation | 3528 | 1.000 |

- Cold compute：3528（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 141120 | 40.000 |
| merge_candidates | 265966 | 75.387 |
| redis_commands | 21168 | 6.000 |
| redis_members | 314067 | 89.021 |
| redis_roundtrips | 7056 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 3528 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 3528 | 0.263 |
| counter | 234 | 0.547 |
| hydrate | 3528 | 0.345 |
| inbox | 3528 | 0.254 |
| merge_dedup | 3528 | 0.012 |
| relation | 3528 | 0.920 |
| route | 3528 | 0.042 |
| total | 3528 | 1.857 |

## Redis 本轮边界增量

- Commands：997496；input：82924865 bytes；output：54242633 bytes
- Hits/Misses：173079/7573；run hit rate：95.81%
- Evicted/Rejected：0/0；ops/s max：19118；safety epoch：28010 -> 28402

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.276 |
| client:loadtest | cpu_percent_total | 4.417 |
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
| docker:zg-canal | cpu_percent | 2.300 |
| docker:zg-canal | memory_percent | 4.800 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.950 |
| docker:zg-es | memory_percent | 12.880 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.260 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 188.120 |
| docker:zg-kafka | memory_percent | 7.540 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 51.190 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 21035.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 21035.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 139746.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 15.911 |
| process:counter | cpu_seconds_total | 34.000 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 45375488.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 4.281 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 42098688.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 86.594 |
| process:knowpost | cpu_seconds_total | 435.312 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 62828544.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.755 |
| process:relation | cpu_seconds_total | 29.484 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 54632448.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 4.644 |
| process:search | cpu_seconds_total | 9.078 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 43212800.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 3.703 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 42291200.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 10967425.000 |
| redis | connected_clients | 80.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 28402.000 |
| redis | hit_rate | 0.857 |
| redis | keys | 572657.000 |
| redis | keyspace_hits | 796712.000 |
| redis | keyspace_misses | 132962.000 |
| redis | net_input_bytes | 872519369.000 |
| redis | net_output_bytes | 330547641.000 |
| redis | ops_per_sec | 19118.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 11975.000 |
| redis | used_memory_bytes | 80941832.000 |

## 停止施压后的恢复

- Kafka drain：6.6847679s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：392
- 测量前恢复：complete=true；耗时=4.9630595s；删除帖子/Outbox=0/0；safety epoch=27932
- 预热后恢复：complete=true；耗时=6.4321771s；删除帖子/Outbox=76/152；safety epoch=28009
- 测量后恢复：complete=true；耗时=5.0038036s；删除帖子/Outbox=392/784；safety epoch=28403

## 说明

- SLA values are reference lines, not pass/fail gates.
