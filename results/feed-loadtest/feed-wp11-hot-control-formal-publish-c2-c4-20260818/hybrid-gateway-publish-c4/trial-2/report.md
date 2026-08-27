# Feed 压测报告：hybrid / gateway / publish-c4

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:31:42+08:00
- 采样时长：1m0.2150431s
- 并发：4
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 300 | 300 | 0 | 0 | 4.98 | 793.732 | 909.350 | 924.459 | 1099.071 | 1105.520 |
| publish_draft | 300 | 300 | 0 | 0 | 4.98 | 4.796 | 6.585 | 6.989 | 199.764 | 200.833 |
| publish_metadata | 300 | 300 | 0 | 0 | 4.98 | 256.694 | 306.392 | 323.024 | 381.444 | 398.389 |
| publish_confirm | 300 | 300 | 0 | 0 | 4.98 | 258.105 | 317.060 | 336.963 | 397.913 | 401.105 |
| publish_commit | 300 | 300 | 0 | 0 | 4.98 | 253.673 | 312.675 | 356.064 | 480.573 | 492.020 |

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

- Commands：954733；input：74619092 bytes；output：17073317 bytes
- Hits/Misses：701/5859；run hit rate：10.69%
- Evicted/Rejected：0/0；ops/s max：18178；safety epoch：15125 -> 15425

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.120 |
| client:loadtest | cpu_percent_total | 1.920 |
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
| docker:zg-canal | cpu_percent | 3.040 |
| docker:zg-canal | memory_percent | 4.950 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.570 |
| docker:zg-es | memory_percent | 12.170 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.460 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 19.000 |
| docker:zg-kafka | cpu_percent | 166.940 |
| docker:zg-kafka | memory_percent | 7.670 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 37.950 |
| docker:zg-zk | memory_percent | 1.550 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 10877.000 |
| kafka | lag_max | 4.000 |
| kafka | lag_total | 4.000 |
| kafka | log_end_offset_total | 10877.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 155743.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 9.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 5.568 |
| process:counter | cpu_seconds_total | 62.344 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45400064.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.100 |
| process:gateway | cpu_seconds_total | 10.844 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 48078848.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 75.074 |
| process:knowpost | cpu_seconds_total | 805.844 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 70615040.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.551 |
| process:relation | cpu_seconds_total | 18.078 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49369088.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.327 |
| process:search | cpu_seconds_total | 13.062 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43401216.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.872 |
| process:user-storage | cpu_seconds_total | 7.422 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41603072.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 41185255.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 15425.000 |
| redis | hit_rate | 0.717 |
| redis | keys | 595668.000 |
| redis | keyspace_hits | 1456596.000 |
| redis | keyspace_misses | 579414.000 |
| redis | net_input_bytes | 3124565229.000 |
| redis | net_output_bytes | 885969115.000 |
| redis | ops_per_sec | 18178.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47549.000 |
| redis | used_memory_bytes | 96362456.000 |

## 停止施压后的恢复

- Kafka drain：6.7175385s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：300
- 测量前恢复：complete=true；耗时=4.8938902s；删除帖子/Outbox=0/0；safety epoch=15066
- 预热后恢复：complete=true；耗时=6.2592771s；删除帖子/Outbox=57/114；safety epoch=15124
- 测量后恢复：complete=true；耗时=4.9187272s；删除帖子/Outbox=300/600；safety epoch=15426

## 说明

- SLA values are reference lines, not pass/fail gates.
