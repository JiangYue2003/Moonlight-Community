# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-publish-c2`
- 开始时间：2026-08-27T05:41:36+08:00
- 采样时长：1m0.1240517s
- 并发：2
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 410 | 410 | 0 | 0 | 6.82 | 285.997 | 342.612 | 362.586 | 460.769 | 508.509 |
| publish_draft | 410 | 410 | 0 | 0 | 6.82 | 3.851 | 5.131 | 5.590 | 8.471 | 16.176 |
| publish_metadata | 410 | 410 | 0 | 0 | 6.82 | 91.055 | 112.505 | 123.546 | 151.999 | 170.824 |
| publish_confirm | 410 | 410 | 0 | 0 | 6.82 | 92.513 | 118.859 | 132.975 | 154.708 | 307.963 |
| publish_commit | 410 | 410 | 0 | 0 | 6.82 | 93.464 | 114.592 | 125.337 | 137.067 | 151.138 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 0 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 0 | 0.000 |
| relation | 0 | 0.000 |

- Cold compute：0（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：941091；input：74246658 bytes；output：16638397 bytes
- Hits/Misses：1876/7372；run hit rate：20.29%
- Evicted/Rejected：0/0；ops/s max：17957；safety epoch：24653 -> 25063

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.125 |
| client:loadtest | cpu_percent_total | 2.001 |
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
| docker-state:zg-kafka | restart_count | 4.000 |
| docker-state:zg-kafka | running | 1.000 |
| docker-state:zg-zk | health_configured | 1.000 |
| docker-state:zg-zk | healthy | 1.000 |
| docker-state:zg-zk | restart_count | 0.000 |
| docker-state:zg-zk | running | 1.000 |
| docker:zg-canal | cpu_percent | 8.580 |
| docker:zg-canal | memory_percent | 4.000 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 7.320 |
| docker:zg-es | memory_percent | 12.530 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.230 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 161.850 |
| docker:zg-kafka | memory_percent | 7.430 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 44.780 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 18431.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 18431.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 75827.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.237 |
| process:counter | cpu_seconds_total | 14.000 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 43913216.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 0.094 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 37912576.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 75.069 |
| process:knowpost | cpu_seconds_total | 100.391 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 62246912.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 16.625 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 54464512.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.101 |
| process:search | cpu_seconds_total | 2.203 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 42848256.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 0.984 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 41828352.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2594458.000 |
| redis | connected_clients | 75.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 25063.000 |
| redis | hit_rate | 0.661 |
| redis | keys | 567952.000 |
| redis | keyspace_hits | 49618.000 |
| redis | keyspace_misses | 31833.000 |
| redis | net_input_bytes | 198533522.000 |
| redis | net_output_bytes | 44775824.000 |
| redis | ops_per_sec | 17957.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 11114.000 |
| redis | used_memory_bytes | 80046416.000 |

## 停止施压后的恢复

- Kafka drain：5.3928906s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：410
- 测量前恢复：complete=true；耗时=5.0149506s；删除帖子/Outbox=0/0；safety epoch=24579
- 预热后恢复：complete=true；耗时=6.3902486s；删除帖子/Outbox=72/144；safety epoch=24652
- 测量后恢复：complete=true；耗时=5.0508564s；删除帖子/Outbox=410/820；safety epoch=25064

## 说明

- SLA values are reference lines, not pass/fail gates.
