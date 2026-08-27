# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-treatment-formal-mixed-c4-rpc`
- 开始时间：2026-08-27T06:17:38+08:00
- 采样时长：1m0.4657065s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1456 | 1456 | 0 | 0 | 24.08 | 1.050 | 2.669 | 3.189 | 4.557 | 6.185 |
| publish_total | 364 | 364 | 0 | 0 | 6.02 | 644.351 | 778.591 | 840.138 | 877.172 | 891.982 |
| publish_draft | 364 | 364 | 0 | 0 | 6.02 | 3.813 | 5.842 | 6.438 | 156.824 | 178.075 |
| publish_metadata | 364 | 364 | 0 | 0 | 6.02 | 205.892 | 246.758 | 258.026 | 378.389 | 445.616 |
| publish_confirm | 364 | 364 | 0 | 0 | 6.02 | 211.235 | 247.821 | 257.428 | 284.972 | 344.702 |
| publish_commit | 364 | 364 | 0 | 0 | 6.02 | 211.339 | 271.312 | 300.852 | 392.263 | 400.173 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1089 | 0.748 |
| counter | 216 | 0.148 |
| mysql | 224 | 0.154 |
| redis | 4373 | 3.003 |
| relation | 216 | 0.148 |

- Cold compute：1045（0.718 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 41800 | 28.709 |
| merge_candidates | 77682 | 53.353 |
| redis_commands | 6270 | 4.306 |
| redis_members | 90115 | 61.892 |
| redis_roundtrips | 1045 | 0.718 |

| page cache source | requests |
|---|---:|
| bypass | 8 |
| l1_fresh | 367 |
| l2_fresh | 1 |
| miss | 1080 |

- L1+L2 Fresh ratio：25.27%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1045 | 0.229 |
| counter | 216 | 0.427 |
| hydrate | 1045 | 0.418 |
| inbox | 1045 | 0.227 |
| merge_dedup | 1045 | 0.008 |
| relation | 216 | 0.998 |
| route | 1045 | 0.304 |
| total | 1456 | 1.043 |

## Redis 本轮边界增量

- Commands：1038323；input：86843015 bytes；output：28327288 bytes
- Hits/Misses：50197/13210；run hit rate：79.17%
- Evicted/Rejected：0/0；ops/s max：19118；safety epoch：42145 -> 42873

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.147 |
| client:loadtest | cpu_percent_total | 2.352 |
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
| docker:zg-canal | cpu_percent | 3.710 |
| docker:zg-canal | memory_percent | 4.980 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.510 |
| docker:zg-es | memory_percent | 13.030 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 4.770 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 23.000 |
| docker:zg-kafka | cpu_percent | 186.800 |
| docker:zg-kafka | memory_percent | 7.560 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 38.900 |
| docker:zg-zk | memory_percent | 1.160 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 25033.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 25033.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 229277.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 14.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.742 |
| process:counter | cpu_seconds_total | 20.844 |
| process:counter | pid | 20740.000 |
| process:counter | process_start_ms | 1787781931297.000 |
| process:counter | rss_bytes | 44220416.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 2.307 |
| process:gateway | cpu_seconds_total | 0.219 |
| process:gateway | pid | 39628.000 |
| process:gateway | process_start_ms | 1787781953868.000 |
| process:gateway | rss_bytes | 37183488.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 84.087 |
| process:knowpost | cpu_seconds_total | 324.469 |
| process:knowpost | pid | 39528.000 |
| process:knowpost | process_start_ms | 1787781942481.000 |
| process:knowpost | rss_bytes | 71458816.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.097 |
| process:relation | cpu_seconds_total | 4.609 |
| process:relation | pid | 21680.000 |
| process:relation | process_start_ms | 1787781936365.000 |
| process:relation | rss_bytes | 45907968.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.325 |
| process:search | cpu_seconds_total | 6.438 |
| process:search | pid | 27196.000 |
| process:search | process_start_ms | 1787781948810.000 |
| process:search | rss_bytes | 42737664.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.538 |
| process:user-storage | cpu_seconds_total | 0.469 |
| process:user-storage | pid | 34404.000 |
| process:user-storage | process_start_ms | 1787781926436.000 |
| process:user-storage | rss_bytes | 34967552.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 24519449.000 |
| redis | connected_clients | 29.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 42873.000 |
| redis | hit_rate | 0.849 |
| redis | keys | 584321.000 |
| redis | keyspace_hits | 1762768.000 |
| redis | keyspace_misses | 316979.000 |
| redis | net_input_bytes | 1980169304.000 |
| redis | net_output_bytes | 742900281.000 |
| redis | ops_per_sec | 19118.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 13276.000 |
| redis | used_memory_bytes | 86694352.000 |

## 停止施压后的恢复

- Kafka drain：5.3616229s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：364
- 测量前恢复：complete=true；耗时=4.9607885s；删除帖子/Outbox=0/0；safety epoch=42007
- 预热后恢复：complete=true；耗时=6.2488653s；删除帖子/Outbox=68/136；safety epoch=42144
- 测量后恢复：complete=true；耗时=4.9855495s；删除帖子/Outbox=364/728；safety epoch=42874

## 说明

- SLA values are reference lines, not pass/fail gates.
