# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:00:56+08:00
- 采样时长：15.1589999s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 320 | 320 | 0 | 0 | 21.11 | 1.670 | 2.713 | 3.186 | 5.288 | 5.814 |
| publish_total | 80 | 80 | 0 | 0 | 5.28 | 762.875 | 822.492 | 827.235 | 831.685 | 831.685 |
| publish_draft | 80 | 80 | 0 | 0 | 5.28 | 3.654 | 4.196 | 4.407 | 6.384 | 6.384 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.28 | 242.694 | 302.717 | 319.098 | 354.709 | 354.709 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.28 | 239.395 | 298.837 | 307.638 | 318.392 | 318.392 |
| publish_commit | 80 | 80 | 0 | 0 | 5.28 | 254.561 | 313.729 | 318.713 | 324.704 | 324.704 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 58 | 0.181 |
| mysql | 40 | 0.125 |
| redis | 1000 | 3.125 |
| relation | 320 | 1.000 |

- Cold compute：320（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 12800 | 40.000 |
| merge_candidates | 15949 | 49.841 |
| redis_commands | 1920 | 6.000 |
| redis_members | 15949 | 49.841 |
| redis_roundtrips | 640 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 320 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 320 | 0.186 |
| counter | 58 | 0.568 |
| hydrate | 320 | 0.335 |
| inbox | 320 | 0.219 |
| merge_dedup | 320 | 0.014 |
| relation | 320 | 0.894 |
| route | 320 | 0.105 |
| total | 320 | 1.770 |

## Redis 本轮边界增量

- Commands：253366；input：20124854 bytes；output：7161434 bytes
- Hits/Misses：15852/2295；run hit rate：87.35%
- Evicted/Rejected：0/0；ops/s max：17836；safety epoch：10134 -> 10214

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.142 |
| client:loadtest | cpu_percent_total | 2.268 |
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
| docker:zg-canal | cpu_percent | 2.730 |
| docker:zg-canal | memory_percent | 4.660 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.940 |
| docker:zg-es | memory_percent | 11.860 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.210 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 167.330 |
| docker:zg-kafka | memory_percent | 7.550 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 40.680 |
| docker:zg-zk | memory_percent | 1.290 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 6923.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 6923.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 67136.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.877 |
| process:counter | cpu_seconds_total | 14.672 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44244992.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.141 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 37421056.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 85.189 |
| process:knowpost | cpu_seconds_total | 144.719 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66883584.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.648 |
| process:relation | cpu_seconds_total | 4.484 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47939584.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.796 |
| process:search | cpu_seconds_total | 1.844 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42602496.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.203 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35250176.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 24285240.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 10214.000 |
| redis | hit_rate | 0.727 |
| redis | keys | 592480.000 |
| redis | keyspace_hits | 897087.000 |
| redis | keyspace_misses | 336044.000 |
| redis | net_input_bytes | 1806934218.000 |
| redis | net_output_bytes | 526286258.000 |
| redis | ops_per_sec | 17836.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45657.000 |
| redis | used_memory_bytes | 95415736.000 |

## 停止施压后的恢复

- Kafka drain：6.6261249s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.9451532s；删除帖子/Outbox=0/0；safety epoch=10112
- 预热后恢复：complete=true；耗时=7.6091419s；删除帖子/Outbox=20/40；safety epoch=10133
- 测量后恢复：complete=true；耗时=4.8495946s；删除帖子/Outbox=80/160；safety epoch=10215

## 说明

- SLA values are reference lines, not pass/fail gates.
