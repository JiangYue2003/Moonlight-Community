# Feed 压测报告：hybrid / gateway / mixed-80-20-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-19T01:21:32+08:00
- 采样时长：1m0.7497638s
- 并发：8
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1300 | 1300 | 0 | 0 | 21.40 | 3.783 | 5.354 | 6.018 | 7.234 | 8.967 |
| publish_total | 325 | 325 | 0 | 0 | 5.35 | 1477.610 | 1609.963 | 1639.456 | 1664.088 | 1672.481 |
| publish_draft | 325 | 325 | 0 | 0 | 5.35 | 4.279 | 5.236 | 5.472 | 7.093 | 14.377 |
| publish_metadata | 325 | 325 | 0 | 0 | 5.35 | 489.279 | 582.431 | 616.497 | 649.753 | 673.896 |
| publish_confirm | 325 | 325 | 0 | 0 | 5.35 | 470.274 | 585.046 | 617.054 | 656.873 | 675.130 |
| publish_commit | 325 | 325 | 0 | 0 | 5.35 | 484.969 | 562.532 | 581.409 | 622.935 | 634.744 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 200 | 0.154 |
| mysql | 114 | 0.088 |
| redis | 4014 | 3.088 |
| relation | 1300 | 1.000 |

- Cold compute：1300（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 52000 | 40.000 |
| merge_candidates | 89540 | 68.877 |
| redis_commands | 7800 | 6.000 |
| redis_members | 95905 | 73.773 |
| redis_roundtrips | 2600 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1300 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1300 | 0.544 |
| counter | 200 | 0.831 |
| hydrate | 1300 | 0.720 |
| inbox | 1300 | 0.570 |
| merge_dedup | 1300 | 0.008 |
| relation | 1300 | 1.276 |
| route | 1300 | 0.134 |
| total | 1300 | 3.288 |

## Redis 本轮边界增量

- Commands：996518；input：79270363 bytes；output：29639839 bytes
- Hits/Misses：64678/8606；run hit rate：88.26%
- Evicted/Rejected：0/0；ops/s max：18228；safety epoch：23404 -> 23729

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.167 |
| client:loadtest | cpu_percent_total | 2.675 |
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
| docker:zg-canal | cpu_percent | 4.670 |
| docker:zg-canal | memory_percent | 5.020 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.210 |
| docker:zg-es | memory_percent | 12.310 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.330 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 168.670 |
| docker:zg-kafka | memory_percent | 7.920 |
| docker:zg-kafka | pids | 133.000 |
| docker:zg-zk | cpu_percent | 0.170 |
| docker:zg-zk | memory_percent | 1.490 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 17397.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 17397.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 341918.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 17.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.971 |
| process:counter | cpu_seconds_total | 206.391 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 44875776.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 7.145 |
| process:gateway | cpu_seconds_total | 39.391 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47394816.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 88.282 |
| process:knowpost | cpu_seconds_total | 2073.141 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 67989504.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.744 |
| process:relation | cpu_seconds_total | 67.469 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47198208.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 13.173 |
| process:search | cpu_seconds_total | 32.062 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42209280.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.970 |
| process:user-storage | cpu_seconds_total | 24.375 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40620032.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 70644529.000 |
| redis | connected_clients | 43.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 23729.000 |
| redis | hit_rate | 0.813 |
| redis | keys | 590739.000 |
| redis | keyspace_hits | 4181617.000 |
| redis | keyspace_misses | 963263.000 |
| redis | net_input_bytes | 5461550590.000 |
| redis | net_output_bytes | 1876133350.000 |
| redis | ops_per_sec | 18228.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54139.000 |
| redis | used_memory_bytes | 95518616.000 |

## 停止施压后的恢复

- Kafka drain：6.6578904s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：325
- 测量前恢复：complete=true；耗时=4.8960106s；删除帖子/Outbox=0/0；safety epoch=23338
- 预热后恢复：complete=true；耗时=7.577536s；删除帖子/Outbox=64/128；safety epoch=23403
- 测量后恢复：complete=true；耗时=4.9273055s；删除帖子/Outbox=325/650；safety epoch=23730

## 说明

- SLA values are reference lines, not pass/fail gates.
