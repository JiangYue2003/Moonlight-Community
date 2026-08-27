# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-mixed-80-20-smoke-20260818`
- 开始时间：2026-08-18T15:54:52+08:00
- 采样时长：3.0145815s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 64 | 64 | 0 | 0 | 21.23 | 1.047 | 3.126 | 6.268 | 6.268 | 6.268 |
| publish_total | 16 | 16 | 0 | 0 | 5.31 | 746.738 | 756.725 | 759.897 | 759.897 | 759.897 |
| publish_draft | 16 | 16 | 0 | 0 | 5.31 | 4.190 | 5.279 | 5.355 | 5.355 | 5.355 |
| publish_metadata | 16 | 16 | 0 | 0 | 5.31 | 233.420 | 246.679 | 248.063 | 248.063 | 248.063 |
| publish_confirm | 16 | 16 | 0 | 0 | 5.31 | 237.896 | 289.650 | 290.210 | 290.210 | 290.210 |
| publish_commit | 16 | 16 | 0 | 0 | 5.31 | 261.899 | 298.018 | 301.193 | 301.193 | 301.193 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 47 | 0.734 |
| counter | 17 | 0.266 |
| mysql | 12 | 0.188 |
| redis | 200 | 3.125 |
| relation | 17 | 0.266 |

- Cold compute：47（0.734 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1880 | 29.375 |
| merge_candidates | 1967 | 30.734 |
| redis_commands | 282 | 4.406 |
| redis_members | 1967 | 30.734 |
| redis_roundtrips | 47 | 0.734 |

| page cache source | requests |
|---|---:|
| l1_fresh | 17 |
| l2_fresh | 0 |
| miss | 47 |

- L1+L2 Fresh ratio：26.56%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 47 | 0.190 |
| counter | 17 | 0.663 |
| hydrate | 47 | 0.424 |
| inbox | 47 | 0.190 |
| merge_dedup | 47 | 0.000 |
| relation | 17 | 1.147 |
| route | 47 | 0.655 |
| total | 64 | 1.302 |

## Redis 本轮边界增量

- Commands：52002；input：4296457 bytes；output：1260859 bytes
- Hits/Misses：2055/955；run hit rate：68.27%
- Evicted/Rejected：0/0；ops/s max：16941；safety epoch：3975 -> 4007

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.130 |
| client:loadtest | cpu_percent_total | 2.073 |
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
| docker:zg-canal | cpu_percent | 0.620 |
| docker:zg-canal | memory_percent | 2.110 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.260 |
| docker:zg-es | memory_percent | 11.240 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 0.480 |
| docker:zg-etcd | memory_percent | 0.250 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 138.930 |
| docker:zg-kafka | memory_percent | 7.440 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 0.100 |
| docker:zg-zk | memory_percent | 0.990 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4187.000 |
| kafka | lag_max | 3.000 |
| kafka | lag_total | 3.000 |
| kafka | log_end_offset_total | 4188.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 2128.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 2.439 |
| process:counter | cpu_seconds_total | 49.844 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 38674432.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.562 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 36720640.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 52.854 |
| process:knowpost | cpu_seconds_total | 53.531 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 64036864.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.626 |
| process:relation | cpu_seconds_total | 1.859 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 41205760.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.219 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 38486016.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.297 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35491840.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1712397.000 |
| redis | connected_clients | 29.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 4007.000 |
| redis | hit_rate | 0.698 |
| redis | keys | 596065.000 |
| redis | keyspace_hits | 72879.000 |
| redis | keyspace_misses | 31602.000 |
| redis | net_input_bytes | 122534403.000 |
| redis | net_output_bytes | 35511834.000 |
| redis | ops_per_sec | 16941.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 20076.000 |
| redis | used_memory_bytes | 94529664.000 |

## 停止施压后的恢复

- Kafka drain：1.0728714s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：16
- 测量前恢复：complete=true；耗时=4.9676352s；删除帖子/Outbox=0/0；safety epoch=3957
- 预热后恢复：complete=true；耗时=7.5729096s；删除帖子/Outbox=8/16；safety epoch=3974
- 测量后恢复：complete=true；耗时=4.9582612s；删除帖子/Outbox=16/32；safety epoch=4008

## 说明

- SLA values are reference lines, not pass/fail gates.
