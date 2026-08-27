# Feed 压测报告：hybrid / gateway / mixed-90-10-c4

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:59:09+08:00
- 采样时长：1m0.2484238s
- 并发：4
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2124 | 2124 | 0 | 0 | 35.25 | 2.695 | 3.920 | 4.656 | 5.992 | 8.877 |
| publish_total | 236 | 236 | 0 | 0 | 3.92 | 975.573 | 1193.379 | 1273.300 | 1385.686 | 1422.710 |
| publish_draft | 236 | 236 | 0 | 0 | 3.92 | 4.231 | 4.923 | 5.495 | 77.090 | 105.117 |
| publish_metadata | 236 | 236 | 0 | 0 | 3.92 | 324.565 | 409.144 | 444.711 | 475.069 | 489.701 |
| publish_confirm | 236 | 236 | 0 | 0 | 3.92 | 327.023 | 425.069 | 456.737 | 539.083 | 549.501 |
| publish_commit | 236 | 236 | 0 | 0 | 3.92 | 310.015 | 404.058 | 445.337 | 531.018 | 543.774 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 216 | 0.102 |
| mysql | 111 | 0.052 |
| redis | 6483 | 3.052 |
| relation | 2124 | 1.000 |

- Cold compute：2124（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 84960 | 40.000 |
| merge_candidates | 138847 | 65.371 |
| redis_commands | 12744 | 6.000 |
| redis_members | 140494 | 66.146 |
| redis_roundtrips | 4248 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2124 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2124 | 0.310 |
| counter | 216 | 0.591 |
| hydrate | 2124 | 0.394 |
| inbox | 2124 | 0.275 |
| merge_dedup | 2124 | 0.013 |
| relation | 2124 | 1.056 |
| route | 2124 | 0.067 |
| total | 2124 | 2.140 |

## Redis 本轮边界增量

- Commands：797459；input：64693687 bytes；output：33152779 bytes
- Hits/Misses：102137/8390；run hit rate：92.41%
- Evicted/Rejected：0/0；ops/s max：15430；safety epoch：20703 -> 20939

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.188 |
| client:loadtest | cpu_percent_total | 3.008 |
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
| docker:zg-canal | memory_percent | 4.990 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.830 |
| docker:zg-es | memory_percent | 12.260 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.080 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 157.710 |
| docker:zg-kafka | memory_percent | 7.900 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 48.670 |
| docker:zg-zk | memory_percent | 1.500 |
| docker:zg-zk | pids | 102.000 |
| kafka | current_offset_total | 15213.000 |
| kafka | lag_max | 4.000 |
| kafka | lag_total | 4.000 |
| kafka | log_end_offset_total | 15213.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 280805.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 7.743 |
| process:counter | cpu_seconds_total | 108.906 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44212224.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 10.847 |
| process:gateway | cpu_seconds_total | 20.641 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46288896.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 71.968 |
| process:knowpost | cpu_seconds_total | 1556.219 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 67170304.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 8.524 |
| process:relation | cpu_seconds_total | 49.891 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47759360.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 4.644 |
| process:search | cpu_seconds_total | 24.375 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42274816.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.958 |
| process:user-storage | cpu_seconds_total | 12.766 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40292352.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 59671623.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 20939.000 |
| redis | hit_rate | 0.799 |
| redis | keys | 597642.000 |
| redis | keyspace_hits | 3325102.000 |
| redis | keyspace_misses | 836895.000 |
| redis | net_input_bytes | 4603847298.000 |
| redis | net_output_bytes | 1535283005.000 |
| redis | ops_per_sec | 15430.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49195.000 |
| redis | used_memory_bytes | 96440976.000 |

## 停止施压后的恢复

- Kafka drain：5.4617689s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：236
- 测量前恢复：complete=true；耗时=4.9706713s；删除帖子/Outbox=0/0；safety epoch=20649
- 预热后恢复：complete=true；耗时=6.4422879s；删除帖子/Outbox=52/104；safety epoch=20702
- 测量后恢复：complete=true；耗时=5.1035738s；删除帖子/Outbox=236/472；safety epoch=20940

## 说明

- SLA values are reference lines, not pass/fail gates.
