# Feed 压测报告：hybrid / gateway / page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:34:48+08:00
- 采样时长：1m0.0114512s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 62719 | 62719 | 0 | 0 | 1045.19 | 13.919 | 22.486 | 25.413 | 31.215 | 55.559 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.004 |
| mysql | 1851 | 0.030 |
| redis | 127289 | 2.030 |
| relation | 240 | 0.004 |

- Cold compute：62719（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 26341980 | 420.000 |
| merge_candidates | 52746679 | 841.000 |
| redis_commands | 376314 | 6.000 |
| redis_members | 57764199 | 921.000 |
| redis_roundtrips | 62719 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 62719 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 62719 | 7.108 |
| counter | 240 | 5.634 |
| hydrate | 62719 | 7.012 |
| inbox | 62719 | 7.106 |
| merge_dedup | 62719 | 0.084 |
| relation | 240 | 6.442 |
| route | 62719 | 0.212 |
| total | 62719 | 14.532 |

## Redis 本轮边界增量

- Commands：459907；input：952139553 bytes；output：6748949917 bytes
- Hits/Misses：26723152/2615；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：8683；safety epoch：3753 -> 3753

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.522 |
| client:loadtest | cpu_percent_total | 24.344 |
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
| docker:zg-canal | cpu_percent | 2.090 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.670 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.600 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 148.630 |
| docker:zg-kafka | memory_percent | 7.620 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 48.400 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 503977.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.388 |
| process:counter | cpu_seconds_total | 280.312 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47001600.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 84.857 |
| process:gateway | cpu_seconds_total | 5574.547 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51507200.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 241.746 |
| process:knowpost | cpu_seconds_total | 17184.625 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 77946880.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.114 |
| process:relation | cpu_seconds_total | 46.219 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 50565120.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.779 |
| process:search | cpu_seconds_total | 5.609 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 40.489 |
| process:user-storage | cpu_seconds_total | 1991.375 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54165504.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 287446117.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3753.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596370.000 |
| redis | keyspace_hits | 2575553505.000 |
| redis | keyspace_misses | 669035.000 |
| redis | net_input_bytes | 108351247530.000 |
| redis | net_output_bytes | 637605766680.000 |
| redis | ops_per_sec | 8683.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 33327.000 |
| redis | used_memory_bytes | 102153960.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
