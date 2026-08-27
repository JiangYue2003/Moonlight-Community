# Feed 压测报告：hybrid / rpc / mixed-90-10-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-18T23:40:56+08:00
- 采样时长：1m0.9814416s
- 并发：8
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 2566 | 2566 | 0 | 0 | 42.08 | 3.201 | 4.482 | 5.035 | 6.206 | 11.392 |
| publish_total | 285 | 285 | 0 | 0 | 4.67 | 1678.996 | 1820.885 | 1848.754 | 1925.245 | 1965.853 |
| publish_draft | 285 | 285 | 0 | 0 | 4.67 | 3.845 | 4.627 | 4.875 | 5.769 | 7.583 |
| publish_metadata | 285 | 285 | 0 | 0 | 4.67 | 556.935 | 676.120 | 719.446 | 749.961 | 757.538 |
| publish_confirm | 285 | 285 | 0 | 0 | 4.67 | 546.749 | 629.863 | 657.987 | 738.754 | 751.690 |
| publish_commit | 285 | 285 | 0 | 0 | 4.67 | 545.584 | 681.497 | 721.493 | 767.011 | 776.457 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 238 | 0.093 |
| mysql | 129 | 0.050 |
| redis | 7827 | 3.050 |
| relation | 2566 | 1.000 |

- Cold compute：2566（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 102640 | 40.000 |
| merge_candidates | 173921 | 67.779 |
| redis_commands | 15396 | 6.000 |
| redis_members | 179718 | 70.038 |
| redis_roundtrips | 5132 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 2566 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2566 | 0.543 |
| counter | 238 | 0.867 |
| hydrate | 2566 | 0.689 |
| inbox | 2566 | 0.567 |
| merge_dedup | 2566 | 0.007 |
| relation | 2566 | 1.224 |
| route | 2566 | 0.083 |
| total | 2566 | 3.135 |

## Redis 本轮边界增量

- Commands：944731；input：76749688 bytes；output：40076132 bytes
- Hits/Misses：122940/9852；run hit rate：92.58%
- Evicted/Rejected：0/0；ops/s max：17905；safety epoch：16817 -> 17102

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.178 |
| client:loadtest | cpu_percent_total | 2.844 |
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
| docker:zg-canal | cpu_percent | 3.590 |
| docker:zg-canal | memory_percent | 4.950 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.820 |
| docker:zg-es | memory_percent | 12.200 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.380 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 202.010 |
| docker:zg-kafka | memory_percent | 7.680 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 49.770 |
| docker:zg-zk | memory_percent | 1.500 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 12204.000 |
| kafka | lag_max | 5.000 |
| kafka | lag_total | 5.000 |
| kafka | log_end_offset_total | 12204.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 194883.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.197 |
| process:counter | cpu_seconds_total | 77.016 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46977024.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 12.094 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 43802624.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 87.468 |
| process:knowpost | cpu_seconds_total | 1033.375 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 70160384.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.744 |
| process:relation | cpu_seconds_total | 28.906 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49340416.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.875 |
| process:search | cpu_seconds_total | 16.547 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43589632.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 8.266 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 37736448.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 46802290.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 17102.000 |
| redis | hit_rate | 0.759 |
| redis | keys | 596298.000 |
| redis | keyspace_hits | 2079802.000 |
| redis | keyspace_misses | 661427.000 |
| redis | net_input_bytes | 3575183314.000 |
| redis | net_output_bytes | 1093945962.000 |
| redis | ops_per_sec | 17905.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48103.000 |
| redis | used_memory_bytes | 96704000.000 |

## 停止施压后的恢复

- Kafka drain：5.4035558s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：285
- 测量前恢复：complete=true；耗时=4.9779327s；删除帖子/Outbox=0/0；safety epoch=16759
- 预热后恢复：complete=true；耗时=7.5479736s；删除帖子/Outbox=56/112；safety epoch=16816
- 测量后恢复：complete=true；耗时=4.9065134s；删除帖子/Outbox=285/570；safety epoch=17103

## 说明

- SLA values are reference lines, not pass/fail gates.
