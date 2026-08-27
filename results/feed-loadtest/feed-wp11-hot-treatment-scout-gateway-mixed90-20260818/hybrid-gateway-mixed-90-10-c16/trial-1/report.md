# Feed 压测报告：hybrid / gateway / mixed-90-10-c16

- Run ID：`feed-wp11-hot-treatment-scout-gateway-mixed90-20260818`
- 开始时间：2026-08-18T22:44:44+08:00
- 采样时长：15.1791827s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 720 | 720 | 0 | 0 | 47.43 | 0.538 | 7.877 | 10.591 | 14.522 | 17.422 |
| publish_total | 80 | 80 | 0 | 0 | 5.27 | 2963.512 | 3213.654 | 3248.360 | 3312.019 | 3312.019 |
| publish_draft | 80 | 80 | 0 | 0 | 5.27 | 5.202 | 7.894 | 9.492 | 14.604 | 14.604 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.27 | 986.123 | 1230.414 | 1244.565 | 1286.977 | 1286.977 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.27 | 959.176 | 1083.612 | 1100.342 | 1149.187 | 1149.187 |
| publish_commit | 80 | 80 | 0 | 0 | 5.27 | 1026.005 | 1106.289 | 1113.514 | 1136.274 | 1136.274 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 129 | 0.179 |
| counter | 60 | 0.083 |
| mysql | 26 | 0.036 |
| redis | 426 | 0.592 |
| relation | 60 | 0.083 |

- Cold compute：100（0.139 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 4000 | 5.556 |
| merge_candidates | 4653 | 6.463 |
| redis_commands | 600 | 0.833 |
| redis_members | 4653 | 6.463 |
| redis_roundtrips | 100 | 0.139 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 591 |
| l2_fresh | 0 |
| miss | 129 |

- L1+L2 Fresh ratio：82.08%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 100 | 1.206 |
| counter | 60 | 1.547 |
| hydrate | 100 | 1.802 |
| inbox | 100 | 1.200 |
| merge_dedup | 100 | 0.000 |
| relation | 60 | 2.018 |
| route | 100 | 2.156 |
| total | 720 | 1.365 |

## Redis 本轮边界增量

- Commands：258213；input：20540399 bytes；output：5348586 bytes
- Hits/Misses：5626/2456；run hit rate：69.61%
- Evicted/Rejected：0/0；ops/s max：19053；safety epoch：8089 -> 8249

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.167 |
| client:loadtest | cpu_percent_total | 2.676 |
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
| docker:zg-canal | cpu_percent | 2.000 |
| docker:zg-canal | memory_percent | 4.130 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.800 |
| docker:zg-es | memory_percent | 11.750 |
| docker:zg-es | pids | 154.000 |
| docker:zg-etcd | cpu_percent | 3.970 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 186.860 |
| docker:zg-kafka | memory_percent | 7.580 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 0.190 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5791.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 5791.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 37863.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 5.421 |
| process:counter | cpu_seconds_total | 494.828 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 44556288.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 15.174 |
| process:gateway | cpu_seconds_total | 8.312 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 46583808.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 82.229 |
| process:knowpost | cpu_seconds_total | 687.219 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 76177408.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.098 |
| process:relation | cpu_seconds_total | 16.375 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 46116864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.776 |
| process:search | cpu_seconds_total | 14.609 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 43851776.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 4.792 |
| process:user-storage | cpu_seconds_total | 12.422 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 40480768.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19085626.000 |
| redis | connected_clients | 90.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 8249.000 |
| redis | hit_rate | 0.699 |
| redis | keys | 594634.000 |
| redis | keyspace_hits | 488987.000 |
| redis | keyspace_misses | 211141.000 |
| redis | net_input_bytes | 1398535181.000 |
| redis | net_output_bytes | 385918259.000 |
| redis | ops_per_sec | 19053.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44685.000 |
| redis | used_memory_bytes | 98198816.000 |

## 停止施压后的恢复

- Kafka drain：6.5268286s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.9183388s；删除帖子/Outbox=0/0；safety epoch=8023
- 预热后恢复：complete=true；耗时=6.2423117s；删除帖子/Outbox=32/64；safety epoch=8088
- 测量后恢复：complete=true；耗时=4.9376395s；删除帖子/Outbox=80/160；safety epoch=8250

## 说明

- SLA values are reference lines, not pass/fail gates.
