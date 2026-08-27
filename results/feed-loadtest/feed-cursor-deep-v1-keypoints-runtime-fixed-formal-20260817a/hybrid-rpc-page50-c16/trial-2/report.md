# Feed 压测报告：hybrid / rpc / page50-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:27:11+08:00
- 采样时长：1m0.0156069s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 35739 | 35739 | 0 | 0 | 595.52 | 26.265 | 32.900 | 35.445 | 42.058 | 70.172 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.007 |
| mysql | 4612 | 0.129 |
| redis | 76090 | 2.129 |
| relation | 240 | 0.007 |

- Cold compute：35739（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 36453780 | 1020.000 |
| merge_candidates | 53572761 | 1499.000 |
| redis_commands | 214434 | 6.000 |
| redis_members | 53608500 | 1500.000 |
| redis_roundtrips | 35739 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 35739 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 35739 | 5.888 |
| counter | 240 | 5.093 |
| hydrate | 35739 | 20.044 |
| inbox | 35739 | 5.885 |
| merge_dedup | 35739 | 0.172 |
| relation | 240 | 5.696 |
| route | 35739 | 0.276 |
| total | 35739 | 26.560 |

## Redis 本轮边界增量

- Commands：279040；input：1295065075 bytes；output：8212557775 bytes
- Hits/Misses：36668809/6866；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：4945；safety epoch：3747 -> 3747

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.882 |
| client:loadtest | cpu_percent_total | 14.111 |
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
| docker:zg-canal | cpu_percent | 1.810 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.500 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.620 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 145.210 |
| docker:zg-kafka | memory_percent | 7.630 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 43.680 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 487487.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.643 |
| process:counter | cpu_seconds_total | 269.672 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 48013312.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 5478.469 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43823104.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 235.402 |
| process:knowpost | cpu_seconds_total | 16142.688 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 86863872.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.321 |
| process:relation | cpu_seconds_total | 44.625 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49459200.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.549 |
| process:search | cpu_seconds_total | 5.469 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38039552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 1947.938 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 46231552.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 263117754.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3747.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2413418539.000 |
| redis | keyspace_misses | 644048.000 |
| redis | net_input_bytes | 101014794641.000 |
| redis | net_output_bytes | 604523892122.000 |
| redis | ops_per_sec | 4945.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 32870.000 |
| redis | used_memory_bytes | 102844800.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
