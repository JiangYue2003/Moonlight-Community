# Feed 压测报告：hybrid / rpc / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:24:42+08:00
- 采样时长：1m0.0204939s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 290939 | 290939 | 0 | 0 | 4848.76 | 3.141 | 4.742 | 5.379 | 6.705 | 16.465 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：311.1829ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 180 | 0.001 |
| redis | 291119 | 1.001 |
| relation | 240 | 0.001 |

- Cold compute：290939（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6109719 | 21.000 |
| merge_candidates | 18620096 | 64.000 |
| redis_commands | 4073146 | 14.000 |
| redis_members | 19492913 | 67.000 |
| redis_roundtrips | 581878 | 2.000 |
| tie_members | 872817 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 290939 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 1.470 |
| cursor_decode | 290939 | 0.009 |
| cursor_seek | 290939 | 2.054 |
| hydrate | 290939 | 1.002 |
| merge_dedup | 290939 | 0.006 |
| relation | 240 | 2.236 |
| route | 290939 | 0.012 |
| total | 290939 | 3.087 |

## Redis 本轮边界增量

- Commands：4400093；input：663591344 bytes；output：1840129133 bytes
- Hits/Misses：10189975/420；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：79246；safety epoch：3745 -> 3745

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.722 |
| client:loadtest | cpu_percent_total | 75.547 |
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
| docker:zg-canal | cpu_percent | 2.400 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.590 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.660 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 149.840 |
| docker:zg-kafka | memory_percent | 7.320 |
| docker:zg-kafka | pids | 115.000 |
| docker:zg-zk | cpu_percent | 52.850 |
| docker:zg-zk | memory_percent | 1.360 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 474911.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.418 |
| process:counter | cpu_seconds_total | 266.047 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46342144.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 5478.469 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43819008.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 242.329 |
| process:knowpost | cpu_seconds_total | 15860.469 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71454720.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 10.839 |
| process:relation | cpu_seconds_total | 44.297 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48537600.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.779 |
| process:search | cpu_seconds_total | 5.406 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38039552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 1947.812 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 46022656.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 262443157.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3745.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 595968.000 |
| redis | keyspace_hits | 2327430768.000 |
| redis | keyspace_misses | 614035.000 |
| redis | net_input_bytes | 97973745870.000 |
| redis | net_output_bytes | 585264474135.000 |
| redis | ops_per_sec | 79246.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 32721.000 |
| redis | used_memory_bytes | 101896368.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
