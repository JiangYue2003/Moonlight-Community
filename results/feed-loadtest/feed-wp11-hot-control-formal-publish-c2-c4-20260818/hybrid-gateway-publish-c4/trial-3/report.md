# Feed 压测报告：hybrid / gateway / publish-c4

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:33:21+08:00
- 采样时长：1m0.0313463s
- 并发：4
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 276 | 276 | 0 | 0 | 4.60 | 873.770 | 985.722 | 1029.450 | 1163.361 | 1170.713 |
| publish_draft | 276 | 276 | 0 | 0 | 4.60 | 4.721 | 6.025 | 6.811 | 9.810 | 10.326 |
| publish_metadata | 276 | 276 | 0 | 0 | 4.60 | 281.054 | 344.728 | 372.966 | 414.220 | 415.079 |
| publish_confirm | 276 | 276 | 0 | 0 | 4.60 | 278.380 | 363.351 | 374.976 | 417.849 | 420.107 |
| publish_commit | 276 | 276 | 0 | 0 | 4.60 | 284.093 | 345.700 | 384.242 | 441.041 | 444.233 |

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

- Commands：882037；input：68926316 bytes；output：15784054 bytes
- Hits/Misses：651/5478；run hit rate：10.62%
- Evicted/Rejected：0/0；ops/s max：16791；safety epoch：15485 -> 15761

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.127 |
| client:loadtest | cpu_percent_total | 2.030 |
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
| docker:zg-canal | cpu_percent | 1.760 |
| docker:zg-canal | memory_percent | 4.950 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.920 |
| docker:zg-es | memory_percent | 12.180 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.190 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 20.000 |
| docker:zg-kafka | cpu_percent | 172.900 |
| docker:zg-kafka | memory_percent | 7.860 |
| docker:zg-kafka | pids | 139.000 |
| docker:zg-zk | cpu_percent | 44.500 |
| docker:zg-zk | memory_percent | 1.330 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 11133.000 |
| kafka | lag_max | 4.000 |
| kafka | lag_total | 4.000 |
| kafka | log_end_offset_total | 11133.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 161055.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 9.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.199 |
| process:counter | cpu_seconds_total | 65.375 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45395968.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 10.841 |
| process:gateway | cpu_seconds_total | 12.062 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 48529408.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 73.616 |
| process:knowpost | cpu_seconds_total | 848.922 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 69783552.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.873 |
| process:relation | cpu_seconds_total | 18.719 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48525312.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.873 |
| process:search | cpu_seconds_total | 13.688 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43212800.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 8.507 |
| process:user-storage | cpu_seconds_total | 8.062 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41455616.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 42273916.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 15761.000 |
| redis | hit_rate | 0.715 |
| redis | keys | 595850.000 |
| redis | keyspace_hits | 1467752.000 |
| redis | keyspace_misses | 590941.000 |
| redis | net_input_bytes | 3209310604.000 |
| redis | net_output_bytes | 905717450.000 |
| redis | ops_per_sec | 16791.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47646.000 |
| redis | used_memory_bytes | 96197432.000 |

## 停止施压后的恢复

- Kafka drain：5.3961039s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：276
- 测量前恢复：complete=true；耗时=4.9339489s；删除帖子/Outbox=0/0；safety epoch=15427
- 预热后恢复：complete=true；耗时=6.3446534s；删除帖子/Outbox=56/112；safety epoch=15484
- 测量后恢复：complete=true；耗时=4.9981463s；删除帖子/Outbox=276/552；safety epoch=15762

## 说明

- SLA values are reference lines, not pass/fail gates.
