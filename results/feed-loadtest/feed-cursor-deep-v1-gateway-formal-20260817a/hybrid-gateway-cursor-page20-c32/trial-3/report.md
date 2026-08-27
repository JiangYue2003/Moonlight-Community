# Feed 压测报告：hybrid / gateway / cursor-page20-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:10:26+08:00
- 采样时长：1m0.0227708s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 102856 | 102856 | 0 | 0 | 1713.84 | 18.055 | 25.087 | 28.330 | 37.065 | 70.457 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：599.8569ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 176 | 0.002 |
| redis | 103032 | 1.002 |
| relation | 240 | 0.002 |

- Cold compute：102856（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2159976 | 21.000 |
| merge_candidates | 6582784 | 64.000 |
| redis_commands | 1439984 | 14.000 |
| redis_members | 6891352 | 67.000 |
| redis_roundtrips | 205712 | 2.000 |
| tie_members | 308568 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 102856 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 6.358 |
| cursor_decode | 102856 | 0.011 |
| cursor_seek | 102856 | 11.351 |
| hydrate | 102856 | 5.831 |
| merge_dedup | 102856 | 0.007 |
| relation | 240 | 7.680 |
| route | 102856 | 0.144 |
| total | 102856 | 17.348 |

## Redis 本轮边界增量

- Commands：1561259；input：234835855 bytes；output：650727536 bytes
- Hits/Misses：3607024/416；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：28793；safety epoch：3692 -> 3692

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.801 |
| client:loadtest | cpu_percent_total | 60.810 |
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
| docker:zg-canal | cpu_percent | 0.210 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.500 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.990 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 120.730 |
| docker:zg-kafka | memory_percent | 7.350 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 48.750 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 266248.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 13.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 16.131 |
| process:counter | cpu_seconds_total | 150.047 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46153728.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 241.969 |
| process:gateway | cpu_seconds_total | 3181.047 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52654080.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 166.112 |
| process:knowpost | cpu_seconds_total | 9741.594 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71794688.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.389 |
| process:relation | cpu_seconds_total | 25.297 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48353280.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 3.656 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38162432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 81.857 |
| process:user-storage | cpu_seconds_total | 1149.312 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 56209408.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 180688731.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3692.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596008.000 |
| redis | keyspace_hits | 1367763321.000 |
| redis | keyspace_misses | 298761.000 |
| redis | net_input_bytes | 59015336925.000 |
| redis | net_output_bytes | 363337735411.000 |
| redis | ops_per_sec | 28793.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28265.000 |
| redis | used_memory_bytes | 103403136.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
