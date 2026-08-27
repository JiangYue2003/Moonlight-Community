# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T22:53:41+08:00
- 采样时长：15.3147073s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 78 | 78 | 0 | 0 | 5.09 | 382.857 | 449.167 | 474.501 | 577.712 | 577.712 |
| publish_draft | 78 | 78 | 0 | 0 | 5.09 | 4.070 | 5.276 | 5.583 | 6.166 | 6.166 |
| publish_metadata | 78 | 78 | 0 | 0 | 5.09 | 124.823 | 155.733 | 183.326 | 190.466 | 190.466 |
| publish_confirm | 78 | 78 | 0 | 0 | 5.09 | 122.151 | 158.433 | 161.807 | 238.595 | 238.595 |
| publish_commit | 78 | 78 | 0 | 0 | 5.09 | 127.637 | 149.224 | 163.431 | 187.503 | 187.503 |

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

- Commands：241988；input：18900576 bytes；output：4333278 bytes
- Hits/Misses：190/1387；run hit rate：12.05%
- Evicted/Rejected：0/0；ops/s max：17168；safety epoch：9165 -> 9243

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.166 |
| client:loadtest | cpu_percent_total | 2.653 |
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
| docker:zg-canal | cpu_percent | 3.220 |
| docker:zg-canal | memory_percent | 4.470 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 7.340 |
| docker:zg-es | memory_percent | 11.840 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 3.890 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 191.610 |
| docker:zg-kafka | memory_percent | 7.540 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 0.200 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 6205.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 6205.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 47221.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 7.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 3.622 |
| process:counter | cpu_seconds_total | 3.172 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 39178240.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.062 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 36978688.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 63.587 |
| process:knowpost | cpu_seconds_total | 17.625 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 58007552.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.875 |
| process:relation | cpu_seconds_total | 0.391 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 42192896.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.797 |
| process:search | cpu_seconds_total | 0.219 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 38363136.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.094 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 34496512.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 21088835.000 |
| redis | connected_clients | 21.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 9243.000 |
| redis | hit_rate | 0.698 |
| redis | keys | 592388.000 |
| redis | keyspace_hits | 598076.000 |
| redis | keyspace_misses | 260415.000 |
| redis | net_input_bytes | 1555699978.000 |
| redis | net_output_bytes | 430458259.000 |
| redis | ops_per_sec | 17168.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45222.000 |
| redis | used_memory_bytes | 94758544.000 |

## 停止施压后的恢复

- Kafka drain：5.2245859s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：78
- 测量前恢复：complete=true；耗时=4.891889s；删除帖子/Outbox=0/0；safety epoch=9145
- 预热后恢复：complete=true；耗时=6.2541313s；删除帖子/Outbox=18/36；safety epoch=9164
- 测量后恢复：complete=true；耗时=4.863235s；删除帖子/Outbox=78/156；safety epoch=9244

## 说明

- SLA values are reference lines, not pass/fail gates.
