# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-treatment-formal-publish-c2-rpc`
- 开始时间：2026-08-27T06:08:54+08:00
- 采样时长：1m0.2722089s
- 并发：2
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 378 | 378 | 0 | 0 | 6.27 | 310.284 | 366.328 | 388.983 | 513.722 | 593.180 |
| publish_draft | 378 | 378 | 0 | 0 | 6.27 | 3.984 | 5.199 | 5.763 | 7.738 | 195.052 |
| publish_metadata | 378 | 378 | 0 | 0 | 6.27 | 101.550 | 124.840 | 138.569 | 167.714 | 271.410 |
| publish_confirm | 378 | 378 | 0 | 0 | 6.27 | 99.796 | 123.209 | 136.515 | 291.176 | 385.849 |
| publish_commit | 378 | 378 | 0 | 0 | 6.27 | 100.820 | 123.296 | 131.187 | 160.352 | 319.975 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 0 | 0.000 |
| counter | 0 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 0 | 0.000 |
| relation | 0 | 0.000 |

- Cold compute：0（0.000 / successful read）

| page cache source | requests |
|---|---:|
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：932295；input：73591677 bytes；output：16584658 bytes
- Hits/Misses：1734/6860；run hit rate：20.18%
- Evicted/Rejected：0/0；ops/s max：17121；safety epoch：37633 -> 38389

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.157 |
| client:loadtest | cpu_percent_total | 2.515 |
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
| docker:zg-canal | cpu_percent | 3.750 |
| docker:zg-canal | memory_percent | 4.920 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.390 |
| docker:zg-es | memory_percent | 13.020 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 4.700 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 160.340 |
| docker:zg-kafka | memory_percent | 7.560 |
| docker:zg-kafka | pids | 120.000 |
| docker:zg-zk | cpu_percent | 42.090 |
| docker:zg-zk | memory_percent | 1.160 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 23264.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 23264.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 193444.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 6.201 |
| process:counter | cpu_seconds_total | 6.625 |
| process:counter | pid | 20740.000 |
| process:counter | process_start_ms | 1787781931297.000 |
| process:counter | rss_bytes | 42237952.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 0.078 |
| process:gateway | pid | 39628.000 |
| process:gateway | process_start_ms | 1787781953868.000 |
| process:gateway | rss_bytes | 36597760.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 66.573 |
| process:knowpost | cpu_seconds_total | 90.969 |
| process:knowpost | pid | 39528.000 |
| process:knowpost | process_start_ms | 1787781942481.000 |
| process:knowpost | rss_bytes | 65638400.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.100 |
| process:relation | cpu_seconds_total | 1.281 |
| process:relation | pid | 21680.000 |
| process:relation | process_start_ms | 1787781936365.000 |
| process:relation | rss_bytes | 43245568.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.093 |
| process:search | cpu_seconds_total | 1.453 |
| process:search | pid | 27196.000 |
| process:search | process_start_ms | 1787781948810.000 |
| process:search | rss_bytes | 42020864.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.322 |
| process:user-storage | cpu_seconds_total | 0.250 |
| process:user-storage | pid | 34404.000 |
| process:user-storage | process_start_ms | 1787781926436.000 |
| process:user-storage | rss_bytes | 34721792.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18307674.000 |
| redis | connected_clients | 21.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 38389.000 |
| redis | hit_rate | 0.866 |
| redis | keys | 573025.000 |
| redis | keyspace_hits | 1373636.000 |
| redis | keyspace_misses | 218495.000 |
| redis | net_input_bytes | 1461031092.000 |
| redis | net_output_bytes | 565395662.000 |
| redis | ops_per_sec | 17121.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 12753.000 |
| redis | used_memory_bytes | 79392496.000 |

## 停止施压后的恢复

- Kafka drain：5.3255239s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：378
- 测量前恢复：complete=true；耗时=4.9403619s；删除帖子/Outbox=0/0；safety epoch=37495
- 预热后恢复：complete=true；耗时=6.3999386s；删除帖子/Outbox=68/136；safety epoch=37632
- 测量后恢复：complete=true；耗时=4.8997592s；删除帖子/Outbox=378/756；safety epoch=38390

## 说明

- SLA values are reference lines, not pass/fail gates.
