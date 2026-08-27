# Feed 压测报告：hybrid / gateway / cursor-page50-c16

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:37:38+08:00
- 采样时长：10.005662s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 22542 | 22542 | 0 | 0 | 2253.28 | 6.876 | 9.986 | 11.142 | 14.375 | 20.385 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：1.3618214s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 0 | 0.000 |
| redis | 22542 | 1.000 |
| relation | 40 | 0.002 |

- Cold compute：22542（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 473382 | 21.000 |
| merge_candidates | 495924 | 22.000 |
| redis_commands | 293046 | 13.000 |
| redis_members | 541008 | 24.000 |
| redis_roundtrips | 45084 | 2.000 |
| tie_members | 45084 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 22542 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 2.504 |
| cursor_decode | 22542 | 0.011 |
| cursor_seek | 22542 | 3.750 |
| hydrate | 22542 | 2.056 |
| merge_dedup | 22542 | 0.005 |
| relation | 40 | 3.548 |
| route | 22542 | 0.037 |
| total | 22542 | 5.864 |

## Redis 本轮边界增量

- Commands：321380；input：49351399 bytes；output：101330776 bytes
- Hits/Misses：767683/15；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：33739；safety epoch：3619 -> 3619

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.202 |
| client:loadtest | cpu_percent_total | 83.234 |
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
| docker:zg-canal | cpu_percent | 0.130 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.390 |
| docker:zg-es | memory_percent | 12.080 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 0.660 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 116.720 |
| docker:zg-kafka | memory_percent | 7.030 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.170 |
| docker:zg-zk | memory_percent | 0.930 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 96650.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 11.576 |
| process:counter | cpu_seconds_total | 11.031 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45408256.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 227.864 |
| process:gateway | cpu_seconds_total | 227.172 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50831360.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 221.348 |
| process:knowpost | cpu_seconds_total | 274.844 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 67100672.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.775 |
| process:relation | cpu_seconds_total | 1.391 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 46030848.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.543 |
| process:search | cpu_seconds_total | 0.531 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42110976.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 110.674 |
| process:user-storage | cpu_seconds_total | 83.875 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 50774016.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 20300083.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3619.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 598274.000 |
| redis | keyspace_hits | 118240021.000 |
| redis | keyspace_misses | 85209.000 |
| redis | net_input_bytes | 5484193913.000 |
| redis | net_output_bytes | 30778355590.000 |
| redis | ops_per_sec | 33739.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22647.000 |
| redis | used_memory_bytes | 102216016.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
