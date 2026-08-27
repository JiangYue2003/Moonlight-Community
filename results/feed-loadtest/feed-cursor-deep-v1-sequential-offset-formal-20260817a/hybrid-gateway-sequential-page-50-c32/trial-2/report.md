# Feed 压测报告：hybrid / gateway / sequential-page-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:59:58+08:00
- 采样时长：1m0.8368457s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 33600 | 33600 | 0 | 0 | 552.32 | 54.617 | 93.046 | 103.425 | 124.816 | 202.931 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：672/33600
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 232 | 0.007 |
| mysql | 8775 | 0.261 |
| redis | 75975 | 2.261 |
| relation | 232 | 0.007 |

- Cold compute：33600（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 17808000 | 530.000 |
| merge_candidates | 30899232 | 919.620 |
| redis_commands | 201600 | 6.000 |
| redis_members | 34233696 | 1018.860 |
| redis_roundtrips | 33600 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 33600 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 33600 | 17.868 |
| counter | 232 | 13.744 |
| hydrate | 33600 | 37.105 |
| inbox | 33600 | 17.865 |
| merge_dedup | 33600 | 0.092 |
| relation | 232 | 15.559 |
| route | 33600 | 0.970 |
| total | 33600 | 56.161 |

## Redis 本轮边界增量

- Commands：258715；input：641938481 bytes；output：4353545310 bytes
- Hits/Misses：18004233/12416；run hit rate：99.93%
- Evicted/Rejected：0/0；ops/s max：5537；safety epoch：3727 -> 3727

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.246 |
| client:loadtest | cpu_percent_total | 19.930 |
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
| docker:zg-canal | cpu_percent | 0.200 |
| docker:zg-canal | memory_percent | 4.260 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.350 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.480 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 157.550 |
| docker:zg-kafka | memory_percent | 7.160 |
| docker:zg-kafka | pids | 99.000 |
| docker:zg-zk | cpu_percent | 45.200 |
| docker:zg-zk | memory_percent | 1.070 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 442471.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 8.494 |
| process:counter | cpu_seconds_total | 233.547 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46260224.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 89.318 |
| process:gateway | cpu_seconds_total | 5332.906 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 53297152.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 196.291 |
| process:knowpost | cpu_seconds_total | 14292.406 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 94605312.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.206 |
| process:relation | cpu_seconds_total | 40.219 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48345088.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.098 |
| process:search | cpu_seconds_total | 4.688 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38055936.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 28.649 |
| process:user-storage | cpu_seconds_total | 1892.125 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54538240.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 240684445.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3727.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2043050599.000 |
| redis | keyspace_misses | 558749.000 |
| redis | net_input_bytes | 86646825078.000 |
| redis | net_output_bytes | 515061984635.000 |
| redis | ops_per_sec | 5537.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31238.000 |
| redis | used_memory_bytes | 104516216.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
