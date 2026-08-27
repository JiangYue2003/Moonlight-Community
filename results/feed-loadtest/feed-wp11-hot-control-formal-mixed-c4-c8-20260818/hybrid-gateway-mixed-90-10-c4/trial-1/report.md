# Feed 压测报告：hybrid / gateway / mixed-90-10-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:55:53+08:00
- 采样时长：1m0.317661s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2484 | 2484 | 0 | 0 | 41.18 | 2.658 | 3.706 | 3.987 | 5.186 | 13.785 |
| publish_total | 276 | 276 | 0 | 0 | 4.58 | 849.679 | 974.490 | 999.356 | 1053.109 | 1074.501 |
| publish_draft | 276 | 276 | 0 | 0 | 4.58 | 3.971 | 4.788 | 4.984 | 6.280 | 6.415 |
| publish_metadata | 276 | 276 | 0 | 0 | 4.58 | 278.786 | 346.311 | 360.601 | 400.325 | 412.628 |
| publish_confirm | 276 | 276 | 0 | 0 | 4.58 | 280.868 | 334.159 | 348.958 | 381.684 | 419.993 |
| publish_commit | 276 | 276 | 0 | 0 | 4.58 | 273.762 | 342.569 | 379.587 | 473.589 | 477.563 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 228 | 0.092 |
| mysql | 120 | 0.048 |
| redis | 7572 | 3.048 |
| relation | 2484 | 1.000 |

- Cold compute：2484（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 99360 | 40.000 |
| merge_candidates | 168320 | 67.762 |
| redis_commands | 14904 | 6.000 |
| redis_members | 173582 | 69.880 |
| redis_roundtrips | 4968 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2484 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2484 | 0.287 |
| counter | 228 | 0.619 |
| hydrate | 2484 | 0.370 |
| inbox | 2484 | 0.265 |
| merge_dedup | 2484 | 0.009 |
| relation | 2484 | 0.991 |
| route | 2484 | 0.062 |
| total | 2484 | 2.011 |

## Redis 本轮边界增量

- Commands：928271；input：75370656 bytes；output：39034174 bytes
- Hits/Misses：119040/9564；run hit rate：92.56%
- Evicted/Rejected：0/0；ops/s max：17352；safety epoch：20035 -> 20311

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.199 |
| client:loadtest | cpu_percent_total | 3.186 |
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
| docker:zg-canal | cpu_percent | 3.220 |
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.720 |
| docker:zg-es | memory_percent | 12.250 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.090 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 19.000 |
| docker:zg-kafka | cpu_percent | 146.240 |
| docker:zg-kafka | memory_percent | 7.750 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 50.290 |
| docker:zg-zk | memory_percent | 1.360 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 14717.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 14717.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 264942.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 10.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.158 |
| process:counter | cpu_seconds_total | 103.031 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 43778048.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 10.070 |
| process:gateway | cpu_seconds_total | 15.062 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 45731840.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 81.337 |
| process:knowpost | cpu_seconds_total | 1465.141 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66195456.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 18.603 |
| process:relation | cpu_seconds_total | 45.109 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47960064.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.100 |
| process:search | cpu_seconds_total | 22.812 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42156032.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 5.382 |
| process:user-storage | cpu_seconds_total | 9.875 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40296448.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 57529468.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 20311.000 |
| redis | hit_rate | 0.791 |
| redis | keys | 597690.000 |
| redis | keyspace_hits | 3038515.000 |
| redis | keyspace_misses | 803790.000 |
| redis | net_input_bytes | 4430731255.000 |
| redis | net_output_bytes | 1447195755.000 |
| redis | ops_per_sec | 17352.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48999.000 |
| redis | used_memory_bytes | 96668384.000 |

## 停止施压后的恢复

- Kafka drain：5.3835902s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：276
- 测量前恢复：complete=true；耗时=4.9342422s；删除帖子/Outbox=0/0；safety epoch=19981
- 预热后恢复：complete=true；耗时=6.3017554s；删除帖子/Outbox=52/104；safety epoch=20034
- 测量后恢复：complete=true；耗时=4.9500373s；删除帖子/Outbox=276/552；safety epoch=20312

## 说明

- SLA values are reference lines, not pass/fail gates.
