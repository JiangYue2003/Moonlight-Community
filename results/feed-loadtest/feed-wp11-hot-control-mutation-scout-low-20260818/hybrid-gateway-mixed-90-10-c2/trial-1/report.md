# Feed 压测报告：hybrid / gateway / mixed-90-10-c2

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:06:38+08:00
- 采样时长：15.2380256s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 630 | 630 | 0 | 0 | 41.34 | 2.068 | 2.592 | 2.696 | 3.249 | 6.328 |
| publish_total | 70 | 70 | 0 | 0 | 4.59 | 418.634 | 458.860 | 503.377 | 533.968 | 533.968 |
| publish_draft | 70 | 70 | 0 | 0 | 4.59 | 3.761 | 3.943 | 4.374 | 6.575 | 6.575 |
| publish_metadata | 70 | 70 | 0 | 0 | 4.59 | 137.068 | 159.432 | 166.432 | 180.522 | 180.522 |
| publish_confirm | 70 | 70 | 0 | 0 | 4.59 | 133.684 | 164.820 | 177.987 | 216.131 | 216.131 |
| publish_commit | 70 | 70 | 0 | 0 | 4.59 | 137.086 | 160.988 | 172.327 | 219.314 | 219.314 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 60 | 0.095 |
| mysql | 31 | 0.049 |
| redis | 1921 | 3.049 |
| relation | 630 | 1.000 |

- Cold compute：630（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 25200 | 40.000 |
| merge_candidates | 30609 | 48.586 |
| redis_commands | 3780 | 6.000 |
| redis_members | 30609 | 48.586 |
| redis_roundtrips | 1260 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 630 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 630 | 0.093 |
| counter | 60 | 0.399 |
| hydrate | 630 | 0.207 |
| inbox | 630 | 0.129 |
| merge_dedup | 630 | 0.012 |
| relation | 630 | 0.756 |
| route | 630 | 0.041 |
| total | 630 | 1.253 |

## Redis 本轮边界增量

- Commands：227009；input：18441487 bytes；output：9230014 bytes
- Hits/Misses：29575/2967；run hit rate：90.88%
- Evicted/Rejected：0/0；ops/s max：15539；safety epoch：10911 -> 10981

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.128 |
| client:loadtest | cpu_percent_total | 2.051 |
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
| docker:zg-canal | cpu_percent | 1.490 |
| docker:zg-canal | memory_percent | 4.670 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.840 |
| docker:zg-es | memory_percent | 11.880 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.900 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 153.920 |
| docker:zg-kafka | memory_percent | 7.710 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 0.170 |
| docker:zg-zk | memory_percent | 1.290 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7490.000 |
| kafka | lag_max | 5.000 |
| kafka | lag_total | 5.000 |
| kafka | log_end_offset_total | 7490.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 81044.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 3.875 |
| process:counter | cpu_seconds_total | 22.812 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45543424.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 8.527 |
| process:gateway | cpu_seconds_total | 2.016 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46055424.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 68.944 |
| process:knowpost | cpu_seconds_total | 243.453 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68247552.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.190 |
| process:relation | cpu_seconds_total | 6.703 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48549888.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 3.625 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42954752.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.874 |
| process:user-storage | cpu_seconds_total | 1.500 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40685568.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 26801044.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 10981.000 |
| redis | hit_rate | 0.729 |
| redis | keys | 593026.000 |
| redis | keyspace_hits | 1049900.000 |
| redis | keyspace_misses | 389678.000 |
| redis | net_input_bytes | 2002647073.000 |
| redis | net_output_bytes | 587009503.000 |
| redis | ops_per_sec | 15539.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46000.000 |
| redis | used_memory_bytes | 95200608.000 |

## 停止施压后的恢复

- Kafka drain：6.649338s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：70
- 测量前恢复：complete=true；耗时=4.8627579s；删除帖子/Outbox=0/0；safety epoch=10893
- 预热后恢复：complete=true；耗时=6.227268s；删除帖子/Outbox=16/32；safety epoch=10910
- 测量后恢复：complete=true；耗时=4.8380445s；删除帖子/Outbox=70/140；safety epoch=10982

## 说明

- SLA values are reference lines, not pass/fail gates.
