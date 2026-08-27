# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-treatment-scout-rpc-publish-low-20260818`
- 开始时间：2026-08-18T22:23:46+08:00
- 采样时长：15.2071211s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 80 | 80 | 0 | 0 | 5.26 | 365.008 | 449.318 | 468.946 | 563.765 | 563.765 |
| publish_draft | 80 | 80 | 0 | 0 | 5.26 | 4.051 | 5.250 | 5.439 | 6.364 | 6.364 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.26 | 124.548 | 149.314 | 177.633 | 224.497 | 224.497 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.26 | 119.810 | 146.365 | 148.299 | 198.566 | 198.566 |
| publish_commit | 80 | 80 | 0 | 0 | 5.26 | 116.606 | 152.322 | 160.788 | 204.785 | 204.785 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 0 | 0.000 |
| counter | 0 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 0 | 0.000 |
| relation | 0 | 0.000 |

- Cold compute：0（0.000 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 0 |
| l2_fresh | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：238197；input：18605174 bytes；output：4262665 bytes
- Hits/Misses：195/1420；run hit rate：12.07%
- Evicted/Rejected：0/0；ops/s max：16276；safety epoch：4227 -> 4387

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.083 |
| client:loadtest | cpu_percent_total | 1.336 |
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
| docker:zg-canal | cpu_percent | 6.450 |
| docker:zg-canal | memory_percent | 2.620 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 8.120 |
| docker:zg-es | memory_percent | 11.460 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.030 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 156.940 |
| docker:zg-kafka | memory_percent | 7.640 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 0.110 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4326.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4326.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 5487.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.102 |
| process:counter | cpu_seconds_total | 466.391 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 40464384.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.125 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 38637568.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 66.622 |
| process:knowpost | cpu_seconds_total | 417.812 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 72544256.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.581 |
| process:relation | cpu_seconds_total | 12.562 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 44462080.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 10.812 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 39780352.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 10.359 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 36507648.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12443634.000 |
| redis | connected_clients | 60.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 4387.000 |
| redis | hit_rate | 0.727 |
| redis | keys | 588403.000 |
| redis | keyspace_hits | 173341.000 |
| redis | keyspace_misses | 66367.000 |
| redis | net_input_bytes | 877377374.000 |
| redis | net_output_bytes | 242238476.000 |
| redis | ops_per_sec | 16276.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43428.000 |
| redis | used_memory_bytes | 94407704.000 |

## 停止施压后的恢复

- Kafka drain：6.471705s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=5.0121811s；删除帖子/Outbox=0/0；safety epoch=4189
- 预热后恢复：complete=true；耗时=6.2967732s；删除帖子/Outbox=18/36；safety epoch=4226
- 测量后恢复：complete=true；耗时=4.96647s；删除帖子/Outbox=80/160；safety epoch=4388

## 说明

- SLA values are reference lines, not pass/fail gates.
