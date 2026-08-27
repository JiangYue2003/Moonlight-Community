# Feed 压测报告：hybrid / gateway / cursor-page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-page5-recheck-c`
- 开始时间：2026-08-17T19:07:08+08:00
- 采样时长：10.0056167s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 16290 | 16290 | 0 | 0 | 1628.25 | 9.160 | 13.517 | 15.914 | 19.495 | 27.720 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：154.461ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 39 | 0.002 |
| redis | 16329 | 1.002 |
| relation | 40 | 0.002 |

- Cold compute：16290（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 342090 | 21.000 |
| merge_candidates | 2443500 | 150.000 |
| redis_commands | 276930 | 17.000 |
| redis_members | 3844440 | 236.000 |
| redis_roundtrips | 32580 | 2.000 |
| tie_members | 2052540 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 16290 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 2.923 |
| cursor_decode | 16290 | 0.012 |
| cursor_seek | 16290 | 5.863 |
| hydrate | 16290 | 2.930 |
| merge_dedup | 16290 | 0.008 |
| relation | 40 | 4.062 |
| route | 16290 | 0.064 |
| total | 16290 | 8.880 |

## Redis 本轮边界增量

- Commands：298076；input：42217327 bytes；output：219001978 bytes
- Hits/Misses：620184/89；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：31735；safety epoch：3736 -> 3736

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.469 |
| client:loadtest | cpu_percent_total | 39.509 |
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
| docker:zg-es | cpu_percent | 0.410 |
| docker:zg-es | memory_percent | 12.350 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.530 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 17.910 |
| docker:zg-kafka | memory_percent | 7.110 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.200 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 455118.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.644 |
| process:counter | cpu_seconds_total | 242.000 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46395392.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 129.401 |
| process:gateway | cpu_seconds_total | 5478.312 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50782208.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 158.193 |
| process:knowpost | cpu_seconds_total | 14536.641 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70926336.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.323 |
| process:relation | cpu_seconds_total | 41.531 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48467968.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.812 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38027264.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 133.641 |
| process:user-storage | cpu_seconds_total | 1946.750 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55820288.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 243024546.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3736.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 595975.000 |
| redis | keyspace_hits | 2078913340.000 |
| redis | keyspace_misses | 577792.000 |
| redis | net_input_bytes | 88039061897.000 |
| redis | net_output_bytes | 525075371141.000 |
| redis | ops_per_sec | 31735.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31617.000 |
| redis | used_memory_bytes | 101873624.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
