# Feed 压测报告：hybrid / rpc / mixed-80-20-c4

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-mixed-c4`
- 开始时间：2026-08-27T06:02:48+08:00
- 采样时长：1m0.039297s
- 并发：4
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1616 | 1616 | 0 | 0 | 26.92 | 2.097 | 3.130 | 3.284 | 4.252 | 9.488 |
| publish_total | 404 | 404 | 0 | 0 | 6.73 | 574.962 | 675.340 | 692.486 | 724.375 | 754.141 |
| publish_draft | 404 | 404 | 0 | 0 | 6.73 | 3.677 | 4.304 | 4.787 | 25.163 | 171.115 |
| publish_metadata | 404 | 404 | 0 | 0 | 6.73 | 188.365 | 229.956 | 247.230 | 293.430 | 301.355 |
| publish_confirm | 404 | 404 | 0 | 0 | 6.73 | 187.318 | 231.583 | 242.010 | 266.642 | 284.487 |
| publish_commit | 404 | 404 | 0 | 0 | 6.73 | 184.876 | 236.836 | 255.330 | 309.394 | 332.768 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 214 | 0.132 |
| mysql | 187 | 0.116 |
| redis | 5035 | 3.116 |
| relation | 1616 | 1.000 |

- Cold compute：1616（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 64640 | 40.000 |
| merge_candidates | 121611 | 75.254 |
| redis_commands | 9696 | 6.000 |
| redis_members | 147352 | 91.183 |
| redis_roundtrips | 3232 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1616 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1616 | 0.223 |
| counter | 214 | 0.514 |
| hydrate | 1616 | 0.366 |
| inbox | 1616 | 0.224 |
| merge_dedup | 1616 | 0.010 |
| relation | 1616 | 0.972 |
| route | 1616 | 0.075 |
| total | 1616 | 1.896 |

## Redis 本轮边界增量

- Commands：1021724；input：82530443 bytes；output：34989685 bytes
- Hits/Misses：82590/7908；run hit rate：91.26%
- Evicted/Rejected：0/0；ops/s max：19856；safety epoch：29954 -> 30358

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.195 |
| client:loadtest | cpu_percent_total | 3.123 |
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
| docker:zg-canal | cpu_percent | 3.790 |
| docker:zg-canal | memory_percent | 4.890 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.710 |
| docker:zg-es | memory_percent | 12.990 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.210 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 149.310 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 45.140 |
| docker:zg-zk | memory_percent | 1.230 |
| docker:zg-zk | pids | 102.000 |
| kafka | current_offset_total | 22569.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 22569.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 179872.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 5.418 |
| process:counter | cpu_seconds_total | 44.547 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 44896256.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 4.391 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 41746432.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 89.954 |
| process:knowpost | cpu_seconds_total | 633.359 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 63660032.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 8.523 |
| process:relation | cpu_seconds_total | 38.812 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 53837824.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.877 |
| process:search | cpu_seconds_total | 12.781 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 43716608.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 3.781 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 42115072.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 15977599.000 |
| redis | connected_clients | 81.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 30358.000 |
| redis | hit_rate | 0.876 |
| redis | keys | 573578.000 |
| redis | keyspace_hits | 1345033.000 |
| redis | keyspace_misses | 190579.000 |
| redis | net_input_bytes | 1278745913.000 |
| redis | net_output_bytes | 523435309.000 |
| redis | ops_per_sec | 19856.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 12387.000 |
| redis | used_memory_bytes | 81117864.000 |

## 停止施压后的恢复

- Kafka drain：5.3703663s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：404
- 测量前恢复：complete=true；耗时=4.9181709s；删除帖子/Outbox=0/0；safety epoch=29872
- 预热后恢复：complete=true；耗时=6.2337134s；删除帖子/Outbox=80/160；safety epoch=29953
- 测量后恢复：complete=true；耗时=5.0314142s；删除帖子/Outbox=404/808；safety epoch=30359

## 说明

- SLA values are reference lines, not pass/fail gates.
