# Feed 压测报告：hybrid / gateway / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:36:20+08:00
- 采样时长：10.006941s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 20719 | 20719 | 0 | 0 | 2070.68 | 7.357 | 10.516 | 12.414 | 15.301 | 26.406 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：408.2905ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 0 | 0.000 |
| redis | 20719 | 1.000 |
| relation | 40 | 0.002 |

- Cold compute：20719（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 435099 | 21.000 |
| merge_candidates | 1326016 | 64.000 |
| redis_commands | 290066 | 14.000 |
| redis_members | 1388173 | 67.000 |
| redis_roundtrips | 41438 | 2.000 |
| tie_members | 62157 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 20719 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 3.036 |
| cursor_decode | 20719 | 0.012 |
| cursor_seek | 20719 | 4.292 |
| hydrate | 20719 | 2.279 |
| merge_dedup | 20719 | 0.007 |
| relation | 40 | 3.784 |
| route | 20719 | 0.045 |
| total | 20719 | 6.640 |

## Redis 本轮边界增量

- Commands：316832；input：47481140 bytes；output：131130237 bytes
- Hits/Misses：726420/18；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：33535；safety epoch：3615 -> 3615

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.533 |
| client:loadtest | cpu_percent_total | 56.523 |
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
| docker:zg-canal | cpu_percent | 1.950 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.270 |
| docker:zg-es | memory_percent | 12.080 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.900 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 133.810 |
| docker:zg-kafka | memory_percent | 7.030 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 10.250 |
| docker:zg-zk | memory_percent | 1.080 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 94662.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 26.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.206 |
| process:counter | cpu_seconds_total | 8.609 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46153728.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 181.102 |
| process:gateway | cpu_seconds_total | 163.375 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51290112.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 167.803 |
| process:knowpost | cpu_seconds_total | 185.281 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 66748416.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.322 |
| process:relation | cpu_seconds_total | 1.000 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 46215168.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.489 |
| process:search | cpu_seconds_total | 0.391 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 41848832.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 82.064 |
| process:user-storage | cpu_seconds_total | 61.875 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 50708480.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19355795.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3615.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597677.000 |
| redis | keyspace_hits | 105321553.000 |
| redis | keyspace_misses | 70363.000 |
| redis | net_input_bytes | 4966162590.000 |
| redis | net_output_bytes | 28016551163.000 |
| redis | ops_per_sec | 33535.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22569.000 |
| redis | used_memory_bytes | 102059168.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
