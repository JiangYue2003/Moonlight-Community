# Feed 压测报告：hybrid / rpc / mixed-90-10-c8

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T22:58:32+08:00
- 采样时长：15.450245s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 720 | 720 | 0 | 0 | 46.60 | 3.131 | 4.265 | 4.744 | 6.102 | 9.640 |
| publish_total | 80 | 80 | 0 | 0 | 5.18 | 1511.187 | 1585.339 | 1599.773 | 1679.477 | 1679.477 |
| publish_draft | 80 | 80 | 0 | 0 | 5.18 | 3.758 | 4.254 | 4.423 | 4.878 | 4.878 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.18 | 509.785 | 627.313 | 678.153 | 700.929 | 700.929 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.18 | 497.315 | 557.154 | 563.154 | 602.183 | 602.183 |
| publish_commit | 80 | 80 | 0 | 0 | 5.18 | 481.789 | 566.948 | 580.209 | 590.921 | 590.921 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 60 | 0.083 |
| mysql | 20 | 0.028 |
| redis | 2180 | 3.028 |
| relation | 720 | 1.000 |

- Cold compute：720（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 28800 | 40.000 |
| merge_candidates | 34861 | 48.418 |
| redis_commands | 4320 | 6.000 |
| redis_members | 34861 | 48.418 |
| redis_roundtrips | 1440 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 720 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 720 | 0.512 |
| counter | 60 | 1.133 |
| hydrate | 720 | 0.616 |
| inbox | 720 | 0.557 |
| merge_dedup | 720 | 0.009 |
| relation | 720 | 1.200 |
| route | 720 | 0.105 |
| total | 720 | 3.017 |

## Redis 本轮边界增量

- Commands：256240；input：20841451 bytes；output：10478159 bytes
- Hits/Misses：33444/3537；run hit rate：90.44%
- Evicted/Rejected：0/0；ops/s max：17710；safety epoch：9825 -> 9905

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.183 |
| client:loadtest | cpu_percent_total | 2.933 |
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
| docker:zg-canal | cpu_percent | 2.200 |
| docker:zg-canal | memory_percent | 4.590 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.610 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.820 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 138.960 |
| docker:zg-kafka | memory_percent | 7.520 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 6.410 |
| docker:zg-zk | memory_percent | 1.430 |
| docker:zg-zk | pids | 102.000 |
| kafka | current_offset_total | 6693.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 6693.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 60379.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 7.746 |
| process:counter | cpu_seconds_total | 11.016 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 43237376.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.125 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 38068224.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 82.153 |
| process:knowpost | cpu_seconds_total | 102.938 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 67584000.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.315 |
| process:relation | cpu_seconds_total | 2.766 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 46125056.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.813 |
| process:search | cpu_seconds_total | 1.469 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42614784.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 0.203 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35201024.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 23258026.000 |
| redis | connected_clients | 60.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 9905.000 |
| redis | hit_rate | 0.717 |
| redis | keys | 592380.000 |
| redis | keyspace_hits | 783729.000 |
| redis | keyspace_misses | 309933.000 |
| redis | net_input_bytes | 1725932624.000 |
| redis | net_output_bytes | 492762585.000 |
| redis | ops_per_sec | 17710.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45513.000 |
| redis | used_memory_bytes | 95466448.000 |

## 停止施压后的恢复

- Kafka drain：5.2196859s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.9097839s；删除帖子/Outbox=0/0；safety epoch=9799
- 预热后恢复：complete=true；耗时=6.2002799s；删除帖子/Outbox=24/48；safety epoch=9824
- 测量后恢复：complete=true；耗时=4.9430146s；删除帖子/Outbox=80/160；safety epoch=9906

## 说明

- SLA values are reference lines, not pass/fail gates.
