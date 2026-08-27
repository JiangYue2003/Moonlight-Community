# Feed 压测报告：hybrid / rpc / publish-c8

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T22:55:15+08:00
- 采样时长：16.1186195s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 88 | 88 | 0 | 0 | 5.46 | 1437.500 | 1577.726 | 1581.694 | 1613.084 | 1613.084 |
| publish_draft | 88 | 88 | 0 | 0 | 5.46 | 4.059 | 5.918 | 6.800 | 7.867 | 7.867 |
| publish_metadata | 88 | 88 | 0 | 0 | 5.46 | 496.683 | 606.990 | 619.396 | 641.498 | 641.498 |
| publish_confirm | 88 | 88 | 0 | 0 | 5.46 | 471.040 | 549.201 | 552.894 | 561.682 | 561.682 |
| publish_commit | 88 | 88 | 0 | 0 | 5.46 | 497.241 | 535.700 | 566.774 | 582.762 | 582.762 |

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

- Commands：271587；input：21221739 bytes；output：4861104 bytes
- Hits/Misses：210/1581；run hit rate：11.73%
- Evicted/Rejected：0/0；ops/s max：17625；safety epoch：9379 -> 9467

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.145 |
| client:loadtest | cpu_percent_total | 2.327 |
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
| docker:zg-canal | cpu_percent | 2.310 |
| docker:zg-canal | memory_percent | 4.550 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.110 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.920 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 206.790 |
| docker:zg-kafka | memory_percent | 7.880 |
| docker:zg-kafka | pids | 150.000 |
| docker:zg-zk | cpu_percent | 0.110 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 6367.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 6367.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 50718.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 13.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.874 |
| process:counter | cpu_seconds_total | 5.797 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 41033728.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.062 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 37003264.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 88.409 |
| process:knowpost | cpu_seconds_total | 46.031 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 63291392.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.550 |
| process:relation | cpu_seconds_total | 0.641 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 44011520.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.531 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 41414656.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.141 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 34877440.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 21811141.000 |
| redis | connected_clients | 32.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 9467.000 |
| redis | hit_rate | 0.694 |
| redis | keys | 592175.000 |
| redis | keyspace_hits | 619389.000 |
| redis | keyspace_misses | 274150.000 |
| redis | net_input_bytes | 1611377918.000 |
| redis | net_output_bytes | 444006637.000 |
| redis | ops_per_sec | 17625.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45318.000 |
| redis | used_memory_bytes | 95016104.000 |

## 停止施压后的恢复

- Kafka drain：6.5772388s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：88
- 测量前恢复：complete=true；耗时=4.8923251s；删除帖子/Outbox=0/0；safety epoch=9353
- 预热后恢复：complete=true；耗时=6.2463895s；删除帖子/Outbox=24/48；safety epoch=9378
- 测量后恢复：complete=true；耗时=4.9653152s；删除帖子/Outbox=88/176；safety epoch=9468

## 说明

- SLA values are reference lines, not pass/fail gates.
