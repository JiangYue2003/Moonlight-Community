# Feed 压测报告：hybrid / gateway / mixed-80-20-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-19T01:23:12+08:00
- 采样时长：1m0.1190777s
- 并发：8
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1121 | 1121 | 0 | 0 | 18.65 | 4.385 | 6.427 | 7.157 | 8.509 | 17.110 |
| publish_total | 280 | 280 | 0 | 0 | 4.66 | 1684.367 | 1891.043 | 1923.610 | 1986.703 | 1992.830 |
| publish_draft | 280 | 280 | 0 | 0 | 4.66 | 4.443 | 5.301 | 5.516 | 6.744 | 6.985 |
| publish_metadata | 280 | 280 | 0 | 0 | 4.66 | 554.838 | 644.084 | 658.634 | 704.873 | 810.173 |
| publish_confirm | 280 | 280 | 0 | 0 | 4.66 | 549.657 | 667.361 | 710.652 | 746.003 | 750.476 |
| publish_commit | 280 | 280 | 0 | 0 | 4.66 | 553.183 | 720.463 | 733.716 | 823.341 | 849.735 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 210 | 0.187 |
| mysql | 112 | 0.100 |
| redis | 3475 | 3.100 |
| relation | 1121 | 1.000 |

- Cold compute：1121（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 44840 | 40.000 |
| merge_candidates | 74792 | 66.719 |
| redis_commands | 6726 | 6.000 |
| redis_members | 77393 | 69.039 |
| redis_roundtrips | 2242 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1121 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1121 | 0.650 |
| counter | 210 | 1.034 |
| hydrate | 1121 | 0.860 |
| inbox | 1121 | 0.672 |
| merge_dedup | 1121 | 0.014 |
| relation | 1121 | 1.446 |
| route | 1121 | 0.201 |
| total | 1121 | 3.870 |

## Redis 本轮边界增量

- Commands：875859；input：69571293 bytes；output：25665287 bytes
- Hits/Misses：56685/7539；run hit rate：88.26%
- Evicted/Rejected：0/0；ops/s max：17329；safety epoch：23789 -> 24069

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.153 |
| client:loadtest | cpu_percent_total | 2.443 |
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
| docker:zg-canal | cpu_percent | 3.430 |
| docker:zg-canal | memory_percent | 5.020 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.810 |
| docker:zg-es | memory_percent | 12.310 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.810 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 198.410 |
| docker:zg-kafka | memory_percent | 7.840 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 37.300 |
| docker:zg-zk | memory_percent | 1.490 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 17662.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 17662.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 348828.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.195 |
| process:counter | cpu_seconds_total | 209.312 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44208128.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 7.742 |
| process:gateway | cpu_seconds_total | 41.438 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47996928.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 82.747 |
| process:knowpost | cpu_seconds_total | 2120.703 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68153344.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.647 |
| process:relation | cpu_seconds_total | 68.641 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47165440.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.876 |
| process:search | cpu_seconds_total | 33.031 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42414080.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 6.977 |
| process:user-storage | cpu_seconds_total | 25.453 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41050112.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 71720428.000 |
| redis | connected_clients | 43.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 24069.000 |
| redis | hit_rate | 0.813 |
| redis | keys | 592018.000 |
| redis | keyspace_hits | 4259263.000 |
| redis | keyspace_misses | 977866.000 |
| redis | net_input_bytes | 5546722925.000 |
| redis | net_output_bytes | 1907425380.000 |
| redis | ops_per_sec | 17329.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54238.000 |
| redis | used_memory_bytes | 95468752.000 |

## 停止施压后的恢复

- Kafka drain：5.5428406s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：280
- 测量前恢复：complete=true；耗时=4.9350497s；删除帖子/Outbox=0/0；safety epoch=23731
- 预热后恢复：complete=true；耗时=6.5007397s；删除帖子/Outbox=56/112；safety epoch=23788
- 测量后恢复：complete=true；耗时=5.1223657s；删除帖子/Outbox=280/560；safety epoch=24070

## 说明

- SLA values are reference lines, not pass/fail gates.
