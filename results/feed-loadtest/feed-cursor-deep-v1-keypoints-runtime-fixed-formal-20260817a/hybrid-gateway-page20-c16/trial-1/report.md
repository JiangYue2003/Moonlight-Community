# Feed 压测报告：hybrid / gateway / page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:33:33+08:00
- 采样时长：1m0.0167018s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 64960 | 64960 | 0 | 0 | 1082.45 | 13.332 | 21.925 | 24.140 | 28.458 | 45.992 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.004 |
| mysql | 1578 | 0.024 |
| redis | 131498 | 2.024 |
| relation | 240 | 0.004 |

- Cold compute：64960（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 27283200 | 420.000 |
| merge_candidates | 54631360 | 841.000 |
| redis_commands | 389760 | 6.000 |
| redis_members | 59828160 | 921.000 |
| redis_roundtrips | 64960 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 64960 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 64960 | 6.906 |
| counter | 240 | 5.386 |
| hydrate | 64960 | 6.710 |
| inbox | 64960 | 6.904 |
| merge_dedup | 64960 | 0.087 |
| relation | 240 | 6.424 |
| route | 64960 | 0.207 |
| total | 64960 | 14.026 |

## Redis 本轮边界增量

- Commands：475998；input：986101838 bytes；output：6990131706 bytes
- Hits/Misses：27678011/2423；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：8700；safety epoch：3752 -> 3752

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.557 |
| client:loadtest | cpu_percent_total | 24.915 |
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
| docker:zg-canal | cpu_percent | 0.170 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.560 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.560 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 159.180 |
| docker:zg-kafka | memory_percent | 7.620 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 52.560 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 500598.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.192 |
| process:counter | cpu_seconds_total | 278.812 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45953024.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 91.342 |
| process:gateway | cpu_seconds_total | 5527.016 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51761152.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 252.317 |
| process:knowpost | cpu_seconds_total | 17042.047 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 76275712.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.390 |
| process:relation | cpu_seconds_total | 45.953 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49799168.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 5.594 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 46.469 |
| process:user-storage | cpu_seconds_total | 1970.672 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54714368.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 286897754.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3752.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596522.000 |
| redis | keyspace_hits | 2543944756.000 |
| redis | keyspace_misses | 664968.000 |
| redis | net_input_bytes | 107224527627.000 |
| redis | net_output_bytes | 629622775640.000 |
| redis | ops_per_sec | 8700.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 33252.000 |
| redis | used_memory_bytes | 102241456.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
