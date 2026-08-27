# Feed 压测报告：hybrid / gateway / mixed-80-20-c2

- Run ID：`feed-wp11-hot-treatment-scout-gateway-mixed80-20260818`
- 开始时间：2026-08-18T22:46:50+08:00
- 采样时长：15.1170326s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 280 | 280 | 0 | 0 | 18.52 | 1.558 | 2.821 | 3.673 | 4.800 | 6.680 |
| publish_total | 70 | 70 | 0 | 0 | 4.63 | 411.348 | 504.171 | 577.935 | 683.158 | 683.158 |
| publish_draft | 70 | 70 | 0 | 0 | 4.63 | 3.704 | 4.258 | 4.332 | 5.379 | 5.379 |
| publish_metadata | 70 | 70 | 0 | 0 | 4.63 | 135.456 | 173.446 | 184.689 | 189.057 | 189.057 |
| publish_confirm | 70 | 70 | 0 | 0 | 4.63 | 134.990 | 165.248 | 174.823 | 375.539 | 375.539 |
| publish_commit | 70 | 70 | 0 | 0 | 4.63 | 132.802 | 168.030 | 184.309 | 227.229 | 227.229 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 234 | 0.836 |
| counter | 56 | 0.200 |
| mysql | 51 | 0.182 |
| redis | 979 | 3.496 |
| relation | 56 | 0.200 |

- Cold compute：232（0.829 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9280 | 33.143 |
| merge_candidates | 11394 | 40.693 |
| redis_commands | 1392 | 4.971 |
| redis_members | 11394 | 40.693 |
| redis_roundtrips | 232 | 0.829 |

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 46 |
| l2_fresh | 0 |
| miss | 234 |

- L1+L2 Fresh ratio：16.43%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 232 | 0.132 |
| counter | 56 | 0.426 |
| hydrate | 232 | 0.367 |
| inbox | 232 | 0.132 |
| merge_dedup | 232 | 0.007 |
| relation | 56 | 0.954 |
| route | 232 | 0.345 |
| total | 280 | 0.948 |

## Redis 本轮边界增量

- Commands：228753；input：18930198 bytes；output：5718869 bytes
- Hits/Misses：10277/3664；run hit rate：73.72%
- Evicted/Rejected：0/0；ops/s max：16094；safety epoch：8351 -> 8491

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.149 |
| client:loadtest | cpu_percent_total | 2.377 |
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
| docker:zg-canal | cpu_percent | 1.630 |
| docker:zg-canal | memory_percent | 4.220 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 6.380 |
| docker:zg-es | memory_percent | 11.760 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.830 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 147.200 |
| docker:zg-kafka | memory_percent | 7.580 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 40.080 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5881.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 5881.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 40018.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.704 |
| process:counter | cpu_seconds_total | 497.531 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 44134400.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 4.791 |
| process:gateway | cpu_seconds_total | 8.844 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 45916160.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 64.338 |
| process:knowpost | cpu_seconds_total | 705.984 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74768384.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.552 |
| process:relation | cpu_seconds_total | 16.625 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 47411200.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.799 |
| process:search | cpu_seconds_total | 15.000 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 42344448.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.396 |
| process:user-storage | cpu_seconds_total | 12.750 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 40214528.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19542286.000 |
| redis | connected_clients | 95.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 8491.000 |
| redis | hit_rate | 0.698 |
| redis | keys | 595150.000 |
| redis | keyspace_hits | 521600.000 |
| redis | keyspace_misses | 225837.000 |
| redis | net_input_bytes | 1434678178.000 |
| redis | net_output_bytes | 396756607.000 |
| redis | ops_per_sec | 16094.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44810.000 |
| redis | used_memory_bytes | 97647392.000 |

## 停止施压后的恢复

- Kafka drain：5.2176515s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：70
- 测量前恢复：complete=true；耗时=4.8798465s；删除帖子/Outbox=0/0；safety epoch=8317
- 预热后恢复：complete=true；耗时=6.365121s；删除帖子/Outbox=16/32；safety epoch=8350
- 测量后恢复：complete=true；耗时=4.9333977s；删除帖子/Outbox=70/140；safety epoch=8492

## 说明

- SLA values are reference lines, not pass/fail gates.
