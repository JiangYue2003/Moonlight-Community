# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:14:55+08:00
- 采样时长：1m0.0226705s
- 并发：2
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 304 | 304 | 0 | 0 | 5.06 | 380.647 | 454.403 | 482.205 | 541.457 | 553.874 |
| publish_draft | 304 | 304 | 0 | 0 | 5.06 | 4.007 | 5.333 | 5.787 | 145.165 | 157.416 |
| publish_metadata | 304 | 304 | 0 | 0 | 5.06 | 127.119 | 157.922 | 173.689 | 192.538 | 199.945 |
| publish_confirm | 304 | 304 | 0 | 0 | 5.06 | 124.445 | 148.399 | 156.985 | 182.471 | 207.258 |
| publish_commit | 304 | 304 | 0 | 0 | 5.06 | 125.409 | 154.731 | 164.136 | 184.984 | 187.214 |

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

- Commands：934808；input：73083173 bytes；output：16699871 bytes
- Hits/Misses：708/5478；run hit rate：11.45%
- Evicted/Rejected：0/0；ops/s max：16844；safety epoch：11787 -> 12091

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.156 |
| client:loadtest | cpu_percent_total | 2.499 |
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
| docker:zg-canal | memory_percent | 4.710 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 12.880 |
| docker:zg-es | memory_percent | 12.140 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.210 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 19.000 |
| docker:zg-kafka | cpu_percent | 166.700 |
| docker:zg-kafka | memory_percent | 7.760 |
| docker:zg-kafka | pids | 146.000 |
| docker:zg-zk | cpu_percent | 38.290 |
| docker:zg-zk | memory_percent | 1.570 |
| docker:zg-zk | pids | 107.000 |
| kafka | current_offset_total | 8327.000 |
| kafka | lag_max | 6.000 |
| kafka | lag_total | 6.000 |
| kafka | log_end_offset_total | 8327.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 103091.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 4.647 |
| process:counter | cpu_seconds_total | 35.562 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45006848.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 6.344 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46690304.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 71.288 |
| process:knowpost | cpu_seconds_total | 388.500 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68952064.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.099 |
| process:relation | cpu_seconds_total | 11.375 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48881664.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.098 |
| process:search | cpu_seconds_total | 5.656 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43364352.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.533 |
| process:user-storage | cpu_seconds_total | 3.812 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40108032.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 30479575.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 12091.000 |
| redis | hit_rate | 0.745 |
| redis | keys | 592391.000 |
| redis | keyspace_hits | 1345031.000 |
| redis | keyspace_misses | 465149.000 |
| redis | net_input_bytes | 2291412629.000 |
| redis | net_output_bytes | 691460565.000 |
| redis | ops_per_sec | 16844.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46541.000 |
| redis | used_memory_bytes | 95937568.000 |

## 停止施压后的恢复

- Kafka drain：6.5472114s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：304
- 测量前恢复：complete=true；耗时=4.8930986s；删除帖子/Outbox=0/0；safety epoch=11731
- 预热后恢复：complete=true；耗时=7.6556461s；删除帖子/Outbox=54/108；safety epoch=11786
- 测量后恢复：complete=true；耗时=4.8949086s；删除帖子/Outbox=304/608；safety epoch=12092

## 说明

- SLA values are reference lines, not pass/fail gates.
