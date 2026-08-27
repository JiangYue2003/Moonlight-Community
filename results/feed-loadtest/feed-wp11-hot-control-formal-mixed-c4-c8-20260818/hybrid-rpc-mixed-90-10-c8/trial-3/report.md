# Feed 压测报告：hybrid / rpc / mixed-90-10-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:44:16+08:00
- 采样时长：1m0.9299829s
- 并发：8
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2637 | 2637 | 0 | 0 | 43.28 | 3.221 | 4.730 | 5.377 | 6.888 | 12.209 |
| publish_total | 293 | 293 | 0 | 0 | 4.81 | 1645.328 | 1715.791 | 1732.632 | 1796.807 | 1807.307 |
| publish_draft | 293 | 293 | 0 | 0 | 4.81 | 3.867 | 4.515 | 4.847 | 5.506 | 9.303 |
| publish_metadata | 293 | 293 | 0 | 0 | 4.81 | 544.504 | 639.119 | 663.936 | 714.947 | 726.457 |
| publish_confirm | 293 | 293 | 0 | 0 | 4.81 | 532.870 | 609.673 | 647.999 | 701.888 | 715.864 |
| publish_commit | 293 | 293 | 0 | 0 | 4.81 | 533.336 | 666.540 | 689.907 | 742.429 | 761.795 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 226 | 0.086 |
| mysql | 142 | 0.054 |
| redis | 8053 | 3.054 |
| relation | 2637 | 1.000 |

- Cold compute：2637（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 105480 | 40.000 |
| merge_candidates | 179664 | 68.132 |
| redis_commands | 15822 | 6.000 |
| redis_members | 186425 | 70.696 |
| redis_roundtrips | 5274 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2637 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2637 | 0.564 |
| counter | 226 | 0.916 |
| hydrate | 2637 | 0.712 |
| inbox | 2637 | 0.603 |
| merge_dedup | 2637 | 0.016 |
| relation | 2637 | 1.240 |
| route | 2637 | 0.084 |
| total | 2637 | 3.241 |

## Redis 本轮边界增量

- Commands：972248；input：79017273 bytes；output：41279544 bytes
- Hits/Misses：125891/10173；run hit rate：92.52%
- Evicted/Rejected：0/0；ops/s max：17858；safety epoch：17518 -> 17811

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.135 |
| client:loadtest | cpu_percent_total | 2.154 |
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
| docker:zg-canal | cpu_percent | 1.950 |
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 8.560 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.190 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 134.580 |
| docker:zg-kafka | memory_percent | 7.680 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 37.240 |
| docker:zg-zk | memory_percent | 1.610 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 12767.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 12767.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 212755.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 11.000 |
| process:counter | cpu_percent | 6.968 |
| process:counter | cpu_seconds_total | 82.359 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44576768.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 12.094 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 40570880.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 90.679 |
| process:knowpost | cpu_seconds_total | 1134.500 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 65581056.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 9.302 |
| process:relation | cpu_seconds_total | 33.641 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47685632.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.100 |
| process:search | cpu_seconds_total | 17.953 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42590208.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 8.359 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 36274176.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 49183036.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 17811.000 |
| redis | hit_rate | 0.775 |
| redis | keys | 596592.000 |
| redis | keyspace_hits | 2399217.000 |
| redis | keyspace_misses | 697587.000 |
| redis | net_input_bytes | 3767921227.000 |
| redis | net_output_bytes | 1193298191.000 |
| redis | ops_per_sec | 17858.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48303.000 |
| redis | used_memory_bytes | 96695648.000 |

## 停止施压后的恢复

- Kafka drain：5.3007711s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：293
- 测量前恢复：complete=true；耗时=4.9717149s；删除帖子/Outbox=0/0；safety epoch=17460
- 预热后恢复：complete=true；耗时=7.61923s；删除帖子/Outbox=56/112；safety epoch=17517
- 测量后恢复：complete=true；耗时=4.9447226s；删除帖子/Outbox=293/586；safety epoch=17812

## 说明

- SLA values are reference lines, not pass/fail gates.
