# Feed 压测报告：hybrid / gateway / mixed-80-20-c16

- Run ID：`feed-wp11-hot-treatment-scout-gateway-mixed80-20260818`
- 开始时间：2026-08-18T22:49:13+08:00
- 采样时长：15.1399494s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 320 | 320 | 0 | 0 | 21.14 | 1.040 | 11.521 | 13.404 | 15.264 | 16.606 |
| publish_total | 80 | 80 | 0 | 0 | 5.28 | 3036.359 | 3140.535 | 3149.778 | 3204.455 | 3204.455 |
| publish_draft | 80 | 80 | 0 | 0 | 5.28 | 5.325 | 8.639 | 11.497 | 31.610 | 31.610 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.28 | 1022.155 | 1071.786 | 1084.054 | 1110.381 | 1110.381 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.28 | 946.617 | 1261.459 | 1277.263 | 1312.151 | 1312.151 |
| publish_commit | 80 | 80 | 0 | 0 | 5.28 | 1035.028 | 1104.250 | 1110.145 | 1135.360 | 1135.360 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 111 | 0.347 |
| counter | 56 | 0.175 |
| mysql | 31 | 0.097 |
| redis | 403 | 1.259 |
| relation | 56 | 0.175 |

- Cold compute：93（0.291 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3720 | 11.625 |
| merge_candidates | 4387 | 13.709 |
| redis_commands | 558 | 1.744 |
| redis_members | 4387 | 13.709 |
| redis_roundtrips | 93 | 0.291 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 209 |
| l2_fresh | 0 |
| miss | 111 |

- L1+L2 Fresh ratio：65.31%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 93 | 1.198 |
| counter | 56 | 1.276 |
| hydrate | 93 | 2.214 |
| inbox | 93 | 1.198 |
| merge_dedup | 93 | 0.005 |
| relation | 56 | 2.364 |
| route | 93 | 2.207 |
| total | 320 | 2.847 |

## Redis 本轮边界增量

- Commands：262459；input：20876247 bytes；output：5357415 bytes
- Hits/Misses：5116/2484；run hit rate：67.32%
- Evicted/Rejected：0/0；ops/s max：19170；safety epoch：8983 -> 9143

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.161 |
| client:loadtest | cpu_percent_total | 2.580 |
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
| docker:zg-canal | cpu_percent | 2.560 |
| docker:zg-canal | memory_percent | 4.450 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.820 |
| docker:zg-es | memory_percent | 11.780 |
| docker:zg-es | pids | 154.000 |
| docker:zg-etcd | cpu_percent | 3.890 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 132.210 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 0.210 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 6132.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 6132.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 45632.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 26.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.096 |
| process:counter | cpu_seconds_total | 500.594 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 45137920.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 7.958 |
| process:gateway | cpu_seconds_total | 10.734 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 47370240.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 91.315 |
| process:knowpost | cpu_seconds_total | 749.422 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 75431936.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.095 |
| process:relation | cpu_seconds_total | 17.484 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 48623616.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.387 |
| process:search | cpu_seconds_total | 15.641 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 43126784.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.322 |
| process:user-storage | cpu_seconds_total | 13.391 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 40800256.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 20674215.000 |
| redis | connected_clients | 95.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 9143.000 |
| redis | hit_rate | 0.699 |
| redis | keys | 596446.000 |
| redis | keyspace_hits | 583836.000 |
| redis | keyspace_misses | 252196.000 |
| redis | net_input_bytes | 1524585437.000 |
| redis | net_output_bytes | 422523017.000 |
| redis | ops_per_sec | 19170.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44955.000 |
| redis | used_memory_bytes | 99095872.000 |

## 停止施压后的恢复

- Kafka drain：6.5033947s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.9327007s；删除帖子/Outbox=0/0；safety epoch=8917
- 预热后恢复：complete=true；耗时=7.7089966s；删除帖子/Outbox=32/64；safety epoch=8982
- 测量后恢复：complete=true；耗时=4.8677997s；删除帖子/Outbox=80/160；safety epoch=9144

## 说明

- SLA values are reference lines, not pass/fail gates.
