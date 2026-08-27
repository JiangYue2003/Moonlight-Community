# Feed 压测报告：hybrid / rpc / mixed-90-10-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:39:14+08:00
- 采样时长：1m0.3139776s
- 并发：4
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2520 | 2520 | 0 | 0 | 41.78 | 1.659 | 2.647 | 2.824 | 3.676 | 5.774 |
| publish_total | 280 | 280 | 0 | 0 | 4.64 | 843.828 | 973.079 | 994.897 | 1086.919 | 1124.026 |
| publish_draft | 280 | 280 | 0 | 0 | 4.64 | 3.667 | 6.402 | 15.890 | 149.483 | 157.394 |
| publish_metadata | 280 | 280 | 0 | 0 | 4.64 | 274.570 | 340.688 | 361.699 | 389.958 | 401.933 |
| publish_confirm | 280 | 280 | 0 | 0 | 4.64 | 274.454 | 343.135 | 372.957 | 456.303 | 459.014 |
| publish_commit | 280 | 280 | 0 | 0 | 4.64 | 270.755 | 337.186 | 351.038 | 394.322 | 412.423 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 223 | 0.088 |
| mysql | 150 | 0.060 |
| redis | 7710 | 3.060 |
| relation | 2520 | 1.000 |

- Cold compute：2520（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 100800 | 40.000 |
| merge_candidates | 171469 | 68.043 |
| redis_commands | 15120 | 6.000 |
| redis_members | 177207 | 70.320 |
| redis_roundtrips | 5040 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2520 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2520 | 0.229 |
| counter | 223 | 0.478 |
| hydrate | 2520 | 0.305 |
| inbox | 2520 | 0.210 |
| merge_dedup | 2520 | 0.009 |
| relation | 2520 | 0.862 |
| route | 2520 | 0.045 |
| total | 2520 | 1.682 |

## Redis 本轮边界增量

- Commands：928117；input：75422205 bytes；output：39362080 bytes
- Hits/Misses：120461/9705；run hit rate：92.54%
- Evicted/Rejected：0/0；ops/s max：17613；safety epoch：16477 -> 16757

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.193 |
| client:loadtest | cpu_percent_total | 3.083 |
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
| docker:zg-canal | cpu_percent | 3.050 |
| docker:zg-canal | memory_percent | 4.950 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 7.770 |
| docker:zg-es | memory_percent | 12.200 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.420 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 146.160 |
| docker:zg-kafka | memory_percent | 7.700 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 45.050 |
| docker:zg-zk | memory_percent | 1.590 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 11929.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 11929.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 186193.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 6.622 |
| process:counter | cpu_seconds_total | 74.297 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46764032.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 12.094 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 43819008.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 75.107 |
| process:knowpost | cpu_seconds_total | 985.375 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68780032.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 9.287 |
| process:relation | cpu_seconds_total | 26.484 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49401856.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.626 |
| process:search | cpu_seconds_total | 15.828 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43593728.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 8.234 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 38064128.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 45643815.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 16757.000 |
| redis | hit_rate | 0.749 |
| redis | keys | 596129.000 |
| redis | keyspace_hits | 1923560.000 |
| redis | keyspace_misses | 643662.000 |
| redis | net_input_bytes | 3481446882.000 |
| redis | net_output_bytes | 1045683405.000 |
| redis | ops_per_sec | 17613.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48001.000 |
| redis | used_memory_bytes | 96410536.000 |

## 停止施压后的恢复

- Kafka drain：6.623777s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：280
- 测量前恢复：complete=true；耗时=4.9546443s；删除帖子/Outbox=0/0；safety epoch=16423
- 预热后恢复：complete=true；耗时=6.2814618s；删除帖子/Outbox=52/104；safety epoch=16476
- 测量后恢复：complete=true；耗时=4.9586506s；删除帖子/Outbox=280/560；safety epoch=16758

## 说明

- SLA values are reference lines, not pass/fail gates.
