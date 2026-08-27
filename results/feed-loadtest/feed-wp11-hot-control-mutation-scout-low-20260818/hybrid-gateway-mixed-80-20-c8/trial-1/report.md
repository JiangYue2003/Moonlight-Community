# Feed 压测报告：hybrid / gateway / mixed-80-20-c8

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:11:30+08:00
- 采样时长：15.1411561s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 320 | 320 | 0 | 0 | 21.13 | 4.150 | 6.313 | 6.963 | 8.415 | 11.207 |
| publish_total | 80 | 80 | 0 | 0 | 5.28 | 1490.963 | 1589.083 | 1632.262 | 1657.561 | 1657.561 |
| publish_draft | 80 | 80 | 0 | 0 | 5.28 | 4.325 | 4.984 | 5.303 | 5.962 | 5.962 |
| publish_metadata | 80 | 80 | 0 | 0 | 5.28 | 520.006 | 616.957 | 633.117 | 646.017 | 646.017 |
| publish_confirm | 80 | 80 | 0 | 0 | 5.28 | 466.704 | 556.351 | 574.608 | 581.340 | 581.340 |
| publish_commit | 80 | 80 | 0 | 0 | 5.28 | 484.307 | 549.782 | 558.275 | 582.676 | 582.676 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 55 | 0.172 |
| mysql | 48 | 0.150 |
| redis | 1008 | 3.150 |
| relation | 320 | 1.000 |

- Cold compute：320（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 12800 | 40.000 |
| merge_candidates | 15669 | 48.966 |
| redis_commands | 1920 | 6.000 |
| redis_members | 15669 | 48.966 |
| redis_roundtrips | 640 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 320 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 320 | 0.566 |
| counter | 55 | 0.974 |
| hydrate | 320 | 0.814 |
| inbox | 320 | 0.576 |
| merge_dedup | 320 | 0.007 |
| relation | 320 | 1.349 |
| route | 320 | 0.177 |
| total | 320 | 3.521 |

## Redis 本轮边界增量

- Commands：254779；input：20242730 bytes；output：7180831 bytes
- Hits/Misses：15704/2355；run hit rate：86.96%
- Evicted/Rejected：0/0；ops/s max：17583；safety epoch：11518 -> 11598

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.135 |
| client:loadtest | cpu_percent_total | 2.167 |
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
| docker:zg-canal | cpu_percent | 1.710 |
| docker:zg-canal | memory_percent | 4.700 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.480 |
| docker:zg-es | memory_percent | 11.890 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.200 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 142.730 |
| docker:zg-kafka | memory_percent | 7.570 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 41.640 |
| docker:zg-zk | memory_percent | 1.300 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7949.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 7949.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 94956.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.970 |
| process:counter | cpu_seconds_total | 30.094 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46751744.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 4.789 |
| process:gateway | cpu_seconds_total | 5.750 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47747072.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 95.252 |
| process:knowpost | cpu_seconds_total | 328.812 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 70463488.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.095 |
| process:relation | cpu_seconds_total | 10.328 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 49471488.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.596 |
| process:search | cpu_seconds_total | 5.062 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43200512.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 6.194 |
| process:user-storage | cpu_seconds_total | 3.500 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40665088.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 28871377.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 11598.000 |
| redis | hit_rate | 0.745 |
| redis | keys | 593179.000 |
| redis | keyspace_hits | 1295378.000 |
| redis | keyspace_misses | 443146.000 |
| redis | net_input_bytes | 2166471151.000 |
| redis | net_output_bytes | 657823019.000 |
| redis | ops_per_sec | 17583.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46292.000 |
| redis | used_memory_bytes | 95599832.000 |

## 停止施压后的恢复

- Kafka drain：6.4964177s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：80
- 测量前恢复：complete=true；耗时=4.884063s；删除帖子/Outbox=0/0；safety epoch=11492
- 预热后恢复：complete=true；耗时=6.2743163s；删除帖子/Outbox=24/48；safety epoch=11517
- 测量后恢复：complete=true；耗时=4.8815373s；删除帖子/Outbox=80/160；safety epoch=11599

## 说明

- SLA values are reference lines, not pass/fail gates.
