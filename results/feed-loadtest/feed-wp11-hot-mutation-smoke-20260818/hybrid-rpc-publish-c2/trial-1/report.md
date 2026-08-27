# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-mutation-smoke-20260818`
- 开始时间：2026-08-18T15:40:25+08:00
- 采样时长：2.5217466s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 12 | 12 | 0 | 0 | 4.76 | 368.714 | 563.207 | 565.312 | 565.312 | 565.312 |
| publish_draft | 12 | 12 | 0 | 0 | 4.76 | 4.157 | 4.812 | 5.457 | 5.457 | 5.457 |
| publish_metadata | 12 | 12 | 0 | 0 | 4.76 | 121.955 | 151.229 | 152.285 | 152.285 | 152.285 |
| publish_confirm | 12 | 12 | 0 | 0 | 4.76 | 132.683 | 223.060 | 224.750 | 224.750 | 224.750 |
| publish_commit | 12 | 12 | 0 | 0 | 4.76 | 125.131 | 199.446 | 200.509 | 200.509 | 200.509 |

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

- Commands：39580；input：3084180 bytes；output：713509 bytes
- Hits/Misses：34/226；run hit rate：13.08%
- Evicted/Rejected：0/0；ops/s max：15185；safety epoch：3821 -> 3845

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.194 |
| client:loadtest | cpu_percent_total | 3.098 |
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
| docker:zg-canal | cpu_percent | 1.580 |
| docker:zg-canal | memory_percent | 2.010 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.420 |
| docker:zg-es | memory_percent | 10.840 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 3.840 |
| docker:zg-etcd | memory_percent | 0.250 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 97.920 |
| docker:zg-kafka | memory_percent | 6.900 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.230 |
| docker:zg-zk | memory_percent | 0.980 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4132.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4132.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 619.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 7.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.428 |
| process:counter | cpu_seconds_total | 34.141 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 38572032.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.484 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 37310464.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 48.757 |
| process:knowpost | cpu_seconds_total | 30.422 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 60243968.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.551 |
| process:relation | cpu_seconds_total | 1.094 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 39755776.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.922 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 37539840.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.969 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 35540992.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1013684.000 |
| redis | connected_clients | 21.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3845.000 |
| redis | hit_rate | 0.696 |
| redis | keys | 595657.000 |
| redis | keyspace_hits | 29538.000 |
| redis | keyspace_misses | 13114.000 |
| redis | net_input_bytes | 71123668.000 |
| redis | net_output_bytes | 19934273.000 |
| redis | ops_per_sec | 15185.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19211.000 |
| redis | used_memory_bytes | 94007944.000 |

## 停止施压后的恢复

- Kafka drain：2.5382699s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：12
- 测量前恢复：complete=true；耗时=5.0380057s；删除帖子/Outbox=0/0；safety epoch=3807
- 预热后恢复：complete=true；耗时=6.2600177s；删除帖子/Outbox=6/12；safety epoch=3820
- 测量后恢复：complete=true；耗时=4.9292363s；删除帖子/Outbox=12/24；safety epoch=3846

## 说明

- SLA values are reference lines, not pass/fail gates.
