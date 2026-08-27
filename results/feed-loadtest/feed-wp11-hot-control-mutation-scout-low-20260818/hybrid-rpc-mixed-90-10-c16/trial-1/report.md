# Feed 压测报告：hybrid / rpc / mixed-90-10-c16

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T22:59:22+08:00
- 采样时长：15.4388099s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 720 | 720 | 0 | 0 | 46.64 | 6.932 | 10.141 | 11.134 | 13.428 | 16.055 |
| publish_total | 80 | 80 | 0 | 0 | 5.18 | 2995.043 | 3150.477 | 3175.319 | 3228.389 | 3228.389 |
| publish_draft | 80 | 80 | 0 | 0 | 5.18 | 4.715 | 5.840 | 5.973 | 7.425 | 7.425 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.18 | 986.491 | 1097.381 | 1118.221 | 1148.337 | 1148.337 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.18 | 960.501 | 1144.547 | 1198.981 | 1210.235 | 1210.235 |
| publish_commit | 80 | 80 | 0 | 0 | 5.18 | 1035.409 | 1092.334 | 1104.054 | 1134.136 | 1134.136 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 60 | 0.083 |
| mysql | 22 | 0.031 |
| redis | 2182 | 3.031 |
| relation | 720 | 1.000 |

- Cold compute：720（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 28800 | 40.000 |
| merge_candidates | 34012 | 47.239 |
| redis_commands | 4320 | 6.000 |
| redis_members | 34012 | 47.239 |
| redis_roundtrips | 1440 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 720 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 720 | 1.426 |
| counter | 60 | 1.812 |
| hydrate | 720 | 1.608 |
| inbox | 720 | 1.369 |
| merge_dedup | 720 | 0.010 |
| relation | 720 | 2.274 |
| route | 720 | 0.169 |
| total | 720 | 6.878 |

## Redis 本轮边界增量

- Commands：256409；input：20861415 bytes；output：10464376 bytes
- Hits/Misses：33168/3843；run hit rate：89.62%
- Evicted/Rejected：0/0；ops/s max：17910；safety epoch：9938 -> 10018

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.196 |
| client:loadtest | cpu_percent_total | 3.137 |
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
| docker:zg-canal | cpu_percent | 2.620 |
| docker:zg-canal | memory_percent | 4.640 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.790 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.880 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 135.360 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 40.170 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 6779.000 |
| kafka | lag_max | 12.000 |
| kafka | lag_total | 12.000 |
| kafka | log_end_offset_total | 6779.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 63206.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 34.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.429 |
| process:counter | cpu_seconds_total | 12.328 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 43646976.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.125 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 38076416.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 84.423 |
| process:knowpost | cpu_seconds_total | 118.094 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68202496.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 13.184 |
| process:relation | cpu_seconds_total | 3.734 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47910912.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.609 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42430464.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.203 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35233792.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 23634861.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 10018.000 |
| redis | hit_rate | 0.724 |
| redis | keys | 592415.000 |
| redis | keyspace_hits | 839176.000 |
| redis | keyspace_misses | 320533.000 |
| redis | net_input_bytes | 1756131451.000 |
| redis | net_output_bytes | 507874146.000 |
| redis | ops_per_sec | 17910.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45563.000 |
| redis | used_memory_bytes | 95944208.000 |

## 停止施压后的恢复

- Kafka drain：5.2696854s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.8360641s；删除帖子/Outbox=0/0；safety epoch=9907
- 预热后恢复：complete=true；耗时=7.7331876s；删除帖子/Outbox=29/58；safety epoch=9937
- 测量后恢复：complete=true；耗时=4.9272354s；删除帖子/Outbox=80/160；safety epoch=10019

## 说明

- SLA values are reference lines, not pass/fail gates.
