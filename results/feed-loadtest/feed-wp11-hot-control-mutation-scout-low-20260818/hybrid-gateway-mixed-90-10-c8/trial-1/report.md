# Feed 压测报告：hybrid / gateway / mixed-90-10-c8

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:08:14+08:00
- 采样时长：15.5885579s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 720 | 720 | 0 | 0 | 46.19 | 3.687 | 4.918 | 5.408 | 6.771 | 7.956 |
| publish_total | 80 | 80 | 0 | 0 | 5.13 | 1509.024 | 1610.774 | 1655.002 | 1693.235 | 1693.235 |
| publish_draft | 80 | 80 | 0 | 0 | 5.13 | 4.255 | 5.158 | 6.389 | 151.222 | 151.222 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.13 | 480.746 | 587.734 | 601.110 | 619.269 | 619.269 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.13 | 491.232 | 578.400 | 633.273 | 675.527 | 675.527 |
| publish_commit | 80 | 80 | 0 | 0 | 5.13 | 491.408 | 595.136 | 602.395 | 674.673 | 674.673 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 60 | 0.083 |
| mysql | 30 | 0.042 |
| redis | 2190 | 3.042 |
| relation | 720 | 1.000 |

- Cold compute：720（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 28800 | 40.000 |
| merge_candidates | 34905 | 48.479 |
| redis_commands | 4320 | 6.000 |
| redis_members | 34905 | 48.479 |
| redis_roundtrips | 1440 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 720 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 720 | 0.512 |
| counter | 60 | 0.811 |
| hydrate | 720 | 0.654 |
| inbox | 720 | 0.559 |
| merge_dedup | 720 | 0.011 |
| relation | 720 | 1.222 |
| route | 720 | 0.078 |
| total | 720 | 3.059 |

## Redis 本轮边界增量

- Commands：258625；input：21028732 bytes；output：10531363 bytes
- Hits/Misses：33427/3562；run hit rate：90.37%
- Evicted/Rejected：0/0；ops/s max：19667；safety epoch：11105 -> 11185

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.263 |
| client:loadtest | cpu_percent_total | 4.210 |
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
| docker:zg-canal | cpu_percent | 2.720 |
| docker:zg-canal | memory_percent | 4.670 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.830 |
| docker:zg-es | memory_percent | 11.880 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.400 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 298.390 |
| docker:zg-kafka | memory_percent | 8.050 |
| docker:zg-kafka | pids | 147.000 |
| docker:zg-zk | cpu_percent | 0.160 |
| docker:zg-zk | memory_percent | 1.290 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7641.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 7641.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 86101.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.646 |
| process:counter | cpu_seconds_total | 25.188 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45830144.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 4.651 |
| process:gateway | cpu_seconds_total | 3.328 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47824896.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 84.353 |
| process:knowpost | cpu_seconds_total | 272.219 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68956160.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.201 |
| process:relation | cpu_seconds_total | 7.984 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48742400.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.379 |
| process:search | cpu_seconds_total | 4.109 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43233280.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.101 |
| process:user-storage | cpu_seconds_total | 2.391 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41086976.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 27489070.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 11185.000 |
| redis | hit_rate | 0.738 |
| redis | keys | 593127.000 |
| redis | keyspace_hits | 1152454.000 |
| redis | keyspace_misses | 408786.000 |
| redis | net_input_bytes | 2057667068.000 |
| redis | net_output_bytes | 614418961.000 |
| redis | ops_per_sec | 19667.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46097.000 |
| redis | used_memory_bytes | 95513704.000 |

## 停止施压后的恢复

- Kafka drain：6.6299482s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.8852066s；删除帖子/Outbox=0/0；safety epoch=11083
- 预热后恢复：complete=true；耗时=7.6534209s；删除帖子/Outbox=20/40；safety epoch=11104
- 测量后恢复：complete=true；耗时=4.9018919s；删除帖子/Outbox=80/160；safety epoch=11186

## 说明

- SLA values are reference lines, not pass/fail gates.
