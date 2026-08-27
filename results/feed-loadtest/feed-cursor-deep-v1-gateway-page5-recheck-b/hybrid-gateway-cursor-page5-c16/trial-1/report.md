# Feed 压测报告：hybrid / gateway / cursor-page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-page5-recheck-b`
- 开始时间：2026-08-17T19:05:47+08:00
- 采样时长：10.0068419s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 14379 | 14379 | 0 | 0 | 1437.08 | 10.544 | 15.336 | 18.250 | 22.612 | 33.126 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：140.689ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.003 |
| mysql | 87 | 0.006 |
| redis | 14466 | 1.006 |
| relation | 40 | 0.003 |

- Cold compute：14379（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 301959 | 21.000 |
| merge_candidates | 2156850 | 150.000 |
| redis_commands | 244443 | 17.000 |
| redis_members | 3393444 | 236.000 |
| redis_roundtrips | 28758 | 2.000 |
| tie_members | 1811754 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 14379 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 3.597 |
| cursor_decode | 14379 | 0.011 |
| cursor_seek | 14379 | 6.658 |
| hydrate | 14379 | 3.322 |
| merge_dedup | 14379 | 0.008 |
| relation | 40 | 4.536 |
| route | 14379 | 0.080 |
| total | 14379 | 10.081 |

## Redis 本轮边界增量

- Commands：263366；input：37285502 bytes；output：193310279 bytes
- Hits/Misses：547521/133；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：27930；safety epoch：3733 -> 3733

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.830 |
| client:loadtest | cpu_percent_total | 45.282 |
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
| docker:zg-canal | cpu_percent | 1.830 |
| docker:zg-canal | memory_percent | 4.270 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.610 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.180 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 111.340 |
| docker:zg-kafka | memory_percent | 7.520 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 41.720 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 454375.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 9.835 |
| process:counter | cpu_seconds_total | 239.875 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45932544.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 236.045 |
| process:gateway | cpu_seconds_total | 5432.016 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50716672.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 137.693 |
| process:knowpost | cpu_seconds_total | 14481.266 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70873088.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.550 |
| process:relation | cpu_seconds_total | 41.141 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48996352.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.781 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38027264.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 49.595 |
| process:user-storage | cpu_seconds_total | 1929.078 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 52768768.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 242322075.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3733.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596083.000 |
| redis | keyspace_hits | 2073067080.000 |
| redis | keyspace_misses | 576338.000 |
| redis | net_input_bytes | 87795559259.000 |
| redis | net_output_bytes | 522959071764.000 |
| redis | ops_per_sec | 27930.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31536.000 |
| redis | used_memory_bytes | 102684024.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
