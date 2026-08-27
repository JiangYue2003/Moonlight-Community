# Feed 压测报告：hybrid / gateway / cursor-page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:49:00+08:00
- 采样时长：1m0.0116009s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 73975 | 73975 | 0 | 0 | 1232.79 | 12.342 | 18.218 | 21.483 | 27.300 | 49.639 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：105.6085ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 118 | 0.002 |
| redis | 74093 | 1.002 |
| relation | 240 | 0.003 |

- Cold compute：73975（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1553475 | 21.000 |
| merge_candidates | 11096250 | 150.000 |
| redis_commands | 1257575 | 17.000 |
| redis_members | 17458100 | 236.000 |
| redis_roundtrips | 147950 | 2.000 |
| tie_members | 9320850 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 73975 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 4.208 |
| cursor_decode | 73975 | 0.011 |
| cursor_seek | 73975 | 7.805 |
| hydrate | 73975 | 3.918 |
| merge_dedup | 73975 | 0.009 |
| relation | 240 | 5.178 |
| route | 73975 | 0.114 |
| total | 73975 | 11.860 |

## Redis 本轮边界增量

- Commands：1355460；input：191759921 bytes；output：994587793 bytes
- Hits/Misses：2818416/118；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：25333；safety epoch：3675 -> 3675

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.482 |
| client:loadtest | cpu_percent_total | 39.706 |
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
| docker:zg-canal | cpu_percent | 1.710 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.830 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.780 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 134.900 |
| docker:zg-kafka | memory_percent | 7.230 |
| docker:zg-kafka | pids | 114.000 |
| docker:zg-zk | cpu_percent | 45.030 |
| docker:zg-zk | memory_percent | 1.000 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 233073.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.961 |
| process:counter | cpu_seconds_total | 116.578 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46784512.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 157.967 |
| process:gateway | cpu_seconds_total | 1652.578 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51335168.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 140.759 |
| process:knowpost | cpu_seconds_total | 7939.312 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70590464.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.079 |
| process:relation | cpu_seconds_total | 17.812 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48394240.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 3.172 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38158336.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 51.881 |
| process:user-storage | cpu_seconds_total | 604.516 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 56242176.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 158330521.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3675.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596915.000 |
| redis | keyspace_hits | 1152365583.000 |
| redis | keyspace_misses | 258375.000 |
| redis | net_input_bytes | 50022968023.000 |
| redis | net_output_bytes | 309069235833.000 |
| redis | ops_per_sec | 25333.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26980.000 |
| redis | used_memory_bytes | 103797992.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
