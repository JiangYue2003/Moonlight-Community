# Feed 压测报告：hybrid / gateway / cursor-page5-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:55:16+08:00
- 采样时长：1m0.0179172s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 84494 | 84494 | 0 | 0 | 1407.96 | 21.994 | 30.062 | 33.822 | 45.816 | 78.187 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：140.6011ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 197 | 0.002 |
| redis | 84691 | 1.002 |
| relation | 240 | 0.003 |

- Cold compute：84494（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1774374 | 21.000 |
| merge_candidates | 12674100 | 150.000 |
| redis_commands | 1436398 | 17.000 |
| redis_members | 19940584 | 236.000 |
| redis_roundtrips | 168988 | 2.000 |
| tie_members | 10646244 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 84494 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 6.895 |
| cursor_decode | 84494 | 0.011 |
| cursor_seek | 84494 | 14.258 |
| hydrate | 84494 | 7.109 |
| merge_dedup | 84494 | 0.009 |
| relation | 240 | 7.920 |
| route | 84494 | 0.193 |
| total | 84494 | 21.582 |

## Redis 本轮边界增量

- Commands：1537494；input：218322101 bytes；output：1135774775 bytes
- Hits/Misses：3218045/197；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：28345；safety epoch：3680 -> 3680

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.745 |
| client:loadtest | cpu_percent_total | 43.919 |
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
| docker:zg-canal | cpu_percent | 0.190 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.350 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.950 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 144.950 |
| docker:zg-kafka | memory_percent | 7.570 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 28.450 |
| docker:zg-zk | memory_percent | 1.010 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 235953.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.421 |
| process:counter | cpu_seconds_total | 126.391 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46034944.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 177.021 |
| process:gateway | cpu_seconds_total | 2106.609 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52547584.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 156.235 |
| process:knowpost | cpu_seconds_total | 8342.766 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71573504.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.750 |
| process:relation | cpu_seconds_total | 20.469 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48635904.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.550 |
| process:search | cpu_seconds_total | 3.344 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38158336.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 55.692 |
| process:user-storage | cpu_seconds_total | 762.000 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 56512512.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 166500099.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3680.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596886.000 |
| redis | keyspace_hits | 1169387534.000 |
| redis | keyspace_misses | 259669.000 |
| redis | net_input_bytes | 51179961821.000 |
| redis | net_output_bytes | 315075851599.000 |
| redis | ops_per_sec | 28345.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27355.000 |
| redis | used_memory_bytes | 105703312.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
