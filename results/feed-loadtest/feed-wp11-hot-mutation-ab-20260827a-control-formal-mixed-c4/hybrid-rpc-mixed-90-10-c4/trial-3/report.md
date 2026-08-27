# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-mixed-c4`
- 开始时间：2026-08-27T05:57:34+08:00
- 采样时长：1m0.269008s
- 并发：4
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 3528 | 3528 | 0 | 0 | 58.54 | 2.089 | 2.772 | 3.188 | 3.954 | 5.735 |
| publish_total | 392 | 392 | 0 | 0 | 6.50 | 587.402 | 698.180 | 729.459 | 783.763 | 793.883 |
| publish_draft | 392 | 392 | 0 | 0 | 6.50 | 3.630 | 4.174 | 4.386 | 5.927 | 12.576 |
| publish_metadata | 392 | 392 | 0 | 0 | 6.50 | 190.466 | 242.467 | 255.317 | 273.772 | 292.142 |
| publish_confirm | 392 | 392 | 0 | 0 | 6.50 | 191.572 | 242.152 | 261.416 | 294.122 | 338.109 |
| publish_commit | 392 | 392 | 0 | 0 | 6.50 | 193.304 | 254.256 | 291.031 | 337.081 | 341.971 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 229 | 0.065 |
| mysql | 165 | 0.047 |
| redis | 10749 | 3.047 |
| relation | 3528 | 1.000 |

- Cold compute：3528（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 141120 | 40.000 |
| merge_candidates | 265919 | 75.374 |
| redis_commands | 21168 | 6.000 |
| redis_members | 314009 | 89.005 |
| redis_roundtrips | 7056 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 3528 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 3528 | 0.269 |
| counter | 229 | 0.560 |
| hydrate | 3528 | 0.336 |
| inbox | 3528 | 0.249 |
| merge_dedup | 3528 | 0.011 |
| relation | 3528 | 0.926 |
| route | 3528 | 0.041 |
| total | 3528 | 1.853 |

## Redis 本轮边界增量

- Commands：1001950；input：83291524 bytes；output：54323497 bytes
- Hits/Misses：172961/7543；run hit rate：95.82%
- Evicted/Rejected：0/0；ops/s max：18761；safety epoch：28482 -> 28874

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.264 |
| client:loadtest | cpu_percent_total | 4.226 |
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
| docker:zg-canal | cpu_percent | 5.400 |
| docker:zg-canal | memory_percent | 4.890 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 6.380 |
| docker:zg-es | memory_percent | 12.880 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 4.560 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 175.990 |
| docker:zg-kafka | memory_percent | 7.510 |
| docker:zg-kafka | pids | 130.000 |
| docker:zg-zk | cpu_percent | 35.110 |
| docker:zg-zk | memory_percent | 1.230 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 21409.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 21409.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 151190.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 7.008 |
| process:counter | cpu_seconds_total | 36.703 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 45502464.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 4.312 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 42102784.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 80.533 |
| process:knowpost | cpu_seconds_total | 484.797 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 63148032.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 9.293 |
| process:relation | cpu_seconds_total | 32.531 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 54063104.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.873 |
| process:search | cpu_seconds_total | 10.078 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 43294720.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 3.734 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 42106880.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12187818.000 |
| redis | connected_clients | 80.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 28874.000 |
| redis | hit_rate | 0.873 |
| redis | keys | 572992.000 |
| redis | keyspace_hits | 1013393.000 |
| redis | keyspace_misses | 146876.000 |
| redis | net_input_bytes | 973609572.000 |
| redis | net_output_bytes | 395179679.000 |
| redis | ops_per_sec | 18761.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 12073.000 |
| redis | used_memory_bytes | 81062944.000 |

## 停止施压后的恢复

- Kafka drain：5.3496825s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：392
- 测量前恢复：complete=true；耗时=4.9021382s；删除帖子/Outbox=0/0；safety epoch=28404
- 预热后恢复：complete=true；耗时=6.3326904s；删除帖子/Outbox=76/152；safety epoch=28481
- 测量后恢复：complete=true；耗时=4.9509836s；删除帖子/Outbox=392/784；safety epoch=28875

## 说明

- SLA values are reference lines, not pass/fail gates.
