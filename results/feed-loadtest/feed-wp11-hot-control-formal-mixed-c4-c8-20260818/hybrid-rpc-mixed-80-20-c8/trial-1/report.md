# Feed 压测报告：hybrid / rpc / mixed-80-20-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:50:54+08:00
- 采样时长：1m0.2578234s
- 并发：8
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1184 | 1184 | 0 | 0 | 19.65 | 3.229 | 4.843 | 5.513 | 7.175 | 10.684 |
| publish_total | 296 | 296 | 0 | 0 | 4.91 | 1615.558 | 1722.983 | 1801.218 | 1887.389 | 1912.998 |
| publish_draft | 296 | 296 | 0 | 0 | 4.91 | 3.868 | 5.021 | 6.352 | 59.967 | 69.897 |
| publish_metadata | 296 | 296 | 0 | 0 | 4.91 | 535.357 | 654.240 | 681.345 | 726.686 | 727.603 |
| publish_confirm | 296 | 296 | 0 | 0 | 4.91 | 513.477 | 625.625 | 660.793 | 687.194 | 688.082 |
| publish_commit | 296 | 296 | 0 | 0 | 4.91 | 515.747 | 651.924 | 669.994 | 756.751 | 763.759 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 194 | 0.164 |
| mysql | 128 | 0.108 |
| redis | 3680 | 3.108 |
| relation | 1184 | 1.000 |

- Cold compute：1184（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 47360 | 40.000 |
| merge_candidates | 80055 | 67.614 |
| redis_commands | 7104 | 6.000 |
| redis_members | 83723 | 70.712 |
| redis_roundtrips | 2368 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1184 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1184 | 0.519 |
| counter | 194 | 0.894 |
| hydrate | 1184 | 0.744 |
| inbox | 1184 | 0.562 |
| merge_dedup | 1184 | 0.011 |
| relation | 1184 | 1.258 |
| route | 1184 | 0.156 |
| total | 1184 | 3.281 |

## Redis 本轮边界增量

- Commands：976161；input：77557978 bytes；output：28157960 bytes
- Hits/Misses：59217/7801；run hit rate：88.36%
- Evicted/Rejected：0/0；ops/s max：18853；safety epoch：18967 -> 19263

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.177 |
| client:loadtest | cpu_percent_total | 2.826 |
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
| docker:zg-canal | cpu_percent | 3.750 |
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.620 |
| docker:zg-es | memory_percent | 12.240 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.990 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 20.000 |
| docker:zg-kafka | cpu_percent | 171.310 |
| docker:zg-kafka | memory_percent | 7.800 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 41.420 |
| docker:zg-zk | memory_percent | 1.560 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 13896.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 13896.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 242113.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.960 |
| process:counter | cpu_seconds_total | 93.469 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 43679744.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 12.172 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 40448000.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 92.938 |
| process:knowpost | cpu_seconds_total | 1319.562 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 67362816.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 9.293 |
| process:relation | cpu_seconds_total | 39.344 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47575040.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.324 |
| process:search | cpu_seconds_total | 20.844 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42242048.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.551 |
| process:user-storage | cpu_seconds_total | 8.516 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35672064.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 54008820.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 19263.000 |
| redis | hit_rate | 0.783 |
| redis | keys | 597445.000 |
| redis | keyspace_hits | 2726521.000 |
| redis | keyspace_misses | 756963.000 |
| redis | net_input_bytes | 4150051744.000 |
| redis | net_output_bytes | 1331983825.000 |
| redis | ops_per_sec | 18853.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48699.000 |
| redis | used_memory_bytes | 96841208.000 |

## 停止施压后的恢复

- Kafka drain：5.3013797s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：296
- 测量前恢复：complete=true；耗时=4.9136565s；删除帖子/Outbox=0/0；safety epoch=18909
- 预热后恢复：complete=true；耗时=6.3172023s；删除帖子/Outbox=56/112；safety epoch=18966
- 测量后恢复：complete=true；耗时=4.9510582s；删除帖子/Outbox=296/592；safety epoch=19264

## 说明

- SLA values are reference lines, not pass/fail gates.
