# Feed 压测报告：hybrid / rpc / publish-c16

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T22:56:07+08:00
- 采样时长：17.5247439s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 96 | 96 | 0 | 0 | 5.48 | 2907.824 | 3020.956 | 3062.065 | 3125.086 | 3125.086 |
| publish_draft | 96 | 96 | 0 | 0 | 5.48 | 4.940 | 8.525 | 8.525 | 11.662 | 11.662 |
| publish_metadata | 96 | 96 | 0 | 0 | 5.48 | 930.780 | 1058.852 | 1069.108 | 1098.366 | 1098.366 |
| publish_confirm | 96 | 96 | 0 | 0 | 5.48 | 988.854 | 1055.309 | 1082.661 | 1108.368 | 1108.368 |
| publish_commit | 96 | 96 | 0 | 0 | 5.48 | 984.524 | 1126.012 | 1146.148 | 1168.739 | 1168.739 |

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

- Commands：296947；input：23202831 bytes；output：5311208 bytes
- Hits/Misses：229/1758；run hit rate：11.52%
- Evicted/Rejected：0/0；ops/s max：18694；safety epoch：9503 -> 9599

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.111 |
| client:loadtest | cpu_percent_total | 1.783 |
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
| docker:zg-canal | cpu_percent | 5.040 |
| docker:zg-canal | memory_percent | 4.620 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.220 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.400 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 185.030 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 39.310 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 6466.000 |
| kafka | lag_max | 14.000 |
| kafka | lag_total | 14.000 |
| kafka | log_end_offset_total | 6466.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 52795.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.746 |
| process:counter | cpu_seconds_total | 7.547 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 41648128.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.062 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 37969920.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 96.968 |
| process:knowpost | cpu_seconds_total | 62.109 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 66170880.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.099 |
| process:relation | cpu_seconds_total | 0.844 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 44208128.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.799 |
| process:search | cpu_seconds_total | 0.766 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42151936.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.550 |
| process:user-storage | cpu_seconds_total | 0.172 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35078144.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 22235327.000 |
| redis | connected_clients | 48.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 9599.000 |
| redis | hit_rate | 0.693 |
| redis | keys | 592168.000 |
| redis | keyspace_hits | 630113.000 |
| redis | keyspace_misses | 281410.000 |
| redis | net_input_bytes | 1644130482.000 |
| redis | net_output_bytes | 451898819.000 |
| redis | ops_per_sec | 18694.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45372.000 |
| redis | used_memory_bytes | 95610240.000 |

## 停止施压后的恢复

- Kafka drain：6.6258461s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：96
- 测量前恢复：complete=true；耗时=4.9408557s；删除帖子/Outbox=0/0；safety epoch=9469
- 预热后恢复：complete=true；耗时=7.641117s；删除帖子/Outbox=32/64；safety epoch=9502
- 测量后恢复：complete=true；耗时=4.9094867s；删除帖子/Outbox=96/192；safety epoch=9600

## 说明

- SLA values are reference lines, not pass/fail gates.
