# Feed 压测报告：hybrid / rpc / mixed-90-10-c16

- Run ID：`feed-wp11-hot-treatment-scout-rpc-mixed90-20260818`
- 开始时间：2026-08-18T22:30:53+08:00
- 采样时长：17.517739s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 864 | 864 | 0 | 0 | 49.32 | 0.000 | 6.924 | 8.812 | 11.551 | 15.415 |
| publish_total | 96 | 96 | 0 | 0 | 5.48 | 2890.028 | 3000.567 | 3034.331 | 3108.998 | 3108.998 |
| publish_draft | 96 | 96 | 0 | 0 | 5.48 | 5.148 | 6.701 | 7.362 | 8.687 | 8.687 |
| publish_metadata | 96 | 96 | 0 | 0 | 5.48 | 947.385 | 1014.201 | 1021.183 | 1055.986 | 1055.986 |
| publish_confirm | 96 | 96 | 0 | 0 | 5.48 | 1007.266 | 1086.709 | 1122.235 | 1165.909 | 1165.909 |
| publish_commit | 96 | 96 | 0 | 0 | 5.48 | 947.528 | 1042.519 | 1066.258 | 1097.379 | 1097.379 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 163 | 0.189 |
| counter | 60 | 0.069 |
| mysql | 23 | 0.027 |
| redis | 503 | 0.582 |
| relation | 60 | 0.069 |

- Cold compute：120（0.139 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 4800 | 5.556 |
| merge_candidates | 5704 | 6.602 |
| redis_commands | 720 | 0.833 |
| redis_members | 5704 | 6.602 |
| redis_roundtrips | 120 | 0.139 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 701 |
| l2_fresh | 0 |
| miss | 163 |

- L1+L2 Fresh ratio：81.13%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 120 | 1.161 |
| counter | 60 | 1.489 |
| hydrate | 120 | 1.616 |
| inbox | 120 | 1.156 |
| merge_dedup | 120 | 0.000 |
| relation | 60 | 1.693 |
| route | 120 | 1.607 |
| total | 864 | 1.315 |

## Redis 本轮边界增量

- Commands：302308；input：24065451 bytes；output：6294649 bytes
- Hits/Misses：6587/2796；run hit rate：70.20%
- Evicted/Rejected：0/0；ops/s max：18876；safety epoch：5603 -> 5795

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.123 |
| client:loadtest | cpu_percent_total | 1.962 |
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
| docker:zg-canal | cpu_percent | 1.870 |
| docker:zg-canal | memory_percent | 3.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.820 |
| docker:zg-es | memory_percent | 11.640 |
| docker:zg-es | pids | 154.000 |
| docker:zg-etcd | cpu_percent | 4.380 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 146.740 |
| docker:zg-kafka | memory_percent | 7.490 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 43.310 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4856.000 |
| kafka | lag_max | 14.000 |
| kafka | lag_total | 14.000 |
| kafka | log_end_offset_total | 4856.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 17108.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.849 |
| process:counter | cpu_seconds_total | 476.109 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 43311104.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.203 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 37552128.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 97.206 |
| process:knowpost | cpu_seconds_total | 513.609 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 76005376.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 13.922 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 44244992.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 12.125 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 42450944.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 10.484 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35049472.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 14785810.000 |
| redis | connected_clients | 83.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 5795.000 |
| redis | hit_rate | 0.708 |
| redis | keys | 592917.000 |
| redis | keyspace_hits | 278731.000 |
| redis | keyspace_misses | 115084.000 |
| redis | net_input_bytes | 1061089767.000 |
| redis | net_output_bytes | 292504982.000 |
| redis | ops_per_sec | 18876.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43855.000 |
| redis | used_memory_bytes | 97562952.000 |

## 停止施压后的恢复

- Kafka drain：5.1166784s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：96
- 测量前恢复：complete=true；耗时=5.0071034s；删除帖子/Outbox=0/0；safety epoch=5537
- 预热后恢复：complete=true；耗时=6.2766396s；删除帖子/Outbox=32/64；safety epoch=5602
- 测量后恢复：complete=true；耗时=4.8428513s；删除帖子/Outbox=96/192；safety epoch=5796

## 说明

- SLA values are reference lines, not pass/fail gates.
