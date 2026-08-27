# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-publish-c2`
- 开始时间：2026-08-27T05:43:14+08:00
- 采样时长：1m0.1771461s
- 并发：2
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 398 | 398 | 0 | 0 | 6.61 | 293.636 | 358.179 | 398.188 | 487.588 | 497.418 |
| publish_draft | 398 | 398 | 0 | 0 | 6.61 | 3.929 | 5.312 | 5.991 | 159.446 | 189.018 |
| publish_metadata | 398 | 398 | 0 | 0 | 6.61 | 95.008 | 120.012 | 129.271 | 186.478 | 292.790 |
| publish_confirm | 398 | 398 | 0 | 0 | 6.61 | 93.019 | 115.769 | 130.272 | 163.669 | 169.500 |
| publish_commit | 398 | 398 | 0 | 0 | 6.61 | 96.089 | 117.655 | 126.214 | 164.137 | 301.343 |

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

- Commands：927699；input：73199051 bytes；output：16425861 bytes
- Hits/Misses：1822/7233；run hit rate：20.12%
- Evicted/Rejected：0/0；ops/s max：17074；safety epoch：25140 -> 25538

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.146 |
| client:loadtest | cpu_percent_total | 2.337 |
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
| docker:zg-canal | cpu_percent | 16.790 |
| docker:zg-canal | memory_percent | 4.090 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.890 |
| docker:zg-es | memory_percent | 12.530 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 4.480 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 23.000 |
| docker:zg-kafka | cpu_percent | 178.410 |
| docker:zg-kafka | memory_percent | 7.510 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 45.360 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 18797.000 |
| kafka | lag_max | 6.000 |
| kafka | lag_total | 6.000 |
| kafka | log_end_offset_total | 18797.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 82885.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 5.424 |
| process:counter | cpu_seconds_total | 16.156 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 45084672.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.094 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 37965824.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 74.362 |
| process:knowpost | cpu_seconds_total | 143.266 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 62410752.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.098 |
| process:relation | cpu_seconds_total | 17.125 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 54366208.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.096 |
| process:search | cpu_seconds_total | 3.062 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 43044864.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.779 |
| process:user-storage | cpu_seconds_total | 1.031 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 41414656.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3715688.000 |
| redis | connected_clients | 75.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 25538.000 |
| redis | hit_rate | 0.614 |
| redis | keys | 569026.000 |
| redis | keyspace_hits | 62216.000 |
| redis | keyspace_misses | 45247.000 |
| redis | net_input_bytes | 286730518.000 |
| redis | net_output_bytes | 64887241.000 |
| redis | ops_per_sec | 17074.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 11212.000 |
| redis | used_memory_bytes | 80119656.000 |

## 停止施压后的恢复

- Kafka drain：5.3927767s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：398
- 测量前恢复：complete=true；耗时=5.069571s；删除帖子/Outbox=0/0；safety epoch=25065
- 预热后恢复：complete=true；耗时=6.3258039s；删除帖子/Outbox=73/146；safety epoch=25139
- 测量后恢复：complete=true；耗时=5.1379546s；删除帖子/Outbox=398/796；safety epoch=25539

## 说明

- SLA values are reference lines, not pass/fail gates.
