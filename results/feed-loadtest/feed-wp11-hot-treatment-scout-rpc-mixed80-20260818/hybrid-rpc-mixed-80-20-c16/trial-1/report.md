# Feed 压测报告：hybrid / rpc / mixed-80-20-c16

- Run ID：`feed-wp11-hot-treatment-scout-rpc-mixed80-20260818`
- 开始时间：2026-08-18T22:35:32+08:00
- 采样时长：15.4267023s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 324 | 324 | 0 | 0 | 21.00 | 0.521 | 10.808 | 12.568 | 15.494 | 18.445 |
| publish_total | 81 | 81 | 0 | 0 | 5.25 | 2994.718 | 3164.525 | 3175.751 | 3253.184 | 3253.184 |
| publish_draft | 81 | 81 | 0 | 0 | 5.25 | 4.812 | 6.384 | 6.844 | 8.069 | 8.069 |
| publish_metadata | 81 | 81 | 0 | 0 | 5.25 | 988.129 | 1164.614 | 1169.946 | 1210.638 | 1210.638 |
| publish_confirm | 81 | 81 | 0 | 0 | 5.25 | 960.634 | 1136.223 | 1148.640 | 1200.643 | 1200.643 |
| publish_commit | 81 | 81 | 0 | 0 | 5.25 | 1007.946 | 1089.071 | 1099.462 | 1155.937 | 1155.937 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 114 | 0.352 |
| counter | 56 | 0.173 |
| mysql | 23 | 0.071 |
| redis | 411 | 1.269 |
| relation | 56 | 0.173 |

- Cold compute：97（0.299 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3880 | 11.975 |
| merge_candidates | 4631 | 14.293 |
| redis_commands | 582 | 1.796 |
| redis_members | 4631 | 14.293 |
| redis_roundtrips | 97 | 0.299 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 210 |
| l2_fresh | 0 |
| miss | 114 |

- L1+L2 Fresh ratio：64.81%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 97 | 1.272 |
| counter | 56 | 1.385 |
| hydrate | 97 | 1.832 |
| inbox | 97 | 1.272 |
| merge_dedup | 97 | 0.017 |
| relation | 56 | 2.257 |
| route | 97 | 2.117 |
| total | 324 | 2.910 |

## Redis 本轮边界增量

- Commands：263811；input：20978769 bytes；output：5424143 bytes
- Hits/Misses：5407/2409；run hit rate：69.18%
- Evicted/Rejected：0/0；ops/s max：19105；safety epoch：6533 -> 6695

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.146 |
| client:loadtest | cpu_percent_total | 2.330 |
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
| docker:zg-canal | cpu_percent | 1.670 |
| docker:zg-canal | memory_percent | 3.580 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.950 |
| docker:zg-es | memory_percent | 11.660 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 3.930 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 141.170 |
| docker:zg-kafka | memory_percent | 7.440 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 9.950 |
| docker:zg-zk | memory_percent | 1.410 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 5206.000 |
| kafka | lag_max | 14.000 |
| kafka | lag_total | 14.000 |
| kafka | log_end_offset_total | 5206.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25015.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.890 |
| process:counter | cpu_seconds_total | 482.766 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 43560960.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.841 |
| process:gateway | cpu_seconds_total | 4.250 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 37507072.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 76.864 |
| process:knowpost | cpu_seconds_total | 575.594 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 76464128.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 15.047 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 45625344.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.441 |
| process:search | cpu_seconds_total | 13.156 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 42504192.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 10.562 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35155968.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 16395695.000 |
| redis | connected_clients | 89.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 6695.000 |
| redis | hit_rate | 0.706 |
| redis | keys | 595782.000 |
| redis | keyspace_hits | 374622.000 |
| redis | keyspace_misses | 156035.000 |
| redis | net_input_bytes | 1188410437.000 |
| redis | net_output_bytes | 329302099.000 |
| redis | ops_per_sec | 19105.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44134.000 |
| redis | used_memory_bytes | 98657856.000 |

## 停止施压后的恢复

- Kafka drain：6.5119403s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：81
- 测量前恢复：complete=true；耗时=4.8991058s；删除帖子/Outbox=0/0；safety epoch=6467
- 预热后恢复：complete=true；耗时=6.2738151s；删除帖子/Outbox=32/64；safety epoch=6532
- 测量后恢复：complete=true；耗时=4.8649193s；删除帖子/Outbox=81/162；safety epoch=6696

## 说明

- SLA values are reference lines, not pass/fail gates.
