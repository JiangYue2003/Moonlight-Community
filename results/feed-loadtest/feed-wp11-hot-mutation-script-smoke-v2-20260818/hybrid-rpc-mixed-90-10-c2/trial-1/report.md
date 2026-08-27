# Feed 压测报告：hybrid / rpc / mixed-90-10-c2

- Run ID：`feed-wp11-hot-mutation-script-smoke-v2-20260818`
- 开始时间：2026-08-18T16:45:46+08:00
- 采样时长：2.0075242s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 90 | 90 | 0 | 0 | 44.83 | 0.532 | 2.246 | 2.714 | 6.371 | 6.371 |
| publish_total | 10 | 10 | 0 | 0 | 4.98 | 373.582 | 452.001 | 458.322 | 458.322 | 458.322 |
| publish_draft | 10 | 10 | 0 | 0 | 4.98 | 3.669 | 4.359 | 5.288 | 5.288 | 5.288 |
| publish_metadata | 10 | 10 | 0 | 0 | 4.98 | 121.087 | 172.004 | 173.924 | 173.924 | 173.924 |
| publish_confirm | 10 | 10 | 0 | 0 | 4.98 | 123.450 | 138.744 | 139.379 | 139.379 | 139.379 |
| publish_commit | 10 | 10 | 0 | 0 | 4.98 | 127.198 | 148.375 | 150.420 | 150.420 | 150.420 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 63 | 0.700 |
| counter | 20 | 0.222 |
| mysql | 8 | 0.089 |
| redis | 256 | 2.844 |
| relation | 20 | 0.222 |

- Cold compute：62（0.689 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2480 | 27.556 |
| merge_candidates | 2628 | 29.200 |
| redis_commands | 372 | 4.133 |
| redis_members | 2628 | 29.200 |
| redis_roundtrips | 62 | 0.689 |

| page cache source | requests |
|---|---:|
| l1_fresh | 27 |
| l2_fresh | 0 |
| miss | 63 |

- L1+L2 Fresh ratio：30.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 62 | 0.161 |
| counter | 20 | 0.491 |
| hydrate | 62 | 0.308 |
| inbox | 62 | 0.161 |
| merge_dedup | 62 | 0.008 |
| relation | 20 | 1.009 |
| route | 62 | 0.484 |
| total | 90 | 0.859 |

## Redis 本轮边界增量

- Commands：33374；input：2848886 bytes；output：1069060 bytes
- Hits/Misses：2891/801；run hit rate：78.30%
- Evicted/Rejected：0/0；ops/s max：14655；safety epoch：4101 -> 4121

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.195 |
| client:loadtest | cpu_percent_total | 3.113 |
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
| docker:zg-canal | cpu_percent | 1.640 |
| docker:zg-canal | memory_percent | 2.170 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.880 |
| docker:zg-es | memory_percent | 11.350 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 0.650 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 123.120 |
| docker:zg-kafka | memory_percent | 7.360 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 36.340 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4223.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4223.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 3272.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.647 |
| process:counter | cpu_seconds_total | 104.719 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 38879232.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1.125 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 38580224.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 67.496 |
| process:knowpost | cpu_seconds_total | 107.906 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 65150976.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 3.531 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 41844736.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.813 |
| process:search | cpu_seconds_total | 2.859 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 38240256.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 2.703 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 36253696.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3270374.000 |
| redis | connected_clients | 29.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 4121.000 |
| redis | hit_rate | 0.709 |
| redis | keys | 588529.000 |
| redis | keyspace_hits | 130352.000 |
| redis | keyspace_misses | 53619.000 |
| redis | net_input_bytes | 233156821.000 |
| redis | net_output_bytes | 68405830.000 |
| redis | ops_per_sec | 14655.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23133.000 |
| redis | used_memory_bytes | 93746848.000 |

## 停止施压后的恢复

- Kafka drain：5.2047987s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：10
- 测量前恢复：complete=true；耗时=4.9182497s；删除帖子/Outbox=0/0；safety epoch=4087
- 预热后恢复：complete=true；耗时=6.2996957s；删除帖子/Outbox=6/12；safety epoch=4100
- 测量后恢复：complete=true；耗时=4.935933s；删除帖子/Outbox=10/20；safety epoch=4122

## 说明

- SLA values are reference lines, not pass/fail gates.
