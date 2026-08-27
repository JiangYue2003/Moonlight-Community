# Feed 压测报告：hybrid / rpc / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:22:09+08:00
- 采样时长：1m0.0197899s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 290962 | 290962 | 0 | 0 | 4849.21 | 3.142 | 4.746 | 5.381 | 6.803 | 11.911 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：366.6542ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 133 | 0.000 |
| redis | 291095 | 1.000 |
| relation | 240 | 0.001 |

- Cold compute：290962（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6110202 | 21.000 |
| merge_candidates | 18621568 | 64.000 |
| redis_commands | 4073468 | 14.000 |
| redis_members | 19494454 | 67.000 |
| redis_roundtrips | 581924 | 2.000 |
| tie_members | 872886 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 290962 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 1.317 |
| cursor_decode | 290962 | 0.010 |
| cursor_seek | 290962 | 2.053 |
| hydrate | 290962 | 1.004 |
| merge_dedup | 290962 | 0.006 |
| relation | 240 | 2.020 |
| route | 290962 | 0.012 |
| total | 290962 | 3.088 |

## Redis 本轮边界增量

- Commands：4400330；input：663628612 bytes；output：1840280737 bytes
- Hits/Misses：10190826/373；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：80035；safety epoch：3743 -> 3743

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.639 |
| client:loadtest | cpu_percent_total | 74.220 |
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
| docker:zg-canal | cpu_percent | 0.150 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.360 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.910 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 152.850 |
| docker:zg-kafka | memory_percent | 7.630 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 45.750 |
| docker:zg-zk | memory_percent | 1.380 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 473598.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.196 |
| process:counter | cpu_seconds_total | 261.781 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46219264.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 5478.469 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43765760.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 236.930 |
| process:knowpost | cpu_seconds_total | 15579.469 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72605696.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.323 |
| process:relation | cpu_seconds_total | 43.656 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48619520.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.549 |
| process:search | cpu_seconds_total | 5.266 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38039552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 1947.734 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 46116864.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 252101379.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3743.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 595974.000 |
| redis | keyspace_hits | 2303498484.000 |
| redis | keyspace_misses | 612324.000 |
| redis | net_input_bytes | 96414660539.000 |
| redis | net_output_bytes | 580938271304.000 |
| redis | ops_per_sec | 80035.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 32568.000 |
| redis | used_memory_bytes | 101938120.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
