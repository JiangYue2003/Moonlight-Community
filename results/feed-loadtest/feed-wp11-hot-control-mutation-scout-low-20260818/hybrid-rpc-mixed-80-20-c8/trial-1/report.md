# Feed 压测报告：hybrid / rpc / mixed-80-20-c8

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:01:45+08:00
- 采样时长：15.1677298s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 320 | 320 | 0 | 0 | 21.10 | 3.144 | 4.781 | 5.729 | 6.485 | 9.437 |
| publish_total | 80 | 80 | 0 | 0 | 5.27 | 1503.635 | 1619.244 | 1638.584 | 1668.481 | 1668.481 |
| publish_draft | 80 | 80 | 0 | 0 | 5.27 | 3.702 | 4.393 | 5.331 | 6.457 | 6.457 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.27 | 484.010 | 563.369 | 591.196 | 601.883 | 601.883 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.27 | 502.808 | 554.543 | 561.275 | 585.485 | 585.485 |
| publish_commit | 80 | 80 | 0 | 0 | 5.27 | 513.263 | 619.672 | 636.771 | 665.971 | 665.971 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 55 | 0.172 |
| mysql | 35 | 0.109 |
| redis | 995 | 3.109 |
| relation | 320 | 1.000 |

- Cold compute：320（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 12800 | 40.000 |
| merge_candidates | 15697 | 49.053 |
| redis_commands | 1920 | 6.000 |
| redis_members | 15697 | 49.053 |
| redis_roundtrips | 640 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 320 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 320 | 0.503 |
| counter | 55 | 0.824 |
| hydrate | 320 | 0.680 |
| inbox | 320 | 0.528 |
| merge_dedup | 320 | 0.010 |
| relation | 320 | 1.247 |
| route | 320 | 0.150 |
| total | 320 | 3.142 |

## Redis 本轮边界增量

- Commands：253447；input：20132149 bytes；output：7160185 bytes
- Hits/Misses：15735/2340；run hit rate：87.05%
- Evicted/Rejected：0/0；ops/s max：18313；safety epoch：10242 -> 10322

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.135 |
| client:loadtest | cpu_percent_total | 2.163 |
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
| docker:zg-canal | cpu_percent | 1.650 |
| docker:zg-canal | memory_percent | 4.660 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.530 |
| docker:zg-es | memory_percent | 11.860 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.710 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 162.560 |
| docker:zg-kafka | memory_percent | 7.550 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 47.490 |
| docker:zg-zk | memory_percent | 1.290 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7005.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 7005.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 69295.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.967 |
| process:counter | cpu_seconds_total | 16.094 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44457984.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.141 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 38334464.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 92.222 |
| process:knowpost | cpu_seconds_total | 159.734 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68935680.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.914 |
| process:relation | cpu_seconds_total | 5.000 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48037888.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.131 |
| process:search | cpu_seconds_total | 2.078 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42483712.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.203 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35299328.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 24642307.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 10322.000 |
| redis | hit_rate | 0.729 |
| redis | keys | 592554.000 |
| redis | keyspace_hits | 927707.000 |
| redis | keyspace_misses | 344234.000 |
| redis | net_input_bytes | 1834878062.000 |
| redis | net_output_bytes | 536368942.000 |
| redis | ops_per_sec | 18313.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45706.000 |
| redis | used_memory_bytes | 95510000.000 |

## 停止施压后的恢复

- Kafka drain：6.6508165s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.990804s；删除帖子/Outbox=0/0；safety epoch=10216
- 预热后恢复：complete=true；耗时=7.576786s；删除帖子/Outbox=24/48；safety epoch=10241
- 测量后恢复：complete=true；耗时=4.9062232s；删除帖子/Outbox=80/160；safety epoch=10323

## 说明

- SLA values are reference lines, not pass/fail gates.
