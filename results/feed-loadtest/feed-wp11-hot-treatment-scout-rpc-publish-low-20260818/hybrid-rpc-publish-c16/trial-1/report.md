# Feed 压测报告：hybrid / rpc / publish-c16

- Run ID：`feed-wp11-hot-treatment-scout-rpc-publish-low-20260818`
- 开始时间：2026-08-18T22:26:10+08:00
- 采样时长：16.8510957s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 96 | 96 | 0 | 0 | 5.70 | 2789.088 | 2912.980 | 2933.112 | 2974.464 | 2974.464 |
| publish_draft | 96 | 96 | 0 | 0 | 5.70 | 4.775 | 10.317 | 10.841 | 13.982 | 13.982 |
| publish_metadata | 96 | 96 | 0 | 0 | 5.70 | 896.225 | 991.149 | 1008.369 | 1070.153 | 1070.153 |
| publish_confirm | 96 | 96 | 0 | 0 | 5.70 | 913.883 | 1097.379 | 1111.140 | 1136.734 | 1136.734 |
| publish_commit | 96 | 96 | 0 | 0 | 5.70 | 923.058 | 1066.687 | 1082.874 | 1141.659 | 1141.659 |

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

- Commands：288786；input：22555169 bytes；output：5145365 bytes
- Hits/Misses：227/1740；run hit rate：11.54%
- Evicted/Rejected：0/0；ops/s max：18992；safety epoch：4903 -> 5095

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.133 |
| client:loadtest | cpu_percent_total | 2.133 |
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
| docker:zg-canal | cpu_percent | 4.120 |
| docker:zg-canal | memory_percent | 2.880 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 7.780 |
| docker:zg-es | memory_percent | 11.550 |
| docker:zg-es | pids | 155.000 |
| docker:zg-etcd | cpu_percent | 4.140 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 159.240 |
| docker:zg-kafka | memory_percent | 7.330 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 40.280 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4590.000 |
| kafka | lag_max | 13.000 |
| kafka | lag_total | 13.000 |
| kafka | log_end_offset_total | 4590.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 11120.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.649 |
| process:counter | cpu_seconds_total | 469.469 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 42844160.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4.156 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 38637568.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 73.099 |
| process:knowpost | cpu_seconds_total | 462.109 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 75767808.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.874 |
| process:relation | cpu_seconds_total | 13.281 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 44978176.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.550 |
| process:search | cpu_seconds_total | 11.562 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 43036672.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 10.438 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 36515840.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 13566262.000 |
| redis | connected_clients | 65.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 5095.000 |
| redis | hit_rate | 0.705 |
| redis | keys | 589705.000 |
| redis | keyspace_hits | 205318.000 |
| redis | keyspace_misses | 87556.000 |
| redis | net_input_bytes | 963961203.000 |
| redis | net_output_bytes | 263170696.000 |
| redis | ops_per_sec | 18992.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43574.000 |
| redis | used_memory_bytes | 95276072.000 |

## 停止施压后的恢复

- Kafka drain：6.5885656s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：96
- 测量前恢复：complete=true；耗时=4.9207545s；删除帖子/Outbox=0/0；safety epoch=4837
- 预热后恢复：complete=true；耗时=6.1804172s；删除帖子/Outbox=32/64；safety epoch=4902
- 测量后恢复：complete=true；耗时=4.9385235s；删除帖子/Outbox=96/192；safety epoch=5096

## 说明

- SLA values are reference lines, not pass/fail gates.
