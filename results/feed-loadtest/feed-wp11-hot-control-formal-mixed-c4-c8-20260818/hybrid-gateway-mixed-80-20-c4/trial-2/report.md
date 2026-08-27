# Feed 压测报告：hybrid / gateway / mixed-80-20-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-19T00:07:29+08:00
- 采样时长：1m0.4013343s
- 并发：4
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1200 | 1200 | 0 | 0 | 19.87 | 2.181 | 3.213 | 3.701 | 4.401 | 5.836 |
| publish_total | 300 | 300 | 0 | 0 | 4.97 | 798.076 | 892.894 | 922.041 | 995.926 | 1009.968 |
| publish_draft | 300 | 300 | 0 | 0 | 4.97 | 3.788 | 4.770 | 6.372 | 131.287 | 149.957 |
| publish_metadata | 300 | 300 | 0 | 0 | 4.97 | 253.351 | 322.698 | 343.812 | 413.830 | 427.687 |
| publish_confirm | 300 | 300 | 0 | 0 | 4.97 | 251.621 | 322.766 | 340.459 | 379.351 | 382.888 |
| publish_commit | 300 | 300 | 0 | 0 | 4.97 | 250.418 | 322.929 | 356.735 | 387.634 | 404.327 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 199 | 0.166 |
| mysql | 151 | 0.126 |
| redis | 3751 | 3.126 |
| relation | 1200 | 1.000 |

- Cold compute：1200（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 48000 | 40.000 |
| merge_candidates | 82034 | 68.362 |
| redis_commands | 7200 | 6.000 |
| redis_members | 86402 | 72.002 |
| redis_roundtrips | 2400 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1200 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1200 | 0.209 |
| counter | 199 | 0.469 |
| hydrate | 1200 | 0.319 |
| inbox | 1200 | 0.193 |
| merge_dedup | 1200 | 0.010 |
| relation | 1200 | 0.874 |
| route | 1200 | 0.084 |
| total | 1200 | 1.717 |

## Redis 本轮边界增量

- Commands：985429；input：78324754 bytes；output：28518452 bytes
- Hits/Misses：60000/7872；run hit rate：88.40%
- Evicted/Rejected：0/0；ops/s max：17948；safety epoch：22288 -> 22588

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.137 |
| client:loadtest | cpu_percent_total | 2.199 |
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
| docker:zg-canal | cpu_percent | 2.970 |
| docker:zg-canal | memory_percent | 5.010 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.130 |
| docker:zg-es | memory_percent | 12.270 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.120 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 138.200 |
| docker:zg-kafka | memory_percent | 7.730 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 40.560 |
| docker:zg-zk | memory_percent | 1.380 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 16512.000 |
| kafka | lag_max | 6.000 |
| kafka | lag_total | 6.000 |
| kafka | log_end_offset_total | 16512.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 318979.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 6.202 |
| process:counter | cpu_seconds_total | 124.953 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44339200.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 10.070 |
| process:gateway | cpu_seconds_total | 33.859 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47493120.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 75.186 |
| process:knowpost | cpu_seconds_total | 1797.484 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66768896.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 15.484 |
| process:relation | cpu_seconds_total | 61.406 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47796224.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.101 |
| process:search | cpu_seconds_total | 27.797 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42487808.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.850 |
| process:user-storage | cpu_seconds_total | 19.922 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40493056.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 65236194.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 22588.000 |
| redis | hit_rate | 0.811 |
| redis | keys | 597054.000 |
| redis | keyspace_hits | 3921046.000 |
| redis | keyspace_misses | 916348.000 |
| redis | net_input_bytes | 5049770969.000 |
| redis | net_output_bytes | 1736289746.000 |
| redis | ops_per_sec | 17948.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49696.000 |
| redis | used_memory_bytes | 96631120.000 |

## 停止施压后的恢复

- Kafka drain：6.6094812s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：300
- 测量前恢复：complete=true；耗时=4.9356583s；删除帖子/Outbox=0/0；safety epoch=22234
- 预热后恢复：complete=true；耗时=7.5561258s；删除帖子/Outbox=52/104；safety epoch=22287
- 测量后恢复：complete=true；耗时=4.9175822s；删除帖子/Outbox=300/600；safety epoch=22589

## 说明

- SLA values are reference lines, not pass/fail gates.
