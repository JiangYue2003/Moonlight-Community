# Feed 压测报告：hybrid / gateway / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 开始时间：2026-08-17T20:02:02+08:00
- 采样时长：1m0.2228745s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 132700 | 132700 | 0 | 0 | 2203.83 | 6.882 | 10.033 | 11.605 | 14.882 | 25.980 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：2654/132700
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 239 | 0.002 |
| mysql | 1127 | 0.008 |
| redis | 136481 | 1.028 |
| relation | 239 | 0.002 |

- Cold compute：132700（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2837126 | 21.380 |
| merge_candidates | 7704562 | 58.060 |
| redis_commands | 1871070 | 14.100 |
| redis_members | 9681792 | 72.960 |
| redis_roundtrips | 262746 | 1.980 |
| tie_members | 2335520 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 132700 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2654 | 2.303 |
| counter | 239 | 2.722 |
| cursor_decode | 130046 | 0.011 |
| cursor_seek | 130046 | 3.995 |
| hydrate | 132700 | 2.109 |
| inbox | 2654 | 2.302 |
| merge_dedup | 132700 | 0.006 |
| relation | 239 | 3.458 |
| route | 132700 | 0.032 |
| total | 132700 | 6.126 |

## Redis 本轮边界增量

- Commands：2040856；input：306504956 bytes；output：879701465 bytes
- Hits/Misses：4714019/1486；run hit rate：99.97%
- Evicted/Rejected：0/0；ops/s max：37442；safety epoch：3773 -> 3773

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.012 |
| client:loadtest | cpu_percent_total | 64.189 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.510 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.520 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 178.950 |
| docker:zg-kafka | memory_percent | 7.650 |
| docker:zg-kafka | pids | 120.000 |
| docker:zg-zk | cpu_percent | 49.140 |
| docker:zg-zk | memory_percent | 1.480 |
| docker:zg-zk | pids | 107.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 572383.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 17.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 9.005 |
| process:counter | cpu_seconds_total | 334.109 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46350336.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 239.427 |
| process:gateway | cpu_seconds_total | 6840.500 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51884032.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 186.793 |
| process:knowpost | cpu_seconds_total | 19779.219 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71688192.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.094 |
| process:relation | cpu_seconds_total | 53.703 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49512448.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 6.344 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38064128.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 92.056 |
| process:user-storage | cpu_seconds_total | 2444.828 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 59289600.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 321757139.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3773.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2990264888.000 |
| redis | keyspace_misses | 773855.000 |
| redis | net_input_bytes | 125127882813.000 |
| redis | net_output_bytes | 731669427180.000 |
| redis | ops_per_sec | 37442.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34961.000 |
| redis | used_memory_bytes | 102459576.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
