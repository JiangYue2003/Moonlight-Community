# Feed 压测报告：hybrid / gateway / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:05:18+08:00
- 采样时长：1m0.0128491s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 89425 | 89425 | 0 | 0 | 1490.25 | 10.406 | 15.036 | 17.531 | 22.409 | 43.486 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：437.967ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 87 | 0.001 |
| redis | 89512 | 1.001 |
| relation | 240 | 0.003 |

- Cold compute：89425（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1877925 | 21.000 |
| merge_candidates | 5723200 | 64.000 |
| redis_commands | 1251950 | 14.000 |
| redis_members | 5991475 | 67.000 |
| redis_roundtrips | 178850 | 2.000 |
| tie_members | 268275 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 89425 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 3.744 |
| cursor_decode | 89425 | 0.011 |
| cursor_seek | 89425 | 6.095 |
| hydrate | 89425 | 3.219 |
| merge_dedup | 89425 | 0.007 |
| relation | 240 | 4.827 |
| route | 89425 | 0.079 |
| total | 89425 | 9.415 |

## Redis 本轮边界增量

- Commands：1369163；input：204954590 bytes；output：566012285 bytes
- Hits/Misses：3137049/327；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：26187；safety epoch：3688 -> 3688

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.162 |
| client:loadtest | cpu_percent_total | 50.588 |
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
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.710 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.790 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 148.050 |
| docker:zg-kafka | memory_percent | 7.580 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 44.030 |
| docker:zg-zk | memory_percent | 1.190 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 263608.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.068 |
| process:counter | cpu_seconds_total | 139.750 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46211072.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 194.184 |
| process:gateway | cpu_seconds_total | 2646.734 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51118080.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 178.352 |
| process:knowpost | cpu_seconds_total | 9321.094 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70037504.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.874 |
| process:relation | cpu_seconds_total | 23.125 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48713728.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.550 |
| process:search | cpu_seconds_total | 3.594 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38174720.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 74.290 |
| process:user-storage | cpu_seconds_total | 972.250 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54853632.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 173243476.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3688.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596160.000 |
| redis | keyspace_hits | 1350649423.000 |
| redis | keyspace_misses | 295330.000 |
| redis | net_input_bytes | 57898393349.000 |
| redis | net_output_bytes | 360240904324.000 |
| redis | ops_per_sec | 26187.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27957.000 |
| redis | used_memory_bytes | 102157432.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
