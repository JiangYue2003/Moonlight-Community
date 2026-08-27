# Feed 压测报告：hybrid / gateway / mixed-80-20-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-19T00:09:08+08:00
- 采样时长：1m0.404334s
- 并发：4
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1200 | 1200 | 0 | 0 | 19.87 | 2.635 | 3.695 | 3.920 | 4.875 | 6.231 |
| publish_total | 300 | 300 | 0 | 0 | 4.97 | 784.529 | 909.060 | 920.483 | 964.247 | 971.841 |
| publish_draft | 300 | 300 | 0 | 0 | 4.97 | 3.866 | 4.431 | 4.779 | 76.895 | 92.695 |
| publish_metadata | 300 | 300 | 0 | 0 | 4.97 | 258.449 | 323.866 | 353.178 | 385.631 | 397.221 |
| publish_confirm | 300 | 300 | 0 | 0 | 4.97 | 255.894 | 321.250 | 343.023 | 363.624 | 372.289 |
| publish_commit | 300 | 300 | 0 | 0 | 4.97 | 250.170 | 314.363 | 349.092 | 413.146 | 423.511 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 201 | 0.168 |
| mysql | 134 | 0.112 |
| redis | 3734 | 3.112 |
| relation | 1200 | 1.000 |

- Cold compute：1200（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 48000 | 40.000 |
| merge_candidates | 81987 | 68.323 |
| redis_commands | 7200 | 6.000 |
| redis_members | 86349 | 71.957 |
| redis_roundtrips | 2400 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1200 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1200 | 0.247 |
| counter | 201 | 0.518 |
| hydrate | 1200 | 0.355 |
| inbox | 1200 | 0.209 |
| merge_dedup | 1200 | 0.010 |
| relation | 1200 | 0.960 |
| route | 1200 | 0.096 |
| total | 1200 | 1.905 |

## Redis 本轮边界增量

- Commands：984922；input：78273928 bytes；output：28521414 bytes
- Hits/Misses：60160/7880；run hit rate：88.42%
- Evicted/Rejected：0/0；ops/s max：17749；safety epoch：22640 -> 22940

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.146 |
| client:loadtest | cpu_percent_total | 2.328 |
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
| docker:zg-canal | cpu_percent | 2.800 |
| docker:zg-canal | memory_percent | 5.010 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.030 |
| docker:zg-es | memory_percent | 12.270 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.500 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 244.620 |
| docker:zg-kafka | memory_percent | 7.720 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 39.550 |
| docker:zg-zk | memory_percent | 1.380 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 16784.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 16784.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 326073.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 6.199 |
| process:counter | cpu_seconds_total | 127.750 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44208128.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 6.977 |
| process:gateway | cpu_seconds_total | 35.719 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47882240.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 81.402 |
| process:knowpost | cpu_seconds_total | 1842.375 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 65433600.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 11.625 |
| process:relation | cpu_seconds_total | 62.812 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47734784.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.325 |
| process:search | cpu_seconds_total | 28.500 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42459136.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.874 |
| process:user-storage | cpu_seconds_total | 20.875 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40443904.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 66405820.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 22940.000 |
| redis | hit_rate | 0.811 |
| redis | keys | 597080.000 |
| redis | keyspace_hits | 4000885.000 |
| redis | keyspace_misses | 930855.000 |
| redis | net_input_bytes | 5142379758.000 |
| redis | net_output_bytes | 1769949553.000 |
| redis | ops_per_sec | 17749.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49793.000 |
| redis | used_memory_bytes | 96682344.000 |

## 停止施压后的恢复

- Kafka drain：5.3591025s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：300
- 测量前恢复：complete=true；耗时=5.0105673s；删除帖子/Outbox=0/0；safety epoch=22590
- 预热后恢复：complete=true；耗时=6.2380023s；删除帖子/Outbox=48/96；safety epoch=22639
- 测量后恢复：complete=true；耗时=4.9427221s；删除帖子/Outbox=300/600；safety epoch=22941

## 说明

- SLA values are reference lines, not pass/fail gates.
