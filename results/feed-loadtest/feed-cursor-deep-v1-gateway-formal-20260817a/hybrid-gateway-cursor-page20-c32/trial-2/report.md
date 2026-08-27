# Feed 压测报告：hybrid / gateway / cursor-page20-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:09:09+08:00
- 采样时长：1m0.020086s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 100623 | 100623 | 0 | 0 | 1676.71 | 18.387 | 26.058 | 29.534 | 37.892 | 63.454 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：551.0601ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 224 | 0.002 |
| redis | 100847 | 1.002 |
| relation | 240 | 0.002 |

- Cold compute：100623（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2113083 | 21.000 |
| merge_candidates | 6439872 | 64.000 |
| redis_commands | 1408722 | 14.000 |
| redis_members | 6741741 | 67.000 |
| redis_roundtrips | 201246 | 2.000 |
| tie_members | 301869 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 100623 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 6.515 |
| cursor_decode | 100623 | 0.011 |
| cursor_seek | 100623 | 11.617 |
| hydrate | 100623 | 5.971 |
| merge_dedup | 100623 | 0.007 |
| relation | 240 | 7.440 |
| route | 100623 | 0.147 |
| total | 100623 | 17.758 |

## Redis 本轮边界增量

- Commands：1527910；input：229778032 bytes；output：636604812 bytes
- Hits/Misses：3528820/464；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：29394；safety epoch：3691 -> 3691

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.581 |
| client:loadtest | cpu_percent_total | 57.299 |
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
| docker:zg-canal | cpu_percent | 0.180 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.530 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.850 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 155.950 |
| docker:zg-kafka | memory_percent | 7.580 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 46.940 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 265585.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 17.104 |
| process:counter | cpu_seconds_total | 147.969 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46661632.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 213.747 |
| process:gateway | cpu_seconds_total | 3045.984 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52719616.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 176.554 |
| process:knowpost | cpu_seconds_total | 9635.031 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71598080.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.651 |
| process:relation | cpu_seconds_total | 24.781 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48349184.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 3.625 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38158336.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 79.640 |
| process:user-storage | cpu_seconds_total | 1105.016 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 59367424.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178835207.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3691.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596032.000 |
| redis | keyspace_hits | 1363493173.000 |
| redis | keyspace_misses | 297907.000 |
| redis | net_input_bytes | 58736940200.000 |
| redis | net_output_bytes | 362565111838.000 |
| redis | ops_per_sec | 29394.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28188.000 |
| redis | used_memory_bytes | 103612360.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
