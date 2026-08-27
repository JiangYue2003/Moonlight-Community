# Feed 压测报告：hybrid / rpc / cursor-page20-c64

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:21:09+08:00
- 采样时长：10.0123933s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 47121 | 47121 | 0 | 0 | 4707.75 | 13.044 | 17.372 | 19.340 | 25.855 | 42.917 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：341.8534ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 47121 | 1.000 |
| relation | 40 | 0.001 |

- Cold compute：47121（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 47121 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 4.948 |
| cursor_decode | 47121 | 0.010 |
| cursor_seek | 47121 | 8.856 |
| hydrate | 47121 | 4.438 |
| merge_dedup | 47121 | 0.007 |
| relation | 40 | 5.604 |
| route | 47121 | 0.055 |
| total | 47121 | 13.370 |

## Redis 本轮边界增量

- Commands：710602；input：107331975 bytes；output：297999291 bytes
- Hits/Misses：1650486/15；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：76847；safety epoch：3600 -> 3600

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.998 |
| client:loadtest | cpu_percent_total | 95.975 |
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
| docker:zg-es | cpu_percent | 2.870 |
| docker:zg-es | memory_percent | 11.970 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.780 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 69.000 |
| docker:zg-kafka | memory_percent | 7.020 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 17.770 |
| docker:zg-zk | memory_percent | 1.100 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 86401.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 26.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 2.323 |
| process:counter | cpu_seconds_total | 36.828 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 47591424.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 37494784.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 281.821 |
| process:knowpost | cpu_seconds_total | 619.422 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 80240640.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.321 |
| process:relation | cpu_seconds_total | 17.344 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 53915648.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 2.359 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42962944.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.344 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40947712.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12774102.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3600.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597364.000 |
| redis | keyspace_hits | 52437026.000 |
| redis | keyspace_misses | 38862.000 |
| redis | net_input_bytes | 2685593386.000 |
| redis | net_output_bytes | 15667753650.000 |
| redis | ops_per_sec | 76847.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21658.000 |
| redis | used_memory_bytes | 105046960.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
