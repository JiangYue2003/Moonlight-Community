# Feed 压测报告：hybrid / gateway / mixed-90-10-c16

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:09:06+08:00
- 采样时长：15.6869157s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 720 | 720 | 0 | 0 | 45.90 | 7.618 | 10.259 | 11.325 | 13.339 | 16.080 |
| publish_total | 80 | 80 | 0 | 0 | 5.10 | 3000.674 | 3235.773 | 3264.440 | 3346.674 | 3346.674 |
| publish_draft | 80 | 80 | 0 | 0 | 5.10 | 5.361 | 6.149 | 6.657 | 7.496 | 7.496 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.10 | 1003.374 | 1077.746 | 1097.030 | 1114.232 | 1114.232 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.10 | 999.677 | 1217.897 | 1246.105 | 1349.930 | 1349.930 |
| publish_commit | 80 | 80 | 0 | 0 | 5.10 | 990.114 | 1174.870 | 1186.854 | 1289.467 | 1289.467 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 61 | 0.085 |
| mysql | 40 | 0.056 |
| redis | 2200 | 3.056 |
| relation | 720 | 1.000 |

- Cold compute：720（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 28800 | 40.000 |
| merge_candidates | 34190 | 47.486 |
| redis_commands | 4320 | 6.000 |
| redis_members | 34190 | 47.486 |
| redis_roundtrips | 1440 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 720 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 720 | 1.519 |
| counter | 61 | 1.680 |
| hydrate | 720 | 1.730 |
| inbox | 720 | 1.499 |
| merge_dedup | 720 | 0.006 |
| relation | 720 | 2.050 |
| route | 720 | 0.153 |
| total | 720 | 6.984 |

## Redis 本轮边界增量

- Commands：258810；input：21053754 bytes；output：10527294 bytes
- Hits/Misses：33219/3813；run hit rate：89.70%
- Evicted/Rejected：0/0；ops/s max：18443；safety epoch：11218 -> 11298

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.137 |
| client:loadtest | cpu_percent_total | 2.191 |
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
| docker:zg-canal | cpu_percent | 3.290 |
| docker:zg-canal | memory_percent | 4.670 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.960 |
| docker:zg-es | memory_percent | 11.880 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.990 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 148.900 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 5.470 |
| docker:zg-zk | memory_percent | 1.440 |
| docker:zg-zk | pids | 101.000 |
| kafka | current_offset_total | 7727.000 |
| kafka | lag_max | 13.000 |
| kafka | lag_total | 13.000 |
| kafka | log_end_offset_total | 7727.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 88948.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 26.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 5.422 |
| process:counter | cpu_seconds_total | 26.625 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46440448.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 11.109 |
| process:gateway | cpu_seconds_total | 4.422 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 48668672.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 100.787 |
| process:knowpost | cpu_seconds_total | 286.406 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 70418432.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.746 |
| process:relation | cpu_seconds_total | 8.922 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48955392.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.793 |
| process:search | cpu_seconds_total | 4.312 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43372544.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 6.202 |
| process:user-storage | cpu_seconds_total | 2.859 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41848832.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 27869253.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 11298.000 |
| redis | hit_rate | 0.742 |
| redis | keys | 593183.000 |
| redis | keyspace_hits | 1207613.000 |
| redis | keyspace_misses | 419525.000 |
| redis | net_input_bytes | 2088154794.000 |
| redis | net_output_bytes | 629548287.000 |
| redis | ops_per_sec | 18443.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46148.000 |
| redis | used_memory_bytes | 95770176.000 |

## 停止施压后的恢复

- Kafka drain：6.5030495s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.8952876s；删除帖子/Outbox=0/0；safety epoch=11187
- 预热后恢复：complete=true；耗时=7.6457286s；删除帖子/Outbox=29/58；safety epoch=11217
- 测量后恢复：complete=true；耗时=4.9819094s；删除帖子/Outbox=80/160；safety epoch=11299

## 说明

- SLA values are reference lines, not pass/fail gates.
