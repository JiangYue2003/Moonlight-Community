# Feed 压测报告：hybrid / rpc / publish-c4

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:20:05+08:00
- 采样时长：1m0.4533842s
- 并发：4
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 283 | 283 | 0 | 0 | 4.68 | 851.909 | 998.602 | 1041.125 | 1120.559 | 1177.768 |
| publish_draft | 283 | 283 | 0 | 0 | 4.68 | 3.837 | 5.614 | 6.544 | 105.431 | 107.577 |
| publish_metadata | 283 | 283 | 0 | 0 | 4.68 | 273.140 | 353.559 | 373.933 | 428.573 | 457.200 |
| publish_confirm | 283 | 283 | 0 | 0 | 4.68 | 276.541 | 336.548 | 362.620 | 402.700 | 433.560 |
| publish_commit | 283 | 283 | 0 | 0 | 4.68 | 283.972 | 346.814 | 362.296 | 419.220 | 424.025 |

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

- Commands：884449；input：69110720 bytes；output：15805576 bytes
- Hits/Misses：665/5292；run hit rate：11.16%
- Evicted/Rejected：0/0；ops/s max：17114；safety epoch：12773 -> 13056

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.152 |
| client:loadtest | cpu_percent_total | 2.430 |
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
| docker:zg-canal | cpu_percent | 1.680 |
| docker:zg-canal | memory_percent | 4.800 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 3.250 |
| docker:zg-es | memory_percent | 12.150 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.410 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 160.720 |
| docker:zg-kafka | memory_percent | 7.640 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 0.160 |
| docker:zg-zk | memory_percent | 1.310 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 9070.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 9070.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 118483.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 10.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.649 |
| process:counter | cpu_seconds_total | 43.750 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46432256.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 6.359 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 43343872.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 70.866 |
| process:knowpost | cpu_seconds_total | 512.516 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 69754880.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.325 |
| process:relation | cpu_seconds_total | 13.266 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48844800.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.324 |
| process:search | cpu_seconds_total | 8.109 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42807296.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.551 |
| process:user-storage | cpu_seconds_total | 4.062 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 37756928.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 33578487.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 13056.000 |
| redis | hit_rate | 0.736 |
| redis | keys | 593621.000 |
| redis | keyspace_hits | 1378451.000 |
| redis | keyspace_misses | 498944.000 |
| redis | net_input_bytes | 2532538344.000 |
| redis | net_output_bytes | 748022978.000 |
| redis | ops_per_sec | 17114.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46851.000 |
| redis | used_memory_bytes | 95895680.000 |

## 停止施压后的恢复

- Kafka drain：5.5446258s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：283
- 测量前恢复：complete=true；耗时=5.0437427s；删除帖子/Outbox=0/0；safety epoch=12715
- 预热后恢复：complete=true；耗时=6.5772657s；删除帖子/Outbox=56/112；safety epoch=12772
- 测量后恢复：complete=true；耗时=5.060184s；删除帖子/Outbox=283/566；safety epoch=13057

## 说明

- SLA values are reference lines, not pass/fail gates.
