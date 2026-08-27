# Feed 压测报告：hybrid / rpc / page5-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:47:57+08:00
- 采样时长：1m0.0139764s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 140348 | 140348 | 0 | 0 | 2339.01 | 6.126 | 10.241 | 11.273 | 13.581 | 23.750 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 541 | 0.004 |
| redis | 281237 | 2.004 |
| relation | 240 | 0.002 |

- Cold compute：140348（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 16841760 | 120.000 |
| merge_candidates | 33823868 | 241.000 |
| redis_commands | 842088 | 6.000 |
| redis_members | 87156108 | 621.000 |
| redis_roundtrips | 140348 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 140348 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 140348 | 3.336 |
| counter | 240 | 3.150 |
| hydrate | 140348 | 3.120 |
| inbox | 140348 | 3.334 |
| merge_dedup | 140348 | 0.032 |
| relation | 240 | 3.561 |
| route | 140348 | 0.053 |
| total | 140348 | 6.589 |

## Redis 本轮边界增量

- Commands：1014117；input：655537253 bytes；output：6429161894 bytes
- Hits/Misses：17690826/634；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：17892；safety epoch：3627 -> 3627

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.812 |
| client:loadtest | cpu_percent_total | 44.990 |
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
| docker:zg-canal | cpu_percent | 2.070 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 6.490 |
| docker:zg-es | memory_percent | 12.200 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.460 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 130.780 |
| docker:zg-kafka | memory_percent | 7.490 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 52.610 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 103743.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 31.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.200 |
| process:counter | cpu_seconds_total | 29.188 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45625344.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.125 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45133824.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 277.141 |
| process:knowpost | cpu_seconds_total | 1422.172 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71041024.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 2.562 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47607808.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.548 |
| process:search | cpu_seconds_total | 1.359 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42414080.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.549 |
| process:user-storage | cpu_seconds_total | 95.172 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44326912.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 37574476.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3627.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597817.000 |
| redis | keyspace_hits | 241067637.000 |
| redis | keyspace_misses | 91960.000 |
| redis | net_input_bytes | 10408830238.000 |
| redis | net_output_bytes | 75660359391.000 |
| redis | ops_per_sec | 17892.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23316.000 |
| redis | used_memory_bytes | 102449096.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
