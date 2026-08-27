# Feed 压测报告：hybrid / gateway / publish-c4

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:30:04+08:00
- 采样时长：1m0.1244417s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 300 | 300 | 0 | 0 | 4.99 | 774.295 | 951.716 | 980.207 | 1134.474 | 1168.330 |
| publish_draft | 300 | 300 | 0 | 0 | 4.99 | 4.404 | 6.347 | 6.923 | 174.987 | 176.058 |
| publish_metadata | 300 | 300 | 0 | 0 | 4.99 | 253.479 | 305.002 | 318.830 | 328.828 | 330.730 |
| publish_confirm | 300 | 300 | 0 | 0 | 4.99 | 253.845 | 336.295 | 364.410 | 402.911 | 409.692 |
| publish_commit | 300 | 300 | 0 | 0 | 4.99 | 253.536 | 331.298 | 354.395 | 434.398 | 454.977 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 0 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 0 | 0.000 |
| relation | 0 | 0.000 |

- Cold compute：0（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：951907；input：74408751 bytes；output：17024197 bytes
- Hits/Misses：699/5814；run hit rate：10.73%
- Evicted/Rejected：0/0；ops/s max：17428；safety epoch：14764 -> 15064

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.104 |
| client:loadtest | cpu_percent_total | 1.663 |
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
| docker:zg-canal | cpu_percent | 2.290 |
| docker:zg-canal | memory_percent | 4.940 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.410 |
| docker:zg-es | memory_percent | 12.160 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.470 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 171.290 |
| docker:zg-kafka | memory_percent | 7.810 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 40.990 |
| docker:zg-zk | memory_percent | 1.500 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 10602.000 |
| kafka | lag_max | 6.000 |
| kafka | lag_total | 6.000 |
| kafka | log_end_offset_total | 10602.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 150066.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 9.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.967 |
| process:counter | cpu_seconds_total | 59.688 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46080000.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 4.646 |
| process:gateway | cpu_seconds_total | 10.078 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47824896.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 75.928 |
| process:knowpost | cpu_seconds_total | 762.109 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 69292032.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 10.071 |
| process:relation | cpu_seconds_total | 17.547 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49217536.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.098 |
| process:search | cpu_seconds_total | 12.141 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43339776.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.871 |
| process:user-storage | cpu_seconds_total | 6.766 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40988672.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 40021334.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 15064.000 |
| redis | hit_rate | 0.720 |
| redis | keys | 595374.000 |
| redis | keyspace_hits | 1445388.000 |
| redis | keyspace_misses | 567455.000 |
| redis | net_input_bytes | 3033925418.000 |
| redis | net_output_bytes | 864886141.000 |
| redis | ops_per_sec | 17428.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47449.000 |
| redis | used_memory_bytes | 96339584.000 |

## 停止施压后的恢复

- Kafka drain：5.3637095s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：300
- 测量前恢复：complete=true；耗时=5.018148s；删除帖子/Outbox=0/0；safety epoch=14706
- 预热后恢复：complete=true；耗时=6.3040855s；删除帖子/Outbox=56/112；safety epoch=14763
- 测量后恢复：complete=true；耗时=4.9350063s；删除帖子/Outbox=300/600；safety epoch=15065

## 说明

- SLA values are reference lines, not pass/fail gates.
