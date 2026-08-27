# Feed 压测报告：hybrid / rpc / mixed-90-10-c2

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T22:56:58+08:00
- 采样时长：15.0753169s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 648 | 648 | 0 | 0 | 42.98 | 1.549 | 1.688 | 2.123 | 3.182 | 11.514 |
| publish_total | 72 | 72 | 0 | 0 | 4.78 | 402.403 | 482.940 | 506.224 | 616.500 | 616.500 |
| publish_draft | 72 | 72 | 0 | 0 | 4.78 | 3.228 | 3.862 | 4.272 | 255.872 | 255.872 |
| publish_metadata | 72 | 72 | 0 | 0 | 4.78 | 128.770 | 159.021 | 183.699 | 248.723 | 248.723 |
| publish_confirm | 72 | 72 | 0 | 0 | 4.78 | 124.567 | 166.022 | 181.361 | 208.417 | 208.417 |
| publish_commit | 72 | 72 | 0 | 0 | 4.78 | 127.685 | 165.093 | 169.391 | 189.035 | 189.035 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 60 | 0.093 |
| mysql | 29 | 0.045 |
| redis | 1973 | 3.045 |
| relation | 648 | 1.000 |

- Cold compute：648（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 25920 | 40.000 |
| merge_candidates | 31591 | 48.752 |
| redis_commands | 3888 | 6.000 |
| redis_members | 31591 | 48.752 |
| redis_roundtrips | 1296 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 648 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 648 | 0.097 |
| counter | 60 | 0.442 |
| hydrate | 648 | 0.186 |
| inbox | 648 | 0.110 |
| merge_dedup | 648 | 0.006 |
| relation | 648 | 0.724 |
| route | 648 | 0.041 |
| total | 648 | 1.190 |

## Redis 本轮边界增量

- Commands：231618；input：18822985 bytes；output：9458288 bytes
- Hits/Misses：30384/3080；run hit rate：90.80%
- Evicted/Rejected：0/0；ops/s max：17205；safety epoch：9621 -> 9693

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.175 |
| client:loadtest | cpu_percent_total | 2.798 |
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
| docker:zg-canal | cpu_percent | 3.410 |
| docker:zg-canal | memory_percent | 4.560 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 10.290 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.410 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 206.770 |
| docker:zg-kafka | memory_percent | 7.660 |
| docker:zg-kafka | pids | 142.000 |
| docker:zg-zk | cpu_percent | 29.880 |
| docker:zg-zk | memory_percent | 1.520 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 6535.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 6535.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 55138.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 10.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.379 |
| process:counter | cpu_seconds_total | 8.828 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 42483712.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 0.078 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 37969920.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 67.481 |
| process:knowpost | cpu_seconds_total | 74.188 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 64995328.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.784 |
| process:relation | cpu_seconds_total | 1.359 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 44421120.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.101 |
| process:search | cpu_seconds_total | 0.969 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42360832.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.188 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35086336.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 22551499.000 |
| redis | connected_clients | 49.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 9693.000 |
| redis | hit_rate | 0.700 |
| redis | keys | 592248.000 |
| redis | keyspace_hits | 678248.000 |
| redis | keyspace_misses | 290461.000 |
| redis | net_input_bytes | 1669366107.000 |
| redis | net_output_bytes | 464497115.000 |
| redis | ops_per_sec | 17205.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45418.000 |
| redis | used_memory_bytes | 94921928.000 |

## 停止施压后的恢复

- Kafka drain：5.1621903s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：72
- 测量前恢复：complete=true；耗时=4.936833s；删除帖子/Outbox=0/0；safety epoch=9601
- 预热后恢复：complete=true；耗时=6.2514579s；删除帖子/Outbox=18/36；safety epoch=9620
- 测量后恢复：complete=true；耗时=4.8555103s；删除帖子/Outbox=72/144；safety epoch=9694

## 说明

- SLA values are reference lines, not pass/fail gates.
