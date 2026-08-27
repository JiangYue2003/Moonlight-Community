# Feed 压测报告：hybrid / gateway / publish-c2

- Run ID：`feed-wp11-hot-treatment-scout-gateway-publish-20260818`
- 开始时间：2026-08-18T22:39:26+08:00
- 采样时长：15.1882358s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 76 | 76 | 0 | 0 | 5.00 | 394.574 | 464.312 | 524.341 | 548.420 | 548.420 |
| publish_draft | 76 | 76 | 0 | 0 | 5.00 | 4.585 | 5.822 | 6.351 | 9.755 | 9.755 |
| publish_metadata | 76 | 76 | 0 | 0 | 5.00 | 136.047 | 156.141 | 164.894 | 187.372 | 187.372 |
| publish_confirm | 76 | 76 | 0 | 0 | 5.00 | 122.590 | 149.892 | 151.069 | 175.567 | 175.567 |
| publish_commit | 76 | 76 | 0 | 0 | 5.00 | 121.184 | 172.154 | 197.984 | 249.339 | 249.339 |

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
| bypass | 0 |
| l1_fresh | 0 |
| l2_fresh | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：236205；input：18448018 bytes；output：4232632 bytes
- Hits/Misses：185/1354；run hit rate：12.02%
- Evicted/Rejected：0/0；ops/s max：16340；safety epoch：6735 -> 6887

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.064 |
| client:loadtest | cpu_percent_total | 1.029 |
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
| docker:zg-canal | cpu_percent | 2.630 |
| docker:zg-canal | memory_percent | 3.650 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 9.320 |
| docker:zg-es | memory_percent | 11.680 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.020 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 187.160 |
| docker:zg-kafka | memory_percent | 7.510 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 40.420 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5278.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 5278.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 26572.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 7.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.518 |
| process:counter | cpu_seconds_total | 487.016 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 41922560.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 6.199 |
| process:gateway | cpu_seconds_total | 4.547 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 41861120.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 63.529 |
| process:knowpost | cpu_seconds_total | 593.766 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74928128.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.596 |
| process:relation | cpu_seconds_total | 15.234 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 45318144.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 13.422 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 40701952.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.550 |
| process:user-storage | cpu_seconds_total | 10.766 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 38580224.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 16801187.000 |
| redis | connected_clients | 89.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 6887.000 |
| redis | hit_rate | 0.705 |
| redis | keys | 592372.000 |
| redis | keyspace_hits | 385419.000 |
| redis | keyspace_misses | 162669.000 |
| redis | net_input_bytes | 1219012214.000 |
| redis | net_output_bytes | 336963313.000 |
| redis | ops_per_sec | 16340.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44366.000 |
| redis | used_memory_bytes | 96309592.000 |

## 停止施压后的恢复

- Kafka drain：5.1184513s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：76
- 测量前恢复：complete=true；耗时=4.9057944s；删除帖子/Outbox=0/0；safety epoch=6697
- 预热后恢复：complete=true；耗时=6.384456s；删除帖子/Outbox=18/36；safety epoch=6734
- 测量后恢复：complete=true；耗时=4.9019812s；删除帖子/Outbox=76/152；safety epoch=6888

## 说明

- SLA values are reference lines, not pass/fail gates.
