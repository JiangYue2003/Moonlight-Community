# Feed 压测报告：hybrid / gateway / same-second-cursor-c1

- Run ID：`feed-cursor-deep-v1-same-second-probe-20260817a`
- 开始时间：2026-08-17T19:03:49+08:00
- 采样时长：5.0012489s
- 并发：1
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4085 | 4085 | 0 | 0 | 816.88 | 1.068 | 1.614 | 1.651 | 2.071 | 12.176 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：2
- 准备请求：20；准备耗时：25.9467ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 20 | 0.005 |
| mysql | 0 | 0.000 |
| redis | 4085 | 1.000 |
| relation | 20 | 0.005 |

- Cold compute：4085（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 85785 | 21.000 |
| merge_candidates | 857850 | 210.000 |
| redis_commands | 69445 | 17.000 |
| redis_members | 964060 | 236.000 |
| redis_roundtrips | 8170 | 2.000 |
| tie_members | 514710 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 4085 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 20 | 0.395 |
| cursor_decode | 4085 | 0.008 |
| cursor_seek | 4085 | 0.461 |
| hydrate | 4085 | 0.096 |
| merge_dedup | 4085 | 0.008 |
| relation | 20 | 0.847 |
| route | 4085 | 0.011 |
| total | 4085 | 0.585 |

## Redis 本轮边界增量

- Commands：76527；input：10698517 bytes；output：54962182 bytes
- Hits/Misses：155838/20；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：15512；safety epoch：3730 -> 3730

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.015 |
| client:loadtest | cpu_percent_total | 16.246 |
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
| docker:zg-canal | cpu_percent | 1.820 |
| docker:zg-canal | memory_percent | 4.260 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.370 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.930 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 14.240 |
| docker:zg-kafka | memory_percent | 7.110 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.100 |
| docker:zg-zk | memory_percent | 1.080 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 453764.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 8.807 |
| process:counter | cpu_seconds_total | 236.781 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46301184.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 60.184 |
| process:gateway | cpu_seconds_total | 5382.781 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50061312.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 49.412 |
| process:knowpost | cpu_seconds_total | 14422.469 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71483392.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.810 |
| process:relation | cpu_seconds_total | 40.969 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48914432.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.750 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38055936.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 29.456 |
| process:user-storage | cpu_seconds_total | 1910.828 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 50864128.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 241301480.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3730.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596333.000 |
| redis | keyspace_hits | 2068102471.000 |
| redis | keyspace_misses | 574368.000 |
| redis | net_input_bytes | 87557647932.000 |
| redis | net_output_bytes | 521171108924.000 |
| redis | ops_per_sec | 15512.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31413.000 |
| redis | used_memory_bytes | 100848048.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
