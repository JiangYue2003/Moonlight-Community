# Feed 压测报告：hybrid / gateway / publish-c2

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:25:06+08:00
- 采样时长：1m0.0076483s
- 并发：2
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 260 | 260 | 0 | 0 | 4.33 | 447.944 | 569.034 | 632.475 | 796.895 | 851.452 |
| publish_draft | 260 | 260 | 0 | 0 | 4.33 | 4.770 | 6.059 | 6.648 | 13.753 | 178.006 |
| publish_metadata | 260 | 260 | 0 | 0 | 4.33 | 139.497 | 199.883 | 247.106 | 322.129 | 334.117 |
| publish_confirm | 260 | 260 | 0 | 0 | 4.33 | 141.127 | 187.726 | 200.203 | 266.862 | 350.809 |
| publish_commit | 260 | 260 | 0 | 0 | 4.33 | 143.771 | 190.560 | 238.522 | 271.126 | 308.235 |

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

- Commands：822251；input：64245798 bytes；output：14725443 bytes
- Hits/Misses：620/5166；run hit rate：10.72%
- Evicted/Rejected：0/0；ops/s max：15748；safety epoch：13770 -> 14030

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.114 |
| client:loadtest | cpu_percent_total | 1.823 |
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
| docker:zg-canal | cpu_percent | 3.520 |
| docker:zg-canal | memory_percent | 4.900 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.640 |
| docker:zg-es | memory_percent | 12.160 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.920 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 140.430 |
| docker:zg-kafka | memory_percent | 7.630 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 41.390 |
| docker:zg-zk | memory_percent | 1.320 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 9814.000 |
| kafka | lag_max | 5.000 |
| kafka | lag_total | 5.000 |
| kafka | log_end_offset_total | 9814.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 133837.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.109 |
| process:counter | cpu_seconds_total | 52.531 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45723648.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.871 |
| process:gateway | cpu_seconds_total | 7.297 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46784512.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 71.304 |
| process:knowpost | cpu_seconds_total | 637.062 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 69988352.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 10.064 |
| process:relation | cpu_seconds_total | 15.453 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48685056.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.858 |
| process:search | cpu_seconds_total | 10.203 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43401216.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 6.116 |
| process:user-storage | cpu_seconds_total | 4.875 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41492480.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 36693657.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 14030.000 |
| redis | hit_rate | 0.728 |
| redis | keys | 594585.000 |
| redis | keyspace_hits | 1411862.000 |
| redis | keyspace_misses | 532667.000 |
| redis | net_input_bytes | 2774888921.000 |
| redis | net_output_bytes | 804553229.000 |
| redis | ops_per_sec | 15748.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47152.000 |
| redis | used_memory_bytes | 95894872.000 |

## 停止施压后的恢复

- Kafka drain：6.5595431s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：260
- 测量前恢复：complete=true；耗时=4.930211s；删除帖子/Outbox=0/0；safety epoch=13718
- 预热后恢复：complete=true；耗时=6.2593176s；删除帖子/Outbox=50/100；safety epoch=13769
- 测量后恢复：complete=true；耗时=4.9300048s；删除帖子/Outbox=260/520；safety epoch=14031

## 说明

- SLA values are reference lines, not pass/fail gates.
