# Feed 压测报告：hybrid / rpc / cursor-page5-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:56:42+08:00
- 采样时长：1m0.0147414s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 160650 | 160650 | 0 | 0 | 2677.35 | 5.532 | 8.541 | 9.893 | 12.213 | 22.909 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：62.4055ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 134 | 0.001 |
| redis | 160784 | 1.001 |
| relation | 240 | 0.001 |

- Cold compute：160650（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3373650 | 21.000 |
| merge_candidates | 24097500 | 150.000 |
| redis_commands | 2731050 | 17.000 |
| redis_members | 37913400 | 236.000 |
| redis_roundtrips | 321300 | 2.000 |
| tie_members | 20241900 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 160650 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 2.258 |
| cursor_decode | 160650 | 0.010 |
| cursor_seek | 160650 | 3.882 |
| hydrate | 160650 | 1.836 |
| merge_dedup | 160650 | 0.008 |
| relation | 240 | 2.812 |
| route | 160650 | 0.030 |
| total | 160650 | 5.768 |

## Redis 本轮边界增量

- Commands：2927870；input：415644927 bytes；output：2159476182 bytes
- Hits/Misses：6112085/144；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：52974；safety epoch：3634 -> 3634

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.691 |
| client:loadtest | cpu_percent_total | 43.062 |
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
| docker:zg-canal | cpu_percent | 2.030 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.810 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.410 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 235.600 |
| docker:zg-kafka | memory_percent | 7.550 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 43.110 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 114169.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 14.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.933 |
| process:counter | cpu_seconds_total | 42.078 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46235648.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.188 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45015040.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 167.364 |
| process:knowpost | cpu_seconds_total | 2371.625 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 69820416.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.097 |
| process:relation | cpu_seconds_total | 4.922 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47513600.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.859 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38330368.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.556 |
| process:user-storage | cpu_seconds_total | 95.906 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44240896.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 50400173.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3634.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597579.000 |
| redis | keyspace_hits | 360240882.000 |
| redis | keyspace_misses | 100873.000 |
| redis | net_input_bytes | 15265637937.000 |
| redis | net_output_bytes | 118826442278.000 |
| redis | ops_per_sec | 52974.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23841.000 |
| redis | used_memory_bytes | 103406168.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
