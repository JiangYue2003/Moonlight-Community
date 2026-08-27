# Feed 压测报告：hybrid / gateway / mixed-80-20-c8

- Run ID：`feed-wp11-hot-treatment-scout-gateway-mixed80-20260818`
- 开始时间：2026-08-18T22:48:23+08:00
- 采样时长：15.1462817s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 320 | 320 | 0 | 0 | 21.13 | 2.074 | 5.307 | 6.029 | 7.623 | 11.200 |
| publish_total | 80 | 80 | 0 | 0 | 5.28 | 1508.931 | 1606.329 | 1674.203 | 1690.188 | 1690.188 |
| publish_draft | 80 | 80 | 0 | 0 | 5.28 | 4.215 | 5.558 | 6.476 | 9.305 | 9.305 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.28 | 499.095 | 543.405 | 680.578 | 697.911 | 697.911 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.28 | 499.758 | 585.452 | 591.748 | 604.014 | 604.014 |
| publish_commit | 80 | 80 | 0 | 0 | 5.28 | 491.959 | 549.743 | 563.873 | 587.956 | 587.956 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 165 | 0.516 |
| counter | 55 | 0.172 |
| mysql | 34 | 0.106 |
| redis | 666 | 2.081 |
| relation | 55 | 0.172 |

- Cold compute：158（0.494 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6320 | 19.750 |
| merge_candidates | 7706 | 24.081 |
| redis_commands | 948 | 2.962 |
| redis_members | 7706 | 24.081 |
| redis_roundtrips | 158 | 0.494 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 155 |
| l2_fresh | 0 |
| miss | 165 |

- L1+L2 Fresh ratio：48.44%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 158 | 0.457 |
| counter | 55 | 0.827 |
| hydrate | 158 | 0.829 |
| inbox | 158 | 0.457 |
| merge_dedup | 158 | 0.004 |
| relation | 55 | 1.307 |
| route | 158 | 0.753 |
| total | 320 | 1.756 |

## Redis 本轮边界增量

- Commands：263005；input：21143673 bytes；output：5918185 bytes
- Hits/Misses：8167/2490；run hit rate：76.64%
- Evicted/Rejected：0/0；ops/s max：17788；safety epoch：8755 -> 8915

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.110 |
| client:loadtest | cpu_percent_total | 1.754 |
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
| docker:zg-canal | cpu_percent | 5.500 |
| docker:zg-canal | memory_percent | 4.370 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.920 |
| docker:zg-es | memory_percent | 11.780 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.470 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 161.410 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 42.790 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 6045.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 6045.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 43671.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.878 |
| process:counter | cpu_seconds_total | 499.641 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 44593152.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 5.413 |
| process:gateway | cpu_seconds_total | 10.062 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 46182400.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 69.601 |
| process:knowpost | cpu_seconds_total | 734.562 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 75472896.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.648 |
| process:relation | cpu_seconds_total | 17.219 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 47239168.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.791 |
| process:search | cpu_seconds_total | 15.344 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 42455040.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 6.190 |
| process:user-storage | cpu_seconds_total | 13.250 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 40964096.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 20279732.000 |
| redis | connected_clients | 95.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 8915.000 |
| redis | hit_rate | 0.700 |
| redis | keys | 596655.000 |
| redis | keyspace_hits | 566763.000 |
| redis | keyspace_misses | 243568.000 |
| redis | net_input_bytes | 1493588782.000 |
| redis | net_output_bytes | 414297725.000 |
| redis | ops_per_sec | 17788.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44905.000 |
| redis | used_memory_bytes | 98879680.000 |

## 停止施压后的恢复

- Kafka drain：6.6670433s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.9371096s；删除帖子/Outbox=0/0；safety epoch=8705
- 预热后恢复：complete=true；耗时=6.228048s；删除帖子/Outbox=24/48；safety epoch=8754
- 测量后恢复：complete=true；耗时=4.9354648s；删除帖子/Outbox=80/160；safety epoch=8916

## 说明

- SLA values are reference lines, not pass/fail gates.
