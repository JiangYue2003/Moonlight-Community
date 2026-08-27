# Feed 压测报告：hybrid / gateway / sequential-page-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 开始时间：2026-08-17T19:59:28+08:00
- 采样时长：1m0.7385113s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 48700 | 48700 | 0 | 0 | 801.84 | 18.247 | 32.460 | 36.980 | 46.258 | 76.558 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：974/48700
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 225 | 0.005 |
| mysql | 4835 | 0.099 |
| redis | 102235 | 2.099 |
| relation | 225 | 0.005 |

- Cold compute：48700（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 25811000 | 530.000 |
| merge_candidates | 44785494 | 919.620 |
| redis_commands | 292200 | 6.000 |
| redis_members | 49618482 | 1018.860 |
| redis_roundtrips | 48700 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 48700 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 48700 | 6.244 |
| counter | 225 | 5.504 |
| hydrate | 48700 | 12.335 |
| inbox | 48700 | 6.241 |
| merge_dedup | 48700 | 0.101 |
| relation | 225 | 5.673 |
| route | 48700 | 0.170 |
| total | 48700 | 18.980 |

## Redis 本轮边界增量

- Commands：365583；input：927937547 bytes；output：6311925593 bytes
- Hits/Misses：26104285/5621；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：7365；safety epoch：3771 -> 3771

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.318 |
| client:loadtest | cpu_percent_total | 21.095 |
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
| docker:zg-canal | cpu_percent | 2.470 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.390 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.510 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 106.040 |
| docker:zg-kafka | memory_percent | 7.560 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 46.110 |
| docker:zg-zk | memory_percent | 1.390 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 564299.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 5.417 |
| process:counter | cpu_seconds_total | 329.219 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45129728.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 106.163 |
| process:gateway | cpu_seconds_total | 6648.219 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51814400.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 236.826 |
| process:knowpost | cpu_seconds_total | 19543.812 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 80093184.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.098 |
| process:relation | cpu_seconds_total | 52.781 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48984064.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 6.328 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38064128.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 34.827 |
| process:user-storage | cpu_seconds_total | 2378.625 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54657024.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 318887531.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3771.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2955046691.000 |
| redis | keyspace_misses | 765058.000 |
| redis | net_input_bytes | 123708660153.000 |
| redis | net_output_bytes | 723465824276.000 |
| redis | ops_per_sec | 7365.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34808.000 |
| redis | used_memory_bytes | 102525264.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
