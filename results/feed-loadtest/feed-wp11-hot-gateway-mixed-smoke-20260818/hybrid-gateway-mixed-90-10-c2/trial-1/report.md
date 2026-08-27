# Feed 压测报告：hybrid / gateway / mixed-90-10-c2

- Run ID：`feed-wp11-hot-gateway-mixed-smoke-20260818`
- 开始时间：2026-08-18T15:59:00+08:00
- 采样时长：2.1166394s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 90 | 90 | 0 | 0 | 42.52 | 1.064 | 2.664 | 3.214 | 6.428 | 6.428 |
| publish_total | 10 | 10 | 0 | 0 | 4.72 | 399.988 | 455.906 | 460.106 | 460.106 | 460.106 |
| publish_draft | 10 | 10 | 0 | 0 | 4.72 | 3.714 | 4.292 | 4.339 | 4.339 | 4.339 |
| publish_metadata | 10 | 10 | 0 | 0 | 4.72 | 138.752 | 171.140 | 172.077 | 172.077 | 172.077 |
| publish_confirm | 10 | 10 | 0 | 0 | 4.72 | 129.940 | 135.613 | 136.163 | 136.163 | 136.163 |
| publish_commit | 10 | 10 | 0 | 0 | 4.72 | 122.072 | 183.013 | 186.207 | 186.207 | 186.207 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 62 | 0.689 |
| counter | 20 | 0.222 |
| mysql | 7 | 0.078 |
| redis | 255 | 2.833 |
| relation | 20 | 0.222 |

- Cold compute：62（0.689 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2480 | 27.556 |
| merge_candidates | 2627 | 29.189 |
| redis_commands | 372 | 4.133 |
| redis_members | 2627 | 29.189 |
| redis_roundtrips | 62 | 0.689 |

| page cache source | requests |
|---|---:|
| l1_fresh | 28 |
| l2_fresh | 0 |
| miss | 62 |

- L1+L2 Fresh ratio：31.11%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 62 | 0.118 |
| counter | 20 | 0.483 |
| hydrate | 62 | 0.292 |
| inbox | 62 | 0.118 |
| merge_dedup | 62 | 0.008 |
| relation | 20 | 0.962 |
| route | 62 | 0.466 |
| total | 90 | 0.776 |

## Redis 本轮边界增量

- Commands：35703；input：3025217 bytes；output：1120121 bytes
- Hits/Misses：2931/764；run hit rate：79.32%
- Evicted/Rejected：0/0；ops/s max：14983；safety epoch：4023 -> 4043

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.138 |
| client:loadtest | cpu_percent_total | 2.215 |
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
| docker:zg-canal | cpu_percent | 0.110 |
| docker:zg-canal | memory_percent | 2.120 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.880 |
| docker:zg-es | memory_percent | 11.280 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 0.540 |
| docker:zg-etcd | memory_percent | 0.250 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 138.920 |
| docker:zg-kafka | memory_percent | 7.450 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 0.190 |
| docker:zg-zk | memory_percent | 0.990 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4199.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4199.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 2497.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.991 |
| process:counter | cpu_seconds_total | 54.219 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 38760448.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 7.184 |
| process:gateway | cpu_seconds_total | 0.750 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 43110400.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 51.083 |
| process:knowpost | cpu_seconds_total | 59.359 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 66158592.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 2.016 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 40660992.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.310 |
| process:search | cpu_seconds_total | 1.406 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 38309888.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.991 |
| process:user-storage | cpu_seconds_total | 1.500 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 38674432.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1899854.000 |
| redis | connected_clients | 29.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 4043.000 |
| redis | hit_rate | 0.702 |
| redis | keys | 595517.000 |
| redis | keyspace_hits | 88252.000 |
| redis | keyspace_misses | 37574.000 |
| redis | net_input_bytes | 136308690.000 |
| redis | net_output_bytes | 40053825.000 |
| redis | ops_per_sec | 14983.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 20327.000 |
| redis | used_memory_bytes | 94496760.000 |

## 停止施压后的恢复

- Kafka drain：5.2477619s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：10
- 测量前恢复：complete=true；耗时=5.0421775s；删除帖子/Outbox=0/0；safety epoch=4009
- 预热后恢复：complete=true；耗时=6.2644204s；删除帖子/Outbox=6/12；safety epoch=4022
- 测量后恢复：complete=true；耗时=4.9585222s；删除帖子/Outbox=10/20；safety epoch=4044

## 说明

- SLA values are reference lines, not pass/fail gates.
