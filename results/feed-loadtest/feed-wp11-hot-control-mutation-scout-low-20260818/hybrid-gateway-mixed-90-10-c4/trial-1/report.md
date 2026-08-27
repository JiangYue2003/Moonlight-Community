# Feed 压测报告：hybrid / gateway / mixed-90-10-c4

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:07:26+08:00
- 采样时长：15.3806748s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 684 | 684 | 0 | 0 | 44.47 | 2.222 | 3.296 | 3.698 | 4.768 | 5.772 |
| publish_total | 76 | 76 | 0 | 0 | 4.94 | 757.314 | 974.361 | 987.790 | 1000.087 | 1000.087 |
| publish_draft | 76 | 76 | 0 | 0 | 4.94 | 3.810 | 4.465 | 4.907 | 5.198 | 5.198 |
| publish_metadata | 76 | 76 | 0 | 0 | 4.94 | 245.817 | 304.067 | 361.664 | 440.216 | 440.216 |
| publish_confirm | 76 | 76 | 0 | 0 | 4.94 | 249.898 | 345.088 | 361.173 | 385.847 | 385.847 |
| publish_commit | 76 | 76 | 0 | 0 | 4.94 | 255.275 | 336.294 | 340.932 | 358.129 | 358.129 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 60 | 0.088 |
| mysql | 36 | 0.053 |
| redis | 2088 | 3.053 |
| relation | 684 | 1.000 |

- Cold compute：684（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 27360 | 40.000 |
| merge_candidates | 33337 | 48.738 |
| redis_commands | 4104 | 6.000 |
| redis_members | 33337 | 48.738 |
| redis_roundtrips | 1368 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 684 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 684 | 0.245 |
| counter | 60 | 0.511 |
| hydrate | 684 | 0.337 |
| inbox | 684 | 0.260 |
| merge_dedup | 684 | 0.006 |
| relation | 684 | 0.885 |
| route | 684 | 0.050 |
| total | 684 | 1.802 |

## Redis 本轮边界增量

- Commands：245653；input：19971434 bytes；output：10004296 bytes
- Hits/Misses：31958/3263；run hit rate：90.74%
- Evicted/Rejected：0/0；ops/s max：17167；safety epoch：11005 -> 11081

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.197 |
| client:loadtest | cpu_percent_total | 3.149 |
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
| docker:zg-canal | cpu_percent | 3.340 |
| docker:zg-canal | memory_percent | 4.670 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.590 |
| docker:zg-es | memory_percent | 11.880 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.830 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 236.940 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 28.050 |
| docker:zg-zk | memory_percent | 1.520 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 7564.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 7564.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 83512.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 5.431 |
| process:counter | cpu_seconds_total | 24.016 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45719552.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 5.422 |
| process:gateway | cpu_seconds_total | 2.672 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47140864.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 78.364 |
| process:knowpost | cpu_seconds_total | 256.641 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68476928.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.964 |
| process:relation | cpu_seconds_total | 7.328 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48136192.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 3.781 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42782720.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 12.380 |
| process:user-storage | cpu_seconds_total | 2.031 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41136128.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 27138034.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 11081.000 |
| redis | hit_rate | 0.734 |
| redis | keys | 593046.000 |
| redis | keyspace_hits | 1100338.000 |
| redis | keyspace_misses | 399027.000 |
| redis | net_input_bytes | 2029590025.000 |
| redis | net_output_bytes | 600415366.000 |
| redis | ops_per_sec | 17167.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46046.000 |
| redis | used_memory_bytes | 95249472.000 |

## 停止施压后的恢复

- Kafka drain：5.1732025s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：76
- 测量前恢复：complete=true；耗时=4.9525009s；删除帖子/Outbox=0/0；safety epoch=10983
- 预热后恢复：complete=true；耗时=6.2315614s；删除帖子/Outbox=20/40；safety epoch=11004
- 测量后恢复：complete=true；耗时=4.9294289s；删除帖子/Outbox=76/152；safety epoch=11082

## 说明

- SLA values are reference lines, not pass/fail gates.
