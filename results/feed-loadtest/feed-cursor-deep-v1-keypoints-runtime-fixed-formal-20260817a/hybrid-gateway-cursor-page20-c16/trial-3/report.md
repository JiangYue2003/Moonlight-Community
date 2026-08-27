# Feed 压测报告：hybrid / gateway / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:39:54+08:00
- 采样时长：1m0.0131806s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 107342 | 107342 | 0 | 0 | 1788.88 | 8.525 | 12.423 | 14.393 | 18.254 | 31.427 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：414.1094ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 140 | 0.001 |
| redis | 107482 | 1.001 |
| relation | 240 | 0.002 |

- Cold compute：107342（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2254182 | 21.000 |
| merge_candidates | 6869888 | 64.000 |
| redis_commands | 1502788 | 14.000 |
| redis_members | 7191914 | 67.000 |
| redis_roundtrips | 214684 | 2.000 |
| tie_members | 322026 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 107342 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 3.379 |
| cursor_decode | 107342 | 0.011 |
| cursor_seek | 107342 | 5.000 |
| hydrate | 107342 | 2.652 |
| merge_dedup | 107342 | 0.007 |
| relation | 240 | 4.168 |
| route | 107342 | 0.059 |
| total | 107342 | 7.733 |

## Redis 本轮边界增量

- Commands：1642511；input：246012113 bytes；output：679363970 bytes
- Hits/Misses：3764112/380；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：31357；safety epoch：3757 -> 3757

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.614 |
| client:loadtest | cpu_percent_total | 57.826 |
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
| docker:zg-canal | cpu_percent | 2.160 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.800 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.450 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 165.910 |
| docker:zg-kafka | memory_percent | 7.810 |
| docker:zg-kafka | pids | 128.000 |
| docker:zg-zk | cpu_percent | 43.820 |
| docker:zg-zk | memory_percent | 1.380 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 509281.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 8.522 |
| process:counter | cpu_seconds_total | 292.828 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46366720.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 224.279 |
| process:gateway | cpu_seconds_total | 5995.281 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51720192.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 170.328 |
| process:knowpost | cpu_seconds_total | 17625.500 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71569408.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.322 |
| process:relation | cpu_seconds_total | 47.875 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48771072.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.548 |
| process:search | cpu_seconds_total | 5.828 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 79.627 |
| process:user-storage | cpu_seconds_total | 2141.375 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 60227584.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 293214507.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3757.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596065.000 |
| redis | keyspace_hits | 2618537458.000 |
| redis | keyspace_misses | 675557.000 |
| redis | net_input_bytes | 110239464688.000 |
| redis | net_output_bytes | 647606502848.000 |
| redis | ops_per_sec | 31357.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 33633.000 |
| redis | used_memory_bytes | 101956392.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
