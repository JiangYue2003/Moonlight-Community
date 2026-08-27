# Feed 压测报告：hybrid / gateway / mixed-80-20-c4

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:10:43+08:00
- 采样时长：15.3858901s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 304 | 304 | 0 | 0 | 19.76 | 2.609 | 3.366 | 3.745 | 5.934 | 5.934 |
| publish_total | 76 | 76 | 0 | 0 | 4.94 | 772.159 | 938.800 | 1086.637 | 1093.618 | 1093.618 |
| publish_draft | 76 | 76 | 0 | 0 | 4.94 | 3.767 | 4.695 | 5.545 | 11.197 | 11.197 |
| publish_metadata | 76 | 76 | 0 | 0 | 4.94 | 242.343 | 313.624 | 370.088 | 392.726 | 392.726 |
| publish_confirm | 76 | 76 | 0 | 0 | 4.94 | 249.705 | 339.077 | 346.356 | 365.277 | 365.277 |
| publish_commit | 76 | 76 | 0 | 0 | 4.94 | 248.442 | 363.302 | 400.591 | 412.214 | 412.214 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 58 | 0.191 |
| mysql | 38 | 0.125 |
| redis | 950 | 3.125 |
| relation | 304 | 1.000 |

- Cold compute：304（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 12160 | 40.000 |
| merge_candidates | 15022 | 49.414 |
| redis_commands | 1824 | 6.000 |
| redis_members | 15022 | 49.414 |
| redis_roundtrips | 608 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 304 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 304 | 0.190 |
| counter | 58 | 0.531 |
| hydrate | 304 | 0.365 |
| inbox | 304 | 0.220 |
| merge_dedup | 304 | 0.010 |
| relation | 304 | 0.928 |
| route | 304 | 0.114 |
| total | 304 | 1.849 |

## Redis 本轮边界增量

- Commands：242270；input：19242750 bytes；output：6836409 bytes
- Hits/Misses：15104/2183；run hit rate：87.37%
- Evicted/Rejected：0/0；ops/s max：17271；safety epoch：11414 -> 11490

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.146 |
| client:loadtest | cpu_percent_total | 2.336 |
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
| docker:zg-canal | cpu_percent | 3.440 |
| docker:zg-canal | memory_percent | 4.700 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.770 |
| docker:zg-es | memory_percent | 11.890 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.390 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 188.710 |
| docker:zg-kafka | memory_percent | 7.720 |
| docker:zg-kafka | pids | 145.000 |
| docker:zg-zk | cpu_percent | 38.110 |
| docker:zg-zk | memory_percent | 1.300 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7867.000 |
| kafka | lag_max | 4.000 |
| kafka | lag_total | 4.000 |
| kafka | log_end_offset_total | 7867.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 92776.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.203 |
| process:counter | cpu_seconds_total | 28.938 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46555136.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 5.428 |
| process:gateway | cpu_seconds_total | 5.391 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47767552.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 83.755 |
| process:knowpost | cpu_seconds_total | 312.766 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68382720.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.421 |
| process:relation | cpu_seconds_total | 9.875 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48685056.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.551 |
| process:search | cpu_seconds_total | 4.875 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42864640.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.562 |
| process:user-storage | cpu_seconds_total | 3.188 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41758720.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 28512418.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 11490.000 |
| redis | hit_rate | 0.744 |
| redis | keys | 593175.000 |
| redis | keyspace_hits | 1264777.000 |
| redis | keyspace_misses | 434941.000 |
| redis | net_input_bytes | 2138373545.000 |
| redis | net_output_bytes | 647708168.000 |
| redis | ops_per_sec | 17271.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46244.000 |
| redis | used_memory_bytes | 95390792.000 |

## 停止施压后的恢复

- Kafka drain：5.5154434s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：76
- 测量前恢复：complete=true；耗时=4.9869392s；删除帖子/Outbox=0/0；safety epoch=11392
- 预热后恢复：complete=true；耗时=7.2983084s；删除帖子/Outbox=20/40；safety epoch=11413
- 测量后恢复：complete=true；耗时=4.9668295s；删除帖子/Outbox=76/152；safety epoch=11491

## 说明

- SLA values are reference lines, not pass/fail gates.
