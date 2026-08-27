# Feed 压测报告：hybrid / gateway / mixed-90-10-c4

- Run ID：`feed-wp11-hot-treatment-scout-gateway-mixed90-20260818`
- 开始时间：2026-08-18T22:43:09+08:00
- 采样时长：15.5197744s
- 并发：4
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 755 | 755 | 0 | 0 | 48.65 | 1.045 | 2.613 | 3.094 | 4.357 | 8.406 |
| publish_total | 83 | 83 | 0 | 0 | 5.35 | 735.831 | 808.216 | 813.096 | 870.507 | 870.507 |
| publish_draft | 83 | 83 | 0 | 0 | 5.35 | 3.748 | 4.371 | 4.436 | 5.830 | 5.830 |
| publish_metadata | 83 | 83 | 0 | 0 | 5.35 | 243.778 | 289.815 | 297.708 | 358.584 | 358.584 |
| publish_confirm | 83 | 83 | 0 | 0 | 5.35 | 237.550 | 285.169 | 290.720 | 294.334 | 294.334 |
| publish_commit | 83 | 83 | 0 | 0 | 5.35 | 227.699 | 296.686 | 314.909 | 327.145 | 327.145 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 368 | 0.487 |
| counter | 60 | 0.079 |
| mysql | 48 | 0.064 |
| redis | 1461 | 1.935 |
| relation | 60 | 0.079 |

- Cold compute：353（0.468 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 14120 | 18.702 |
| merge_candidates | 17525 | 23.212 |
| redis_commands | 2118 | 2.805 |
| redis_members | 17525 | 23.212 |
| redis_roundtrips | 353 | 0.468 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 387 |
| l2_fresh | 1 |
| miss | 367 |

- L1+L2 Fresh ratio：51.39%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 353 | 0.204 |
| counter | 60 | 0.544 |
| hydrate | 353 | 0.330 |
| inbox | 353 | 0.204 |
| merge_dedup | 353 | 0.012 |
| relation | 60 | 1.005 |
| route | 353 | 0.266 |
| total | 755 | 0.598 |

## Redis 本轮边界增量

- Commands：265488；input：22169010 bytes；output：7435220 bytes
- Hits/Misses：16246/3795；run hit rate：81.06%
- Evicted/Rejected：0/0；ops/s max：17526；safety epoch：7627 -> 7793

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.164 |
| client:loadtest | cpu_percent_total | 2.618 |
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
| docker:zg-canal | cpu_percent | 1.700 |
| docker:zg-canal | memory_percent | 3.970 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 21.720 |
| docker:zg-es | memory_percent | 11.750 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.870 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 162.390 |
| docker:zg-kafka | memory_percent | 7.430 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 0.110 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5616.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 5616.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 33946.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.195 |
| process:counter | cpu_seconds_total | 492.219 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 43704320.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 6.206 |
| process:gateway | cpu_seconds_total | 6.422 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 45522944.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 69.733 |
| process:knowpost | cpu_seconds_total | 654.562 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 76496896.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 15.922 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 47448064.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.552 |
| process:search | cpu_seconds_total | 14.359 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 43302912.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.877 |
| process:user-storage | cpu_seconds_total | 11.609 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 39911424.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18312789.000 |
| redis | connected_clients | 90.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 7793.000 |
| redis | hit_rate | 0.698 |
| redis | keys | 593542.000 |
| redis | keyspace_hits | 447739.000 |
| redis | keyspace_misses | 193532.000 |
| redis | net_input_bytes | 1337273523.000 |
| redis | net_output_bytes | 368525228.000 |
| redis | ops_per_sec | 17526.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44590.000 |
| redis | used_memory_bytes | 96920912.000 |

## 停止施压后的恢复

- Kafka drain：5.2149434s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：83
- 测量前恢复：complete=true；耗时=4.9934507s；删除帖子/Outbox=0/0；safety epoch=7585
- 预热后恢复：complete=true；耗时=6.3462029s；删除帖子/Outbox=20/40；safety epoch=7626
- 测量后恢复：complete=true；耗时=4.9076745s；删除帖子/Outbox=83/166；safety epoch=7794

## 说明

- SLA values are reference lines, not pass/fail gates.
