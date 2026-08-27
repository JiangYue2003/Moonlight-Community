# Feed 压测报告：hybrid / gateway / mixed-80-20-c2

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:09:54+08:00
- 采样时长：15.1467511s
- 并发：2
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 288 | 288 | 0 | 0 | 19.01 | 2.095 | 2.724 | 3.185 | 4.354 | 5.937 |
| publish_total | 72 | 72 | 0 | 0 | 4.75 | 410.117 | 468.231 | 514.039 | 518.804 | 518.804 |
| publish_draft | 72 | 72 | 0 | 0 | 4.75 | 3.696 | 4.314 | 4.403 | 6.446 | 6.446 |
| publish_metadata | 72 | 72 | 0 | 0 | 4.75 | 134.177 | 162.506 | 176.863 | 185.684 | 185.684 |
| publish_confirm | 72 | 72 | 0 | 0 | 4.75 | 126.598 | 160.340 | 166.949 | 204.619 | 204.619 |
| publish_commit | 72 | 72 | 0 | 0 | 4.75 | 130.242 | 165.437 | 168.940 | 199.987 | 199.987 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 55 | 0.191 |
| mysql | 33 | 0.115 |
| redis | 897 | 3.115 |
| relation | 288 | 1.000 |

- Cold compute：288（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 11520 | 40.000 |
| merge_candidates | 14198 | 49.299 |
| redis_commands | 1728 | 6.000 |
| redis_members | 14198 | 49.299 |
| redis_roundtrips | 576 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 288 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 288 | 0.103 |
| counter | 55 | 0.401 |
| hydrate | 288 | 0.233 |
| inbox | 288 | 0.129 |
| merge_dedup | 288 | 0.013 |
| relation | 288 | 0.803 |
| route | 288 | 0.082 |
| total | 288 | 1.384 |

## Redis 本轮边界增量

- Commands：229389；input：18220269 bytes；output：6478708 bytes
- Hits/Misses：14344/2028；run hit rate：87.61%
- Evicted/Rejected：0/0；ops/s max：15878；safety epoch：11318 -> 11390

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.110 |
| client:loadtest | cpu_percent_total | 1.754 |
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
| docker:zg-canal | cpu_percent | 2.960 |
| docker:zg-canal | memory_percent | 4.680 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.960 |
| docker:zg-es | memory_percent | 11.880 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.000 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 184.430 |
| docker:zg-kafka | memory_percent | 7.760 |
| docker:zg-kafka | pids | 143.000 |
| docker:zg-zk | cpu_percent | 0.140 |
| docker:zg-zk | memory_percent | 1.300 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7793.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 7793.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 90786.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.188 |
| process:counter | cpu_seconds_total | 27.656 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 46526464.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.873 |
| process:gateway | cpu_seconds_total | 4.906 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47951872.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 62.750 |
| process:knowpost | cpu_seconds_total | 298.719 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 68390912.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.867 |
| process:relation | cpu_seconds_total | 9.375 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48594944.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 4.609 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42876928.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.323 |
| process:user-storage | cpu_seconds_total | 3.047 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41549824.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 28177919.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 11390.000 |
| redis | hit_rate | 0.743 |
| redis | keys | 593186.000 |
| redis | keyspace_hits | 1235443.000 |
| redis | keyspace_misses | 427075.000 |
| redis | net_input_bytes | 2112238221.000 |
| redis | net_output_bytes | 638273340.000 |
| redis | ops_per_sec | 15878.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 46195.000 |
| redis | used_memory_bytes | 95227792.000 |

## 停止施压后的恢复

- Kafka drain：5.1898383s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：72
- 测量前恢复：complete=true；耗时=4.9151733s；删除帖子/Outbox=0/0；safety epoch=11300
- 预热后恢复：complete=true；耗时=6.2274646s；删除帖子/Outbox=16/32；safety epoch=11317
- 测量后恢复：complete=true；耗时=4.9020614s；删除帖子/Outbox=72/144；safety epoch=11391

## 说明

- SLA values are reference lines, not pass/fail gates.
