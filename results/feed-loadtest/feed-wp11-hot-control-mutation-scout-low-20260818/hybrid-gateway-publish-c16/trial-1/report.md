# Feed 压测报告：hybrid / gateway / publish-c16

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:05:51+08:00
- 采样时长：16.7215925s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 89 | 89 | 0 | 0 | 5.32 | 2995.640 | 3083.816 | 3109.087 | 3190.172 | 3190.172 |
| publish_draft | 89 | 89 | 0 | 0 | 5.32 | 5.766 | 16.191 | 29.351 | 88.537 | 88.537 |
| publish_metadata | 89 | 89 | 0 | 0 | 5.32 | 979.927 | 1073.401 | 1083.933 | 1111.914 | 1111.914 |
| publish_confirm | 89 | 89 | 0 | 0 | 5.32 | 917.122 | 1101.332 | 1108.314 | 1158.714 | 1158.714 |
| publish_commit | 89 | 89 | 0 | 0 | 5.32 | 1067.513 | 1120.752 | 1128.027 | 1195.735 | 1195.735 |

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

- Commands：276359；input：21601086 bytes；output：4950603 bytes
- Hits/Misses：216/1600；run hit rate：11.89%
- Evicted/Rejected：0/0；ops/s max：18427；safety epoch：10802 -> 10891

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.117 |
| client:loadtest | cpu_percent_total | 1.869 |
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
| docker:zg-canal | cpu_percent | 1.760 |
| docker:zg-canal | memory_percent | 4.690 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 1.280 |
| docker:zg-es | memory_percent | 11.870 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.040 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 234.730 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 37.940 |
| docker:zg-zk | memory_percent | 1.290 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7424.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 7424.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 78789.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.479 |
| process:counter | cpu_seconds_total | 21.750 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45416448.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.866 |
| process:gateway | cpu_seconds_total | 1.172 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46178304.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 72.002 |
| process:knowpost | cpu_seconds_total | 230.797 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 70180864.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.097 |
| process:relation | cpu_seconds_total | 6.172 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48132096.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 3.484 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42860544.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.325 |
| process:user-storage | cpu_seconds_total | 1.031 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40583168.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 26494831.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 10891.000 |
| redis | hit_rate | 0.726 |
| redis | keys | 592922.000 |
| redis | keyspace_hits | 1003358.000 |
| redis | keyspace_misses | 380809.000 |
| redis | net_input_bytes | 1978240157.000 |
| redis | net_output_bytes | 574883618.000 |
| redis | ops_per_sec | 18427.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45953.000 |
| redis | used_memory_bytes | 96105088.000 |

## 停止施压后的恢复

- Kafka drain：5.2063437s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：89
- 测量前恢复：complete=true；耗时=4.908139s；删除帖子/Outbox=0/0；safety epoch=10768
- 预热后恢复：complete=true；耗时=6.2230408s；删除帖子/Outbox=32/64；safety epoch=10801
- 测量后恢复：complete=true；耗时=4.9164174s；删除帖子/Outbox=89/178；safety epoch=10892

## 说明

- SLA values are reference lines, not pass/fail gates.
