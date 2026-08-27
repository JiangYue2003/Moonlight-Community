# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-treatment-formal-publish-c2-rpc`
- 开始时间：2026-08-27T06:07:17+08:00
- 采样时长：1m0.0559398s
- 并发：2
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 380 | 380 | 0 | 0 | 6.33 | 305.396 | 372.209 | 402.144 | 465.591 | 647.474 |
| publish_draft | 380 | 380 | 0 | 0 | 6.33 | 3.787 | 5.061 | 5.436 | 7.151 | 73.342 |
| publish_metadata | 380 | 380 | 0 | 0 | 6.33 | 100.718 | 123.281 | 135.183 | 155.032 | 172.222 |
| publish_confirm | 380 | 380 | 0 | 0 | 6.33 | 99.592 | 122.978 | 129.826 | 145.462 | 203.981 |
| publish_commit | 380 | 380 | 0 | 0 | 6.33 | 99.702 | 130.625 | 148.072 | 189.526 | 428.505 |

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

- Commands：936872；input：73984103 bytes；output：16698067 bytes
- Hits/Misses：1740/6195；run hit rate：21.93%
- Evicted/Rejected：0/0；ops/s max：18000；safety epoch：36733 -> 37493

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.153 |
| client:loadtest | cpu_percent_total | 2.446 |
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
| docker:zg-canal | cpu_percent | 27.450 |
| docker:zg-canal | memory_percent | 4.950 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 13.340 |
| docker:zg-es | memory_percent | 13.020 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 4.210 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 230.160 |
| docker:zg-kafka | memory_percent | 7.540 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 22.790 |
| docker:zg-zk | memory_percent | 1.160 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 22917.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 22917.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 186755.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 10.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 5.419 |
| process:counter | cpu_seconds_total | 4.234 |
| process:counter | pid | 20740.000 |
| process:counter | process_start_ms | 1787781931297.000 |
| process:counter | rss_bytes | 41013248.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 0.062 |
| process:gateway | pid | 39628.000 |
| process:gateway | process_start_ms | 1787781953868.000 |
| process:gateway | rss_bytes | 36417536.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 68.178 |
| process:knowpost | cpu_seconds_total | 49.703 |
| process:knowpost | pid | 39528.000 |
| process:knowpost | process_start_ms | 1787781942481.000 |
| process:knowpost | rss_bytes | 64540672.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.875 |
| process:relation | cpu_seconds_total | 0.656 |
| process:relation | pid | 21680.000 |
| process:relation | process_start_ms | 1787781936365.000 |
| process:relation | rss_bytes | 42336256.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.325 |
| process:search | cpu_seconds_total | 0.656 |
| process:search | pid | 27196.000 |
| process:search | process_start_ms | 1787781948810.000 |
| process:search | rss_bytes | 41607168.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.551 |
| process:user-storage | cpu_seconds_total | 0.141 |
| process:user-storage | pid | 34404.000 |
| process:user-storage | process_start_ms | 1787781926436.000 |
| process:user-storage | rss_bytes | 34570240.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 17183085.000 |
| redis | connected_clients | 21.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 37493.000 |
| redis | hit_rate | 0.872 |
| redis | keys | 573052.000 |
| redis | keyspace_hits | 1361144.000 |
| redis | keyspace_misses | 205543.000 |
| redis | net_input_bytes | 1372544541.000 |
| redis | net_output_bytes | 545134533.000 |
| redis | ops_per_sec | 18000.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 12656.000 |
| redis | used_memory_bytes | 79368776.000 |

## 停止施压后的恢复

- Kafka drain：5.2659884s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：380
- 测量前恢复：complete=true；耗时=4.9765659s；删除帖子/Outbox=0/0；safety epoch=36595
- 预热后恢复：complete=true；耗时=6.4115463s；删除帖子/Outbox=68/136；safety epoch=36732
- 测量后恢复：complete=true；耗时=4.9607763s；删除帖子/Outbox=380/760；safety epoch=37494

## 说明

- SLA values are reference lines, not pass/fail gates.
