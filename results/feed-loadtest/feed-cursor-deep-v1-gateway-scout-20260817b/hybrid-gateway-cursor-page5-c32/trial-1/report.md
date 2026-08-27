# Feed 压测报告：hybrid / gateway / cursor-page5-c32

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:35:22+08:00
- 采样时长：10.0216013s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 11378 | 11378 | 0 | 0 | 1135.41 | 28.256 | 40.383 | 45.684 | 59.833 | 83.567 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：141.0272ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.004 |
| mysql | 40 | 0.004 |
| redis | 11418 | 1.004 |
| relation | 40 | 0.004 |

- Cold compute：11378（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 238938 | 21.000 |
| merge_candidates | 1706700 | 150.000 |
| redis_commands | 193426 | 17.000 |
| redis_members | 2685208 | 236.000 |
| redis_roundtrips | 22756 | 2.000 |
| tie_members | 1433628 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 11378 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 7.557 |
| cursor_decode | 11378 | 0.012 |
| cursor_seek | 11378 | 17.612 |
| hydrate | 11378 | 8.855 |
| merge_dedup | 11378 | 0.010 |
| relation | 40 | 8.032 |
| route | 11378 | 0.267 |
| total | 11378 | 26.757 |

## Redis 本轮边界增量

- Commands：207296；input：29416504 bytes；output：152954162 bytes
- Hits/Misses：433572/55；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：21871；safety epoch：3612 -> 3612

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.855 |
| client:loadtest | cpu_percent_total | 45.683 |
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
| docker:zg-canal | cpu_percent | 0.140 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.320 |
| docker:zg-es | memory_percent | 12.070 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.240 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 121.720 |
| docker:zg-kafka | memory_percent | 7.030 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 28.410 |
| docker:zg-zk | memory_percent | 1.190 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 93889.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.871 |
| process:counter | cpu_seconds_total | 7.016 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 44867584.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 154.790 |
| process:gateway | cpu_seconds_total | 122.766 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52457472.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 132.378 |
| process:knowpost | cpu_seconds_total | 121.859 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 66842624.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.300 |
| process:relation | cpu_seconds_total | 0.656 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 45912064.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.860 |
| process:search | cpu_seconds_total | 0.344 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 41611264.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 50.319 |
| process:user-storage | cpu_seconds_total | 44.828 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 50761728.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18737294.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3612.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597359.000 |
| redis | keyspace_hits | 94737560.000 |
| redis | keyspace_misses | 64415.000 |
| redis | net_input_bytes | 4557237178.000 |
| redis | net_output_bytes | 25413066245.000 |
| redis | ops_per_sec | 21871.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22512.000 |
| redis | used_memory_bytes | 102475728.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
