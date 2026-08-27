# Feed 压测报告：hybrid / gateway / cursor-page50-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:25:42+08:00
- 采样时长：1m0.0206806s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 119554 | 119554 | 0 | 0 | 1992.18 | 15.443 | 21.687 | 24.311 | 31.492 | 62.876 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：1.3602087s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 208 | 0.002 |
| redis | 119762 | 1.002 |
| relation | 240 | 0.002 |

- Cold compute：119554（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2510634 | 21.000 |
| merge_candidates | 2630188 | 22.000 |
| redis_commands | 1554202 | 13.000 |
| redis_members | 2869296 | 24.000 |
| redis_roundtrips | 239108 | 2.000 |
| tie_members | 239108 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 119554 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 5.747 |
| cursor_decode | 119554 | 0.011 |
| cursor_seek | 119554 | 9.581 |
| hydrate | 119554 | 4.894 |
| merge_dedup | 119554 | 0.005 |
| relation | 240 | 6.590 |
| route | 119554 | 0.105 |
| total | 119554 | 14.603 |

## Redis 本轮边界增量

- Commands：1694428；input：260998892 bytes；output：537188663 bytes
- Hits/Misses：4071886/448；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：33300；safety epoch：3704 -> 3704

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.901 |
| client:loadtest | cpu_percent_total | 78.410 |
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
| docker:zg-canal | cpu_percent | 2.380 |
| docker:zg-canal | memory_percent | 4.240 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 0.750 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.770 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 160.370 |
| docker:zg-kafka | memory_percent | 7.610 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 50.540 |
| docker:zg-zk | memory_percent | 1.220 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 327028.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 25.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.066 |
| process:counter | cpu_seconds_total | 181.609 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45277184.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 267.172 |
| process:gateway | cpu_seconds_total | 4321.891 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52441088.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 230.739 |
| process:knowpost | cpu_seconds_total | 11223.516 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71892992.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 9.989 |
| process:relation | cpu_seconds_total | 30.703 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48922624.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.016 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38035456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 95.032 |
| process:user-storage | cpu_seconds_total | 1533.312 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55033856.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 193641840.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3704.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596557.000 |
| redis | keyspace_hits | 1576691284.000 |
| redis | keyspace_misses | 403041.000 |
| redis | net_input_bytes | 67206133394.000 |
| redis | net_output_bytes | 407636122688.000 |
| redis | ops_per_sec | 33300.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 29181.000 |
| redis | used_memory_bytes | 102707672.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
